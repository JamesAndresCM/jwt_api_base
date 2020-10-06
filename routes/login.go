package routes

import (
  "github.com/gorilla/mux"
  "github.com/JamesAndresCM/jwt_api_base/controllers"
)

// router to login
func SetLoginRouter(router *mux.Router) {
  router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
