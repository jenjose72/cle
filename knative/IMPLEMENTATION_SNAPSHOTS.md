# Knative Demonstration: High-End AIOps Incident Triage Service

## 1) Application Selected
High-end application: **AIOps Incident Triage API**

Why this is high-end:
- Performs structured incident intelligence from raw incident text.
- Produces urgency scoring, sentiment signal, entity extraction, and team-routing recommendation.
- Supports single and batch inference patterns.
- Fits perfectly with Knative serverless autoscaling for bursty incident traffic.

## 2) Implementation Overview
### Backend service
- Platform: Flask API in `app.py`
- Endpoints:
  - `GET /` : service metadata
  - `GET /health` : liveness check
  - `POST /analyze` : single incident analysis
  - `POST /batch-analyze` : batch incident analysis

### Containerization
- File: `Dockerfile`
- Runtime: Python 3.11 slim image
- Server: Gunicorn (`2 workers`, `4 threads`) on port `8080`

### Knative serving
- File: `service.yaml`
- Service name: `aiops-triage-service`
- Autoscaling:
  - `minScale: 0` (scale-to-zero)
  - `maxScale: 20`
  - Concurrency metric with target-based scaling

## 3) Step-by-Step Deployment With Snapshots
Use these exact steps and capture screenshots for each snapshot label.

### Snapshot 1: Build and push container
Command:
```bash
docker build -t 2023bcs0053jefinfrancis/knative:latest .
docker push 2023bcs0053jefinfrancis/knative
```
Capture:
- Successful image build lines
- Successful push with digest

### Snapshot 2: Update image in Knative manifest
Edit `service.yaml` image field:
```yaml
image: docker.io/2023bcs0053jefinfrancis/knative:latest
```
Capture:
- `service.yaml` open in editor showing updated image

### Snapshot 3: Deploy Knative Service
Command:
```bash
kubectl apply -f service.yaml
kubectl get ksvc aiops-triage-service
```
Expected:
- `READY` should become `True`
- URL assigned to service

Capture:
- Terminal output of both commands

### Snapshot 4: Functional API test
If your service URL looks like `http://*.svc.cluster.local`, it is cluster-internal and not directly resolvable from your laptop.

Start local access tunnel to Knative ingress (keep this terminal running):
```bash
kubectl port-forward -n kourier-system svc/kourier 8080:80
```

Open a second terminal and run:
```bash
KSVC_HOST=$(kubectl get ksvc aiops-triage-service -o jsonpath='{.status.url}' | sed 's#http://##')

curl -X POST "http://127.0.0.1:8080/analyze" \
  -H "Host: ${KSVC_HOST}" \
  -H "Content-Type: application/json" \
  -d '{
    "incident_id":"INC-90012",
    "severity":"critical",
    "text":"Payment gateway outage in production. Multiple timeout errors from nginx at 10.2.1.50. Escalate to P1 immediately."
  }'
```

PowerShell equivalent:
```powershell
$KSVC_HOST = ((kubectl get ksvc aiops-triage-service -o jsonpath="{.status.url}") -replace "http://", "")
curl.exe -X POST "http://127.0.0.1:8080/analyze" -H "Host: $KSVC_HOST" -H "Content-Type: application/json" -d '{"incident_id":"INC-90012","severity":"critical","text":"Payment gateway outage in production. Multiple timeout errors from nginx at 10.2.1.50. Escalate to P1 immediately."}'
```

If your shell variable is empty or ingress returns `404`, run a guaranteed in-cluster test:
```bash
kubectl run curl-test --rm -it --restart=Never --image=curlimages/curl -- \
  curl -sS -X POST "http://aiops-triage-service.default.svc.cluster.local/analyze" \
  -H "Content-Type: application/json" \
  -d '{"incident_id":"INC-90012","severity":"critical","text":"Payment gateway outage in production. Multiple timeout errors from nginx at 10.2.1.50. Escalate to P1 immediately."}'
```
Expected response fields:
- `analysis.urgency`
- `analysis.priority`
- `analysis.recommended_owner`
- `analysis.entities`

Capture:
- Request and JSON response in terminal

### Snapshot 5: Scale-to-zero proof
Commands:
```bash
kubectl get pods -w
# Wait for idle period and observe pod termination (scale to zero)
```
Then trigger new request:
```bash
KSVC_HOST=$(kubectl get ksvc aiops-triage-service -o jsonpath='{.status.url}' | sed 's#http://##')
curl -H "Host: ${KSVC_HOST}" "http://127.0.0.1:8080/health"
```
Capture:
- Pod count dropping to zero
- New pod spin-up after fresh request

### Snapshot 6: Autoscaling under load
Use any load tool (hey/ab/fortio). Example with hey:
```bash
KSVC_HOST=$(kubectl get ksvc aiops-triage-service -o jsonpath='{.status.url}' | sed 's#http://##')

hey -n 1000 -c 50 -m POST -H "Content-Type: application/json" \
  -host "${KSVC_HOST}" \
  -d '{"incident_id":"INC-1","severity":"high","text":"critical outage and latency spike in gateway"}' \
  "http://127.0.0.1:8080/analyze"
```
Capture:
- Load test results
- Parallel `kubectl get pods` output showing scale-out

## 4) Discussion of Knative Features Demonstrated
1. Serverless deployment:
   - No manual pod management required.
2. Scale-to-zero:
   - Saves compute when there is no traffic.
3. Rapid scale-up on demand:
   - Handles burst traffic for incident storms.
4. HTTP routing and revision management:
   - Each deployment creates a revision, simplifying rollouts.

## 5) Conclusion
This project demonstrates how a high-value, AI-style inference API can be deployed on Knative with minimal operational overhead while still supporting production-like behavior (health checks, concurrency-aware autoscaling, and burst handling).
