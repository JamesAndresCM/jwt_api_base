package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/JamesAndresCM/jwt_api_base/lib"
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
	var db database
	file, err := os.Open("./.env")
	lib.Fatal(err)
	defer file.Close()

	gotenv.Load()
	db.Database = os.Getenv("DB_NAME")
	db.Password = os.Getenv("DB_PASS")
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.User = os.Getenv("DB_USER")

	if db.Database == "" || db.Password == "" || db.Host == "" || db.Port == "" || db.User == "" {
		log.Fatal("env vars not defined")
	}

	return db
}

func GetConnection() *gorm.DB {
	c := getConfiguration()
	fmt.Println("new postgresql connection")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Host, c.Port, c.User, c.Database, c.Password)
	db, err := gorm.Open("postgres", psqlInfo)
	lib.Fatal(err)

	return db
}
