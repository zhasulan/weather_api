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

func ReadAllCitiesEndpoint(writer http.ResponseWriter, request *http.Request) {
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

	result = cities
}

func ReadOneCityEndpoint(writer http.ResponseWriter, request *http.Request) {
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
	idString, found := vars["id"]
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

	result = city
}

func CreateCityEndpoint(writer http.ResponseWriter, request *http.Request) {
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
	cityName, found := vars["name"]
	if !found {
		status = http.StatusBadRequest

		return
	}

	cities, err := services.WEATHER_API.Search(ctx, cityName)
	if err != nil {
		status = http.StatusInternalServerError

		return
	}

	// last city id
	var citiesID []int

	// TODO change to bulk insert

	for _, city := range cities {
		var cityID int
		cityID, err = repo.CreateCity(ctx, city)
		if err != nil {
			status = http.StatusInternalServerError

			return
		}

		citiesID = append(citiesID, cityID)
	}

	result = model.Response{
		Status: "Cities successfully created",
		IDs:    citiesID,
	}
}

func DeleteCityEndpoint(writer http.ResponseWriter, request *http.Request) {
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
	idString, found := vars["id"]
	if !found {
		status = http.StatusBadRequest

		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		status = http.StatusBadRequest

		return
	}

	err = repo.DeleteCity(ctx, id)
	if err != nil {
		status = http.StatusInternalServerError

		return
	}

	result = model.Response{
		Status: "City successfully deleted",
		ID:     id,
	}
}
