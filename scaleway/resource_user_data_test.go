package scaleway

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccScalewayUserData_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckScalewayUserDataDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckScalewayUserDataConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckScalewayServerExists("scaleway_server.base"),
					testAccCheckScalewayUserDataExists("scaleway_user_data.base"),
					resource.TestCheckResourceAttr("scaleway_user_data.base", "value", "supersecret"),
					resource.TestCheckResourceAttr("scaleway_user_data.base", "key", "gcp_username"),
				),
			},
		},
	})
}

func testAccCheckScalewayUserDataExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		client := testAccProvider.Meta().(*Client).scaleway
		_, err := client.GetUserdata(rs.Primary.Attributes["server"], rs.Primary.Attributes["key"], false)

		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckScalewayUserDataDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client).scaleway

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "scaleway" {
			continue
		}

		_, err := client.GetUserdata(rs.Primary.Attributes["server"], rs.Primary.Attributes["key"], false)

		if err == nil {
			return fmt.Errorf("UserData still exists")
		}
	}

	return nil
}

var testAccCheckScalewayUserDataConfig = fmt.Sprintf(`
resource "scaleway_server" "base" {
  name = "test"
  # ubuntu 14.04
  image = "%s"
  type = "C1"

  tags = [ "terraform-test", "user-data" ]
}

resource "scaleway_user_data" "base" {
	server = "${scaleway_server.base.id}"
	key = "gcp_username"
	value = "supersecret"
}
`, armImageIdentifier)
