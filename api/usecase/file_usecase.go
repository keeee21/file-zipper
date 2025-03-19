package usecase

import (
	"bytes"
	"context"
	"file-zipper-api/model"
	"file-zipper-api/repository"
	"file-zipper-api/util"
	"fmt"
	"log"
	"os"

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

	// 4 ダウンロードURLを返す
	return model.FileResponse{ID: file.ID, Name: file.Name}, nil
}
