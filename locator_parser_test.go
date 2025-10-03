package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func Test_ParseLocator_basic_example_container_name(t *testing.T) {
	is := is.New(t)

	// Given: A locator string with container_name field
	locator := "locator://confetti-sites-confetti-cms_local_vendor-confetti-cms-monitor_8609-development-cmd?container_name=image/container"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Container name should be parsed correctly
	is.Equal(result["container_name"], "image/container")
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

func Test_ParseLocator_internal_name(t *testing.T) {
	is := is.New(t)

	// Given: A locator string with internal_name field
	locator := "locator://confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd"

	// When: ParseLocator parses the locator
	result := ParseLocator(locator)

	// Then: Internal name should be parsed correctly
	is.Equal(result["internal_name"], "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
}
