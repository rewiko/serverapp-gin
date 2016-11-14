package gorillamux

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rewiko/app/libs/jsonapi/api2go/routing"
	"net/http"
	"strings"
)

type gorillamuxRouter struct {
	router *mux.Router
}

func (gm gorillamuxRouter) Handler() http.Handler {
	return gm.router
}

func (gm gorillamuxRouter) Handle(protocol, route string, handler routing.HandlerFunc) {
	wrappedHandler := func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, mux.Vars(r))
	}

	// The request path will have parameterized segments indicated as :name.  Convert
	// that notation to the {name} notation used by Gorilla mux.
	orig := strings.Split(route, "/")
	var mod []string
	for _, s := range orig {
		if len(s) > 0 && s[0] == ':' {
			s = fmt.Sprintf("{%s}", s[1:])
		}
		mod = append(mod, s)
	}
	modroute := strings.Join(mod, "/")

	gm.router.HandleFunc(modroute, wrappedHandler).Methods(protocol)
}

//New creates a new api2go router to use with the Gorilla mux framework
func New(gm *mux.Router) routing.Routeable {
	return &gorillamuxRouter{router: gm}
}
