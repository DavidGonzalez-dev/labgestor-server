package middleware

import (
	"fmt"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)


func RequireAuth(repo repository.UsuarioRepository, rolPermitido string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context) error {
			// Obtener la cookie del request
			tokenCookie, err := c.Cookie("authToken")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Ingreso no autorizado"})
			}

			// Hacer el Decode y Validarla
			tokenString := tokenCookie.Value
	
			// Decode Token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				return []byte(os.Getenv("SECRET")), nil
			}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	
			if err != nil {
				return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al decodificar el token", Error: err.Error()})
			}
	
			// Verificar Token
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				// Verificar Expiracion
				if time.Now().Unix() > int64(claims["exp"].(float64)) {
					return c.JSON(http.StatusUnauthorized, response.Response{Message: "Ingreso no autorizado: Token expirado"})
				}
	
				// Obtenemos el ID del token
				userID, ok := claims["userID"].(string)
				if !ok {
					return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al verificar token: campo 'userID' no encontrado"})
				}
	
				// Verificamos que el usuario exista o que este activo
				usuario, err := repo.ObtenerUsuarioID(userID)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener usuario", Error: err.Error()})
				}
				if usuario.ID == "0" || !usuario.Estado {
					return c.JSON(http.StatusUnauthorized, response.Response{Message: "Ingreso no autorizado: El usuario esta inhabilitado "})
				}

				// Verificar que el rol que esta intentando acceder este permitido
				if usuario.Rol.NombreRol != rolPermitido{

					return c.JSON(http.StatusForbidden, response.Response{Message: "Ingreso no permitido", Error: fmt.Sprintf("Rol requerido para acceso %s", rolPermitido)})
				}
	
			} else {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Ingreso no permitido: Token expirado"})
			}
	
			// Continuar
			return next(c)
		}
	}
}
