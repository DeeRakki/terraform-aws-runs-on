terraform {
  # control_plane_fleet uses a `removed` block, which needs Terraform or OpenTofu 1.7.
  required_version = ">= 1.7.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 6.45"
    }
    random = {
      source  = "hashicorp/random"
      version = ">= 3.7"
    }
  }
}
