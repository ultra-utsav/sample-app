package http

import (
	"github.com/ultra-utsav/sample-app/models"
	"github.com/ultra-utsav/sample-app/stores"
	"log"
	"net/http"
)

type Message struct {
	store *stores.Message
}

func New(store *stores.Message) *Message {
	return &Message{store: store}
}

func (m *Message) 	Create(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println("error in parsing form")

		return
	}

	msg := &models.Message{Message: req.Form.Get("message"), Name: req.Form.Get("name")}

	err = m.store.Create(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Println("message added successfully")
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}
