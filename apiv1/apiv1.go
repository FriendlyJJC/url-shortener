package apiv1

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func GenerateID() string {
	const length int8 = 7
	id := make([]byte, length)
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890_?!*+&"
	for i := 0; i < int(length); i++ {
		id[i] = characters[rand.Intn(len(characters))]
	}
	return string(id)
}

type AddURLBody struct {
	LongURL  string  `json:"longurl"`
	ShortURL *string `json:"shorturl"`
}

func AddURL(w http.ResponseWriter, r *http.Request) {
	var ReqBody AddURLBody
	err := json.NewDecoder(r.Body).Decode(&ReqBody)
	if err != nil {
		http.Error(w, "JSON could not be decoded, try again", http.StatusInternalServerError)
		return
	}
	generated_id := GenerateID()
	ReqBody.ShortURL = &generated_id
	fmt.Println(ReqBody)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ShortURL was successfully created"))
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It works"))
}

func APIHandleV1(w http.ResponseWriter, r *http.Request) {
	v1_mux := http.NewServeMux()
	v1_mux.HandleFunc("/shorturl/add", AddURL)
	v1_mux.HandleFunc("/test", Test)
	v1_mux.ServeHTTP(w, r)
}
