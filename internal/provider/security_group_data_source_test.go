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

func TestAcc_SecurityGroupDataSource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	securityGroupID := "sg-data-source-" + testRunID

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
	  }`, securityGroupID)

	testAccSecurityGroupDataSourceConfig := fmt.Sprintf(`
	  data "cudo_security_group" "sg" {
		  id = "%s"
	  }`, securityGroupID)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(state *terraform.State) error {
			_, cl := getClients(t)

			getRes, err := cl.GetSecurityGroup(ctx, &network.GetSecurityGroupRequest{
				Id:        securityGroupID,
				ProjectId: projectID,
			})
			if helper.IsErrCode(err, codes.NotFound) {
				// successfully destroyed already
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not get security group after resource create %s, %v", securityGroupID, err)
			}
			deleteParams := &network.DeleteSecurityGroupRequest{
				Id:        securityGroupID,
				ProjectId: projectID,
			}
			_, err = cl.DeleteSecurityGroup(ctx, deleteParams)
			if err != nil {
				return fmt.Errorf("security group not deleted %s , %s", getRes.Id, err)
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
					resource.TestCheckResourceAttr("cudo_security_group.sg", "id", securityGroupID),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_security_group.sg", "project_id", projectID),
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
