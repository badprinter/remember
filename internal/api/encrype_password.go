package api

import "golang.org/x/crypto/bcrypt"

func encrypt_passsword(p string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 30)
	if err != nil {
		return ""
	}
	return string(hash)
}
