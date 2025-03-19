package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword は平文のパスワードを bcrypt でハッシュ化する
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("❌ パスワードのハッシュ化に失敗:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword は入力されたパスワードとハッシュ化されたパスワードを比較する
func ComparePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
