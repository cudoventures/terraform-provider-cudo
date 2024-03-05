package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"

	// "github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var apiKey = os.Getenv("TF_TEST_CUDO_API_KEY")
var projectID = os.Getenv("TF_TEST_CUDO_PROJECT_ID")
var remoteAddr = "grpc.staging.compute.cudo.org:443"
var billingAccountID = os.Getenv("TF_TEST_CUDO_BILLING_ACCOUNT_ID")

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cudo": providerserver.NewProtocol6WithError(New("test", remoteAddr)()),
}

func getProviderConfig() string {

	return fmt.Sprintf(`
provider "cudo" {
  api_key            = "%s"
  remote_addr        = "%s"
  project_id         = "%s"
  billing_account_id = "%s"
}
`, apiKey, remoteAddr, projectID, billingAccountID)
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

var testRunID, _ = helper.NewNanoID(6)

func getClients(t *testing.T) (vm.VMServiceClient, network.NetworkServiceClient) {
	var config CudoProviderModel
	if apiKey == "" {
		t.Error("no api key for tests")
	}

	conn, err := config.dial(context.Background(), remoteAddr, apiKey, "test")
	if err != nil {
		t.Error("Dial err", err)
		// TODO: sort this out
		// resp.Diagnostics.AddAttributeError(
		// 	path.Root("project_id"),
		// 	"Missing Cudo project ID",
		// 	"The provider cannot create the client without a project_id please pass it or set the CUDO_PROJECT_ID environment variable or set it in your cudo config file.",
		// )
	}

	// TODO: it would be nice to plug the debug logging into t.Log
	// tx.Debug = true
	return vm.NewVMServiceClient(conn), network.NewNetworkServiceClient(conn)
}
