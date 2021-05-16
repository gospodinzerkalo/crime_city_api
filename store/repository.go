package store

import "github.com/gospodinzerkalo/crime_city_api/domain"

type Repository interface {
	CreateCrime(crime *domain.Crime) (*domain.Crime, error)
	GetCrimes() (*[]domain.Crime, error)
	UpdateCrime(id int64, crime *domain.Crime) (*domain.Crime, error)
	DeleteCrime(id int64) error
	GetCrime(id int64) (*domain.Crime, error)

	CreateHome(home *domain.Home) (*domain.Home, error)
	GetHome(id int64) (*domain.Home, error)
	DeleteHome(id int64) error
}