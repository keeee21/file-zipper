package controller

import (
	"net/http"

	"file-zipper-api/usecase"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(u *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

type googleLoginRequest struct {
	IDToken string `json:"idToken"`
}

func (h *AuthHandler) GoogleLogin(c echo.Context) error {
	var req googleLoginRequest
	if err := c.Bind(&req); err != nil || req.IDToken == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user, err := h.usecase.GoogleLogin(req.IDToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    user,
	})
}
