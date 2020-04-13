package main

import (
	"flag"
	"log"

	"github.com/JamesAndresCM/jwt_api_base/migration"
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
}
