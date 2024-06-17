package config

import (
	"database/sql"
	"weather_api/internal/repository"
)

type Env struct {
	DB         *sql.DB
	Repository repository.Repository
}

var env Env // private

func SetEnv(value Env) {
	env = value
}

func GetEnv() Env {
	return env
}
