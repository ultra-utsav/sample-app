package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ultra-utsav/sample-app/driver"
	"github.com/ultra-utsav/sample-app/stores"
	"log"
	"net/http"
	"os"

	http2 "github.com/ultra-utsav/sample-app/http"
)

func main() {
	cfg := &driver.MySQLConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Database: os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	db, err := driver.InitConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./static")))

	messageStore := stores.New(db)
	messageHTTP := http2.New(messageStore)

	r.HandleFunc("/message", messageHTTP.Create).Methods(http.MethodPost)

	log.Println("listening on port :8000")
	http.ListenAndServe(":8000", r)
}
