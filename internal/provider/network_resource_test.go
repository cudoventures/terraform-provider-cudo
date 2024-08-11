package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAcc_NetworkResource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "network-resource-" + testRunID

	networkConfig := fmt.Sprintf(`
resource "cudo_network" "network" {
	data_center_id = "black-mesa"
	id             = "%s"
	ip_range       = "192.168.0.0/24"
 }`, name)

	resource.ParallelTest(t, resource.TestCase{
		CheckDestroy: func(state *terraform.State) error {
			_, cl := getClients(t)

			getParams := &network.GetNetworkRequest{
				Id:        name,
				ProjectId: projectID,
			}

			getRes, err := cl.GetNetwork(ctx, getParams)
			if err == nil && (getRes.Network.State != network.Network_DELETING && getRes.Network.State != network.Network_STOPPING && getRes.Network.State != network.Network_SUSPENDING) {
				stopParams := &network.StopNetworkRequest{
					Id:        name,
					ProjectId: projectID,
				}
				stopRes, err := cl.StopNetwork(ctx, stopParams)
				t.Logf("(%s) %#v: %v", getRes.Network.State.String(), stopRes, err)
				if err != nil {
					return fmt.Errorf("network resource not stopped %s , %s , %s", getRes.Network.Id, getRes.Network.State.String(), err)
				}

				if _, err := waitForNetworkStop(ctx, projectID, name, cl); err != nil {
					return fmt.Errorf("error waiting for network stopped %s , %s , %s", getRes.Network.Id, getRes.Network.State.String(), err)
				}

				terminateParams := &network.DeleteNetworkRequest{
					Id:        name,
					ProjectId: projectID,
				}
				res, err := cl.DeleteNetwork(ctx, terminateParams)
				t.Logf("(%s) %#v: %v", getRes.Network.State.String(), res, err)

				return fmt.Errorf("network resource not deleted %s , %s", getRes.Network.Id, getRes.Network.State.String())
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
					resource.TestCheckResourceAttr("cudo_network.network", "id", name),
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
