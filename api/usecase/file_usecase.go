package usecase

import (
	"bytes"
	"context"
	"file-zipper-api/model"
	"file-zipper-api/repository"
	"file-zipper-api/util"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type IFileUsecase interface {
	GetFileNamesByRoomId(roomID string) ([]string, error)
	GetFileByRoomId(roomID string) ([]model.File, error)
	Upload(file *model.File, fileData []byte, fileExt string) (model.FileResponse, error)
	CreateDownloadRoom(file *model.File, password string, expirationDays int) (*model.DownloadRoom, error)
	CreateRoomFile(roomID string, fileID uint) error
	GetSignedUrl(fileId string) (string, error)
	VerifyRoomPassword(roomID, password string) (bool, error)
}

type fileUsecase struct {
	fileRepo     repository.IFileRepository
	roomRepo     repository.IDownloadRoomRepository
	roomFileRepo repository.IRoomFilesRepository
	minioClient  *minio.Client
}

func NewFileUsecase(
	fileRepo repository.IFileRepository,
	roomRepo repository.IDownloadRoomRepository,
	roomFileRepo repository.IRoomFilesRepository,
) IFileUsecase {
	// MinIO クライアントを作成
	minioClient, err := minio.New("s3:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("MinIO クライアントの作成に失敗しました: %v", err)
	}

	return &fileUsecase{
		fileRepo:     fileRepo,
		roomRepo:     roomRepo,
		roomFileRepo: roomFileRepo,
		minioClient:  minioClient,
	}
}

func (fu *fileUsecase) GetFileNamesByRoomId(roomId string) ([]string, error) {
	fileIds, err := fu.roomFileRepo.GetFileIdsByRoomId(roomId)
	if err != nil {
		return nil, fmt.Errorf("failed to get fileIds: %w", err)
	}
	if len(fileIds) == 0 {
		return nil, fmt.Errorf("no files associated with roomId: %s", roomId)
	}

	var names []string
	for _, id := range fileIds {
		file, err := fu.fileRepo.GetFileById(id)
		if err != nil {
			return nil, fmt.Errorf("failed to get file %d: %w", id, err)
		}
		names = append(names, file.OriginalName)
	}
	return names, nil
}

func (fu *fileUsecase) GetFileByRoomId(roomId string) ([]model.File, error) {
	fileIds, err := fu.roomFileRepo.GetFileIdsByRoomId(roomId)
	if err != nil {
		return nil, fmt.Errorf("failed to get fileIds: %w", err)
	}
	if len(fileIds) == 0 {
		return nil, fmt.Errorf("no files associated with roomId: %s", roomId)
	}

	var files []model.File
	for _, id := range fileIds {
		file, err := fu.fileRepo.GetFileById(id)
		if err != nil {
			return nil, fmt.Errorf("failed to get file %d: %w", id, err)
		}
		files = append(files, file)
	}
	return files, nil
}

func (fu *fileUsecase) Upload(file *model.File, fileData []byte, fileExt string) (model.FileResponse, error) {
	// 1 一意のファイル名を生成
	fileName := fmt.Sprintf("%s", util.GenerateULID())

	// 2 MinIO にアップロード
	ctx := context.Background()
	bucketName := os.Getenv("BUCKET_NAME")
	_, err := fu.minioClient.PutObject(ctx, bucketName, fileName,
		bytes.NewReader(fileData), int64(len(fileData)), minio.PutObjectOptions{})
	if err != nil {
		return model.FileResponse{}, err
	}

	// 3 DB に保存
	file.Name = fmt.Sprintf("%s", fileName)
	file.OriginalName = fileExt
	err = fu.fileRepo.CreateFile(file)
	if err != nil {
		return model.FileResponse{}, err
	}

	return model.FileResponse{ID: file.ID, Name: file.Name}, nil
}

func (fu *fileUsecase) CreateDownloadRoom(file *model.File, pass string, expirationDays int) (*model.DownloadRoom, error) {
	roomID := util.GenerateULID()

	var hashedPassword *string
	if pass != "" {
		hashed, err := util.HashPassword(pass)
		if err != nil {
		} else {
			hashedPassword = &hashed
		}
	}

	// ルームの有効期限を設定
	expiredAt := time.Now().Add(time.Duration(expirationDays) * 24 * time.Hour)

	room := &model.DownloadRoom{
		ID:        roomID,
		Password:  hashedPassword,
		ExpiredAt: expiredAt,
	}

	err := fu.roomRepo.CreateRoom(room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (fu *fileUsecase) GetSignedUrl(fileName string) (string, error) {
	ctx := context.Background()
	bucketName := os.Getenv("BUCKET_NAME")

	reqParams := make(url.Values)
	presignedUrl, err := fu.minioClient.PresignedGetObject(
		ctx,
		bucketName,
		fileName,
		time.Minute*15,
		reqParams,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %w", err)
	}

	// ✅ URLのhost部分を書き換える
	presignedUrl.Host = "localhost:9000"

	// ✅ 必要ならSchemeも明示的にhttpに
	presignedUrl.Scheme = "http"

	return presignedUrl.String(), nil
}

func (fu *fileUsecase) CreateRoomFile(roomID string, fileID uint) error {
	err := fu.roomFileRepo.Create(roomID, fileID)
	if err != nil {
		return fmt.Errorf("failed to create room_file: %w", err)
	}
	return nil
}

func (fu *fileUsecase) VerifyRoomPassword(roomID, password string) (bool, error) {
	room, err := fu.roomRepo.GetRoomByID(roomID)
	if err != nil {
		return false, fmt.Errorf("failed to get room: %w", err)
	}
	if room == nil {
		return false, fmt.Errorf("room not found")
	}

	if room.Password == nil {
		return true, nil
	}

	isValid := util.ComparePassword(*room.Password, password)
	if !isValid {
		return false, fmt.Errorf("invalid password")
	}

	return true, nil
}
