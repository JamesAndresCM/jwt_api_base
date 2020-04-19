package commons

import (
	"encoding/json"
	"net/http"

	"github.com/JamesAndresCM/jwt_api_base/lib"
	"github.com/JamesAndresCM/jwt_api_base/models"
)

func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	lib.Fatal(err)
	w.WriteHeader(m.Statuscode)
	w.Write(j)
}
