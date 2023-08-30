terraform {
  required_providers {
    opswat = {
      source = "opswat"
    }
  }
}

provider "opswat" {
  host   = "https://opswat.dev.av.swissre.cn"
  apikey = "8d47ca941d0cce586ea6c878d28c7fdb84c3"
}

data "opswat_globalSync" "current" {}

output "opswat_globalSync_timeout" {
  value = data.opswat_globalSync.current
}