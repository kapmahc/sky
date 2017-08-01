package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Mount mount web points
func (p *Plugin) Mount(rt *mux.Router) {
	ung := rt.PathPrefix("/users").Subrouter()
	ung.HandleFunc("/sign-in", p.Wrapper.JSON(p.postUsersSignIn)).Methods(http.MethodPost)
	ung.HandleFunc("/sign-up", p.Wrapper.JSON(p.postUsersSignUp)).Methods(http.MethodPost)
	ung.HandleFunc("/confirm", p.Wrapper.JSON(p.postUsersConfirm)).Methods(http.MethodPost)
	ung.HandleFunc("/confirm/{token}", p.Wrapper.JSON(p.getUsersConfirm)).Methods(http.MethodGet)
	ung.HandleFunc("/unlock", p.Wrapper.JSON(p.postUsersUnlock)).Methods(http.MethodPost)
	ung.HandleFunc("/unlock/{token}", p.Wrapper.JSON(p.getUsersUnlock)).Methods(http.MethodGet)

	umg := mux.NewRouter()
	umg.HandleFunc("/info", p.Wrapper.JSON(p.getUsersInfo)).Methods(http.MethodGet)
	umg.HandleFunc("/info", p.Wrapper.JSON(p.postUsersInfo)).Methods(http.MethodPost)
	rt.PathPrefix("/users").Handler(negroni.New(
		p.MustSignInMiddleware,
		negroni.Wrap(umg),
	))
}
