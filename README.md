# Go DevOps Platform – Automated CI/CD Pipeline

> A production-oriented DevOps project demonstrating Continuous Integration and Continuous Deployment of a containerized Go application using GitHub Actions, Docker, Terraform, and Microsoft Azure.

[CI Badge]
[CD Badge]

---

## Overview

This repository demonstrates how a Go web application can be automatically built, validated, containerized, and deployed to Microsoft Azure whenever code is pushed to GitHub.

The project extends my previous Go deployment project by replacing the manual release process with a fully automated CI/CD pipeline.

The pipeline performs:

- Continuous Integration using GitHub Actions
- Static code analysis
- Unit testing
- Docker image creation
- Docker image publishing to Azure Container Registry
- Automatic deployment to Azure Container Instances
- Health verification after deployment

The objective is to demonstrate production-oriented DevOps workflows rather than application complexity.

---

# Live Application

Application

http://prajwal-devops-platform.eastus.azurecontainer.io:8080

Health Endpoint

http://prajwal-devops-platform.eastus.azurecontainer.io:8080/health

Version Endpoint

http://prajwal-devops-platform.eastus.azurecontainer.io:8080/version

---

# Project Architecture

(Insert High-Level Architecture Diagram)

Developer
        │
        ▼
Local Development (WSL)
        │
        ▼
Git Commit
        │
        ▼
GitHub Repository
        │
        ▼
GitHub Actions
        │
        ├── Continuous Integration
        │       │
        │       ├── Checkout Repository
        │       ├── Setup Go
        │       ├── Download Dependencies
        │       ├── Verify Formatting
        │       ├── Go Vet
        │       ├── Unit Tests
        │       └── Build Application
        │
        ▼
Continuous Deployment
        │
        ├── Azure Authentication
        ├── Docker Build
        ├── Tag Image (latest + Commit SHA)
        ├── Push to Azure Container Registry
        ├── Delete Existing Container
        ├── Create Updated Container
        └── Health Verification
        │
        ▼
Azure Container Instance
        │
        ▼
Users

---

# CI/CD Pipeline

## Continuous Integration

Every push to the `main` branch automatically triggers the Continuous Integration workflow.

The CI pipeline is responsible for validating every code change before deployment.

Checks performed:

- Repository checkout
- Go environment setup
- Dependency download
- Code formatting verification
- Static analysis using Go Vet
- Unit testing
- Application build

Deployment only proceeds if every stage completes successfully.

---

## Continuous Deployment

After a successful CI run, the Continuous Deployment workflow executes automatically.

The deployment pipeline performs the following tasks:

1. Authenticate with Microsoft Azure
2. Authenticate with Azure Container Registry
3. Build a Docker image
4. Tag the image using:
   - latest
   - Git Commit SHA
5. Push both images to Azure Container Registry
6. Delete the existing Azure Container Instance
7. Deploy a new Azure Container Instance
8. Poll the `/health` endpoint until the application becomes healthy
9. Publish a deployment summary in GitHub Actions

This enables fully automated deployments with immutable image versions.

---

# Technology Stack

| Category | Technology |
|-----------|------------|
| Language | Go |
| CI | GitHub Actions |
| CD | GitHub Actions |
| Containers | Docker |
| Registry | Azure Container Registry |
| Cloud | Microsoft Azure |
| Hosting | Azure Container Instances |
| IaC | Terraform |
| Version Control | Git & GitHub |

---

# Repository Structure

```text
go-devops-platform-cicd/
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── app/
│   ├── cmd/
│   ├── internal/
│   ├── Dockerfile
│   └── go.mod
│
├── infrastructure/
│   └── terraform/
│
├── docs/
│   ├── project-1/
│   └── project-2/
│
└── README.md
```

| Directory | Purpose |
|-----------|---------|
| `.github/workflows` | CI/CD pipelines |
| `app` | Go web application |
| `infrastructure` | Terraform infrastructure |
| `docs` | Engineering blog source files |

---

# Infrastructure

Infrastructure is provisioned using Terraform.

Azure resources include:

- Azure Resource Group
- Azure Container Registry
- Azure Container Instance

Terraform provisions the infrastructure once.

GitHub Actions continuously deploys new application versions to the existing infrastructure.

---

# Engineering Blogs

The complete engineering journey is documented separately from the repository documentation.

## Project 2 — Part 1

Building Continuous Integration with GitHub Actions

https://prajwalsinghkhatri.com.np/content/posts/project2_cicd_github_actions/part1/

## Project 2 — Part 2

Building an Automated Continuous Deployment Pipeline

https://prajwalsinghkhatri.com.np/content/posts/project2_cicd_github_actions/part2/

---

# Future Roadmap

- ✅ Project 1 — Go Web Application
- ✅ Project 2 — CI/CD Pipeline
- ⏳ Project 3 — Monitoring (Prometheus & Grafana)
- ⏳ Project 4 — Kubernetes (AKS)
- ⏳ Project 5 — GitOps (Argo CD)
- ⏳ Project 6 — Observability
- ⏳ Project 7 — Production Platform

Each project builds upon the previous one to demonstrate progressively more advanced DevOps engineering practices.

---

# License

This repository is published for educational and portfolio purposes.
