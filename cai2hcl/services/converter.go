package services

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services/resourcemanager"
)

var UberConverter = common.UberConverter{
	ConverterMatchersByAssetType: joinMaps(
		compute.Converter.ConverterMatchersByAssetType,
		resourcemanager.Converter.ConverterMatchersByAssetType),
	ConverterByAssetType: joinMaps(
		compute.Converter.ConverterByAssetType,
		resourcemanager.Converter.ConverterByAssetType),
	Converters: joinMaps(
		compute.Converter.Converters,
		resourcemanager.Converter.Converters),
}

func joinMaps[V interface{}](arr ...map[string]V) map[string]V {
	result := make(map[string]V)

	for _, m := range arr {
		for key, value := range m {
			if _, hasKey := result[key]; hasKey {
				panic(fmt.Sprintf("Map keys are not unique. Duplicate: %s", key))
			}

			result[key] = value
		}
	}

	return result
}
