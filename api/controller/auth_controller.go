package controller

import (
	"net/http"

	"file-zipper-api/usecase"
	"file-zipper-api/util"

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

func (h *AuthHandler) GetUserInfo(c echo.Context) error {
	// ミドルウェアでセットされた userID を取得
	userIDInterface := c.Get("userID")
	if userIDInterface == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	userID, err := util.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid user id"})
	}

	// DBからユーザー情報を取得
	user, err := h.usecase.UserRepo.FindByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch user"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
