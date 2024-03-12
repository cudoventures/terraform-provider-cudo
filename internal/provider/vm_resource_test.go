package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
     image_id = "alpine-linux-317"
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
			if err == nil && getRes.VM.ShortState != "epil" {
				terminateParams := &vm.TerminateVMRequest{
					Id:        name,
					ProjectId: projectID,
				}
				res, err := cl.TerminateVM(ctx, terminateParams)
				t.Logf("(%s) %#v: %v", getRes.VM.ShortState, res, err)

				return fmt.Errorf("vm resource not destroyed %s , %s", getRes.VM.Id, getRes.VM.ShortState)
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
					resource.TestCheckResourceAttr("cudo_vm.vm", "boot_disk.image_id", "alpine-linux-317"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "boot_disk.size_gib", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.vm", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.vm", "memory_gib", "2"),
					resource.TestCheckResourceAttrSet("cudo_vm.vm", "price_hr"),
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
     image_id = "alpine-linux-317"
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

			if err == nil && ins.VM.ShortState != "epil" {
				terminateParams := &vm.TerminateVMRequest{
					Id:        name,
					ProjectId: projectID,
				}
				res, err := cl.TerminateVM(ctx, terminateParams)
				t.Logf("(%s) %#v: %v", ins.VM.ShortState, res, err)

				return fmt.Errorf("vm resource not destroyed %s, %s", ins.VM.Id, ins.VM.ShortState)
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
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "boot_disk.image_id", "alpine-linux-317"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "boot_disk.size_gib", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.vm-minimal", "memory_gib", "2"),
					resource.TestCheckResourceAttrSet("cudo_vm.vm-minimal", "price_hr"),
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
     image_id = "alpine-linux-317"
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

			if err == nil && ins.VM.ShortState != "epil" {
				terminateParams := &vm.TerminateVMRequest{
					Id:        name,
					ProjectId: projectID,
				}
				res, err := cl.TerminateVM(ctx, terminateParams)
				t.Logf("(%s) %#v: %v", ins.VM.ShortState, res, err)

				return fmt.Errorf("vm resource not destroyed %s, %s", ins.VM.Id, ins.VM.ShortState)
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
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "boot_disk.image_id", "alpine-linux-317"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "boot_disk.size_gib", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "memory_gib", "2"),
					resource.TestCheckResourceAttrSet("cudo_vm.vm-oob-delete", "price_hr"),
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
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "boot_disk.image_id", "alpine-linux-317"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "boot_disk.size_gib", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "memory_gib", "2"),
					resource.TestCheckResourceAttrSet("cudo_vm.vm-oob-delete", "price_hr"),
					resource.TestCheckResourceAttrSet("cudo_vm.vm-oob-delete", "internal_ip_address"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "renewable_energy", "true"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "vcpus", "1"),
					resource.TestCheckResourceAttr("cudo_vm.vm-oob-delete", "id", name),
				),
			},
		},
	})
}
