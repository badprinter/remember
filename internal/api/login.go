package api

import (
	"log"
	"net/http"
)

func (a *App) login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	userName := r.Form.Get("username")
	password := r.Form.Get("password")

	pass := encrypt_passsword(password)
	u := a.Repo.GetUserByUserName(userName)

	if pass == u.PasswordEncrypted {
		jwt, err := createJWT(u)

		if err != nil {
			log.Println("jwt error")
			return
		}
		// TODO create table Sessions and write it
		r.Header.Set("jwt", jwt)
	}

}
