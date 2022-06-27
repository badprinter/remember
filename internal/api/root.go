package api

import "net/http"

func (a *App) root(w http.ResponseWriter, r *http.Request) {
	w
	w.Write("Hello root")
}
