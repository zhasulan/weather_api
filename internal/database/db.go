package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Config struct {
	Host        string      `json:"host"`
	Port        string      `json:"port"`
	User        string      `json:"user"`
	Pass        string      `json:"pass"`
	Name        string      `json:"name"`
	Scheme      string      `json:"scheme"`
	Connections Connections `json:"connections"`
}

type Connections struct {
	MaxOpen  int `json:"max_open"`
	MaxIdle  int `json:"max_idle"`
	IdleLife int `json:"idle_life"`
	MaxRecon int `json:"max_recon"`
}

func Connection(config Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s TimeZone=Asia/Almaty",
		config.Host, config.User, config.Pass, config.Name, config.Port, config.Scheme,
	)

	db, err := sql.Open("postgres", connectionString)

	db.SetMaxOpenConns(config.Connections.MaxOpen)
	db.SetMaxIdleConns(config.Connections.MaxIdle)
	db.SetConnMaxIdleTime(time.Duration(config.Connections.IdleLife) * time.Minute)

	if err != nil {
		return nil, err
	}

	return db, nil
}
