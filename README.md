Terraform Provider For OPSWAT
==================

- Tutorials: [learn.hashicorp.com](https://learn.hashicorp.com/terraform?track=getting-started#getting-started)
- Documentation: https://www.terraform.io/docs/providers/alicloud/index.html
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

Supported Versions
------------------

| Terraform version | minimum provider version |maximum provider version
|-------------------|--------------------------| ----| 
| >= 1.5.x	         | 1.0.0	                   | latest |

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 1.5+
-	[Go](https://golang.org/doc/install) 1.19+ (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/gerbil/terraform-provider-opswat`

```sh
$ git clone https://github.com/gerbil/terraform-provider-opswat.git
$ cd terraform-provider-opswat/
```

Enter the provider directory and build/install the provider

```sh
$ cd terraform-provider-opswat/
$ go install .
```

Using the provider
----------------------
?


## Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.19+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `go install .`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.


On windows
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