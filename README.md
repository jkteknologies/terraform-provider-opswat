Terraform Provider For OPSWAT Metadefender Core
==================

- Official API - https://docs.opswat.com/mdcore/metadefender-core (OAS 3)
- Some additional APIs (not provided in the official api doc) supported (workflows, userdirectories for sso)
- Documentation: https://registry.terraform.io/providers/jkteknologies/opswat/latest/docs
- [![Release Go project](https://github.com/jkteknologies/terraform-provider-opswat/actions/workflows/release.yaml/badge.svg)](https://github.com/jkteknologies/terraform-provider-opswat/actions/workflows/release.yaml)

Supported Versions
------------------

| Terraform version | Minimum Core version   | Maximum Core version |
|-------------------|------------------------|----------------------| 
| >= 1.10.x	        | 5.13.0	               | latest               |

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 1.10+
-	[Go](https://golang.org/doc/install) 1.23+ (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/jkteknologies/terraform-provider-opswat`

```sh
$ git clone https://github.com/jkteknologies/terraform-provider-opswat.git
$ cd terraform-provider-opswat/
```

Enter the provider directory and build/install the provider

```sh
$ cd terraform-provider-opswat/
$ go install .
```

Using the provider
----------------------
```hcl
terraform {
  required_providers {
    opswat = {
      source = "jkteknologies/opswat"
    }
  }
}
```


## Developing the Provider
---------------------------

If you are new to plugin development, study the [Terraform Plugin Framework tutorial](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider) first.

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.23+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

### Windows development

To compile the provider, run `go install .`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```shell
$ cd C:\Users\xxx\AppData\Roaming
```

To test local build - create/edit `terraform.rc` file in %APPDATA% Roaming:

```text
provider_installation {

  dev_overrides {
    "opswat" = "C:/Users/xxx/go/bin" #GOBIN location
  }

  direct {}
}
```

### Linux development

To compile the provider, run `go install .`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```shell
cd ~/go/bin
```

To test local build - create/edit `.terraformrc` file in home directory:

```text
provider_installation {

  dev_overrides {
    "opswat" = "/home/xxx/go/bin" #GOBIN location
  }

  direct {}
}
```
