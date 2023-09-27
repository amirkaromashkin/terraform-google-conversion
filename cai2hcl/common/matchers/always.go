package matchers

import "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

func Always(converterName string) ConverterMatcher {
	return _TrueMatcher{ConverterName: converterName}
}

type _TrueMatcher struct {
	ConverterName string
}

func (inst _TrueMatcher) GetConverterName() string {
	return inst.ConverterName
}

func (inst _TrueMatcher) Match(asset *caiasset.Asset) bool {
	return true
}
