package dockerconfig

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestFilterMatchingRequests_MultipleMatches_Targets(t *testing.T) {
	is := is.New(t)

	// Given: Multiple requested objects with two that should match
	requested := []map[string]string{
		{"action": "write", "container_name": "image"},
		{"action": "read", "target": "cmd"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "read", "target": "cmd"},
	}

	// When: FilterMatchingRequests processes the objects
	result := filterMatchingRequests(requested, granted)

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

var internal_name = "confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd"
var environment_name = "local"
var environment_stage = "development"
var target = "cmd"
var umbrella_organization = "confetti-sites"
var umbrella_repository = "confetti-cms"
var container_name = "vendor/confetti-cms/image/container"

var locator = fmt.Sprintf("locator://%s?environment_name=%s&environment_stage=%s&target=%s&umbrella_organization=%s&umbrella_repository=%s&container_name=%s",
	internal_name, environment_name, environment_stage, target, umbrella_organization, umbrella_repository, container_name)

func TestFilterMatchingRequestsWithLocator_container_match(t *testing.T) {
	is := is.New(t)

	// Given: Multiple requested objects with one that should match
	requested := []map[string]string{
		{"container_name": "image"},
	}

	granted := []map[string]string{
		{"container_name": "image"},
	}

	// When: FilterMatchingRequests processes the objects
	result := CanSync(requested, granted, locator)

	// Then: Check that we got the expected match
	is.Equal(len(result), 1)
	is.Equal(result[0]["container_name"], "image")
}
