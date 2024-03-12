package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAcc_VMInstanceDataSource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "tf-ds-test-" + testRunID

	resourcesConfig := fmt.Sprintf(`
resource "cudo_vm" "my-vm" {
   machine_type       = "standard"
   data_center_id     = "black-mesa"
   vcpus              = 1
   boot_disk = {
     image_id = "alpine-linux-317"
     size_gib = 1
   }
   memory_gib         = 2
   metadata = {
	testkey1 = "testval1"
	testkey2 = "testval2"
   }
   id                 = "%s"
   networks = [
    {
      network_id         = "tf-test"
    }
  ]
 }`, name)

	testAccVMInstanceDataSourceConfig := fmt.Sprintf(`
data "cudo_vm" "test" {
	id = "%s"
}`, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(state *terraform.State) error {
			cl, _ := getClients(t)

			ins, err := cl.GetVM(ctx, &vm.GetVMRequest{
				Id:        name,
				ProjectId: projectID,
			})
			if err == nil && ins.VM.ShortState != "epil" {
				res, err := cl.TerminateVM(ctx, &vm.TerminateVMRequest{
					Id:        name,
					ProjectId: projectID,
				})
				t.Log(res, err)

				return fmt.Errorf("vm resource not destroyed %s, %s", ins.VM.Id, ins.VM.ShortState)
			}
			return nil
		},

		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: getProviderConfig() + resourcesConfig,
			},
			{
				Config: getProviderConfig() + resourcesConfig + testAccVMInstanceDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_vm.test", "id", name)),
			},
		},
	})
}
