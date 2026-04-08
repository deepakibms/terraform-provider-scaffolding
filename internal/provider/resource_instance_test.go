// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceInstance(t *testing.T) {
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInstance,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"scaffolding_instance.test", "name", regexp.MustCompile("^test")),
					resource.TestCheckResourceAttr(
						"scaffolding_instance.test", "instance_type", "t2.micro"),
				),
			},
		},
	})
}

const testAccResourceInstance = `
resource "scaffolding_instance" "test" {
  name          = "test-instance"
  instance_type = "t2.micro"
  tags = {
    Environment = "test"
    Owner       = "terraform"
  }
}
`

// Made with Bob
