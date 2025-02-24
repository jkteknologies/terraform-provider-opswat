// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"log"
	"os"
	"os/signal"
	"syscall"

	"terraform-provider-opswat/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "v0.1.0"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// NOTE: This is not a typical Terraform Registry provider address,
		// such as registry.terraform.io/jkteknologies/terraform-provider-opswat". This specific
		// provider address is used in these tutorials in conjunction with a
		// specific Terraform CLI configuration for manual development testing
		// of this provider.
		Address: "registry.terraform.io/jkteknologies/terraform-provider-opswat",
		Debug:   debug,
	}

	// Create a context that will be canceled when a termination signal is received
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancellation when done

	// Set up a signal handler to capture termination signals and cancel the context
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		// Received a termination signal; cancel the context
		cancel()
	}()

	err := providerserver.Serve(ctx, opswatProvider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}

}
