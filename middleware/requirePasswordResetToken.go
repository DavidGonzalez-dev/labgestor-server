package middleware

import (
	"fmt"
	"labgestor-server/utils/response"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RequireResetPasswordToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// Obtenemos la cookie del navegaodor
			tokenCookie, err := c.Cookie("resetPasswordToken")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Accion no valida", Error: "Token de autorizacion no encontrado"})
			}

			// Hacemos el decode del token
			tokenString := tokenCookie.Value
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

			if err != nil {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Accion no valida", Error: "Token invalido"})
			}

			// Obtenemos los claims del token
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				fmt.Println(claims)
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Accion no valida", Error: "Token invalido"})
			}

			// Verificamos que el token no haya expirado
			exp, ok := claims["exp"].(float64)
			if !ok {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Accion no valida", Error: "Token invalido"})
			}
			if int64(exp) < time.Now().Unix() {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Accion no valida", Error: "Token expirado"})
			}

			// Verificamos que el token no halla sido usado
			if claims["used"] == true {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Accion no valida", Error: "Este token ya ha sido usado"})
			}

			// Guardamos el id de usario en el contexto
			userId, ok := claims["userID"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Accion no valida", Error: "Token Invalido"})
			}

			c.Set("passwordTokenUserId", userId)

			return next(c)
		}
	}
}
