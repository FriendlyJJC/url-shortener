package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

var (
	ShortURLs ShortURLS
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

func RemoveItem(arr []AddURLBody, index int) ([]AddURLBody, error) {
	var urls []AddURLBody
	if arr == nil {
		return nil, errors.New("Given Array or Slice does not have elements")
	}
	for i, item := range arr {
		if i != index {
			urlbody := AddURLBody{ShortURL: item.ShortURL, LongURL: item.LongURL}
			urls = append(urls, urlbody)
		}
	}
	return urls, nil
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
		http.Error(w, "No ID Parameter was given. Please pass it for example like this /shorturl/get/wwhrgf", http.StatusBadRequest)
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

func DeleteURL(w http.ResponseWriter, r *http.Request) {
	itemIndex := 0
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "No ID Parameter was given. Please pass it for example like this /shorturl/get/wwhrgf", http.StatusBadRequest)
		return
	}

	for i, item := range ShortURLs.Data {
		if *item.ShortURL == id {
			itemIndex = i
		}
	}
	updatedURL, err := RemoveItem(ShortURLs.Data, itemIndex)
	if err != nil {
		res := fmt.Sprintf("Error Occurred: %v", err)
		http.Error(w, res, http.StatusInternalServerError)
	}
	ShortURLs.Data = updatedURL
	w.WriteHeader(http.StatusOK)
	res := fmt.Sprintf("URL was correctly deleted")
	w.Write([]byte(res))
}

func UpdateURL(w http.ResponseWriter, r *http.Request) {
	var userBody AddURLBody
	id := r.PathValue("id") // Correctly retrieve the "id" parameter from the query string

	if id == "" {
		http.Error(w, "No ID Parameter was given. Please pass it for example like this /shorturl/update?id=wwhrgf", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		http.Error(w, "JSON could not be decoded, try again", http.StatusBadRequest)
		return
	}

	var url_item *AddURLBody
	for i := range ShortURLs.Data {
		if *ShortURLs.Data[i].ShortURL == id {
			url_item = &ShortURLs.Data[i] // Update the reference directly in the slice
			break
		}
	}

	if url_item == nil {
		res := fmt.Sprintf("ShortURL with ID %v does not exist. Please use another id", id)
		http.Error(w, res, http.StatusNotFound)
		return
	}

	if len(userBody.LongURL) > 0 {
		url_item.LongURL = userBody.LongURL
	}
	if userBody.ShortURL != nil && len(*userBody.ShortURL) > 0 && len(*userBody.ShortURL) < 8 {
		url_item.ShortURL = userBody.ShortURL
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL was successfully updated"))
}

func APIHandleV1(w http.ResponseWriter, r *http.Request) {
	v1_mux := http.NewServeMux()
	v1_mux.HandleFunc("POST /shorturl/add", AddURL)
	v1_mux.HandleFunc("GET /shorturl/get", GetAll)
	v1_mux.HandleFunc("GET /shorturl/get/{id}", GetShortURL)
	v1_mux.HandleFunc("DELETE /shorturl/delete/{id}", DeleteURL)
	v1_mux.HandleFunc("PUT /shorturl/update/{id}", UpdateURL)
	v1_mux.ServeHTTP(w, r)
}
