package router

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"weather_api/internal/controller"
)

func NewRouter() *mux.Router {
	var routes = make(Routes, 0)

	// City CRUD endpoints
	{
		routes = append(routes, Route{
			Name:          "ReadAllCities",
			Method:        http.MethodGet,
			Pattern:       "/city/{page}/{pageSize}",
			HandlerFunc:   controller.ReadAllCitiesEndpoint,
			RouteBehavior: TRACE_LOGGED_ENDPOINT,
		})

		routes = append(routes, Route{
			Name:          "ReadOneCities",
			Method:        http.MethodGet,
			Pattern:       "/city/{id}",
			HandlerFunc:   controller.ReadOneCityEndpoint,
			RouteBehavior: TRACE_LOGGED_ENDPOINT,
		})

		routes = append(routes, Route{
			Name:          "CreateCity",
			Method:        http.MethodPost,
			Pattern:       "/city/{name}",
			HandlerFunc:   controller.CreateCityEndpoint,
			RouteBehavior: TRACE_LOGGED_ENDPOINT,
		})

		routes = append(routes, Route{
			Name:          "DeleteCity",
			Method:        http.MethodDelete,
			Pattern:       "/city/{id}",
			HandlerFunc:   controller.DeleteCityEndpoint,
			RouteBehavior: TRACE_LOGGED_ENDPOINT,
		})
	}

	// GET WEATHER
	{
		routes = append(routes, Route{
			Name:          "WeatherForAllCities",
			Method:        http.MethodGet,
			Pattern:       "/weather/all/{page}/{pageSize}",
			HandlerFunc:   controller.GetAllWeatherEndpoint,
			RouteBehavior: TRACE_LOGGED_ENDPOINT,
		})

		routes = append(routes, Route{
			Name:          "WeatherForOneCity",
			Method:        http.MethodGet,
			Pattern:       "/weather/city/{city_id}",
			HandlerFunc:   controller.GetCityWeatherEndpoint,
			RouteBehavior: TRACE_LOGGED_ENDPOINT,
		})
	}

	// System endpoints
	{
		routes = append(routes, Route{
			Name:          "healthz",
			Method:        http.MethodGet,
			Pattern:       "/healthz",
			HandlerFunc:   controller.HealthCheck,
			RouteBehavior: SYSTEM_ENDPOINT,
		})

		routes = append(routes, Route{
			Name:          "Meta Build",
			Method:        http.MethodGet,
			Pattern:       "/meta/build",
			HandlerFunc:   controller.MetaBuild,
			RouteBehavior: SYSTEM_ENDPOINT,
		})

		routes = append(routes, Route{
			Name:          "Metrics",
			Method:        http.MethodGet,
			Pattern:       "/metrics",
			HandlerFunc:   promhttp.Handler().ServeHTTP,
			RouteBehavior: SYSTEM_ENDPOINT,
		})
	}

	return CreateRouter(routes)
}
