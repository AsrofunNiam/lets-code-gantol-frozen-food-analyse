package service

import (
	"context"

	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/model"
	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/repository"
)

type ForecastService struct {
	Pg *repository.Postgres
}

func (s *ForecastService) GetNextWeeksForecast(ctx context.Context, warungID, productID int, weeks int) ([]model.Forecast, error) {
	all, err := s.Pg.GetForecast(ctx, warungID, productID)
	if err != nil {
		return nil, err
	}

	var result []model.Forecast
	// now := time.Now()
	for _, f := range all {
		// if f.Week.After(now) && len(result) < weeks {
		result = append(result, f)
		// }
	}
	return result, nil
}
