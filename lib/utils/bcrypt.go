package utils

import (
	"github.com/artnoi43/todong/enums"
	"golang.org/x/crypto/bcrypt"
)

func EncodeBcrypt(plain []byte) ([]byte, error) {
	if len(plain) < 6 {
		return nil, enums.ErrPwTooShort
	}
	return bcrypt.GenerateFromPassword(plain, 14)
}

func DecodeBcrypt(hashed, plain []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, plain)
}
