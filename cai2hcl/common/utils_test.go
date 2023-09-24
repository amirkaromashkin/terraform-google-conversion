package common

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tpg_provider "github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"
)

func TestSubsetOfFieldsMapsToCtyValue(t *testing.T) {
	schema := createSchema("google_compute_forwarding_rule")

	outputMap := map[string]interface{}{
		"name": "forwarding-rule-1",
	}

	val, err := MapToCtyValWithSchema(outputMap, schema)

	assert.Nil(t, err)
	assert.Equal(t, "forwarding-rule-1", val.GetAttr("name").AsString())
}

func TestMissingFieldDoesNotBreakConversion(t *testing.T) {
	provider := tpg_provider.Provider()
	var schema map[string]*schema.Schema = provider.ResourcesMap["google_compute_forwarding_rule"].Schema

	outputMap := map[string]interface{}{
		"unknownField": "unknownValue",
	}

	_, err := MapToCtyValWithSchema(outputMap, schema)

	assert.Nil(t, err)
}

func TestExperiment(t *testing.T) {
	provider := tpg_provider.Provider()
	resourceSchema := provider.ResourcesMap["google_compute_forwarding_rule"]

	outputMap := map[string]interface{}{
		"name":          "fr-1",
		"unknown_field": "unknownValue",
	}

	val, err := MapToCtyValWithSchema(outputMap, resourceSchema.Schema)

	assert.Nil(t, err)

	assert.True(t, val.Type().HasAttribute("name"))
	assert.Equal(t, "fr-1", val.GetAttr("name").AsString())

	assert.False(t, val.Type().HasAttribute("unknown_field"))
}

func createSchema(name string) map[string]*schema.Schema {
	provider := tpg_provider.Provider()

	return provider.ResourcesMap["google_compute_forwarding_rule"].Schema
}
