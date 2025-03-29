package usecase

import (
	"file-zipper-api/model"
	"file-zipper-api/repository"
)

type AuthUsecase struct {
	UserRepo       repository.IUserRepository
	GoogleAuthRepo repository.IGoogleAuthRepository
}

func NewAuthUsecase(userRepo repository.IUserRepository, googleAuthRepo repository.IGoogleAuthRepository) *AuthUsecase {
	return &AuthUsecase{
		UserRepo:       userRepo,
		GoogleAuthRepo: googleAuthRepo,
	}
}

func (u *AuthUsecase) GoogleLogin(idToken string) (*model.UserResponse, error) {
	// 1. IDトークンの検証
	payload, err := u.GoogleAuthRepo.VerifyIDToken(idToken)
	if err != nil {
		return nil, err
	}

	// 2. ユーザーが存在するかチェック
	user, err := u.UserRepo.FindByGoogleSub(payload.Sub)
	if err != nil {
		// 存在しないなら新規作成
		user = &model.User{
			GoogleSub: payload.Sub,
			Email:     payload.Email,
			Name:      payload.Name,
		}
		if err := u.UserRepo.Create(user); err != nil {
			return nil, err
		}
	} else {
		// Optional: email や name を毎回更新したい場合
		user.Email = payload.Email
		user.Name = payload.Name
		_ = u.UserRepo.Update(user) // Note: 更新に失敗しても致命的じゃないのでエラー握りつぶしても良いかも
	}

	return &model.UserResponse{
		GoogleSub: user.GoogleSub,
		Email:     user.Email,
		Name:      user.Name,
	}, nil
}
