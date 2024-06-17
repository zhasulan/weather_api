package router

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"weather_api/internal/logger"
)

var (
	SYSTEM_ENDPOINT       = RouteBehavior{System: true}
	TRACE_LOGGED_ENDPOINT = RouteBehavior{System: false, TraceLog: true}
	ONLY_TRACED_ENDPOINT  = RouteBehavior{System: false, TraceLog: false}
)

type RouteBehavior struct {
	System   bool // true - endpoint not logged and traced
	TraceLog bool // true - endpoint traced if System is true
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc

	RouteBehavior
}

type Routes []Route

func CreateRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		zap.S().Infof("%s endpoint -> %s path -> name %s ...", route.Method, route.Pattern, route.Name)

		var handler http.Handler = route.HandlerFunc

		// log tracing non system endpoints
		if !route.System {
			handler = logger.Logger(handler, route.Name, route.TraceLog)
		}
		handler = gzipHandlerMiddleware(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
