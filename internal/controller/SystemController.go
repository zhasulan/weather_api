package controller

import (
	"net/http"
	"weather_api/internal/server"
	"weather_api/meta"
)

func HealthCheck(writer http.ResponseWriter, request *http.Request) {
	server.SendJSON(writer, "Server is healthy", http.StatusOK)
}

func MetaBuild(writer http.ResponseWriter, request *http.Request) {
	meta := meta.MetaInfo{
		GitBranch:    meta.GitBranch,
		GitHash:      meta.GitHash,
		GitAuthor:    meta.GitAuthor,
		BuildDate:    meta.BuildDate,
		BuildDocker:  meta.BuildDocker,
		BuildVersion: meta.BuildVersion,
	}

	server.SendJSON(writer, meta, http.StatusOK)
}
