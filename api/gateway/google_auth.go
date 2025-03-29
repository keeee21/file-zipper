package gateway

import (
	"context"
	"fmt"
	"os"

	"file-zipper-api/model"

	"google.golang.org/api/idtoken"
)

type googleAuthRepo struct{}

func NewGoogleAuthRepository() model.GoogleAuthRepository {
	return &googleAuthRepo{}
}

func (r *googleAuthRepo) VerifyIDToken(idToken string) (*model.GoogleAuthPayload, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	payload, err := idtoken.Validate(context.Background(), idToken, clientID)
	if err != nil {
		return nil, fmt.Errorf("token validation failed: %w", err)
	}

	return &model.GoogleAuthPayload{
		Email: fmt.Sprintf("%v", payload.Claims["email"]),
		Name:  fmt.Sprintf("%v", payload.Claims["name"]),
		Sub:   payload.Subject,
	}, nil
}
