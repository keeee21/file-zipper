package usecase

import (
	"file-zipper-api/model"
)

type AuthUsecase struct {
	googleRepo model.GoogleAuthRepository
}

func NewAuthUsecase(repo model.GoogleAuthRepository) *AuthUsecase {
	return &AuthUsecase{googleRepo: repo}
}

func (u *AuthUsecase) GoogleLogin(idToken string) (*model.GoogleAuthPayload, error) {
	return u.googleRepo.VerifyIDToken(idToken)
}
