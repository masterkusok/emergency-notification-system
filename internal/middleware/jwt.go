package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/masterkusok/emergency-notification-system/internal/handlers"
	"net/http"
)

func AuthJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customContext := new(handlers.AuthContext)
		customContext.Context = c

		tokenString := c.Request().Header.Get("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret-string"), nil
		})
		if err != nil {
			return c.JSON(http.StatusForbidden, handlers.GetForbiddenResponse())
		}

		claims := token.Claims.(jwt.MapClaims)
		if claims.Valid() != nil {
			return c.JSON(http.StatusForbidden, handlers.GetForbiddenResponse())
		}
		customContext.Id = uint(claims["id"].(float64))
		customContext.IsAuthenticated = true
		return next(customContext)
	}
}
