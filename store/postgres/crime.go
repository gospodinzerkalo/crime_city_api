package postgres

import (
	"errors"
	"github.com/gospodinzerkalo/crime_city_api/domain"
)

func (s *store) CreateCrime(crime *domain.Crime) (*domain.Crime, error) {
	err := s.db.QueryRow("INSERT INTO crimes (location_name, longitude, latitude, description, image) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at",
		crime.LocationName, crime.Longitude, crime.Latitude, crime.Description, crime.Image,
		).Scan(&crime.ID, &crime.Date)

	return crime, err
}

func (s *store) GetCrimes() (*[]domain.Crime, error) {
	rows, err := s.db.Query("SELECT id, location_name, longitude, latitude, created_at, description, image FROM crimes")
	if err != nil {
		return nil, err
	}

	var res []domain.Crime

	for rows.Next() {
		var crime domain.Crime
		if err := rows.Scan(&crime.ID, &crime.LocationName, &crime.Longitude, &crime.Latitude, &crime.Date, &crime.Description, &crime.Image); err != nil {
			return nil, err
		}

		res = append(res, crime)
	}

	return &res, nil
}

func (s *store) GetCrime(id int64) (*domain.Crime, error) {
	var crime domain.Crime
	if err := s.db.QueryRow("SELECT id, location_name, longitude, latitude, created_at, description, image FROM crimes WHERE id = $1").Scan(
		&crime.ID, &crime.LocationName, &crime.Longitude, &crime.Latitude, &crime.Date, &crime.Description, &crime.Image,
		); err != nil {
			return nil, err
	}

	return &crime, nil
}

func (s *store) UpdateCrime(id int64, crime *domain.Crime) (*domain.Crime, error) {
	return nil, nil
}

func (s *store) DeleteCrime(id int64) error {
	res, err := s.db.Exec("DELETE FROM crimes WHERE id=$1", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New("Not able to delete")
	}

	return nil
}


