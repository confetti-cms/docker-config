package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func Test_ParseLocator_basic_example_environment_name(t *testing.T) {
	is := is.New(t)

	// Given: A locator string with environment_name field
	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_name=local"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Environment name should be parsed correctly
	is.Equal(result["environment_name"], "local")
}

func Test_ParseLocator_basic_example_environment_stage(t *testing.T) {
	is := is.New(t)

	// Given: A locator string with environment_stage field
	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?environment_stage=development"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Environment stage should be parsed correctly
	is.Equal(result["environment_stage"], "development")
}

func Test_ParseLocator_basic_example_target(t *testing.T) {
	is := is.New(t)

	// Given: A locator string with target field
	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?target=cmd"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Target should be parsed correctly
	is.Equal(result["target"], "cmd")
}

func Test_ParseLocator_basic_example_umbrella_organization(t *testing.T) {
	is := is.New(t)

	// Given: A locator string with umbrella_organization field
	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?umbrella_organization=confetti-sites"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Umbrella organization should be parsed correctly
	is.Equal(result["umbrella_organization"], "confetti-sites")
}

func Test_ParseLocator_basic_example_umbrella_repository(t *testing.T) {
	is := is.New(t)

	// Given: A locator string with umbrella_repository field
	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?umbrella_repository=confetti-cms"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Umbrella repository should be parsed correctly
	is.Equal(result["umbrella_repository"], "confetti-cms")
}

func Test_ParseLocator_basic_example_container_name(t *testing.T) {
	is := is.New(t)

	// Given: A locator string with container_name field
	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?container_name=image/container"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Container name should be parsed correctly
	is.Equal(result["container_name"], "image/container")
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

func Test_ParseLocator_complex_path_umbrella_organization(t *testing.T) {
	is := is.New(t)

	// Given: A locator with complex path structure focusing on umbrella organization
	locator := "locator://myorg-myrepo_prod_vendor-monitor_v1-stable-api?umbrella_organization=myorg"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Umbrella organization should be parsed correctly
	is.Equal(result["umbrella_organization"], "myorg")
}

func Test_ParseLocator_complex_path_umbrella_repository(t *testing.T) {
	is := is.New(t)

	// Given: A locator with complex path structure focusing on umbrella repository
	locator := "locator://myorg-myrepo_prod_vendor-monitor_v1-stable-api?umbrella_repository=myrepo"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Umbrella repository should be parsed correctly
	is.Equal(result["umbrella_repository"], "myrepo")
}

func Test_ParseLocator_complex_path_environment_name(t *testing.T) {
	is := is.New(t)

	// Given: A locator with complex path structure focusing on environment name
	locator := "locator://myorg-myrepo_prod_vendor-monitor_v1-stable-api?environment_name=prod"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Environment name should be parsed correctly
	is.Equal(result["environment_name"], "prod")
}

func Test_ParseLocator_complex_path_environment_stage(t *testing.T) {
	is := is.New(t)

	// Given: A locator with complex path structure focusing on environment stage
	locator := "locator://myorg-myrepo_prod_vendor-monitor_v1-stable-api?environment_stage=stable"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Environment stage should be parsed correctly
	is.Equal(result["environment_stage"], "stable")
}

func Test_ParseLocator_complex_path_target(t *testing.T) {
	is := is.New(t)

	// Given: A locator with complex path structure focusing on target
	locator := "locator://myorg-myrepo_prod_vendor-monitor_v1-stable-api?target=api"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Target should be parsed correctly
	is.Equal(result["target"], "api")
}
