package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/subosito/gotenv"
)

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func getConfiguration() database {
	var c database
	file, err := os.Open("./.env")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gotenv.Load()
	fmt.Println(file)
	c.Database = os.Getenv("DB_NAME")
	c.Password = os.Getenv("DB_PASS")
	c.Host = os.Getenv("DB_HOST")
	c.Port = os.Getenv("DB_PORT")
	c.User = os.Getenv("DB_USER")

	if c.Database == "" || c.Password == "" || c.Host == "" || c.Port == "" || c.User == "" {
		log.Fatal("env vars not defined")
	}

	return c
}

func GetConnection() *gorm.DB {
	c := getConfiguration()
	fmt.Println("new postgresql connection")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Host, c.Port, c.User, c.Database, c.Password)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
