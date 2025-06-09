package apiv1

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

var ShortURLs ShortURLS

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
type ShortURLS struct {
	Data []AddURLBody
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
	w.WriteHeader(http.StatusOK)
	res_text := fmt.Sprintf("The LongURl is %s and the ShortURl is %s", ReqBody.LongURL, *ReqBody.ShortURL)
	w.Write([]byte(res_text))
	ShortURLs.Data = append(ShortURLs.Data, ReqBody)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(ShortURLs)
	if err != nil {
		http.Error(w, "Something went Wrong while Encoding the Response. Try again", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetShortURL(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "No ID Parameter was given. Please pass it for example like this /shorturl/get?id=wwhrgf", http.StatusBadRequest)
		return
	}

	var searchedItem *AddURLBody
	for _, data := range ShortURLs.Data {
		if *data.ShortURL == id {
			searchedItem = &data
			break
		}
	}

	if searchedItem == nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(searchedItem)
	if err != nil {
		http.Error(w, "Something went wrong while encoding the response. Try again", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func APIHandleV1(w http.ResponseWriter, r *http.Request) {
	v1_mux := http.NewServeMux()
	v1_mux.HandleFunc("POST /shorturl/add", AddURL)
	v1_mux.HandleFunc("GET /shorturl/get", GetAll)
	v1_mux.HandleFunc("GET /shorturl/get/{id}", GetShortURL)
	v1_mux.ServeHTTP(w, r)
}
