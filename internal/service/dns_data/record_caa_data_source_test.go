package dns_data_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dnsdata"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccRecordCAADataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_caa_records.test"
	resourceName := "bloxone_dns_caa_record.test"
	var v dnsdata.Record
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	niz := acctest.RandomNameWithPrefix("caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordCAADataSourceConfigFilters(zoneFqdn, niz),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccRecordCAADataSourceConfigFilters(zoneFqdn, nameInZone string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_caa_record" "test" {
  name_in_zone = %[1]q
  zone = bloxone_dns_auth_zone.test.id
    rdata = {
        tag = "issue"
        value = "ca.example.com"
	}
}

data "bloxone_dns_caa_records" "test" {
  filters = {
    name_in_zone = %[1]q
    zone = bloxone_dns_auth_zone.test.id
  }
  depends_on = [bloxone_dns_caa_record.test]
}`, nameInZone)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}
