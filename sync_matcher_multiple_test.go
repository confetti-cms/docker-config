package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func TestFilterMatchingRequests_MultipleMatches_Targets(t *testing.T) {
	is := is.New(t)

	// Given: Multiple requested objects with two that should match
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "write", "container_name": "image"},
		{"action": "read", "container_name": "video"},
		{"action": "read", "target": "cmd"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "read", "target": "cmd"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: Check that we got the expected target matches
	matchedTargets := make(map[string]bool)
	for _, match := range result {
		if target, exists := match["target"]; exists {
			matchedTargets[target] = true
		}
	}

	is.Equal(len(matchedTargets), 1)
	is.Equal(matchedTargets["cmd"], true)
}
