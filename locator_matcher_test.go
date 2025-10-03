package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

// func Test_LocatorMatcher_exact_match(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator string and a granted object that should match exactly
// 	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&container_name=image/container"

// 	granted := map[string]string{
// 		"environment_name":      "local",
// 		"environment_stage":     "development",
// 		"target":                "cmd",
// 		"umbrella_organization": "confetti-sites",
// 		"umbrella_repository":   "confetti-cms",
// 		"container_name":        "image/container",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: The objects should match
// 	is.Equal(result, true)
// }

// func Test_LocatorMatcher_wildcard_match(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator string and a granted object with wildcard for some fields
// 	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&container_name=image/container"

// 	granted := map[string]string{
// 		"environment_name":      "*",
// 		"environment_stage":     "development",
// 		"target":                "*",
// 		"umbrella_organization": "confetti-sites",
// 		"umbrella_repository":   "confetti-cms",
// 		"container_name":        "image/container",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: The objects should match via wildcards
// 	is.Equal(result, true)
// }

// func Test_LocatorMatcher_no_match(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator string and a granted object that should not match
// 	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&container_name=image/container"

// 	granted := map[string]string{
// 		"environment_name":      "production",
// 		"environment_stage":     "development",
// 		"target":                "cmd",
// 		"umbrella_organization": "confetti-sites",
// 		"umbrella_repository":   "confetti-cms",
// 		"container_name":        "image/container",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: The objects should not match
// 	is.Equal(result, false)
// }

// func Test_LocatorMatcher_no_match_environment_mismatch(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator string and a granted object that should not match due to environment_name mismatch
// 	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&container_name=image/container"

// 	granted := map[string]string{
// 		"environment_name":      "production", // Different from locator's "local"
// 		"environment_stage":     "development",
// 		"target":                "cmd",
// 		"umbrella_organization": "confetti-sites",
// 		"umbrella_repository":   "confetti-cms",
// 		"container_name":        "image/container",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: The objects should not match due to environment_name mismatch
// 	is.Equal(result, false)
// }

// func Test_LocatorMatcher_debug_parsing(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator string for debugging parsing behavior
// 	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&container_name=image/container"

// 	granted := map[string]string{
// 		"environment_name":      "production",
// 		"environment_stage":     "development",
// 		"target":                "cmd",
// 		"umbrella_organization": "confetti-sites",
// 		"umbrella_repository":   "confetti-cms",
// 		"container_name":        "image/container",
// 	}

// 	// Debug: Check what the locator parses to and test SyncMatcher directly
// 	parsedLocator := ParseLocator(locator)
// 	t.Logf("Parsed locator: %+v", parsedLocator)
// 	t.Logf("Granted object: %+v", granted)

// 	// Test SyncMatcher directly
// 	syncResult := SyncMatcher(parsedLocator, granted)
// 	t.Logf("SyncMatcher result: %v", syncResult)

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: The objects should not match
// 	is.Equal(result, false)
// }

// func Test_LocatorMatcher_partial_match_fails(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator string that matches in some fields but not others
// 	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&container_name=image/container"

// 	granted := map[string]string{
// 		"environment_name":      "local",
// 		"environment_stage":     "development",
// 		"target":                "web", // Different target
// 		"umbrella_organization": "confetti-sites",
// 		"umbrella_repository":   "confetti-cms",
// 		"container_name":        "image/container",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: The objects should not match due to target mismatch
// 	is.Equal(result, false)
// }

// func Test_LocatorMatcher_empty_granted(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator string and an empty granted object
// 	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&container_name=image/container"

// 	granted := map[string]string{}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: Should match because empty granted means no restrictions
// 	is.Equal(result, true)
// }

// func Test_LocatorMatcher_invalid_locator(t *testing.T) {
// 	is := is.New(t)

// 	// Given: An invalid locator string and a granted object
// 	locator := "invalid-locator-string"

// 	granted := map[string]string{
// 		"environment_name": "local",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: Should not match because locator parsing will return empty map
// 	is.Equal(result, false)
// }

// func Test_LocatorMatcher_query_parameters_override(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator where query parameters should override path parsing
// 	locator := "locator://different-org-different-repo_env-stage-target?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms"

// 	granted := map[string]string{
// 		"environment_name":      "local",
// 		"environment_stage":     "development",
// 		"target":                "cmd",
// 		"umbrella_organization": "confetti-sites",
// 		"umbrella_repository":   "confetti-cms",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: Should match using query parameter values (which override path parsing)
// 	is.Equal(result, true)
// }

// func Test_LocatorMatcher_query_parameters_override_organization(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator where query parameters override path parsing for organization fields
// 	locator := "locator://different-org-different-repo_env-stage-target?umbrella_organization=confetti-sites&umbrella_repository=confetti-cms"

// 	granted := map[string]string{
// 		"umbrella_organization": "confetti-sites",
// 		"umbrella_repository":   "confetti-cms",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: Should match using query parameter values for organization fields
// 	is.Equal(result, true)
// }

// func Test_LocatorMatcher_query_parameters_override_environment(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator where query parameters override path parsing for environment fields
// 	locator := "locator://different-org-different-repo_env-stage-target?environment_name=local&environment_stage=development"

// 	granted := map[string]string{
// 		"environment_name":  "local",
// 		"environment_stage": "development",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: Should match using query parameter values for environment fields
// 	is.Equal(result, true)
// }

func Test_LocatorMatcher_query_parameters_override_target(t *testing.T) {
	is := is.New(t)

	// Given: A locator where query parameters override path parsing for target field
	locator := "locator://different-org-different-repo_env-stage-target?target=cmd"

	granted := map[string]string{
		"target": "cmd",
	}

	// When: LocatorMatcher compares the locator with granted object
	result := LocatorMatcher(locator, granted)

	// Then: Should match using query parameter value for target field
	is.Equal(result, true)
}

// func Test_LocatorMatcher_subset_match(t *testing.T) {
// 	is := is.New(t)

// 	// Given: A locator string and a granted object that only specifies some fields
// 	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&container_name=image/container"

// 	granted := map[string]string{
// 		"target":                "cmd", // Only specifying target
// 		"umbrella_organization": "confetti-sites",
// 	}

// 	// When: LocatorMatcher compares the locator with granted object
// 	result := LocatorMatcher(locator, granted)

// 	// Then: Should match because only the specified fields need to match
// 	is.Equal(result, true)
// }
