package models

type User struct {
	ID                uint
	UserName          string
	PasswordEncrypted string
}
