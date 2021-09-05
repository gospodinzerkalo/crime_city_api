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

func (s *store) CheckHome(id int64) (*domain.HomeCrime, *domain.Home, error) {
	myHome := domain.Home{}
	err := s.db.QueryRow("SELECT id, first_name, last_name, user_name, longitude, latitude, image FROM homes WHERE id = $1", id).
		Scan(&myHome.ID, &myHome.FirstName, &myHome.LastName, &myHome.UserName, &myHome.Longitude, &myHome.Latitude, &myHome.Image)
	if err != nil {
		return nil, nil, err
	}

	crime := domain.HomeCrime{}

	err = s.db.QueryRow("SELECT id, location_name,longitude,latitude,created_at,description,image,calculate_distance($1, $2, latitude, longitude, 'K') as distance FROM crimes ORDER BY distance",
		myHome.Latitude, myHome.Longitude).
		Scan(&crime.ID,
			&crime.LocationName,
			&crime.Longitude,
			&crime.Latitude,
			&crime.Date,
			&crime.Description,
			&crime.Image,
			&crime.Distance,
		)

	if err != nil {
		return nil, nil, err
	}

	return &crime, &myHome, nil
}

