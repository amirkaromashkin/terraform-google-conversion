package common

import "regexp"

// A pair of asset name format and converter name.
// Used to identify converters not based on asset type, but based on the name (i.e. selfLink).
type RegexpNamePair struct {
	AssetNameRegexp string
	ConverterName   string
}

func FindMatchingConverterName(assetName string, regexpNamePairs []RegexpNamePair) (string, bool) {
	for _, pair := range regexpNamePairs {
		match, _ := regexp.MatchString(pair.AssetNameRegexp, assetName)
		if match {
			return pair.ConverterName, true
		}
	}

	return "", false
}
