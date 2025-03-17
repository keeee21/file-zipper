package usecase

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"file-zipper-api/model"
	"file-zipper-api/repository"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type IFileUsecase interface {
	Upload(file *model.File, fileData []byte, fileExt string) (model.FileResponse, error)
}

type fileUsecase struct {
	fileRepo    repository.IFileRepository
	minioClient *minio.Client
}

func NewFileUsecase(fileRepo repository.IFileRepository) IFileUsecase {
	// MinIO クライアントを作成
	minioClient, err := minio.New("s3:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("MinIO クライアントの作成に失敗しました: %v", err)
	}

	return &fileUsecase{
		fileRepo:    fileRepo,
		minioClient: minioClient,
	}
}

func (fu *fileUsecase) Upload(file *model.File, fileData []byte, fileExt string) (model.FileResponse, error) {
	// 1️⃣ パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(file.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.FileResponse{}, err
	}
	file.Password = string(hash)

	// 2️⃣ 一意のファイル名を生成
	randomBytes := make([]byte, 16)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return model.FileResponse{}, err
	}
	fileName := fmt.Sprintf("%s%s", hex.EncodeToString(randomBytes), fileExt)

	// 3️⃣ MinIO にアップロード
	ctx := context.Background()
	bucketName := os.Getenv("BUCKET_NAME")
	_, err = fu.minioClient.PutObject(ctx, bucketName, fileName,
		bytes.NewReader(fileData), int64(len(fileData)), minio.PutObjectOptions{})
	if err != nil {
		return model.FileResponse{}, err
	}

	// 4️⃣ DB に保存
	file.FilePath = fmt.Sprintf("%s/%s", bucketName, fileName)
	file.DownloadUrl = fmt.Sprintf("/download/%d", file.ID)
	file.ExpiredAt = time.Now().Add(24 * time.Hour) // 24時間後に失効
	err = fu.fileRepo.CreateFile(file)
	if err != nil {
		return model.FileResponse{}, err
	}

	// 5️⃣ ダウンロードURLを返す
	return model.FileResponse{ID: file.ID, DownloadUrl: file.DownloadUrl}, nil
}
