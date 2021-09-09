package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gospodinzerkalo/crime_city_api/domain"
	"time"
)

// Logging Middleware in application level

type LoggingMiddleWare struct {
	Logger 	log.Logger
	Next 	Service
}

func (m LoggingMiddleWare) GetCrimes(ctx context.Context) (res *[]domain.Crime, err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "GetCrimes",
			"input", nil,
			"output", fmt.Sprintf("%v", res),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	res, err = m.Next.GetCrimes(ctx)
	return
}

func (m LoggingMiddleWare) GetCrime(ctx context.Context, id int64) (res *domain.Crime, err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "GetCrime",
			"input", id,
			"output", fmt.Sprintf("%v", res),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	res, err = m.Next.GetCrime(ctx, id)
	return
}

func (m LoggingMiddleWare) CreateCrime(ctx context.Context, crime *domain.Crime) (res *domain.Crime, err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "CreateCrime",
			"input", fmt.Sprintf("%v", crime),
			"output", fmt.Sprintf("%v", res),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	res, err = m.Next.CreateCrime(ctx, crime)
	return
}

func (m LoggingMiddleWare) UpdateCrime(ctx context.Context, crime *domain.Crime) (res *domain.Crime, err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "UpdateCrime",
			"input", fmt.Sprintf("%v", crime),
			"output", fmt.Sprintf("%v", res),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	res, err = m.Next.UpdateCrime(ctx, crime)
	return
}

func (m LoggingMiddleWare) DeleteCrime(ctx context.Context, id int64) (err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "DeleteCrime",
			"input", id,
			"output", nil,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = m.Next.DeleteCrime(ctx, id)
	return
}

func (m LoggingMiddleWare) CreateHome(ctx context.Context, req *domain.Home) (home *domain.Home, err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "CreateHome",
			"input", fmt.Sprintf("%v", req),
			"output", fmt.Sprintf("%v", home),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	home, err = m.Next.CreateHome(ctx, req)
	return
}

func (m LoggingMiddleWare) DeleteHome(ctx context.Context, id int64) (err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "DeleteHome",
			"input", id,
			"output", nil,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = m.Next.DeleteHome(ctx, id)
	return
}

func (m LoggingMiddleWare) CheckHome(ctx context.Context, id int64) (home *domain.HomeCrime, err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "CheckHome",
			"input", id,
			"output", fmt.Sprintf("%v", home),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	home, err = m.Next.CheckHome(ctx, id)
	return
}

func (m LoggingMiddleWare) GetHome(ctx context.Context, id int64) (home *domain.Home, err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "GetHome",
			"input", id,
			"output", fmt.Sprintf("%v", home),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	home, err = m.Next.GetHome(ctx, id)
	return
}