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
	fmt.Println(c)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"dbname=%s password=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Database, c.Password)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
