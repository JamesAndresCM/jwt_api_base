package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JamesAndresCM/jwt_api_base/commons"
	"github.com/JamesAndresCM/jwt_api_base/configuration"
	"github.com/JamesAndresCM/jwt_api_base/lib"
	"github.com/JamesAndresCM/jwt_api_base/models"
)

const (
	gravatarURI = "https://gravatar.com/avatar"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	lib.Fatal(err)

	db := configuration.GetConnection()
	defer db.Close()

	c := sha256.Sum256([]byte(user.Password))
	pwd := base64.URLEncoding.EncodeToString(c[:32])

	db.Where("email = (?) AND password = (?)", user.Email, pwd).First(&user)

	if user.ID > 0 {
		user.Password = ""
		token := commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})
		lib.Fatal(err)
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message:    "User or password are not valid",
			Statuscode: http.StatusUnauthorized,
		}
		commons.DisplayMessage(w, m)
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error user is wrong %s", err)
		m.Statuscode = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	if user.Password != user.ConfirmPassword {
		m.Message = "Password not mismatch"
		m.Statuscode = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)
	user.Password = pwd

	picmd5 := md5.Sum([]byte(user.Email))
	picstr := fmt.Sprintf("%x", picmd5)
	user.Picture = gravatarURI + picstr + "?s=100"

	db := configuration.GetConnection()
	defer db.Close()
	err = db.Create(&user).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error to register user %s", err)
		m.Statuscode = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "User create successfully"
	m.Statuscode = http.StatusCreated
	commons.DisplayMessage(w, m)
}
