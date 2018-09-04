package handlers

import (
	"net/http"
)

type ContactHandler struct {
	BaseHandler
}


func (t ContactHandler) Submit(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	w.WriteHeader(http.StatusOK)
}
