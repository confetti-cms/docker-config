package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func Test_ParseLocator_basic_example(t *testing.T) {
	is := is.New(t)

	// Given: A locator string with the expected format
	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&container_name=image/container"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Debug output
	t.Logf("Parsed result: %+v", result)

	// Then: The result should contain the expected fields (query parameters override path parsing)
	is.Equal(result["environment_name"], "local")
	is.Equal(result["environment_stage"], "development")
	is.Equal(result["target"], "cmd")
	is.Equal(result["umbrella_organization"], "confetti-sites")
	is.Equal(result["umbrella_repository"], "confetti-cms")
	is.Equal(result["container_name"], "image/container")
}

func Test_ParseLocator_with_query_parameters(t *testing.T) {
	is := is.New(t)

	// Given: A locator with query parameters
	locator := "locator://test-path?environment_name=test&container_name=test-container&custom_field=custom-value"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Should extract both path and query parameters
	is.Equal(result["environment_name"], "test")
	is.Equal(result["container_name"], "test-container")
	is.Equal(result["custom_field"], "custom-value")
}

func Test_ParseLocator_no_scheme(t *testing.T) {
	is := is.New(t)

	// Given: A string without the locator:// scheme
	locator := "confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Should return empty map for invalid scheme
	is.Equal(len(result), 0)
}

func Test_ParseLocator_empty_string(t *testing.T) {
	is := is.New(t)

	// Given: An empty string
	locator := ""

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Should return empty map
	is.Equal(len(result), 0)
}

func Test_ParseLocator_query_only(t *testing.T) {
	is := is.New(t)

	// Given: A locator with only query parameters
	locator := "locator://?environment_name=test&target=cmd"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Should extract query parameters but no path fields
	is.Equal(result["environment_name"], "test")
	is.Equal(result["target"], "cmd")
	is.Equal(result["umbrella_organization"], "")
}

func Test_ParseLocator_complex_path(t *testing.T) {
	is := is.New(t)

	// Given: A locator with complex path structure
	locator := "locator://myorg-myrepo_prod_vendor-monitor_v1-stable-api?environment_name=prod&environment_stage=stable&target=api&umbrella_organization=myorg&umbrella_repository=myrepo"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Should correctly parse all components
	is.Equal(result["umbrella_organization"], "myorg")
	is.Equal(result["umbrella_repository"], "myrepo")
	is.Equal(result["environment_name"], "prod")
	is.Equal(result["environment_stage"], "stable")
	is.Equal(result["target"], "api")
	// Query parameters should override path parsing for the same fields
	is.Equal(result["umbrella_organization"], "myorg") // from query
	is.Equal(result["umbrella_repository"], "myrepo")  // from query
}

func Test_ParseLocator_debug(t *testing.T) {
	is := is.New(t)

	// Given: Simple test case to debug parsing
	locator := "locator://test-path?environment_name=test&container_name=test-container"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Debug output
	t.Logf("Parsed result: %+v", result)

	// Then: Should extract query parameters correctly
	is.Equal(result["environment_name"], "test")
	is.Equal(result["container_name"], "test-container")
}
