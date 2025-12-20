import pandas as pd
from prophet import Prophet
from sqlalchemy import create_engine, text

# =========================
# CONFIG
# =========================
DB_URL = "postgresql://postgres:password@localhost:5432/gantol_frozen"
MODEL_VERSION = "order_forecast_v1"
FORECAST_WEEKS = 4

# =========================
# CONNECT DB
# =========================
engine = create_engine(DB_URL)

# =========================
# LOAD DATA 
# =========================
min_data = 6
warungs = pd.read_sql("SELECT id FROM warungs", engine)
products = pd.read_sql("SELECT id FROM products", engine)

for _, w_row in warungs.iterrows():
    warung_id = w_row['id']
    for _, p_row in products.iterrows():
        product_id = p_row['id']
        query = f"""
        SELECT
          date_trunc('week', o.order_date)::date AS ds,
          SUM(oi.qty) AS y
        FROM orders o
        JOIN order_items oi ON o.id = oi.order_id
        WHERE o.warung_id = {warung_id}
          AND oi.product_id = {product_id}
        GROUP BY ds
        ORDER BY ds;
        """
        df = pd.read_sql(query, engine)

        if len(df) < min_data:
            print(f"warung {warung_id}, product {product_id} has less than {min_data} weeks of data")
            continue

# =========================
# TRAIN MODEL
# =========================
        model = Prophet(weekly_seasonality=True, daily_seasonality=False)
        model.fit(df)

        # forecast
        future = model.make_future_dataframe(periods=FORECAST_WEEKS, freq="W")
        forecast = model.predict(future)
        result = forecast[["ds", "yhat"]].tail(FORECAST_WEEKS)
        result["yhat"] = result["yhat"].round().astype(int)

# =========================
# SAVE TO DB
# =========================
        with engine.begin() as conn:
            for _, row in result.iterrows():

                print(f"saving forecast for warung {warung_id} - product {product_id} {row['ds']} - {row['yhat']}")
                conn.execute(
                    text("""
                        INSERT INTO order_forecasts
                        (warung_id, product_id, week, predicted_qty, model_version)
                        VALUES (:warung_id, :product_id, :week, :qty, :model_version)
                        ON CONFLICT (warung_id, product_id, week)
                        DO UPDATE SET
                        predicted_qty = EXCLUDED.predicted_qty,
                        model_version = EXCLUDED.model_version,
                        created_at = now();
                    """),
                    {
                        "warung_id": int(warung_id),   
                        "product_id": int(product_id),
                        "week": row["ds"].to_pydatetime().date(),
                        "qty": int(row["yhat"]),     
                        "model_version": MODEL_VERSION,
                    }
                )

        print(f"forecast saved for warung {warung_id} - product {product_id}")
