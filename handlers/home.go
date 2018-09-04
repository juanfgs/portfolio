package handlers

import (
	"net/http"
)

type HomeHandler struct {
	BaseHandler
}


func (t HomeHandler ) Index(w http.ResponseWriter, r *http.Request){
	t.AddTemplate( "layout.tpl")
	t.AddTemplate( "index.tpl")
	t.Render(w, "")
}
