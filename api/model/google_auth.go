package model

type GoogleAuthPayload struct {
	Email string
	Name  string
	Sub   string
}

type GoogleAuthRepository interface {
	VerifyIDToken(idToken string) (*GoogleAuthPayload, error)
}
