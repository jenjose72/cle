# 🚀 OpenFaaS + Minikube + Prometheus + Terraform + GitHub Actions

## Complete End-to-End DevOps Lab Guide

---

# 🧠 Project Overview

This project demonstrates:

* Deploying **serverless functions** using OpenFaaS
* Running on a **Kubernetes cluster (Minikube)**
* Monitoring using **Prometheus**
* Automating deployment & destruction using **Terraform + GitHub Actions (CI/CD)**

---

# 🧱 1. Prerequisites Installation

## 🐳 Docker

```bash
sudo apt update
sudo apt install docker.io -y
sudo usermod -aG docker $USER
newgrp docker
```

---

## ☸️ kubectl

```bash
sudo apt install kubectl -y
```

---

## 📦 Minikube

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

---

## ⛵ Helm

```bash
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

---

## ⚡ OpenFaaS CLI

```bash
curl -sSL https://cli.openfaas.com | sudo sh
```

---

# ⚙️ 2. Start Minikube Cluster

```bash
minikube start --driver=docker
minikube addons enable metrics-server
```

---

# 📦 3. Install OpenFaaS (with Prometheus)

## Add Helm repo

```bash
helm repo add openfaas https://openfaas.github.io/faas-netes/
helm repo update
```

---

## Create namespaces

```bash
kubectl create namespace openfaas
kubectl create namespace openfaas-fn
```

---

## Install OpenFaaS

```bash
helm upgrade openfaas openfaas/openfaas \
  --install \
  --namespace openfaas \
  --set functionNamespace=openfaas-fn \
  --set generateBasicAuth=true \
  --set prometheus.enabled=true
```

---

# 🔐 4. Get OpenFaaS Credentials

## Password

```bash
kubectl -n openfaas get secret basic-auth \
-o jsonpath="{.data.basic-auth-password}" | base64 --decode
```

---

## Gateway Access (Local)

```bash
kubectl port-forward -n openfaas svc/gateway 8081:8080
```

Open:

```
http://127.0.0.1:8081
```

---

# 🔑 5. Login to OpenFaaS

```bash
faas-cli login --username admin --password <PASSWORD> --gateway http://127.0.0.1:8081
```

---

# ⚠️ 6. Fix Python Template Error

## ❌ Error

```
Template: "python" not found
```

## Fix

```bash
faas-cli template pull https://github.com/openfaas/python-flask-template.git
```

---

# ⚡ 7. Create Function

```bash
faas-cli new hello-fn --lang python3-flask
```

---

# ⚠️ 8. IMPORTANT: Correct Function Signature

Edit:

```bash
nano hello-fn/handler.py
```

```python
def handle(event):
    return "Hello from OpenFaaS 🚀"
```

---

# ⚙️ 9. Update Image Name

Edit:

```bash
nano stack.yaml
```

```yaml
image: <your-docker-username>/hello-fn:v1
```

---

# 🐳 10. Build, Push, Deploy

```bash
docker login

faas-cli build -f stack.yaml --no-cache
faas-cli push -f stack.yaml
faas-cli deploy -f stack.yaml
```

---

# 🧪 11. Test Function

```bash
curl http://127.0.0.1:8081/function/hello-fn
```

---

# 📊 12. Monitor with Prometheus

```bash
kubectl port-forward -n openfaas svc/prometheus 9090:9090
```

Open:

```
http://127.0.0.1:9090
```

Query:

```
gateway_function_invocation_total
```

---

# 🌐 13. Expose OpenFaaS (for CI/CD using ngrok)

## Start port-forward

```bash
kubectl port-forward -n openfaas svc/gateway 8081:8080
```

---

## Run ngrok

```bash
ngrok http 8081
```

Copy URL:

```
https://xxxxx.ngrok-free.app
```

---

# 🏗️ 14. Terraform Setup

## main.tf

```hcl
provider "null" {}

resource "null_resource" "deploy_function" {
  provisioner "local-exec" {
    command = <<EOT
      faas-cli build -f stack.yaml --no-cache
      faas-cli push -f stack.yaml
      faas-cli deploy -f stack.yaml
    EOT
  }
}

resource "null_resource" "destroy_function" {
  provisioner "local-exec" {
    when = destroy
    command = "faas-cli remove hello-fn --gateway $OPENFAAS_URL"
  }
}
```

---

## Run locally

```bash
terraform init
terraform apply -auto-approve
terraform destroy -auto-approve
```

---

# 🤖 15. GitHub Actions CI/CD

## Folder Structure

```
.github/workflows/
    deploy.yml
    destroy.yml
```

---

# 🚀 Deploy Workflow (`deploy.yml`)

```yaml
name: OpenFaaS Deploy

on:
  push:
    branches:
      - main

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3

    - name: Install faas-cli
      run: curl -sSL https://cli.openfaas.com | sudo sh

    - name: Setup Terraform
      uses: hashicorp/setup-terraform@v3

    - name: Login to Docker
      run: echo "${{ secrets.DOCKER_PASSWORD }}" | docker login -u ${{ secrets.DOCKER_USERNAME }} --password-stdin

    - name: Login to OpenFaaS
      run: echo "${{ secrets.OPENFAAS_PASSWORD }}" | faas-cli login --username admin --password-stdin --gateway ${{ secrets.OPENFAAS_URL }}

    - name: Terraform Init
      run: terraform init

    - name: Terraform Apply
      run: terraform apply -auto-approve
```

---

# 💣 Destroy Workflow (`destroy.yml`)

```yaml
name: Destroy OpenFaaS Function

on:
  workflow_dispatch:

jobs:
  destroy:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3

    - name: Install faas-cli
      run: curl -sSL https://cli.openfaas.com | sudo sh

    - name: Setup Terraform
      uses: hashicorp/setup-terraform@v3

    - name: Login to OpenFaaS
      run: echo "${{ secrets.OPENFAAS_PASSWORD }}" | faas-cli login --username admin --password-stdin --gateway ${{ secrets.OPENFAAS_URL }}

    - name: Terraform Init
      run: terraform init

    - name: Terraform Destroy
      env:
        OPENFAAS_URL: ${{ secrets.OPENFAAS_URL }}
      run: terraform destroy -auto-approve
```

---

# 🔐 16. GitHub Secrets

Add in repo → Settings → Secrets:

* `DOCKER_USERNAME`
* `DOCKER_PASSWORD`
* `OPENFAAS_URL` (ngrok URL)
* `OPENFAAS_PASSWORD`

---

# ⚠️ Important Notes

* Minikube is local → not accessible from GitHub
* Use **ngrok** or cloud deployment for CI/CD
* Keep ngrok running during CI/CD execution

---

# 🧠 Key Concepts

| Tool           | Purpose                   |
| -------------- | ------------------------- |
| OpenFaaS       | Serverless functions      |
| Minikube       | Local Kubernetes          |
| Prometheus     | Monitoring                |
| Terraform      | Infrastructure automation |
| GitHub Actions | CI/CD                     |

---

# ⚡ Common Errors & Fixes

| Error                 | Fix                       |
| --------------------- | ------------------------- |
| Template not found    | Pull python template      |
| 500 error             | Fix handler signature     |
| Old code running      | Change image tag          |
| Terraform yml error   | Use stack.yaml            |
| Gateway not reachable | Use ngrok                 |
| Destroy fails         | Pass OPENFAAS_URL via env |

---

# 🎯 Final Outcome

You successfully built:

✅ Serverless platform
✅ Kubernetes deployment
✅ Monitoring system
✅ CI/CD pipeline
✅ Infrastructure lifecycle (deploy + destroy)

---

# 🔥 Conclusion

This project demonstrates a complete **DevOps lifecycle** from development to deployment and teardown using modern cloud-native tools.

---

**You are now officially DevOps-ready 🚀**
