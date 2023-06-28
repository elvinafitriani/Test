package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"test/connection"
	"test/model"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.Connection()
}

func Salam(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var kata = "Hello, "

		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			salam := kata + "World!!"
			w.Header().Set("Content-Type", "application/json")
			w.Write(Decode(salam))
			w.WriteHeader(200)
			return
		}
		salam := kata + id[2]
		w.Header().Set("Content-Type", "application/json")
		w.Write(Decode(salam))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Foundd", 404)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []model.Product
		DB.Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}

	http.Error(w, "Error Not Found", 404)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data model.Product
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Error Create", 500)
			return
		}

		response := map[string]interface{}{
			"message": "Success Insert",
			"data":    data,
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}

		w.Write(responseJSON)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func Decode(a string) []byte {
	data := a
	datajson, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return datajson
}
