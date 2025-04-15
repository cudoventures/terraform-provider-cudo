package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAcc_NetworkDataSource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "tf-test"

	testAccSecurityGroupDataSourceConfig := fmt.Sprintf(`
	  data "cudo_network" "network" {
		  id = "%s"
	  }`, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },

		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: getProviderConfig() + testAccSecurityGroupDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_network.network", "id", name),
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
