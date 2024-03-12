package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAcc_NetworkDataSource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "network-data-source-" + testRunID

	sgConfig := fmt.Sprintf(`
	resource "cudo_network" "network" {
		data_center_id = "black-mesa"
		id             = "%s"
		ip_range       = "192.168.0.0/24"
	 }`, name)

	testAccSecurityGroupDataSourceConfig := fmt.Sprintf(`
	  data "cudo_network" "network" {
		  id = "%s"
	  }`, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(state *terraform.State) error {
			_, cl := getClients(t)

			getRes, err := cl.GetNetwork(ctx, &network.GetNetworkRequest{
				Id:        name,
				ProjectId: projectID,
			})
			if err == nil && getRes.Network.ShortState != "epil" {
				stopParams := &network.StopNetworkRequest{
					Id:        name,
					ProjectId: projectID,
				}
				stopRes, err := cl.StopNetwork(ctx, stopParams)
				t.Logf("(%s) %#v: %v", getRes.Network.ShortState, stopRes, err)
				if err != nil {
					return fmt.Errorf("network resource not stopped %s , %s , %s", getRes.Network.Id, getRes.Network.ShortState, err)
				}

				if _, err := waitForNetworkStop(ctx, projectID, name, cl); err != nil {
					return fmt.Errorf("error waiting for network stopped %s , %s , %s", getRes.Network.Id, getRes.Network.ShortState, err)
				}

				terminateParams := &network.DeleteNetworkRequest{
					Id:        name,
					ProjectId: projectID,
				}
				res, err := cl.DeleteNetwork(ctx, terminateParams)
				t.Logf("(%s) %#v: %v", getRes.Network.ShortState, res, err)

				return fmt.Errorf("network resource not deleted %s , %s", getRes.Network.Id, getRes.Network.ShortState)
			}
			return nil
		},

		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: getProviderConfig() + sgConfig,
			},
			{
				Config: getProviderConfig() + sgConfig + testAccSecurityGroupDataSourceConfig,
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
