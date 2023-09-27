package services

import (
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/stretchr/testify/assert"
)

func TestJoinMaps(t *testing.T) {
	joinedMap := joinMaps(
		map[string]string{"1": "a"},
		map[string]string{"2": "b"})

	assert.Equal(t, joinedMap, map[string]string{"1": "a", "2": "b"})
}

func TestJoinMapsConflict(t *testing.T) {
	assert.Panics(t, func() {
		joinMaps(
			map[string][]common.RegexpNamePair{"compute.googleapis.com/Instance": {{ConverterName: "compute_instance_1"}}},
			map[string][]common.RegexpNamePair{"compute.googleapis.com/Instance": {{ConverterName: "compute_instance_2"}}})
	})
}
