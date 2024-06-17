package services

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"
	"weather_api/internal/config"
	"weather_api/internal/model"
)

type IWeatherApi interface {
	Search(ctx context.Context, query string) (cities []model.City, err error)
	GetCurrent(ctx context.Context, cityName string) (weather model.CurrentResponse, err error)
}

var WEATHER_API IWeatherApi = &DummyWeatherApi{}

func NewWeatherService() IWeatherApi {
	conf := config.Config.WeatherApi

	client := NewClient(conf.HostURL, nil, conf.Timeout)

	return &WeatherApi{Client: client, Key: conf.Key}
}

type WeatherApi struct {
	Client *Client
	Key    string
}

func (w WeatherApi) GetCurrent(ctx context.Context, cityName string) (weather model.CurrentResponse, err error) {
	var urlQueryParams url.Values = map[string][]string{
		"q":   {cityName},
		"key": {w.Key},
	}

	responseBody, err := w.Client.ExecuteRequest(ctx, Current, urlQueryParams.Encode(), nil, &weather, true)
	if err != nil {
		err = parseError(responseBody)
	}

	return
}

func parseError(responseBody []byte) error {
	var weatherStatus model.WeatherStatus
	err := json.Unmarshal(responseBody, &weatherStatus)
	if err != nil {
		return err
	}

	return errors.New(weatherStatus.Message)
}

func (w WeatherApi) Search(ctx context.Context, query string) (cities []model.City, err error) {
	var urlQueryParams url.Values = map[string][]string{
		"q":   {query},
		"key": {w.Key},
	}

	responseBody, err := w.Client.ExecuteRequest(ctx, Search, urlQueryParams.Encode(), nil, &cities, true)
	if err != nil {
		err = parseError(responseBody)
	}

	if len(cities) == 0 {
		err = errors.New("no cities found")
	}

	return
}
