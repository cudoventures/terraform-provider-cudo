package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAcc_SecurityGroupResource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "sg-resource-" + testRunID

	sgConfig := fmt.Sprintf(`
	resource "cudo_security_group" "sg" {
		id             = "%s"
		data_center_id = "black-mesa"
		description    = "security group for a web server"
		rules = [
		  {
			ports     = "80"
			rule_type = "outbound"
			protocol  = "tcp"
		  },
		  {
			ip_range  = "192.168.0.0/24"
			ports     = "22,80,443"
			rule_type = "inbound"
			protocol  = "tcp"
		  }
		]
	  }`, name)

	resource.ParallelTest(t, resource.TestCase{
		CheckDestroy: func(state *terraform.State) error {
			_, cl := getClients(t)

			getParams := &network.GetSecurityGroupRequest{
				Id:        name,
				ProjectId: projectID,
			}

			getRes, err := cl.GetSecurityGroup(ctx, getParams)
			if err == nil {
				deleteParams := &network.DeleteSecurityGroupRequest{
					Id:        name,
					ProjectId: projectID,
				}
				res, err := cl.DeleteSecurityGroup(ctx, deleteParams)
				t.Logf("%#v: %v", res, err)

				return fmt.Errorf("security groupd not deleted %s", getRes.Id)
			}
			return nil
		},
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: getProviderConfig() + sgConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_security_group.sg", "id", name),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "description", "security group for a web server"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "rules.#", "2"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "rules.0.ports", "80"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "rules.0.rule_type", "outbound"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "rules.0.protocol", "tcp"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "rules.1.ip_range", "192.168.0.0/24"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "rules.1.ports", "22,80,443"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "rules.1.rule_type", "inbound"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "rules.1.protocol", "tcp"),
				),
			},
		},
	})
}
