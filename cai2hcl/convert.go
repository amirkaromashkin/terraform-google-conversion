package cai2hcl

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/caiasset"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/cai2hcl/common"
	"go.uber.org/zap"
)

// Struct for options so that adding new options does not
// require updating function signatures all along the pipe.
type Options struct {
	ErrorLogger *zap.Logger
}

var converters = GetConverters()
var converterMapByAssetType = getConverterMapByAssetType()

// Converts CAI Assets into HCL string.
func Convert(assets []*caiasset.Asset, options *Options) ([]byte, error) {
	if options == nil || options.ErrorLogger == nil {
		return nil, fmt.Errorf("logger is not initialized")
	}

	convertersMap := make(map[string]common.Converter)
	for _, converter := range converters {
		convertersMap[converter.GetTerraformResourceType()] = converter
	}

	// Group resources from the same tf resource type for convert.
	// tf -> cai has 1:N mappings occasionally
	groups := make(map[string][]*caiasset.Asset)
	for _, asset := range assets {
		for _, converter := range converterMapByAssetType[asset.Type] {
			name := converter.GetTerraformResourceType()

			groups[name] = append(groups[name], asset)
		}
	}

	allBlocks := []*common.HCLResourceBlock{}
	for name, assets := range groups {
		converter, ok := convertersMap[name]
		if !ok {
			continue
		}
		newBlocks, err := converter.Convert(assets)
		if err != nil {
			return nil, err
		}

		allBlocks = append(allBlocks, newBlocks...)
	}

	t, err := common.HclWriteBlocks(allBlocks)

	options.ErrorLogger.Debug(string(t))

	return t, err
}

func getConverterMapByAssetType() map[string][]common.Converter {
	var converterMapByAssetType = make(map[string][]common.Converter)
	for _, converter := range converters {
		for _, assetType := range converter.GetAssetTypes() {
			converterMapByAssetType[assetType] = append(converterMapByAssetType[assetType], converter)
		}
	}
	return converterMapByAssetType
}
