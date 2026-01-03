package service

import (
	"context"
	"errors"

	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/database"
	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/model"
	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/repository"
)

type ForecastService struct {
	Pg     *repository.Postgres
	Logger database.Logger
}

func (s *ForecastService) GetNextWeeksForecast(
	ctx context.Context,
	warungID, productID, weeks int,
) ([]model.Forecast, error) {

	if weeks <= 0 {
		err := errors.New("weeks must be > 0")
		s.Logger.Error(ctx, "invalid weeks", err,
			"weeks", weeks,
		)
		return nil, err
	}

	return s.Pg.GetForecast(ctx, warungID, productID)
}

// func (s *ForecastService) GetNextWeeksForecast(ctx context.Context, warungID, productID int, weeks int) ([]model.Forecast, error) {
// 	all, err := s.Pg.GetForecast(ctx, warungID, productID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result []model.Forecast
// 	// now := time.Now()
// 	for _, f := range all {
// 		// if f.Week.After(now) && len(result) < weeks {
// 		result = append(result, f)
// 		// }
// 	}
// 	return result, nil
// }
