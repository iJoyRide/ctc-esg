import joblib
import pandas as pd

model = joblib.load("output/isolation_forest.pkl")
scaler = joblib.load("output/scaler.pkl")

FEATURE_COLS = [
    "CHL_CD_FLOW_1", "CHL_COMP_SPD_CTRL_1", "CHL_CW_FLOW_1",
    "CHL_POW_1", "CHL_RWCD_TEMP_1", "CHL_RW_TEMP_1",
    "CHL_STA_1", "CHL_SWCD_TEMP_1", "CHL_SW_TEMP_1"
]

THRESHOLD = -0.05

sample = pd.DataFrame([{
    "CHL_CD_FLOW_1": 1126.6584,
    "CHL_COMP_SPD_CTRL_1": 0.49726182,
    "CHL_CW_FLOW_1": 978.6087,
    "CHL_POW_1": 69.50739796,
    #"CHL_POW_1": 999,
    "CHL_RWCD_TEMP_1": 69.17861,
    "CHL_RW_TEMP_1": 57.255424,
    "CHL_STA_1": 1,
    "CHL_SWCD_TEMP_1": 62.26277,
    "CHL_SW_TEMP_1": 51.432564
}])

X = scaler.transform(sample)
score = model.decision_function(X)[0]

print(f"Anomaly score: {score:.4f}")

if score > THRESHOLD:
    print("Prediction: NORMAL")
else:
    print("Prediction: FAULT")
    
    # Show contributing sensors
    print("\nSensor contributions (how far from normal):")
    contributions = sorted(
        zip(FEATURE_COLS, abs(X[0])),
        key=lambda x: x[1],
        reverse=True
    )
    for sensor, contrib in contributions:
        print(f"  {sensor}: {contrib:.4f}")
    print("(Higher value = further from normal = more likely cause)")