resource "azurerm_resource_group" "main" {
  name     = var.resource_group_name
  location = var.location
}

resource "azurerm_container_registry" "acr" {
  name                = var.acr_name
  resource_group_name = azurerm_resource_group.main.name
  location            = azurerm_resource_group.main.location

  sku           = "Basic"
  admin_enabled = true
}

resource "azurerm_container_group" "app" {
  name                = "devops-platform-app"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  ip_address_type     = "Public"
  os_type             = "Linux"

  image_registry_credential {
    server   = azurerm_container_registry.acr.login_server
    username = azurerm_container_registry.acr.admin_username
    password = azurerm_container_registry.acr.admin_password
  }

  container {
    name   = "devops-platform"
    image  = "${azurerm_container_registry.acr.login_server}/devops-platform:v1"
    cpu    = 1
    memory = 1

    ports {
      port     = 8080
      protocol = "TCP"
    }

    environment_variables = {
      APP_VERSION = "1.0.0"
    }
  }

  exposed_port {
    port     = 8080
    protocol = "TCP"
  }

  dns_name_label = "prajwal-devops-platform"
}