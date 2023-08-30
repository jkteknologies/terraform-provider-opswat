terraform {
  required_providers {
    opswat = {
      source = "opswat"
    }
  }
  backend "local" {
    path = "terraform.tfstate"
  }
}

provider "opswat" {
  host   = "https://opswat.dev.av.swissre.cn"
  apikey = "8d47ca941d0cce586ea6c878d28c7fdb84c3"
}

resource "opswat_globalSync" "new" {
  timeout = 30
}

#data "opswat_globalSyncDS" "current" {}
#
#output "opswat_globalSyncDS" {
#  value = data.opswat_globalSyncDS.current
#}