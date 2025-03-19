package util

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// ULIDを生成して返す
func GenerateULID() string {
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}
