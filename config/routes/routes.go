package routes

import (
//	"../../config"
	"net/http"
)

func RedirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}
