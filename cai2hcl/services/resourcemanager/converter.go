package resourcemanager

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
)

var Converter = common.UberConverter{
	ConverterNamesPerAssetType: map[string]string{
		ProjectAssetType:        "google_project",
		ProjectBillingAssetType: "google_project",
	},
	ConverterMap: common.CreateConverterMap(map[string]common.ConverterFactory{
		"google_project": NewProjectConverter,
	}),
}
