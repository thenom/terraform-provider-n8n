terraform {
  required_providers {
    n8n = {
      source = "hashicorp.com/thenom/n8n"
    }
  }
}

provider "n8n" {}

data "n8n_workflow" "example" {
  id = "flibble"
}
