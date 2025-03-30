package controller

import (
	"file-zipper-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DownloadRoomController struct {
	roomUsecase usecase.IDownloadRoomUsecase
}

func NewDownloadRoomController(roomUsecase usecase.IDownloadRoomUsecase) *DownloadRoomController {
	return &DownloadRoomController{roomUsecase: roomUsecase}
}

func (rc *DownloadRoomController) CheckDownloadRoomValidity(c echo.Context) error {
	roomID := c.Param("roomID")
	isValid, err := rc.roomUsecase.IsDownloadRoomValid(roomID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"data": map[string]bool{"isValid": false},
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]bool{"isValid": isValid},
	})
}
