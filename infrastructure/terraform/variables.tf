variable "resource_group_name" {
  description = "Azure Resource Group name"
  type        = string
  default     = "devops-platform-rg"
}

variable "location" {
  description = "Azure region"
  type        = string
  default     = "East US"
}

variable "acr_name" {
  description = "Azure Container Registry name"
  type        = string
}