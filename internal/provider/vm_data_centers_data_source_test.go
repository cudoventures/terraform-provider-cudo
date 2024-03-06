package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAcc_RegionsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: getProviderConfig() + testAccRegionsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_vm_data_centers.test", "data_centers.#", "1"),
					resource.TestCheckResourceAttr("data.cudo_vm_data_centers.test", "data_centers.0.id", "black-mesa"),
					resource.TestCheckResourceAttr("data.cudo_vm_data_centers.test", "data_centers.0.region_id", "gb-bournemouth"),
				),
			},
		},
	})
}

const testAccRegionsDataSourceConfig = `
data "cudo_vm_data_centers" "test" {
}`
