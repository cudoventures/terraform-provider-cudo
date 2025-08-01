package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"google.golang.org/grpc/codes"
)

func TestAcc_VMInstanceDataSource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	vmID := "tf-ds-test-" + testRunID

	resourcesConfig := fmt.Sprintf(`
resource "cudo_vm" "my-vm" {
   machine_type       = "standard"
   data_center_id     = "black-mesa"
   vcpus              = 1
   boot_disk = {
     image_id = "alpine-linux-319"
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
 }`, vmID)

	testAccVMInstanceDataSourceConfig := fmt.Sprintf(`
data "cudo_vm" "test" {
	id = "%s"
}`, vmID)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(state *terraform.State) error {
			cl, _ := getClients(t)

			getRes, err := cl.GetVM(ctx, &vm.GetVMRequest{
				Id:        vmID,
				ProjectId: projectID,
			})
			if helper.IsErrCode(err, codes.NotFound) {
				// successfully destroyed already
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not get vm after resource create %s, %v", vmID, err)
			}

			if getRes.VM.State != vm.VM_DELETING {
				_, err = cl.TerminateVM(ctx, &vm.TerminateVMRequest{
					Id:        vmID,
					ProjectId: projectID,
				})
				if err != nil {
					return fmt.Errorf("vm resource not destroyed %s, %s", getRes.VM.Id, getRes.VM.State.String())
				}
			}
			if getRes, err := waitForVmDelete(ctx, cl, projectID, vmID); err != nil {
				return fmt.Errorf("error waiting for vm delete %s, %s, %v", getRes.VM.Id, getRes.VM.State.String(), err)
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
					resource.TestCheckResourceAttr("data.cudo_vm.test", "id", vmID)),
			},
		},
	})
}
