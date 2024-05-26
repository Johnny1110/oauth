package utils

import (
	"golang.org/x/crypto/bcrypt"
	"oauth/sys"
)

func checkPassword(plaintext string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintext))
	if err != nil {
		sys.Logger().Warning("Password does not match:", err)
		return false
	}
	sys.Logger().Debug("Password matched!")
	return true
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)

}
