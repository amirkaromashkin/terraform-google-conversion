package matchers

import (
	"github.com/amirkaromashkin/terraform-google-conversion/v5/caiasset"
)

type ConverterMatcher interface {
	GetConverterName() string

	Match(asset *caiasset.Asset) bool
}

func Matcher(regexp string) _AssetNameMatcher {
	return _AssetNameMatcher{Regexp: regexp}
}
