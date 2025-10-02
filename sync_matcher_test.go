package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func Test_container_name_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching container_name values
	requested := map[string]string{
		"container_name": "vendor/confetti-cms/image/container",
	}

	granted := map[string]string{
		"container_name": "vendor/confetti-cms/image/container",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_container_name_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different container_name values
	requested := map[string]string{
		"container_name": "vendor/confetti-cms/image/container",
	}

	granted := map[string]string{
		"container_name": "vendor/confetti-cms/different/container",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_container_name_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with both fields and granted object with "*" wildcard for container_name only
	requested := map[string]string{
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	granted := map[string]string{
		"container_name": "*",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Both fields should match (container_name via wildcard, target exactly)
	is.Equal(result, true)
}

func Test_target_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching target values
	requested := map[string]string{
		"target": "cmd",
	}

	granted := map[string]string{
		"target": "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_target_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different target values
	requested := map[string]string{
		"target": "cmd",
	}

	granted := map[string]string{
		"target": "all_up",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_target_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with both fields and granted object with "*" wildcard for target only
	requested := map[string]string{
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	granted := map[string]string{
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Both fields should match (container_name exactly, target via wildcard)
	is.Equal(result, true)
}

func Test_partial_match_fails(t *testing.T) {
	is := is.New(t)

	// Given: Container name matches but target doesn't
	requested := map[string]string{
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	granted := map[string]string{
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "all_up",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because target doesn't match
	is.Equal(result, false)
}

func Test_container_name_wildcard_but_target_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Container name matches via wildcard but target doesn't
	requested := map[string]string{
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	granted := map[string]string{
		"container_name": "*",
		"target":         "all_up",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because target doesn't match
	is.Equal(result, false)
}
