package provider

import (
	"context"
	"crypto/x509"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

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

	// dial without block handle errors in the rpcs
	dialOptions := []grpc.DialOption{}
	pool, err := x509.SystemCertPool()
	if err != nil {
		t.Error("Error getting system cerificate: ", err.Error())
	}
	creds := credentials.NewClientTLSFromCert(pool, "")
	dialOptions = append(dialOptions, grpc.WithTransportCredentials(creds))
	dialOptions = append(dialOptions,
		grpc.WithPerRPCCredentials(&apiKeyCallOption{
			disableTransportSecurity: config.DisableTLS.ValueBool(),
			key:                      apiKey,
		}),
	)

	dialTimeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(dialTimeoutCtx, remoteAddr, dialOptions...)
	if err != nil {
		t.Error("Dialing compute service failed: " + err.Error())
	}

	// TODO: it would be nice to plug the debug logging into t.Log
	// tx.Debug = true
	return vm.NewVMServiceClient(conn), network.NewNetworkServiceClient(conn)
}
