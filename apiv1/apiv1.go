package apiv1

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FriendlyJJC/api_server/db"
	"gorm.io/gorm"
)

var (
	ShortURLs ShortURLS
	DB        *gorm.DB
)

// Error Messages for HTTP Return
const (
	JSON_CONVERT_ERROR = "Something went wrong while converting to JSON. Please try again later"
	DB_WRITE_ERROR     = "Something went wrong while writing to DB. Please try again later"
	JSON_ENCODE_ERROR  = "JSON could not be encoded. Please try again later"
	JSON_DECODE_ERROR  = "JSON could not be decoded. Please try again later"
	NO_ENTRIES_ERROR   = "No Entries found in Database"
	DB_ERROR           = "Something went wrong while interacting with the Database. Please try again later"
	DB_DELETE_ERROR    = "Something went wrong while deleting from DB. Please try again later"
	DB_UPDATE_ERROR    = "Something went wrong while updating the DB. Please try again later"
)

// Initialize DB Connection and store connection into a variable
func init() {
	database := db.InitializeDB()
	if ok := db.Migrate(database); ok != true {
		fmt.Println("Something went wrong with the Schema Migration")
	}
	DB = database
}

// Function that uses Crypto.Rand to randomly assert random bytes and return the byte slice
func GenerateID() (id []byte, is_err bool) {
	const length = 4
	id = make([]byte, length)
	_, err := rand.Read(id)
	if err != nil {
		return nil, true
	}

	return id, false
}

func AddURL(w http.ResponseWriter, r *http.Request) {
	var ReqBody AddURLBody
	err := json.NewDecoder(r.Body).Decode(&ReqBody)
	if err != nil {
		http.Error(w, JSON_DECODE_ERROR, http.StatusInternalServerError)
		return
	}
	generated_id, _ := GenerateID()
	new_url := db.ShortUrls{Longurl: ReqBody.LongURL, Shorturl: fmt.Sprintf("%x", generated_id)}
	result := DB.Exec("INSERT INTO short_urls (longurl, shorturl) VALUES (?, ?)", new_url.Longurl, new_url.Shorturl)
	if result.Error != nil {
		http.Error(w, DB_WRITE_ERROR, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	res_text := fmt.Sprintf("LongURL: %s and ShortURL: %v were added", ReqBody.LongURL, fmt.Sprintf("%x", generated_id))
	w.Write([]byte(res_text))
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var urls []db.ShortUrls
	DB.Raw("SELECT * FROM short_urls").Scan(&urls)
	if len(urls) < 1 {
		http.Error(w, NO_ENTRIES_ERROR, http.StatusNotFound)
		return
	}
	res_json, json_err := json.Marshal(urls)
	if json_err != nil {
		http.Error(w, JSON_CONVERT_ERROR, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res_json)
}

func GetShortURL(w http.ResponseWriter, r *http.Request) {
	var url db.ShortUrls
	id := r.PathValue("id")
	if result := DB.Exec("SELECT * FROM short_urls WHERE shorturl = ?", id).Scan(&url); result.Error != nil {
		error_message := fmt.Sprintf("The URL with the ID: %s, was not found", id)
		http.Error(w, error_message, http.StatusNotFound)
		return
	}
	json_res, json_err := json.Marshal(url)
	if json_err != nil {
		http.Error(w, JSON_CONVERT_ERROR, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_res)
}

func DeleteURL(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	tx := DB.Begin()
	var url db.ShortUrls
	search_result := DB.Exec("SELECT * FROM short_urls WHERE shorturl = ?", id).Scan(&url)
	if search_result.Error != nil {
		http.Error(w, DB_ERROR, http.StatusInternalServerError)
		return
	}
	if url.Shorturl != "" || url.Longurl != "" {
		delete_result := tx.Exec("DELETE FROM short_urls WHERE shorturl = ?", id)
		if delete_result.Error != nil {
			tx.Rollback()
			http.Error(w, DB_DELETE_ERROR, http.StatusInternalServerError)
			return
		}
		tx.Commit()
		res_text := fmt.Sprintf("URL with Shorturl %v was deleted", id)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(res_text))
	}
}

func UpdateURL(w http.ResponseWriter, r *http.Request) {
	var reqBody AddURLBody
	var searched_url db.ShortUrls
	tx := DB.Begin()
	id := r.PathValue("id")
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, JSON_DECODE_ERROR, http.StatusInternalServerError)
		return
	}
	search_result := DB.Exec("SELECT * FROM short_urls WHERE shorturl = ?", id).Scan(&searched_url)
	if search_result.Error != nil {
		http.Error(w, DB_ERROR, http.StatusInternalServerError)
		return
	}
	//checks if both inputs are not empty
	if *reqBody.ShortURL != "" || reqBody.LongURL != "" {
		//if Shorturl is not empty, longer than 7 characters and is not the same as already
		if *reqBody.ShortURL != "" && len(*reqBody.ShortURL) < 12 && *reqBody.ShortURL != searched_url.Shorturl {
			update_shorturl := tx.Exec("UPDATE short_urls SET shorturl = ?", *reqBody.ShortURL)
			if update_shorturl.Error != nil {
				tx.Rollback()
				http.Error(w, DB_UPDATE_ERROR, http.StatusInternalServerError)
				return
			}
			tx.Commit()
			res_text := fmt.Sprintf("The Shorturl was succesfully updated to %v", *reqBody.ShortURL)
			w.Header().Set("Content-Type", "application")
			w.WriteHeader(200)
			w.Write([]byte(res_text))
		}
		//if Longurl is not empty and not the same as already
		if reqBody.LongURL != "" && reqBody.LongURL != searched_url.Longurl {
			update_longurl := tx.Exec("UPDATE short_urls SET shorturl = ?", *reqBody.ShortURL)
			if update_longurl.Error != nil {
				tx.Rollback()
				http.Error(w, DB_UPDATE_ERROR, http.StatusInternalServerError)
				return
			}
			tx.Commit()
			res_text := fmt.Sprintf("The Longurl was succesfully updated to %v", reqBody.LongURL)
			w.Header().Set("Content-Type", "application")
			w.WriteHeader(200)
			w.Write([]byte(res_text))
		}
		tx.Rollback()
		http.Error(w, "The Data does not fit the requirements. Please check the Docs for further Information", http.StatusNotAcceptable)
		return

	} else {
		http.Error(w, "No Data provided", http.StatusNoContent)
		return
	}
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
