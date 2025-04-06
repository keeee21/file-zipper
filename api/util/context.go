package util

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetUserID(c echo.Context) (uint, error) {
	val := c.Get("userID")
	switch v := val.(type) {
	case float64:
		return uint(v), nil
	case int:
		return uint(v), nil
	case uint:
		return v, nil
	default:
		return 0, fmt.Errorf("invalid user id type: %T", v)
	}
}
