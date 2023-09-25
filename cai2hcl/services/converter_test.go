package services

import (
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services/compute"
)

func TestConvertPanicsOnConverterNamesConflict(t *testing.T) {
	assertPanic(t, func() {
		joinStringMaps([]map[string]string{
			{"compute.googleapis.com/Instance": "compute_instance_1"},
			{"compute.googleapis.com/Instance": "compute_instance_2"},
		})
	})
}

func TestConvertPanicsOnConverterMapConflict(t *testing.T) {
	assertPanic(t, func() {
		joinConverterMaps([]map[string]common.Converter{
			common.CreateConverterMap(map[string]common.ConverterFactory{
				"google_compute_instance": compute.NewComputeInstanceConverter,
			}),
			common.CreateConverterMap(map[string]common.ConverterFactory{
				"google_compute_instance": compute.NewComputeForwardingRuleConverter,
			}),
		})
	})
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
