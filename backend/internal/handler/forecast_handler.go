package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/database"
	// "github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/httpx"
	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/service"
)

type ForecastHandler struct {
	Service *service.ForecastService
	Logger  database.Logger
}

// func (h *ForecastHandler) GetForecast(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	warungID, _ := strconv.Atoi(r.URL.Query().Get("warung_id"))
// 	productID, _ := strconv.Atoi(r.URL.Query().Get("product_id"))
// 	weeks := 4

// 	forecasts, err := h.Service.GetNextWeeksForecast(ctx, warungID, productID, weeks)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(forecasts)
// }

func (h *ForecastHandler) GetForecast(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	warungID, _ := strconv.Atoi(r.URL.Query().Get("warung_id"))
	productID, _ := strconv.Atoi(r.URL.Query().Get("product_id"))
	weeks := 4
	result, err := h.Service.GetNextWeeksForecast(ctx, warungID, productID, weeks)
	if err != nil {
		h.Logger.Error(ctx, "get_forecast_failed", err)
		// httpx.WriteError(w, err)
		return
	}

	json.NewEncoder(w).Encode(result)
}
