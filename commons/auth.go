package commons

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/JamesAndresCM/jwt_api_base/lib"
	"github.com/JamesAndresCM/jwt_api_base/models"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "keys/private.rsa"
	pubKeyPath  = "keys/public.rsa.pub"
)

var (
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func init() {

	privateBytes, err := ioutil.ReadFile(privKeyPath)
	lib.Fatal(err)

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	lib.Fatal(err)

	publicBytes, err := ioutil.ReadFile(pubKeyPath)
	lib.Fatal(err)

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	lib.Fatal(err)
}

func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer: "Jwt Api Base",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	lib.Fatal(err)

	return result
}
