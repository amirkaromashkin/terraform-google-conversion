package common

// Configuration object to map CAI Assets to Converters.
type ConverterMap struct {
	// Mapping from CAI Asset Type to converter name.
	AssetTypeToConverterName map[string]string

	// Mapping from converter name and converter.
	ConverterNameToConverter map[string]Converter
}
