package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/database"
	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/handler"
	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/repository"
	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open(
		"postgres",
		"postgres://postgres:password@localhost:5432/gantol_frozen?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	postgresRepo := repository.New(db)

	logger := database.StdLogger{
		Service: "forecast-api",
	}

	forecastService := &service.ForecastService{
		Pg:     postgresRepo,
		Logger: logger,
	}

	forecastHandler := &handler.ForecastHandler{
		Service: forecastService,
	}

	http.HandleFunc("/forecast", forecastHandler.GetForecast)
	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
