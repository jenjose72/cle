from flask import Flask, request, jsonify
import numpy as np

app = Flask(__name__)

# Dummy model (y = 2x)
@app.route('/predict', methods=['POST'])
def predict():
    data = request.json['input']
    result = [2 * x for x in data]
    return jsonify({'prediction': result})

@app.route('/')
def home():
    return "ML Service Running"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)