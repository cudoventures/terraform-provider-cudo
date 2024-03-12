package main

import (
	"context"
	"flag"
	"log"

	"github.com/CudoVentures/terraform-provider-cudo/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

// set by goreleaser at build time
var (
	version           = "dev"
	commit            = "unknown"
	date              = "unknown"
	defaultRemoteAddr = "grpc.compute.cudo.org:443"
)

func main() {

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/cudoventures/cudo",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version, defaultRemoteAddr), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
