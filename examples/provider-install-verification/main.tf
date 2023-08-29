terraform {
  required_providers {
    opswat = {
      source = "opswat"
    }
  }
}

provider "opswat" {
  host   = "opswat.dev.av.swissre.cn"
  apikey = "8d47ca941d0cce586ea6c878d28c7fdb84c3"
}

data "globalSync" "example" {}