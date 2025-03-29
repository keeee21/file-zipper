package middleware

import (
	"file-zipper-api/util"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const UserIDKey = "userID"

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token format"})
		}

		tokenStr := parts[1]

		claims, err := util.ValidateJWT(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		// ユーザー情報をContextにセット
		c.Set(UserIDKey, claims.UserID)
		return next(c)
	}
}
