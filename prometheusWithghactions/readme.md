#  Kubernetes + Prometheus Monitoring Setup (Minikube)

## 📌 Overview

This project demonstrates how to:

* Build a Docker image for a Python app
* Deploy it to Kubernetes using Minikube
* Expose it via a Service
* Monitor it using Prometheus

---

# 🧱 1. Build Docker Image

### 👉 Why?

Kubernetes runs containers, not raw code. So we first package our app into a Docker image.

### 🔧 Command:

```bash
eval $(minikube docker-env)
```

👉 Switch Docker to Minikube’s internal environment
👉 Ensures Kubernetes can access the image

```bash
docker build -t myapp:v1 .
```

👉 Builds Docker image from Dockerfile

---

# ☸️ 2. Deploy Application

### 👉 Why?

A Deployment manages pods and ensures your app stays running.

### 🔧 Command:

```bash
kubectl apply -f deployment.yaml
```

### 📄 Key Concepts:

* `replicas: 3` → runs 3 instances (scaling)
* `imagePullPolicy: Never` → use local image (important for Minikube)
* `containerPort: 5000` → app runs on port 5000

---

# 🌐 3. Expose via Service

### 👉 Why?

Pods are temporary. Services give a stable way to access them.

### 🔧 Command:

```bash
kubectl apply -f service.yaml
```

### 📄 Key Concepts:

* `type: NodePort` → exposes app outside cluster
* `port: 80` → external port
* `targetPort: 5000` → container port

---

# 🧪 4. Test Application

### 🔧 Command:

```bash
minikube service myapp-service --url
```

```bash
curl <URL>
curl <URL>/metrics
```

### 👉 Why?

* Verify app is running
* Ensure `/metrics` endpoint works (critical for Prometheus)

---

# 📊 5. Setup Prometheus

### 👉 Why?

Prometheus collects and stores metrics from your app.

---

## Create Namespace

```bash
kubectl create namespace monitoring
```

👉 Keeps monitoring tools separate from app

---

## Apply ConfigMap

```bash
kubectl apply -f prometheus-config.yaml
```

### 📄 Key Concept:

```yaml
targets: ['myapp-service.default.svc.cluster.local:80']
```

👉 Full DNS format:

```
<service>.<namespace>.svc.cluster.local
```

👉 Required because Prometheus runs in a different namespace

---

# 🚀 6. Deploy Prometheus

```bash
kubectl apply -f prometheus-deployment.yaml
kubectl apply -f prometheus-service.yaml
```

---

# 🌍 7. Access Prometheus UI

```bash
minikube service prometheus-service -n monitoring
```

---

# 🔍 8. Verify Targets

Go to:

```
Status → Targets
```

Expected:

```
myapp → UP ✅
```

---

# 📈 9. Query Metrics

### ❌ Wrong:

```
request_count
```

### ✅ Correct:

```
request_count_total
```

---

# 🧠 Why `_total`?

Prometheus automatically converts:

```
Counter → *_total
```

---

# ⚡ Useful Queries

### Total requests:

```
request_count_total
```

### Requests per second:

```
rate(request_count_total[1m])
```

---

# 🧪 10. Generate Traffic

```bash
curl <app-url>
curl <app-url>
curl <app-url>
```

👉 Metrics increase only when app is hit

---

# 🧠 Key Learnings

## 🔹 Docker

* Packages app into image
* Must be built inside Minikube

## 🔹 Kubernetes

* Deployment → runs pods
* Service → exposes pods

## 🔹 Networking

* Services use DNS:

```
myapp-service.default.svc.cluster.local
```

## 🔹 Prometheus

* Pull-based system (scrapes data)
* Needs correct target + reachable endpoint

---

# 🚨 Common Errors & Fixes

## ❌ ImagePullBackOff

✔ Fix:

```bash
eval $(minikube docker-env)
docker build -t myapp:v1 .
```

---

## ❌ Target DOWN (DNS error)

✔ Fix:
Use correct service name + namespace:

```
myapp-service.default.svc.cluster.local
```

---

## ❌ No metrics shown

✔ Check:

```bash
curl <URL>/metrics
```

---

## ❌ Wrong metric name

✔ Use:

```
request_count_total
```

---

# 🧠 Final Mental Model

```
Docker → Image
Kubernetes → Runs it
Service → Makes it reachable
Prometheus → Scrapes metrics
```

---

# 🏁 Conclusion

You have successfully built a full monitoring pipeline:

✅ Containerized app
✅ Deployed on Kubernetes
✅ Exposed via Service
✅ Monitored using Prometheus

---
---
