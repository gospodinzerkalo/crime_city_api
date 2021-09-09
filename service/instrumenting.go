package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/gospodinzerkalo/crime_city_api/domain"
	"time"
)

// Instrumenting Middleware in application level

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}

func (i InstrumentingMiddleware) GetCrimes(ctx context.Context) (res *[]domain.Crime, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetCrimes", "error", fmt.Sprint(err != nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	res, err = i.Next.GetCrimes(ctx)
	return
}

func (i InstrumentingMiddleware) GetCrime(ctx context.Context, id int64) (res *domain.Crime, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetCrime", "error", fmt.Sprint(err != nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	res, err = i.Next.GetCrime(ctx, id)
	return
}

func (i InstrumentingMiddleware) CreateCrime(ctx context.Context, crime *domain.Crime) (res *domain.Crime, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateCrime", "error", fmt.Sprint(err != nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	res, err = i.Next.CreateCrime(ctx, crime)
	return
}

func (i InstrumentingMiddleware) UpdateCrime(ctx context.Context, crime *domain.Crime) (res *domain.Crime, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateCrime", "error", fmt.Sprint(err != nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	res, err = i.Next.UpdateCrime(ctx, crime)
	return
}

func (i InstrumentingMiddleware) DeleteCrime(ctx context.Context, id int64) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "DeleteCrime", "error", fmt.Sprint(err != nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = i.Next.DeleteCrime(ctx, id)
	return
}

func (i InstrumentingMiddleware) CreateHome(ctx context.Context, home *domain.Home) (res *domain.Home, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateHome", "error", fmt.Sprint(err != nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	res, err = i.Next.CreateHome(ctx, home)
	return
}

func (i InstrumentingMiddleware) GetHome(ctx context.Context, id int64) (res *domain.Home, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetHome", "error", fmt.Sprint(err != nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	res, err = i.Next.GetHome(ctx, id)
	return
}

func (i InstrumentingMiddleware) DeleteHome(ctx context.Context, id int64) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "DeleteHome", "error", fmt.Sprint(err != nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = i.Next.DeleteHome(ctx, id)
	return
}

func (i InstrumentingMiddleware) CheckHome(ctx context.Context, id int64) (res *domain.HomeCrime, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CheckHome", "error", fmt.Sprint(err != nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	res, err = i.Next.CheckHome(ctx, id)
	return
}


