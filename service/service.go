package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gospodinzerkalo/crime_city_api/domain"
	"github.com/gospodinzerkalo/crime_city_api/store"
	"net/http"
)

type Service interface {
	GetCrimes(ctx context.Context) (*[]domain.Crime, error)
	GetCrime(ctx context.Context, id int64) (*domain.Crime, error)
	CreateCrime(ctx context.Context, crime *domain.Crime) (*domain.Crime, error)
	UpdateCrime(ctx context.Context, crime *domain.Crime) (*domain.Crime, error)
	DeleteCrime(ctx context.Context, id int64) error

	CreateHome(ctx context.Context, home *domain.Home) (*domain.Home, error)
	GetHome(ctx context.Context, id int64) (*domain.Home, error)
	DeleteHome(ctx context.Context, id int64) error
	CheckHome(ctx context.Context, id int64) (*domain.HomeCrime, error)
}


func NewService(store store.Repository, logger log.Logger) Service {
	return service{store: store, log: logger}
}

type service struct {
	store  store.Repository
	log    log.Logger
}

func(s service) GetCrimes(ctx context.Context) (*[]domain.Crime, error) {
	return s.store.GetCrimes()
}

func(s service) GetCrime(ctx context.Context, id int64) (*domain.Crime, error) {
	return s.store.GetCrime(id)
}

func(s service) CreateCrime(ctx context.Context, crime *domain.Crime) (*domain.Crime, error) {
	return s.store.CreateCrime(crime)
}

func(s service) UpdateCrime(ctx context.Context, crime *domain.Crime) (*domain.Crime, error) {
	return s.store.UpdateCrime(crime.ID, crime)
}

func(s service) DeleteCrime(ctx context.Context, id int64) error {
	return s.store.DeleteCrime(id)
}

func (s service) CreateHome(ctx context.Context, home *domain.Home) (*domain.Home, error) {
	return s.store.CreateHome(home)
}

func (s service) GetHome(ctx context.Context, id int64) (*domain.Home, error) {
	return s.store.GetHome(id)
}

func (s service) DeleteHome(ctx context.Context, id int64) error {
	return s.store.DeleteHome(id)
}

func (s service) CheckHome(ctx context.Context, id int64) (*domain.HomeCrime, error) {
	crime, home, err := s.store.CheckHome(id)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(fmt.Sprintf("https://static-maps.yandex.ru/1.x/?ll=%f,%f&size=450,450&z=14&l=map&pt=%f,%f,home~%f,%f,flag",
		home.Longitude, home.Latitude, home.Longitude, home.Latitude, crime.Longitude, crime.Latitude))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body []byte
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	crime.MapImage = body

	return crime, nil
}