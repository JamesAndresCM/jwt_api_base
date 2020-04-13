package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Configuration struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)

	if err != nil {
		log.Fatal(err)
	}

	return c
}

func GetConnection() *gorm.DB {
	c := GetConfiguration()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Database)
	fmt.Println(psqlInfo)
	db, err := gorm.Open("jwt_api_base", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
