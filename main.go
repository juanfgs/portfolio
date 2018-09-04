package main

import (
	"github.com/gorilla/mux"
	"github.com/juanfgs/portfolio/routes"
	"net/http"
)

func main() {
	mux := mux.NewRouter()

	/* Handle CSS and JS paths */
	jsHandler := http.FileServer(http.Dir("./static/js/"))
	cssHandler := http.FileServer(http.Dir("./static/css/"))
	imageHandler := http.FileServer(http.Dir("./static/img/"))
	vendorHandler := http.FileServer(http.Dir("./static/vendor/"))

	mux.PathPrefix("/js/").Handler(http.StripPrefix("/js/", jsHandler))
	mux.PathPrefix("/img/").Handler(http.StripPrefix("/img/", imageHandler))
	mux.PathPrefix("/vendor/").Handler(http.StripPrefix("/vendor/", vendorHandler))
	mux.PathPrefix("/css/").Handler(http.StripPrefix("/css/", cssHandler))
	/* Load Routes */
	for _, route := range routes.Get() {
		mux.HandleFunc(route.Destination,
			route.Handler).Methods(route.Method)
	}
	http.Handle("/", mux)
	http.ListenAndServe("localhost:6060", mux)

}
