import pandas as pd
from sklearn.ensemble import IsolationForest
from sklearn.preprocessing import StandardScaler
import joblib
import os

FEATURE_COLS = [
    "CHL_CD_FLOW_1",
    "CHL_COMP_SPD_CTRL_1",
    "CHL_CW_FLOW_1",
    "CHL_POW_1",
    "CHL_RWCD_TEMP_1",
    "CHL_RW_TEMP_1",
    "CHL_STA_1",
    "CHL_SWCD_TEMP_1",
    "CHL_SW_TEMP_1"
]

def main():
    print("Loading dataset...")
    df = pd.read_excel("Dataset1.xlsx")
    
    print(f"Columns found: {df.columns.tolist()}")
    print(f"Total rows: {len(df)}")

    # Check all required columns exist
    missing_cols = [col for col in FEATURE_COLS if col not in df.columns]
    if missing_cols:
        raise ValueError(f"Missing columns in dataset: {missing_cols}")

    # Drop rows with missing values
    df = df.dropna(subset=FEATURE_COLS)
    print(f"Rows after dropping nulls: {len(df)}")

    # Scale
    print("Scaling features...")
    scaler = StandardScaler()
    X = scaler.fit_transform(df[FEATURE_COLS])

    # Train
    print("Training Isolation Forest...")
    model = IsolationForest(
        contamination=0.05,
        random_state=42,
        n_estimators=100
    )
    model.fit(X)

    # Sanity check
    predictions = model.predict(X)
    normal = (predictions == 1).sum()
    anomaly = (predictions == -1).sum()
    print(f"Normal: {normal} ({normal/len(df)*100:.1f}%)")
    print(f"Anomaly: {anomaly} ({anomaly/len(df)*100:.1f}%)")

    # Save
    os.makedirs("/output", exist_ok=True)
    joblib.dump(model, "/output/isolation_forest.pkl")
    joblib.dump(scaler, "/output/scaler.pkl")
    print("Saved isolation_forest.pkl and scaler.pkl to /output")

if __name__ == "__main__":
    main()