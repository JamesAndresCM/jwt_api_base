package main

import (
	"flag"
	"github.com/JamesAndresCM/jwt_api_base/migration"
	"github.com/JamesAndresCM/jwt_api_base/routes"
	"github.com/urfave/negroni"
	"log"
  "net/http"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Generating migrations")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Start migrate process")
		migration.Migrate()
		log.Println("migration finished")
	}

	// initialize routes
	router := routes.InitRoutes()

	// initialize middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}

	log.Println("Server Start http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Fin execute")

}
