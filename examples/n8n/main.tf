terraform {
  required_providers {
    n8n = {
      source = "hashicorp.com/thenom/n8n"
    }
  }
}

provider "n8n" {
  host_url = "http://localhost:5678"
  api_key  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyOTA3MWJjNC02ZDRlLTRmZjAtOTdhOS1iYjQ0M2VmNDI0N2QiLCJpc3MiOiJuOG4iLCJhdWQiOiJwdWJsaWMtYXBpIiwiaWF0IjoxNzUyNDE1MDE5LCJleHAiOjE3NTQ5NzEyMDB9.FdH20ZN4fHvJBlTQrJRpqRoSmbbvR-5WI5hkzl9ztkE"
}

data "n8n_workflows" "nz" {}

output "nz_workflow" {
  value = data.n8n_workflows.nz
}
