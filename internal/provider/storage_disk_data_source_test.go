package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAcc_StorageDiskDataSource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "tf-ds-test-" + testRunID

	resourcesConfig := fmt.Sprintf(`
resource "cudo_storage_disk" "disk" {
	data_center_id = "black-mesa"
	id = "%s"
	size_gib = 15
}`, name)

	testAccStorageDiskDataSourceConfig := fmt.Sprintf(`
data "cudo_storage_disk" "test" {
	id = "%s"
}`, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(state *terraform.State) error {
			cl, _ := getClients(t)

			ins, err := cl.GetDisk(ctx, &vm.GetDiskRequest{
				Id:        name,
				ProjectId: projectID,
			})

			if err == nil && ins.Disk.DiskState != vm.Disk_DISK_STATE_DELETE {
				res, err := cl.DeleteStorageDisk(ctx, &vm.DeleteStorageDiskRequest{
					Id:        name,
					ProjectId: projectID,
				})
				t.Logf("(%s) %#v: %v", ins.Disk.DiskState, res, err)

				return fmt.Errorf("disk resource not destroyed %s , %s", ins.Disk.Id, ins.Disk.DiskState)
			}
			return nil
		},

		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: getProviderConfig() + resourcesConfig,
			},
			{
				Config: getProviderConfig() + resourcesConfig + testAccStorageDiskDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_storage_disk.test", "id", name),
					resource.TestCheckResourceAttr("data.cudo_storage_disk.test", "size_gib", "15"),
				),
			},
		},
	})
}
