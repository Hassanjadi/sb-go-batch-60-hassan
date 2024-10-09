package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type NilaiMahasiswa struct {
	ID, Nilai          uint   
	Name, MataKuliah, IndeksNilai  string
	 
	       
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func hitungIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

func checkBasicAuth(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	if !strings.HasPrefix(authHeader, "Basic ") {
		return false
	}

	encodedCreds := strings.TrimPrefix(authHeader, "Basic ")
	decodedCreds, err := base64.StdEncoding.DecodeString(encodedCreds)
	if err != nil {
		return false
	}

	creds := strings.SplitN(string(decodedCreds), ":", 2)
	if len(creds) != 2 {
		return false
	}

	username, password := creds[0], creds[1]

	if username == "admin" && password == "admin" {
		return true
	}

	return false
}

func GetNilaiMahasiswa(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		dataNilai, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(rw, "Unable to marshal data", http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(dataNilai)
		return
	}

	http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
}

func PostNilaiMahasiswa(rw http.ResponseWriter, r *http.Request) {
	if !checkBasicAuth(r) {
		rw.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	var newNilaiMahasiswa = NilaiMahasiswa{ID: uint(len(nilaiNilaiMahasiswa) + 1)}

	if r.Method == http.MethodPost {
		if r.Header.Get("Content-Type") == "application/json" {
			// Parse dari JSON
			decodedJson := json.NewDecoder(r.Body)
			if err := decodedJson.Decode(&newNilaiMahasiswa); err != nil {
				http.Error(rw, "Invalid JSON format", http.StatusBadRequest)
				return
			}
		} else {
			// Parse dari form data
			name := r.PostFormValue("name")
			mataKuliah := r.PostFormValue("mataKuliah")
			getNilai := r.PostFormValue("nilai")

			nilai, err := strconv.Atoi(getNilai)
			if err != nil {
				http.Error(rw, "Invalid value for 'nilai'", http.StatusBadRequest)
				return
			}

			newNilaiMahasiswa.Name = name
			newNilaiMahasiswa.MataKuliah = mataKuliah
			newNilaiMahasiswa.Nilai = uint(nilai)
		}

		if newNilaiMahasiswa.Nilai > 100 {
			http.Error(rw, "Nilai cannot exceed 100", http.StatusBadRequest)
			return
		}

		newNilaiMahasiswa.IndeksNilai = hitungIndeksNilai(newNilaiMahasiswa.Nilai)

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilaiMahasiswa)
		dataNilai, _ := json.Marshal(newNilaiMahasiswa)

		rw.WriteHeader(http.StatusCreated)
		rw.Write(dataNilai)
		return
	}

	http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/nilai", GetNilaiMahasiswa)
	http.HandleFunc("/post-nilai", PostNilaiMahasiswa)

	fmt.Println("Server is running at 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
