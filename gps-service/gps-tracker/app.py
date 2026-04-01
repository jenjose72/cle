from flask import Flask, jsonify, request
from flask_cors import CORS
from datetime import datetime

app = Flask(__name__)
CORS(app)

CLASSROOM_LOCATION = {
    "latitude": 9.755029630830405,
    "longitude": 76.64991585959865,
    "room": "2023 Batch 2 - AA103"
}

# In-memory storage for last-reported client location(s).
LAST_LOCATION = None
RECEIVED_LOCATIONS = []

@app.route("/")
def home():
    return "GPS Microservice Running"

@app.route("/location", methods=["GET"])
def location():
    """Return the last received location (if any) otherwise the classroom location."""
    if LAST_LOCATION:
        return jsonify(LAST_LOCATION)
    return jsonify(CLASSROOM_LOCATION)


@app.route("/location", methods=["POST"])
def receive_location():
    """Receive a location from the frontend. Expected JSON: { latitude, longitude, timestamp? }

    Stores the location in memory (RECEIVED_LOCATIONS) and updates LAST_LOCATION.
    """
    global LAST_LOCATION
    data = request.get_json(silent=True)
    if not data:
        return jsonify({"error": "invalid or missing JSON payload"}), 400

    latitude = data.get("latitude")
    longitude = data.get("longitude")
    timestamp = data.get("timestamp") or datetime.utcnow().isoformat() + "Z"

    try:
        latitude = float(latitude)
        longitude = float(longitude)
    except (TypeError, ValueError):
        return jsonify({"error": "latitude and longitude must be numbers"}), 400

    record = {
        "latitude": latitude,
        "longitude": longitude,
        "timestamp": timestamp,
        "source": "frontend"
    }

    RECEIVED_LOCATIONS.append(record)
    LAST_LOCATION = record

    app.logger.info("Received location: %s", record)
    return jsonify({"status": "ok", "received": record}), 201

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)