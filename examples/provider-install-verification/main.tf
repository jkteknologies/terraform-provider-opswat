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

resource "opswat_global_sync" "new" {
  timeout = 10
}

data "opswat_global_sync" "current" {}

output "opswat_global_sync" {
  value = data.opswat_global_sync.current
}