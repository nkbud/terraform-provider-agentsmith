terraform {
  required_providers {
    agentsmith = {
      source = "artifactory.nisc.coop/nisc/agentsmith"
    }
  }
}

provider "agentsmith" {}

data "agentsmith_claude" "example" {}
