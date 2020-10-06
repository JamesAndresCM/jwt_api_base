package routes

import (
	"github.com/JamesAndresCM/jwt_api_base/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter :=
		mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
