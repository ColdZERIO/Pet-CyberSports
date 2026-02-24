package service

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type ArgonParams struct {
	Memory      uint32
	Time        uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var DefaultParams = ArgonParams{
	Memory:      64 * 1024,
	Time:        3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func HashPassword(password string, ap ArgonParams) (string, error) {
	if password == "" {
		return "", errors.New("empty password")
	}

	salt := make([]byte, ap.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("read salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), salt, ap.Time, ap.Memory, ap.Parallelism, ap.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s", ap.Memory, ap.Time, ap.Parallelism, b64Salt, b64Hash)

	return encoded, nil
}

func VerifyPassword(password, encoded string) (bool, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return false, errors.New("invalid encoded hash format")
	}

	var (
		mem  uint32
		time uint32
		par  uint8
	)

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &mem, &time, &par)
	if err != nil {
		return false, errors.New("invalid params section")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, errors.New("invalid salt b64")
	}

	want, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, errors.New("invalid hash b64")
	}

	got := argon2.IDKey([]byte(password), salt, time, mem, par, uint32(len(want)))

	if subtle.ConstantTimeCompare(got, want) == 1 {
		return true, nil
	}

	return false, nil
}
