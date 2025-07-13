terraform {
  required_providers {
    n8n = {
      source = "hashicorp.com/thenom/n8n"
    }
  }
}

provider "n8n" {
  host_url = "http://localhost:5678"
  api_key  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1YTY1NTQ4ZS1mNWI3LTRlMmUtOGJmMi03ODU2M2E1ZTgxYWQiLCJpc3MiOiJuOG4iLCJhdWQiOiJwdWJsaWMtYXBpIiwiaWF0IjoxNzUyNDA4MTQ0LCJleHAiOjE3NTQ5NzEyMDB9.Ra93Tm9b4XUX9Xw9rgNnpvUKDB478ywJp27sZFYdJjM"
}

data "n8n_workflow" "nz" {
  id = "LRUL1Yajg85Au9yv"
}

output "nz_workflow" {
  value = data.n8n_workflow.nz
}
