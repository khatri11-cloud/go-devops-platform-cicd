# Go DevOps Platform

> A cloud-native Go web service built with Go, containerized using Docker, and deployed to Microsoft Azure using Terraform.
> 

---

---

## Project Overview

Go DevOps Platform is a hands-on engineering project designed to demonstrate the complete lifecycle of deploying a cloud-native application using modern DevOps practices.

The project begins with a production-ready Go web application and follows the same deployment workflow commonly used in production environments. The application is containerized using Docker, stored inside Azure Container Registry, and deployed to Microsoft Azure using Terraform.

Rather than learning individual technologies independently, the goal of this project was to understand how they work together as one deployment pipeline.

The repository focuses on the practical implementation of:

- Application Development
- Containerization
- Infrastructure as Code
- Cloud Deployment
- Documentation
- Engineering Best Practices

This repository represents the implementation, while the engineering journey is documented separately through a three-part blog series.

---

## Why This Project?

Although I hold the AWS Solutions Architect – Associate certification, my available AWS credits had already been exhausted.

Instead of delaying the project, I chose Microsoft Azure as the deployment platform. The primary objective was never tied to a specific cloud provider. I wanted to gain practical experience with concepts that apply across cloud platforms, including:

- Containers
- Infrastructure as Code
- Image Registries
- Cloud Deployments
- Production-oriented application design

Building this project helped bridge the gap between learning individual technologies and integrating them into a complete deployment workflow.

---

# Features

- Production-ready Go HTTP server
- Health endpoint
- Version endpoint
- Graceful shutdown support
- Multi-stage Docker build
- Small production-ready Docker image
- Infrastructure as Code using Terraform
- Azure Resource Group provisioning
- Azure Container Registry integration
- Azure Container Instance deployment
- Public application endpoint
- Organized project documentation
- Engineering blog series

---

# High-Level Architecture

> **Figure 1 — High-Level System Architecture**
> 

> 
> 
> 
> ![dockerarchiteture.png](/docs/extra/dockerarchiteture.png)
> 

The project follows a straightforward cloud-native deployment workflow.

Development begins locally where the Go application is implemented and tested. Once the application is ready, Docker packages it into a lightweight container image using a multi-stage build. The image is then pushed to Azure Container Registry, which serves as the central storage location for container images.

Terraform provisions all required Azure infrastructure, including the Resource Group, Azure Container Registry, and Azure Container Instance. During deployment, Azure Container Instances pull the application image directly from Azure Container Registry and expose the application through a public Fully Qualified Domain Name (FQDN).

This approach separates application development from infrastructure provisioning while ensuring deployments remain reproducible and easy to maintain.

---

# Technology Stack

| Category | Technology |
| --- | --- |
| Programming Language | Go |
| HTTP Server | net/http |
| Containerization | Docker |
| Container Registry | Azure Container Registry |
| Infrastructure as Code | Terraform |
| Cloud Platform | Microsoft Azure |
| Compute Platform | Azure Container Instances |
| Version Control | Git |
| Repository Hosting | GitHub |
| Documentation | Markdown |
| Architecture Diagrams | [draw.io](http://draw.io/) |

---

# Repository Structure

```
go-devops-platform/

├── app/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   │
│   ├── internal/
│   │   ├── config/
│   │   └── handlers/
│   │
│   ├── Dockerfile
│   ├── Makefile
│   └── go.mod
│
├── infrastructure/
│   └── terraform/
│       ├── provider.tf
│       ├── main.tf
│       ├── variables.tf
│       ├── outputs.tf
│       └── versions.tf
│
├── docs/
│   ├── project-1/
│   │   ├── part-1/
│   │   ├── part-2/
│   │   └── part-3/
│   │
│   └── architecture/
│
├── README.md
├── LICENSE
└── .gitignore
```

### Repository Organization

The repository is intentionally divided into independent sections to improve maintainability.

### app/

Contains the Go application source code.

Responsibilities include:

- HTTP server
- Configuration
- Request handlers
- Docker build
- Dependency management

---

### infrastructure/

Contains Terraform configuration responsible for provisioning Azure resources.

Resources created include:

- Azure Resource Group
- Azure Container Registry
- Azure Container Instance

Infrastructure is completely separated from application code, allowing changes to infrastructure without modifying the application.

---

### docs/

Contains project documentation.

This directory stores:

- Engineering blog posts
- Screenshots
- Architecture diagrams
- Project notes

The documentation is intentionally separated from the application source code to keep the repository organized.

---

## Engineering Blog Series

The implementation process for this project is documented as a three-part engineering blog series.

## Engineering Blog Series

This project is documented as a three-part engineering blog series where I explain the design decisions, implementation process, challenges encountered, and lessons learned throughout the project.

- **[Part 1 – Building a Production-Ready Go Web Service](https://prajwalsinghkhatri.com.np/content/posts/project1_cloud-native_go_deployement/part1/)**
- **[Part 2 – Containerizing a Go Application with Docker](https://prajwalsinghkhatri.com.np/content/posts/project1_cloud-native_go_deployement/part2/)**
- **[Part 3 – Deploying a Go Application to Azure with Terraform](https://prajwalsinghkhatri.com.np/content/posts/project1_cloud-native_go_deployement/part3/)**

The README provides a technical overview of the repository, while the blog series explains the design decisions, implementation process, challenges encountered, and lessons learned throughout the project.

---