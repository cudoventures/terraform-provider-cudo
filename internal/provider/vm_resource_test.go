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

func TestAcc_VMResource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "vm-resource-" + testRunID

	vmConfig := fmt.Sprintf(`
resource "cudo_vm" "vm" {
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
   id                = "%s"
   networks = [
    {
      network_id         = "tf-test"
    }
  ]
 }`, name)

	resource.ParallelTest(t, resource.TestCase{
		CheckDestroy: func(state *terraform.State) error {
			cl, _ := getClients(t)

			getParams := &vm.GetVMRequest{
				Id:        name,
				ProjectId: projectID,
			}

			getRes, err := cl.GetVM(ctx, getParams)
			if err == nil && getRes.VM.State != vm.VM_DELETING {
				terminateParams := &vm.TerminateVMRequest{
					Id:        name,
					ProjectId: projectID,
				}
				res, err := cl.TerminateVM(ctx, terminateParams)
				t.Logf("(%s) %#v: %v", getRes.VM.State.String(), res, err)

				return fmt.Errorf("vm resource not destroyed %s , %s", getRes.VM.Id, getRes.VM.State.String())
			}
			return nil
		},
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: getProviderConfig() + vmConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_vm.vm", "boot_disk.image_id", "alpine-linux-319"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "boot_disk.size_gib", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.vm", "memory_gib", "2"),
					resource.TestCheckResourceAttrSet("cudo_vm.vm", "internal_ip_address"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "renewable_energy", "true"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "vcpus", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "id", name),
				),
			},
		},
	})
}

func TestAcc_VMResourceMinimal(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "vm-resource-minimal-" + testRunID

	vmConfig := fmt.Sprintf(`
resource "cudo_vm" "vm-minimal" {
   boot_disk = {
     image_id = "alpine-linux-319"
   }
   id                = "%s"
   machine_type       = "standard"
   data_center_id     = "black-mesa"
   vcpus              = 1
   memory_gib         = 2
   networks = [
    {
      network_id         = "tf-test"
    }
  ]
 }`, name)

	resource.ParallelTest(t, resource.TestCase{
		CheckDestroy: func(state *terraform.State) error {
			cl, _ := getClients(t)

			getParams := &vm.GetVMRequest{
				Id:        name,
				ProjectId: projectID,
			}

			ins, err := cl.GetVM(ctx, getParams)

			if err == nil && ins.VM.State != vm.VM_DELETING {
				terminateParams := &vm.TerminateVMRequest{
					Id:        name,
					ProjectId: projectID,
				}
				res, err := cl.TerminateVM(ctx, terminateParams)
				t.Logf("(%s) %#v: %v", ins.VM.State.String(), res, err)

				return fmt.Errorf("vm resource not destroyed %s, %s", ins.VM.Id, ins.VM.State.String())
			}
			return nil
		},
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: getProviderConfig() + vmConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "boot_disk.image_id", "alpine-linux-319"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "boot_disk.size_gib", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "memory_gib", "2"),
					resource.TestCheckResourceAttrSet("cudo_vm.vm-minimal", "internal_ip_address"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "renewable_energy", "true"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "vcpus", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "id", name),
				),
			},
		},
	})
}

func TestAcc_VMResourceOOBDelete(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "vm-resource-oob-delete-" + testRunID

	vmConfig := fmt.Sprintf(`
resource "cudo_vm" "vm-oob-delete" {
   boot_disk = {
     image_id = "alpine-linux-319"
   }
   id                = "%s"
   machine_type       = "standard"
   data_center_id     = "black-mesa"
   vcpus              = 1
   memory_gib         = 2
   networks = [
    {
      network_id         = "tf-test"
    }
  ]
 }`, name)

	resource.ParallelTest(t, resource.TestCase{
		CheckDestroy: func(state *terraform.State) error {
			cl, _ := getClients(t)

			getParams := &vm.GetVMRequest{
				Id:        name,
				ProjectId: projectID,
			}

			ins, err := cl.GetVM(ctx, getParams)

			if err == nil && ins.VM.State != vm.VM_DELETING {
				terminateParams := &vm.TerminateVMRequest{
					Id:        name,
					ProjectId: projectID,
				}
				res, err := cl.TerminateVM(ctx, terminateParams)
				t.Logf("(%s) %#v: %v", ins.VM.State.String(), res, err)

				return fmt.Errorf("vm resource not destroyed %s, %s", ins.VM.Id, ins.VM.State.String())
			}
			return nil
		},
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: getProviderConfig() + vmConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "boot_disk.image_id", "alpine-linux-319"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "boot_disk.size_gib", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "memory_gib", "2"),
					resource.TestCheckResourceAttrSet("cudo_vm.vm-oob-delete", "internal_ip_address"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "renewable_energy", "true"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "vcpus", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "id", name),
				),
				Destroy: false,
			},
			// {
			// 	PreConfig: func() {
			// 		terminateParams := vm.TerminateVMRequest(ctx)
			// 		terminateParams.Id = name
			// 		terminateParams.ProjectId = projectId
			// 		res, err := getClients(t).TerminateVM(terminateParams)
			// 		t.Log(res, err)
			// 	},
			// 	Config: getProvIderConfig() + vmConfig,
			// },
			{
				RefreshState: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "boot_disk.image_id", "alpine-linux-319"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "boot_disk.size_gib", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "memory_gib", "2"),
					resource.TestCheckResourceAttrSet("cudo_vm.vm-oob-delete", "internal_ip_address"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "renewable_energy", "true"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "vcpus", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "id", name),
				),
			},
		},
	})
}

func TestAcc_VMStorageDiskResource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	vmID := "vm-storage-disk-resource-" + testRunID
	// important that the disks are in ascending order as that is how the api returns them
	diskID1 := "vm-resource-disk1" + testRunID
	diskID2 := "vm-resource-disk2" + testRunID

	diskConfig := fmt.Sprintf(`
resource "cudo_storage_disk" "vm_resource_disk1" {
	data_center_id = "black-mesa"
	id = "%s"
	size_gib = 10
}

resource "cudo_storage_disk" "vm_resource_disk2" {
	data_center_id = "black-mesa"
	id = "%s"
	size_gib = 20
}`, diskID1, diskID2)

	vmConfig := fmt.Sprintf(`
resource "cudo_vm" "vm_disk_resource" {
   depends_on =  [cudo_storage_disk.vm_resource_disk1, cudo_storage_disk.vm_resource_disk2]
   machine_type       = "standard"
   data_center_id     = "black-mesa"
   vcpus              = 1
   boot_disk = {
     image_id = "alpine-linux-319"
     size_gib = 1
   }
   memory_gib         = 2
   id                = "%s"
   networks = [
    {
      network_id         = "tf-test"
    }
  ]
  storage_disks = [ 
        {
            disk_id = "%s"
        },
        {
            disk_id = "%s"
        }
    ]
 }`, vmID, diskID1, diskID2)

	resource.ParallelTest(t, resource.TestCase{
		CheckDestroy: func(state *terraform.State) error {
			cl, _ := getClients(t)

			getParams := &vm.GetVMRequest{
				Id:        vmID,
				ProjectId: projectID,
			}

			getRes, err := cl.GetVM(ctx, getParams)
			if err == nil || !helper.IsErrCode(err, codes.NotFound) {
				// successfully destroyed already
				if getRes.VM.State != vm.VM_DELETING {
					terminateParams := &vm.TerminateVMRequest{
						Id:        vmID,
						ProjectId: projectID,
					}
					_, err = cl.TerminateVM(ctx, terminateParams)
					if err != nil {
						return fmt.Errorf("vm resource not destroyed %s, %s", getRes.VM.Id, getRes.VM.State.String())
					}

					if getRes, err := waitForVmDelete(ctx, cl, projectID, vmID); err != nil {
						return fmt.Errorf("error waiting for vm delete %s, %s, %v", getRes.VM.Id, getRes.VM.State.String(), err)
					}
				}
			}

			_, err = cl.GetDisk(ctx, &vm.GetDiskRequest{
				ProjectId: projectID,
				Id:        diskID1,
			})
			if !helper.IsErrCode(err, codes.NotFound) {
				if _, err := cl.DeleteStorageDisk(ctx, &vm.DeleteStorageDiskRequest{
					ProjectId: projectID,
					Id:        diskID1,
				}); err != nil {
					return fmt.Errorf("error deleting storage disk %s, %s, %v", projectID, diskID1, err)
				}
				if err := waitForDiskDelete(ctx, cl, projectID, diskID1); err != nil {
					return fmt.Errorf("error waiting for disk delete %s, %s, %v", projectID, diskID1, err)
				}
			}

			_, err = cl.GetDisk(ctx, &vm.GetDiskRequest{
				ProjectId: projectID,
				Id:        diskID2,
			})
			if !helper.IsErrCode(err, codes.NotFound) {
				if _, err := cl.DeleteStorageDisk(ctx, &vm.DeleteStorageDiskRequest{
					ProjectId: projectID,
					Id:        diskID2,
				}); err != nil {
					return fmt.Errorf("error deleting storage disk %s, %s, %v", projectID, diskID2, err)
				}
				if err := waitForDiskDelete(ctx, cl, projectID, diskID2); err != nil {
					return fmt.Errorf("error waiting for disk delete %s, %s, %v", projectID, diskID2, err)
				}
			}
			return nil
		},
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: getProviderConfig() + diskConfig + vmConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "boot_disk.image_id", "alpine-linux-319"),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "boot_disk.size_gib", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "project_id", projectID),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "memory_gib", "2"),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "storage_disks.0.disk_id", diskID1),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "storage_disks.1.disk_id", diskID2),
					resource.TestCheckResourceAttrSet("cudo_vm.vm_disk_resource", "internal_ip_address"),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "renewable_energy", "true"),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "vcpus", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm_disk_resource", "id", vmID),
				),
			},
		},
	})
}
