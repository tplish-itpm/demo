package utils

import (
	"github.com/tplish-itpm/demo/pkg/e"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(pwd string) (string, e.Error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", e.ErrNewCode(e.ErrEncrypt)
	}
	return string(hash), nil
}

func Compare(pwd string, hash string) e.Error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)); err != nil {
		return e.ErrNewErr(err)
	}
	return nil
}
