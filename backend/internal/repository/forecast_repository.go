package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/model"
)

func (r *Postgres) GetForecast(
	ctx context.Context,
	warungID int,
	productID int,
) ([]model.Forecast, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.
		Select(
			"warung_id",
			"product_id",
			"week",
			"predicted_qty",
			"model_version",
		).
		From("order_forecasts").
		Where(sq.Eq{
			"warung_id":  warungID,
			"product_id": productID,
		}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var results []model.Forecast

	mapResult := func(r *sql.Rows) error {
		var v model.Forecast
		if err := r.Scan(
			&v.WarungID,
			&v.ProductID,
			&v.Week,
			&v.PredictedQty,
			&v.ModelVersion,
		); err != nil {
			return err
		}
		results = append(results, v)
		return nil
	}

	if err := r.secondary.RunQuery(ctx, query, mapResult, args...); err != nil {
		return nil, err
	}

	return results, nil

}

func (p *Postgres) GetWeeklyOrderHistory(
	ctx context.Context,
	warungID int64,
) ([]model.Forecast, error) {

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.
		Select(
			"o.warung_id",
			"oi.product_id",
			"date_trunc('week', o.order_date)::date AS week",
			"SUM(oi.qty) AS qty",
		).
		From("orders o").
		Join("order_items oi ON o.id = oi.order_id").
		Where(sq.Eq{
			"o.warung_id": warungID,
		}).
		GroupBy("o.warung_id", "oi.product_id", "week").
		OrderBy("oi.product_id", "week").
		ToSql()

	if err != nil {
		return nil, err
	}

	var results []model.Forecast

	mapResult := func(r *sql.Rows) error {
		var v model.Forecast
		if err := r.Scan(
			&v.WarungID,
			&v.ProductID,
			&v.Week,
			&v.PredictedQty,
		); err != nil {
			return err
		}
		results = append(results, v)
		return nil
	}

	err = p.secondary.RunQuery(ctx, query, mapResult, args...)
	if err != nil {
		return nil, err
	}

	return results, nil
}
