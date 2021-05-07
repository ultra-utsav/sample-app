package stores

import (
	"database/sql"
	"github.com/ultra-utsav/sample-app/models"
)

type Message struct {
	db *sql.DB
}

func New(db *sql.DB) *Message {
	return &Message{db: db}
}

func (m *Message) Create(msg *models.Message) error {
	_, err := m.db.Exec("insert into messages values(?,?)", msg.Name, msg.Message)

	return err
}
