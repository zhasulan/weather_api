package repository

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"weather_api/internal/model"
)

type IRepository struct {
	DB *sql.DB
}

func (r IRepository) ReadAllCities(ctx context.Context, page int, pageSize int) (cities []model.City, err error) {
	const query = `SELECT id, name, region, country, lat, lon FROM city LIMIT $1 OFFSET $2`

	offset := (page - 1) * pageSize

	rows, err := r.DB.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var city model.City

		err = rows.Scan(&city.ID, &city.Name, &city.Region, &city.Country, &city.Lat, &city.Lon)
		if err != nil {
			return
		}

		cities = append(cities, city)
	}

	return
}

func (r IRepository) ReadOneCity(ctx context.Context, id int) (city model.City, err error) {
	const query = `SELECT id, name, region, country, lat, lon FROM city WHERE id = $1`

	err = r.DB.QueryRowContext(ctx, query, id).Scan(&city.ID, &city.Name, &city.Region, &city.Country, &city.Lat, &city.Lon)

	return
}

func (r IRepository) CreateCity(ctx context.Context, city model.City) (id int, err error) {
	const query = `INSERT INTO city (name, region, country, lat, lon) VALUES ($1, $2, $3, $4, $5) returning id`

	err = r.DB.QueryRowContext(ctx, query, city.Name, city.Region, city.Country, city.Lat, city.Lon).Scan(&id)

	return
}

func (r IRepository) DeleteCity(ctx context.Context, id int) (err error) {
	const query = `DELETE FROM city WHERE id = $1`

	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return
}
