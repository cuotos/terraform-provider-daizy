terraform {
  required_providers {
    daizy = {
      source = "daizy/daizy"
    }
  }
}

provider "daizy" {
  authtoken    = "auth_token"
  organisation = "1"
}

data "daizy_project" "project1" {
  name    = "terra-project-2"
  user_id = "99"
}

output "all" {
  value = daizy_project.project1
}
