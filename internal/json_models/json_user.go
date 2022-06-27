package jsonmodels

import "fmt"

type JsonUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u JsonUser) String() string {
	return fmt.Sprintf("%s %s", u.UserName, u.Password)
}
