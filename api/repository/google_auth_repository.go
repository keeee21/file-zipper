package repository

import "file-zipper-api/model"

type IGoogleAuthRepository interface {
	VerifyIDToken(idToken string) (*model.GoogleAuthPayload, error)
}
