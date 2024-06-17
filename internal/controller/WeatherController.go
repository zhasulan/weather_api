package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"weather_api/internal/config"
	"weather_api/internal/model"
	"weather_api/internal/server"
	"weather_api/internal/services"
	"weather_api/internal/utils"
)

func GetAllWeatherEndpoint(writer http.ResponseWriter, request *http.Request) {
	var err error
	var result interface{}

	status := http.StatusOK
	repo := config.GetEnv().Repository
	ctx := request.Context()

	defer func() {
		if e := recover(); e != nil {
			err = utils.AnyError(e)
		}

		if err != nil {
			result = model.Response{
				Status: "Failed",
				Error:  err.Error(),
			}
		}

		server.SendJSON(writer, result, status)
	}()

	vars := mux.Vars(request)
	pageString, found := vars["page"]
	if !found {
		status = http.StatusBadRequest

		return
	}
	pageSizeString, found := vars["pageSize"]
	if !found {
		status = http.StatusBadRequest

		return
	}
	page, err := strconv.Atoi(pageString)
	if err != nil {
		status = http.StatusBadRequest

		return
	}
	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {
		status = http.StatusBadRequest

		return
	}

	cities, err := repo.ReadAllCities(ctx, page, pageSize)
	if err != nil {
		status = http.StatusInternalServerError

		return
	}

	weathers := make([]model.WeatherResponse, 0)
	for _, city := range cities {
		weather, err := services.WEATHER_API.GetCurrent(ctx, city.Name)
		if err != nil {
			status = http.StatusInternalServerError

			return
		}

		weathers = append(weathers, model.WeatherResponse{
			CityName:       weather.Location.Name,
			CelsiusTemp:    weather.Current.TempC,
			FahrenheitTemp: weather.Current.TempF,
			Humidity:       int(weather.Current.Humidity),
			Condition:      weather.Current.Condition.Text,
		})
	}

	result = weathers
}

func GetCityWeatherEndpoint(writer http.ResponseWriter, request *http.Request) {
	var err error
	var result interface{}

	status := http.StatusOK
	repo := config.GetEnv().Repository
	ctx := request.Context()

	defer func() {
		if e := recover(); e != nil {
			err = utils.AnyError(e)
		}

		if err != nil {
			result = model.Response{
				Status: "Failed",
				Error:  err.Error(),
			}
		}

		server.SendJSON(writer, result, status)
	}()

	vars := mux.Vars(request)
	idString, found := vars["city_id"]
	if !found {
		status = http.StatusBadRequest

		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		status = http.StatusBadRequest

		return
	}

	city, err := repo.ReadOneCity(ctx, id)
	if err != nil {
		status = http.StatusInternalServerError

		return
	}

	weather, err := services.WEATHER_API.GetCurrent(ctx, city.Name)
	if err != nil {
		status = http.StatusInternalServerError

		return
	}

	result = model.WeatherResponse{
		CityName:       weather.Location.Name,
		CelsiusTemp:    weather.Current.TempC,
		FahrenheitTemp: weather.Current.TempF,
		Humidity:       int(weather.Current.Humidity),
		Condition:      weather.Current.Condition.Text,
	}
}
