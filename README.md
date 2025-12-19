# lets-code-gantol-frozen-food-analyse

AI-powered analysis system for a real frozen food business.

## Background
This project is built from a real frozen food operation where procurement, marketing,
and delivery are handled directly by the owner. The goal is to apply data engineering
and classical machine learning to support better business decisions.

## Tech Stack
- Backend: Golang
- Database: PostgreSQL
- Data & Machine Learning: Python (Pandas, Prophet, scikit-learn)
- Approach: Classical ML (Time Series Forecasting & Regression)

## Features
- Weekly demand forecasting per warung
- Delivery time analysis and prediction
- Stock recommendation logic
- Model versioning & retraining workflow
- Real business data (not synthetic use cases)

## Architecture
- Golang handles API & business logic
- Python handles data processing & ML
- PostgreSQL acts as the single source of truth

## Roadmap
- [x] Database schema & dummy data
- [ ] Demand forecasting model
- [ ] Delivery time prediction
- [ ] Automation & retraining
- [ ] Dashboard & reporting
