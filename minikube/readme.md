🚀 Cloud Lab: Minikube + Kubernetes + ML-as-a-Service

🧰 Step 1: Install Required Tools
1. Install Docker
sudo apt update
sudo apt install docker.io -y
sudo systemctl start docker
sudo systemctl enable docker

Add your user to Docker group:

sudo usermod -aG docker $USER

👉 Logout and login again

2. Install kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s \
https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

chmod +x kubectl
sudo mv kubectl /usr/local/bin/

Verify:

kubectl version --client
3. Install Minikube
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64

sudo install minikube-linux-amd64 /usr/local/bin/minikube

Verify:

minikube version
▶️ Task 1: Start Minikube

Start cluster using Docker driver:

minikube start --driver=docker

Check status:

minikube status
📊 Open Kubernetes Dashboard
minikube dashboard

👉 This opens a web UI showing:

Pods
Deployments
Services
Nodes
🚀 Task 2: Deploy an Application
📄 Step 1: Create Deployment File
nano deployment.yaml

Paste:

apiVersion: apps/v1
kind: Deployment

metadata:
  name: nginx-deployment

spec:
  replicas: 3

  selector:
    matchLabels:
      app: nginx

  template:
    metadata:
      labels:
        app: nginx

    spec:
      containers:
      - name: nginx-container
        image: nginx:latest
        ports:
        - containerPort: 80

Save and exit.

▶️ Step 2: Apply Deployment
kubectl apply -f deployment.yaml
📊 Step 3: Verify Deployment
kubectl get deployments
kubectl get pods
kubectl describe deployment nginx-deployment
🌐 Step 4: Expose Application
kubectl expose deployment nginx-deployment --type=NodePort --port=80

Get URL:

minikube service nginx-deployment
🧠 YAML Explanation (Important for Viva)
apiVersion → API version used (apps/v1)
kind → Resource type (Deployment)
metadata → Name of resource
spec → Configuration:
replicas → number of pods
selector → identifies pods
template → pod blueprint
containers → image, ports
🚀 Task 3: ML-as-a-Service
🧠 Step 1: Create ML App
nano app.py
from flask import Flask, request, jsonify

app = Flask(__name__)

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
📦 Step 2: Create Dockerfile
nano Dockerfile
FROM python:3.9

WORKDIR /app

COPY . .

RUN pip install flask numpy

CMD ["python", "app.py"]
🐳 Step 3: Build Docker Image in Minikube
eval $(minikube docker-env)
docker build -t ml-app .
📄 Step 4: Create ML Deployment
nano ml-deployment.yaml
apiVersion: apps/v1
kind: Deployment

metadata:
  name: ml-deployment

spec:
  replicas: 1

  selector:
    matchLabels:
      app: ml-app

  template:
    metadata:
      labels:
        app: ml-app

    spec:
      containers:
      - name: ml-container
        image: ml-app
        imagePullPolicy: Never
        ports:
        - containerPort: 5000
📄 Step 5: Create Service
nano ml-service.yaml
apiVersion: v1
kind: Service

metadata:
  name: ml-service

spec:
  type: NodePort

  selector:
    app: ml-app

  ports:
    - port: 5000
      targetPort: 5000
▶️ Step 6: Deploy ML Service
kubectl apply -f ml-deployment.yaml
kubectl apply -f ml-service.yaml
📊 Step 7: Verify
kubectl get pods
kubectl get svc
🌐 Step 8: Access Service
minikube service ml-service
🧪 Step 9: Test Prediction API
curl -X POST http://<URL>/predict \
-H "Content-Type: application/json" \
-d '{"input":[1,2,3]}'
✅ Expected Output
{
  "prediction": [2,4,6]
}
⚠️ Common Errors & Fixes
❌ File not found
error: the path "ml-deployment.yaml" does not exist

✅ Fix:

ls

Ensure file exists or recreate it.

❌ NodePort already allocated
Invalid value: provided port is already allocated

✅ Fix:

Remove nodePort from YAML (recommended)
OR use another port (30000–32767)
❌ ImagePullBackOff
kubectl get pods

✅ Fix:

eval $(minikube docker-env)
docker build -t ml-app .