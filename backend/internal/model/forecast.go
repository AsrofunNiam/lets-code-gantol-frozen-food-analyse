package model

import "time"

type Forecast struct {
	WarungID     int       `json:"warung_id"`
	ProductID    int       `json:"product_id"`
	Week         time.Time `json:"week"`
	PredictedQty int       `json:"predicted_qty"`
	ModelVersion string    `json:"model_version"`
}
