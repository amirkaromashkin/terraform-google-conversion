package matchers

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
)

func ByAssetName(converterName string, regexp string) ConverterMatcher {
	return _AssetNameMatcher{ConverterName: converterName, Regexp: regexp}
}

type _AssetNameMatcher struct {
	ConverterName string
	Regexp        string
}

func (inst _AssetNameMatcher) Match(asset *caiasset.Asset) bool {
	matched, _ := regexp.MatchString(inst.Regexp, asset.Name)
	return matched
}

func (inst _AssetNameMatcher) GetConverterName() string {
	return inst.ConverterName
}
