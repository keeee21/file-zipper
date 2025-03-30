package usecase

import (
	"file-zipper-api/repository"
	"time"
)

type IDownloadRoomUsecase interface {
	IsDownloadRoomValid(roomID string) (bool, error)
}

type DownloadRoomUsecase struct {
	roomRepo repository.IDownloadRoomRepository
}

func NewDownloadRoomUsecase(repo repository.IDownloadRoomRepository) IDownloadRoomUsecase {
	return &DownloadRoomUsecase{roomRepo: repo}
}

func (u *DownloadRoomUsecase) IsDownloadRoomValid(roomID string) (bool, error) {
	room, err := u.roomRepo.GetRoomByID(roomID)
	if err != nil {
		return false, err
	}

	// 期限切れのチェック
	// Note: isDeletedも将来的には追加して判定したい
	if room.ExpiredAt.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}
