package services

import (
	"context"
	"github.com/pkg/errors"
	"weather_api/internal/model"
)

type DummyWeatherApi struct{}

func (d DummyWeatherApi) GetCurrent(ctx context.Context, cityName string) (weather model.CurrentResponse, err error) {
	return model.CurrentResponse{}, errors.New("api not configured properly")
}

func (d DummyWeatherApi) Search(ctx context.Context, query string) (cities []model.City, err error) {
	return nil, errors.New("api not configured properly")
}
