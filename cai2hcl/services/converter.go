package services

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common/matchers"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services/resourcemanager"
)

var uberConverters = []common.UberConverter{
	compute.UberConverter,
	resourcemanager.UberConverter,
}

var UberConverter common.UberConverter

func init() {
	var converterByAssetType map[string]string
	var converterMatchersByAssetType map[string][]matchers.ConverterMatcher
	var converters map[string]common.Converter

	for _, uberConverter := range uberConverters {
		converterByAssetType = joinMaps(converterByAssetType, uberConverter.ConverterByAssetType)
		converterMatchersByAssetType = joinMaps(converterMatchersByAssetType, uberConverter.ConverterMatchersByAssetType)
		converters = joinMaps(converters, uberConverter.Converters)
	}

	UberConverter = common.UberConverter{
		ConverterByAssetType:         converterByAssetType,
		ConverterMatchersByAssetType: converterMatchersByAssetType,
		Converters:                   converters,
	}
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
