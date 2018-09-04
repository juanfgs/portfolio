package routes

import(
	"github.com/juanfgs/portfolio/handlers"
	"net/http"
)

type Route struct {
	Destination string
	Handler func( w http.ResponseWriter, r *http.Request) ()
	Method string
}

//Regular routes
func Get() []Route{
	return []Route{
		{ "/", handlers.HomeHandler{}.Index,"GET"},
		{ "/contact", handlers.ContactHandler{}.Submit,"POST"},
	}
}
