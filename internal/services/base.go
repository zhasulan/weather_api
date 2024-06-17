package services

import "net/http"

const (
	// WEATHER API
	Search  = 1
	Current = 2
)

type httpReq struct {
	name     string
	method   string
	endpoint string
}

var methodMap = map[int]httpReq{
	// WEATHER API
	Search: {
		name:     "Search",
		method:   http.MethodGet,
		endpoint: "/v1/search.json",
	},
	Current: {
		name:     "GetCurrent",
		method:   http.MethodGet,
		endpoint: "/v1/current.json",
	},
}
