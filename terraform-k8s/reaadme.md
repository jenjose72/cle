# 🚀 Cloud Lab Setup Guide

## Terraform + Kubernetes + MongoDB + GitHub Actions CI/CD

---

# 📌 Overview

This project demonstrates:

* Infrastructure as Code using Terraform
* Kubernetes deployment using Minikube
* MongoDB setup and CRUD operations
* CI/CD pipeline using GitHub Actions

---

# 🧰 Prerequisites Installation

## 1. Update System

```bash
sudo apt update && sudo apt upgrade -y
```

---

## 2. Install Docker

```bash
sudo apt install -y docker.io
sudo systemctl start docker
sudo systemctl enable docker
```

Verify:

```bash
docker --version
```

---

## 3. Install Minikube

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

Start cluster:

```bash
minikube start
```

---

## 4. Install kubectl

```bash
sudo apt install -y kubectl
```

Verify:

```bash
kubectl get nodes
```

---

## 5. Install Terraform

```bash
sudo apt install -y gnupg software-properties-common curl
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
sudo apt-add-repository "deb https://apt.releases.hashicorp.com $(lsb_release -cs) main"
sudo apt update
sudo apt install terraform
```

Verify:

```bash
terraform -v
```

---

## 6. Install MongoDB

### Add key

```bash
curl -fsSL https://pgp.mongodb.com/server-7.0.asc | sudo gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor
```

### Add repo

```bash
echo "deb [ arch=amd64 signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/ubuntu noble/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list
```

### Install

```bash
sudo apt update
sudo apt install -y mongodb-org
```

### Start service

```bash
sudo systemctl start mongod
sudo systemctl enable mongod
```

---

# 🗂️ Project Structure

```
project-root/
│
├── terraform/
│   └── main.tf
│
└── .github/
    └── workflows/
        └── deploy.yml
```

---

# 🧱 Terraform Configuration (terraform/main.tf)

```hcl
provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_deployment" "mongodb" {
  metadata {
    name = "mongodb"
    labels = { app = "mongodb" }
  }

  spec {
    replicas = 1

    selector {
      match_labels = { app = "mongodb" }
    }

    template {
      metadata {
        labels = { app = "mongodb" }
      }

      spec {
        container {
          name  = "mongodb"
          image = "mongo:latest"

          port {
            container_port = 27017
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "mongodb_service" {
  metadata {
    name = "mongodb-service"
  }

  spec {
    selector = { app = "mongodb" }

    port {
      port        = 27017
      target_port = 27017
    }

    type = "NodePort"
  }
}
```

---

# ⚙️ Run Terraform Locally

```bash
cd terraform
terraform init
terraform apply
```

Verify:

```bash
kubectl get pods
kubectl get svc
```

---

# 🍃 MongoDB Commands (CRUD)

Open shell:

```bash
mongosh
```

Create DB:

```js
use btech_journey
```

Insert:

```js
db.stories.insertOne({ title: "Internship", year: 3 })
```

Read:

```js
db.stories.find()
```

Delete:

```js
db.stories.deleteOne({ title: "Internship" })
```

Drop:

```js
db.stories.drop()
```

---

# 🤖 GitHub Actions CI/CD

## Create Workflow File

```
.github/workflows/deploy.yml
```

## Workflow Code

```yaml
name: Terraform Deploy

on:
  push:
    branches: [ "main" ]

jobs:
  terraform:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout
      uses: actions/checkout@v3

    - name: Setup Terraform
      uses: hashicorp/setup-terraform@v2

    - name: Setup Kubeconfig
      run: |
        mkdir -p $HOME/.kube
        echo "${{ secrets.KUBECONFIG }}" > $HOME/.kube/config

    - name: Terraform Init
      run: terraform init
      working-directory: ./terraform

    - name: Terraform Apply
      run: terraform apply -auto-approve
      working-directory: ./terraform
```

---

# 🔐 Setup GitHub Secret

## Get kubeconfig (IMPORTANT)

```bash
kubectl config view --flatten
```

Copy full output.

---

## Add Secret in GitHub

* Go to Repo → Settings → Secrets → Actions
* Add:

```
Name: KUBECONFIG
Value: (paste kubeconfig content)
```

---

# 🚀 Push Code

```bash
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin <repo-url>
git push -u origin main
```

---

# 📊 Monitor Pipeline

* Go to GitHub → Actions tab
* Check workflow execution

---

# 🧠 Architecture Flow

```
Git Push → GitHub Actions → Terraform → Kubernetes → MongoDB Pod
```

---

# ⚠️ Common Errors

## 1. Working directory not found

Fix folder structure.

## 2. kubeconfig error

Use:

```bash
kubectl config view --flatten
```

## 3. Minikube not accessible

CI cannot access local cluster.

---

# 🎯 Viva Key Points

* Terraform = Infrastructure as Code tool
* Kubernetes = Container orchestration
* MongoDB = NoSQL database
* CI/CD = Automation pipeline

---

# ✅ Conclusion

This setup automates deployment of MongoDB on Kubernetes using Terraform and GitHub Actions.
