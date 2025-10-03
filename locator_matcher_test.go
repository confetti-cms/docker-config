package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

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
