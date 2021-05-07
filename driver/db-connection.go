package driver

import (
	"database/sql"
	"fmt"
	"log"
)

type MySQLConfig struct {
	Host     string
	Port     string
	Username string
	Database string
	Password string
}

func InitConnection(cfg *MySQLConfig) (*sql.DB, error) {
	log.Println(cfg)
	connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Println("error in making database connection")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("not able to ping, ",err)
	}

	return db, err
}
