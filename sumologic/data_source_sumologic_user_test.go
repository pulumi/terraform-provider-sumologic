package sumologic

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourcSumologicUser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceAccSumologicUserConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceUserCheck("data.sumologic_user.by_id", "sumologic_user.test_user"),
					testAccDataSourceUserCheck("data.sumologic_user.by_email", "sumologic_user.test_user"),
				),
			},
		},
	})
}

func testAccDataSourceUserCheck(email, reference string) resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet(email, "id"),
		resource.TestCheckResourceAttrPair(email, "id", reference, "id"),
		resource.TestCheckResourceAttrPair(email, "email", reference, "email"),
		resource.TestCheckResourceAttrPair(email, "first_name", reference, "first_name"),
		resource.TestCheckResourceAttrPair(email, "last_name", reference, "last_name"),
		resource.TestCheckResourceAttrPair(email, "is_active", reference, "is_active"),
	)
}

var testDataSourceAccSumologicUserConfig = fmt.Sprintf(`
resource "sumologic_user" "test_user" {
  first_name = "Test"
  last_name = "User"
  email = "%s"
  is_active = "true"
  role_ids = ["${sumologic_role.test_role.id}"]
  transfer_to = ""
}

resource "sumologic_role" "test_role" {
	name = "Test_role_user_data"
	description = "My_SumoRoleDesc"
	filter_predicate = "_sourceCategory=Test"
	capabilities = ["viewCollectors"]
  }

data "sumologic_user" "by_email" {
  email = "${sumologic_user.test_user.email}"
}

data "sumologic_user" "by_id" {
  id = "${sumologic_user.test_user.id}"
}
`, FieldsMap["User"]["email"])
