package provider

import (
	"cudo.org/v1/terraform-provider-cudo/internal/helper"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccExampleResource(t *testing.T) {

	name, err := helper.NewNanoID(12)

	if err != nil {
		return
	}

	vmConfig := fmt.Sprintf(`
resource "cudo_vm" "test-vm" {
   config_id          = "oaml6hca4fb0"
   vcpu_quantity      = 1
   boot_disk_size_gib = 50
   image_id           = "ubuntu-minimal-2004"
   memory_gib         = 4
   vm_id              = "%s"
   boot_disk_class    = "network"
 }`, name)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: getProviderConfig() + vmConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "boot_disk_class", "network"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "boot_disk_size_gib", "50"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "config_id", "oaml6hca4fb0"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "cpu_class", ""),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "cpu_model", "Intel(R) Xeon(R) CPU E5-2690 v3 @ 2.60GHz"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "datacenter_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "gpu_mem", "0"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "image_desc", "Ubuntu 20.04 LTS (focal) Minimal"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "image_id", "ubuntu-minimal-2004"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "image_name", "Ubuntu Minimal 20.04"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "memory_gib", "4"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "one_state", "INIT"),
					resource.TestCheckResourceAttrSet("cudo_vm.test-vm", "price_hr"),
					resource.TestCheckResourceAttrSet("cudo_vm.test-vm", "public_ip_address"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "region_id", "gb-bournemouth"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "region_name", "Bournemouth, United Kingdom"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "renewable_energy", "true"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "vcpu_quantity", "1"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "vm_id", name),
				),
			},
		},
	})
}