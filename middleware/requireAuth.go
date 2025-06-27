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
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			//? ----------------------------------------------------------------------------------
			//? Se obtiene la cookie del navegador del cliente
			//? ----------------------------------------------------------------------------------
			tokenCookie, err := c.Cookie("authToken")
			println("COOKIES ENVIADAS", c.Cookies())
			// Caso: La cookie no existe y se devuelve un error de autenticacion
			if err != nil {
				fmt.Println(err)
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Ingreso no autorizado", Error: "Token no encontrado"})
			}

			//? ----------------------------------------------------------------------------------
			//? Hacemos el decode del token
			//? ----------------------------------------------------------------------------------
			tokenString := tokenCookie.Value
			// Decode Token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

			// Se verifica que no hallan habido errores al decodificar el token
			if err != nil {
				return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al decodificar el token", Error: err.Error()})
			}

			//? ----------------------------------------------------------------------------------
			//? Se hace la verificacion del token
			//? ----------------------------------------------------------------------------------
			// Se accede a los claims para validarlos
			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Ingreso no permitido: Token invalido"})
			}

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
				return c.JSON(http.StatusInternalServerError, response.Response{Message: "Este usuario no existe", Error: err.Error()})
			}
			if usuario.ID == "0" || !usuario.Estado {
				return c.JSON(http.StatusUnauthorized, response.Response{Message: "Ingreso no autorizado: El usuario esta inhabilitado "})
			}

			// Verificar si se pide un rol especifico para entrar a la ruta
			if rolPermitido != "" {

				// En caso de que se pase un rol se verifica si el usuario tiene este rol
				if usuario.Rol.NombreRol != rolPermitido {
					return c.JSON(http.StatusUnauthorized, response.Response{Message: "Ingreso no permitido", Error: fmt.Sprintf("Rol requerido para acceso %s", rolPermitido)})
				}
			}

			// Aqui se guardan variables en el contexto de echo para que puedan ser accedidas en los handlers
			c.Set("userID", usuario.ID) // Guardamos el ID del usuario en el contexto

			// Continua con el siguiente handler en caso de que el token sea valido
			return next(c)
		}
	}
}
