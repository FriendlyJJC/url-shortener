package apiv1

import (
	"encoding/json"
	"net/http"
)

func GenerateID() (string, error) {

}

type AddURLBody struct {
	LongURl  string  `json:"longurl"`
	ShortURL *string `json:"shorturl"`
}

func AddURL(w http.ResponseWriter, r *http.Request) {
	var ReqBody AddURLBody
	err := json.NewDecoder(r.Body).Decode(&ReqBody)
	if err != nil {
		http.Error(w, "JSON could not be Encoded, try again", http.StatusInternalServerError)
	}

}

func APIHandleV1() {
	v1_mux := http.NewServeMux()
}
