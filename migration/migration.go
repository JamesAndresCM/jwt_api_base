package migration

import (
	"fmt"

	"github.com/JamesAndresCM/jwt_api_base/configuration"
	"github.com/JamesAndresCM/jwt_api_base/models"
)

func Migrate() {
	db := configuration.GetConnection()
	fmt.Println(db)
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")
}
