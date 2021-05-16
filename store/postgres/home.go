package postgres

import (
	"errors"
	"github.com/gospodinzerkalo/crime_city_api/domain"
)

func (s *store) CreateHome(home *domain.Home) (*domain.Home, error) {
	_, err := s.db.Exec("INSERT INTO homes (id, first_name, last_name, user_name, longitude, latitude, image) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		home.ID, home.FirstName, home.LastName, home.UserName, home.Longitude, home.Latitude, home.Image,
		)
	return home, err
}

func (s *store) GetHome(id int64) (*domain.Home, error) {
	var home domain.Home
	err := s.db.QueryRow("SELECT id, first_name, last_name, user_name, longitude, latitude, image FROM homes WHERE id=$1", id).Scan(
		&home.ID, &home.FirstName, &home.LastName, &home.UserName, &home.Longitude, &home.Latitude, &home.Image,
		)

	return &home, err
}

func (s *store) DeleteHome(id int64) error {
	res, err := s.db.Exec("DELETE FROM homes WHERE id=$1", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("Not able to delete home")
	}

	return nil
}

