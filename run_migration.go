package main

import (
  "flag"
  "github.com/JamesAndresCM/jwt_api_base/migration"
  "log"
)

func main() {
  var migrate string
  flag.StringVar(&migrate, "migrate", "no", "Generating migrations")
  flag.Parse()

  if migrate == "yes" {
    log.Println("Start migrate process")
    migration.Migrate()
    log.Println("migration finished")
  }
}
