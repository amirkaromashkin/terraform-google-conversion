package cai2hcl

import (
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services"
	cai2hclTesting "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/testing"
)

func TestConvertCompute(t *testing.T) {
	cai2hclTesting.AssertTestFiles(
		t,
		services.Converter,
		"./services/compute/testdata",
		[]string{
			"full_compute_instance",
		})
}

func TestConvertResourcemanager(t *testing.T) {
	cai2hclTesting.AssertTestFiles(
		t,
		services.Converter,
		"./services/resourcemanager/testdata",
		[]string{
			"project_create",
		})
}
