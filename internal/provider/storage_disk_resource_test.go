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

func TestAcc_StorageDiskResource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	diskID := "disk-resource-" + testRunID

	diskConf := fmt.Sprintf(`
resource "cudo_storage_disk" "storage_disk_resource" {
	data_center_id = "black-mesa"
	id = "%s"
	size_gib = 15
}`, diskID)

	resource.ParallelTest(t, resource.TestCase{
		CheckDestroy: func(state *terraform.State) error {
			cl, _ := getClients(t)

			getRes, err := cl.GetDisk(ctx, &vm.GetDiskRequest{
				Id:        diskID,
				ProjectId: projectID,
			})
			if helper.IsErrCode(err, codes.NotFound) {
				// successfully destroyed already
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not get disk after resource create %s, %v", diskID, err)
			}

			if getRes.Disk.DiskState != vm.Disk_DELETING {
				_, err := cl.DeleteStorageDisk(ctx, &vm.DeleteStorageDiskRequest{
					Id:        diskID,
					ProjectId: projectID,
				})
				if err != nil {
					return fmt.Errorf("disk resource not destroyed %s, %s, %v", getRes.Disk.Id, getRes.Disk.DiskState, err)
				}

				if err := waitForDiskDelete(ctx, cl, projectID, diskID); err != nil {
					return fmt.Errorf("error waiting for disk delete %s, %s, %v", getRes.Disk.Id, getRes.Disk.DiskState.String(), err)
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
				Config: getProviderConfig() + diskConf,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_storage_disk.storage_disk_resource", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_storage_disk.storage_disk_resource", "size_gib", "15"),
					resource.TestCheckResourceAttr("cudo_storage_disk.storage_disk_resource", "id", diskID),
					resource.TestCheckResourceAttr("cudo_storage_disk.storage_disk_resource", "project_id", projectID),
				),
			},
		},
	})
}
