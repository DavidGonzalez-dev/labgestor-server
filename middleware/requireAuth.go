package middleware

import (
	"labgestor-server/internal/repository"
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
			tokenCookie, err := c.Cookie("sesionUsuario")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token no encontrado, ingreso no permitido"})
			}

			// Hacer el Decode y Validarla
			tokenString := tokenCookie.Value
	
			// Decode Token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				return []byte(os.Getenv("SECRET")), nil
			}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error el analizar token", "error": err.Error()})
			}
	
			// Verificar Token
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				// Verificar Expiracion
				if time.Now().Unix() > int64(claims["exp"].(float64)) {
					return c.JSON(http.StatusUnauthorized, map[string]string{"message": "El token ya expiro"})
				}
	
				// Obtenemos el ID del token
				userID, ok := claims["userID"].(string)
				if !ok {
					return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error al analizar el userID del token"})
				}
	
				// Verificamos que el usuario exista o que este activo
				usuario := repo.ObtenerUsuarioID(userID)
				if usuario.ID == "0" || !usuario.Estado {
					return c.JSON(http.StatusUnauthorized, map[string]string{"message": "El usuario no esta activo"})
				}

				// Verificar que el rol que esta intentando acceder este permitido
				if usuario.Rol.NombreRol != rolPermitido{
					return c.JSON(http.StatusForbidden, map[string]string{"message": "No tienes permitido el acceso", "rol requerido": rolPermitido})
				}
	
			} else {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "El token ya expiro"})
			}
	
			// Continuar
			return next(c)
		}
	}
}
