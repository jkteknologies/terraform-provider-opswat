terraform {
  required_providers {
    opswat = {
      source = "opswat"
    }
  }
}

provider "opswat" {}

data "opswat_test" "example" {}