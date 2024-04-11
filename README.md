# Cudo Compute Terraform provider

This repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework). 

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

## Building the provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Using the provider

See docs directory

## Developing the provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

The Cudo API client is generated in the cudo-compute-market repo using the `make tf` command, copy the generated "compute" folder from cudo-compute-market/clients/go-grpc/ to terraform-provider-cudo/internal/.

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To run locally add the code below to ~/.terraformrc

```
provider_installation {

  dev_overrides {
      "cudoventures/cudo" = "/home/<USER-DIR>/go/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}

The [Vercel terraform provider](https://github.com/vercel/terraform-provider-vercel) is a useful reference as it uses the same plugin framework.
```

## Documentation

To generate or update documentation, run `make docs`.

or to change the name:

go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --rendered-provider-name Cudo

## Testing

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```

## Releases

To make a release:

```shell
git tag v0.2.1
git push origin v0.2.1
```
