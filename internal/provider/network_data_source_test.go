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

func TestAcc_NetworkDataSource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	networkID := "tf-test"

	testAccSecurityGroupDataSourceConfig := fmt.Sprintf(`
	  data "cudo_network" "network" {
		  id = "%s"
	  }`, networkID)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },

		CheckDestroy: func(state *terraform.State) error {
			_, cl := getClients(t)

			getParams := &network.GetNetworkRequest{
				Id:        networkID,
				ProjectId: projectID,
			}

			_, err := cl.GetNetwork(ctx, getParams)
			if !helper.IsErrCode(err, codes.OK) {
				return fmt.Errorf("required tf-test network missing from test project")
			}
			return nil
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: getProviderConfig() + testAccSecurityGroupDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_network.network", "id", networkID),
					resource.TestCheckResourceAttr("data.cudo_network.network", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("data.cudo_network.network", "ip_range", "192.168.0.0/24"),
					resource.TestCheckResourceAttrSet("data.cudo_network.network", "internal_ip_address"),
					resource.TestCheckResourceAttrSet("data.cudo_network.network", "external_ip_address"),
					resource.TestCheckResourceAttrSet("data.cudo_network.network", "gateway"),
				),
			},
		},
	})
}
