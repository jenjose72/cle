# GPS Service

This folder contains a microservice with two parts:

- `gps-tracker` — a small Flask backend that returns a classroom location and accepts client-submitted locations.
- `gps-frontend` — a React app (created with create-react-app) that shows a Leaflet map, requests the browser geolocation, and POSTs it to the backend.

This README shows the commands to install and run both parts locally, how to configure the frontend to talk to the backend, and troubleshooting tips.

---

## Prerequisites

- Node.js (>=14) and npm for the frontend
- Python 3.8+ for the backend
- Recommended: create a virtual environment for Python

## Backend (gps-tracker)

Path: `gps-tracker/`

1. Create and activate a Python virtual environment (optional but recommended):

```bash
cd gps-tracker
python3 -m venv .venv
source .venv/bin/activate
```

2. Install dependencies:

```bash
pip install flask flask-cors
```

3. Run the backend:

```bash
# command runs Flask app on 0.0.0.0:5000
python app.py
```

The backend exposes:
- GET /location — returns the last reported location (if any) or the classroom default
- POST /location — accepts JSON { "latitude": <number>, "longitude": <number>, "timestamp": <string optional> }

Example POST payload:

```json
{ "latitude": 40.12345, "longitude": -73.98765, "timestamp": "2026-04-01T12:00:00Z" }
```

The service stores received locations in memory for the running process.

## Frontend (gps-frontend)

Path: `gps-frontend/`

1. Install dependencies and start the dev server:

```bash
cd gps-frontend
npm install
# If you want the frontend to talk to a backend running elsewhere, set REACT_APP_BACKEND_URL
# Example: REACT_APP_BACKEND_URL=http://localhost:5000 npm start
npm start
```

2. By default, the frontend will attempt to POST the browser geolocation to `http://localhost:5000/location`.

If your backend runs on a different host/port (for example in Kubernetes or a VM), set `REACT_APP_BACKEND_URL` before starting the frontend:

```bash
REACT_APP_BACKEND_URL=http://192.168.49.2:30008 npm start
```

The UI will show a small status box with the detected user coordinates and whether sending succeeded.

## Notes & Troubleshooting

- The backend has CORS enabled, so the frontend on a different origin can POST to it.
- The backend stores locations in memory only. For production you should persist locations to a database.
- If geolocation permission is denied in the browser the frontend will display `denied` in the status box.
- To view logs from the backend, run it in a terminal and watch the INFO lines when the frontend posts a location.

## Next steps / Improvements

- Persist received locations to a database (Postgres, DynamoDB, etc.).
- Add basic authentication if the data is sensitive.
- Add an audit log and UI page listing recent locations.

---

If you want, I can also add an example `docker-compose.yml` or Kubernetes manifests to run both parts together locally. Below are concrete examples and a full walkthrough to build images, deploy to Kubernetes, and set up CI/CD with GitHub Actions.

## Containerization (Docker)

Below are example Dockerfiles and quick steps to build and run both services as containers. Example files are provided in this repository (see paths). These are minimal and intended as a starting point.

### Build & run locally with Docker

1. Build the backend image:

```bash
cd gps-tracker
# create a minimal requirements file if you don't have one
echo "flask\nflask-cors" > requirements.txt
docker build -t gps-tracker:local .
```

2. Build the frontend image:

```bash
cd ../gps-frontend
docker build -t gps-frontend:local .
```

3. Run them together (example using Docker network):

```bash
docker network create gps-net || true
docker run -d --name backend --network gps-net -p 5000:5000 gps-tracker:local
docker run -d --name frontend --network gps-net -p 3000:80 -e REACT_APP_BACKEND_URL=http://backend:5000 gps-frontend:local
```

Open http://localhost:3000 in your browser. The frontend is configured with `REACT_APP_BACKEND_URL` so it can reach the backend inside the same Docker network.

---

## Kubernetes (example manifests)

The `k8s/` folder contains example manifests for deploying both services to a Kubernetes cluster. These are simple Deployment + Service manifests and an example Ingress. They use placeholder image names — replace them with your built image locations (for example `ghcr.io/<owner>/<repo>/gps-tracker:latest`).

High-level steps to deploy to a cluster:

1. Build and push images to a container registry (DockerHub, GitHub Container Registry, GCR, ECR, etc.).
	- Example image names used in manifests: `ghcr.io/<owner>/<repo>/gps-tracker:latest` and `ghcr.io/<owner>/<repo>/gps-frontend:latest`.

2. Update `k8s/*.yaml` to use your image names (or pass them with `kubectl set image`).

3. Apply the manifests:

```bash
kubectl apply -f k8s/backend-deployment.yaml
kubectl apply -f k8s/backend-service.yaml
kubectl apply -f k8s/frontend-deployment.yaml
kubectl apply -f k8s/frontend-service.yaml
# optional ingress
kubectl apply -f k8s/ingress.yaml
```

4. If you created an Ingress, configure DNS to point to the Ingress controller's external IP and set `REACT_APP_BACKEND_URL` inside the frontend deployment (or use an internal URL and an Ingress rule to proxy).

Notes:
- The backend listens on port 5000. The frontend serves static files on port 80 in the example.
- Services are created as ClusterIP by default; change to `LoadBalancer` or add an Ingress for external access.

---

## GitHub Actions CI/CD (build & deploy)

An example GitHub Actions workflow is included at `.github/workflows/ci-deploy.yml`. It demonstrates:

- Building and pushing backend and frontend images to GitHub Container Registry (GHCR) using the `docker/build-push-action`.
- Optionally deploying the `k8s/` manifests to a Kubernetes cluster if you provide a `KUBE_CONFIG` secret.

You will need to configure the following repository secrets in GitHub (Repository -> Settings -> Secrets):

- `CR_PAT` (optional): a personal access token with `write:packages` & `read:packages` if you prefer to use a PAT instead of `GITHUB_TOKEN` for pushing images.
- `KUBE_CONFIG` (optional): base64-encoded kubeconfig used to authenticate to the target cluster. If provided, the workflow will apply the `k8s/` manifests after pushing images.

Quick notes for GHCR (GitHub Container Registry):
- You can use the automatically provided `GITHUB_TOKEN` for publishing to GHCR in many orgs. If you run into permission issues, create `CR_PAT` and set it in the workflow as described in the example.

---

## Files added as examples

The repository now contains example artifacts you can use or adapt:

- `gps-tracker/Dockerfile` — builds the Flask backend
- `gps-tracker/requirements.txt` — minimal Python deps
- `gps-frontend/Dockerfile` — multi-stage build for React app -> nginx
- `k8s/` — example Kubernetes manifests (backend/frontend deployments & services, ingress)
- `.github/workflows/ci-deploy.yml` — example workflow to build/push images and optionally deploy

If you'd like, I can also modify the frontend `Deployment` to inject a runtime environment variable for `REACT_APP_BACKEND_URL` using a ConfigMap, or add a Helm chart to parameterize deployments. Tell me which option you prefer and I'll add it.
