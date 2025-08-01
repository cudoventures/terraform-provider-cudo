package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"google.golang.org/grpc/codes"
)

func TestAcc_NetworkResource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	networkID := "network-resource-" + testRunID

	networkConfig := fmt.Sprintf(`
resource "cudo_network" "network" {
	data_center_id = "black-mesa"
	id             = "%s"
	ip_range       = "192.168.0.0/24"
 }`, networkID)

	resource.ParallelTest(t, resource.TestCase{
		CheckDestroy: func(state *terraform.State) error {
			_, cl := getClients(t)

			getParams := &network.GetNetworkRequest{
				Id:        networkID,
				ProjectId: projectID,
			}

			getRes, err := cl.GetNetwork(ctx, getParams)
			if helper.IsErrCode(err, codes.NotFound) {
				// successfully destroyed already
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not get network after resource create %s, %v", networkID, err)
			}
			if getRes.State != network.Network_DELETING {
				terminateParams := &network.DeleteNetworkRequest{
					Id:        networkID,
					ProjectId: projectID,
				}
				_, err := cl.DeleteNetwork(ctx, terminateParams)
				if err != nil {
					return fmt.Errorf("network resource not deleted %s , %s , %s", getRes.Id, getRes.State.String(), err)
				}

				if _, err := waitForNetworkDelete(ctx, cl, projectID, networkID); err != nil {
					return fmt.Errorf("error waiting for network delete %s , %s , %s", getRes.Id, getRes.State.String(), err)
				}

				return nil
			}
			return nil
		},
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: getProviderConfig() + networkConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_network.network", "id", networkID),
					resource.TestCheckResourceAttr("cudo_network.network", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_network.network", "ip_range", "192.168.0.0/24"),
					resource.TestCheckResourceAttrSet("cudo_network.network", "internal_ip_address"),
					resource.TestCheckResourceAttrSet("cudo_network.network", "external_ip_address"),
					resource.TestCheckResourceAttrSet("cudo_network.network", "gateway"),
				),
			},
		},
	})
}
