package repository

import (
	"context"
	"weather_api/internal/model"
)

type Repository interface {
	ReadAllCities(ctx context.Context, page int, pageSize int) (cities []model.City, err error)
	ReadOneCity(ctx context.Context, id int) (city model.City, err error)
	CreateCity(ctx context.Context, city model.City) (id int, err error)
	DeleteCity(ctx context.Context, id int) (err error)
}
