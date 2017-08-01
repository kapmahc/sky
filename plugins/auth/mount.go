package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Mount mount web points
func (p *Plugin) Mount(rt *mux.Router) {
	users := rt.PathPrefix("/users").Subrouter()
	users.HandleFunc("/sign-in", p.getUsersSignIn).Methods(http.MethodGet)

}
