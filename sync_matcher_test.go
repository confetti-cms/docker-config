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

func Test_host_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching host values
	requested := map[string]string{
		"host": "localhost",
	}

	granted := map[string]string{
		"host": "localhost",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_host_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with host field and granted object with "*" wildcard for host
	requested := map[string]string{
		"host": "localhost",
	}

	granted := map[string]string{
		"host": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The host field should match via wildcard
	is.Equal(result, true)
}

func Test_host_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different host values
	requested := map[string]string{
		"host": "localhost",
	}

	granted := map[string]string{
		"host": "remotehost",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_host_matches_but_container_name_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Host matches but container_name doesn't
	requested := map[string]string{
		"host":           "localhost",
		"container_name": "vendor/confetti-cms/image/container",
	}

	granted := map[string]string{
		"host":           "localhost",
		"container_name": "vendor/confetti-cms/different/container",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because container_name doesn't match
	is.Equal(result, false)
}

func Test_host_matches_but_target_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Host matches but target doesn't
	requested := map[string]string{
		"host":   "localhost",
		"target": "cmd",
	}

	granted := map[string]string{
		"host":   "localhost",
		"target": "all_up",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because target doesn't match
	is.Equal(result, false)
}

func Test_all_three_fields_match(t *testing.T) {
	is := is.New(t)

	// Given: All three fields match exactly
	requested := map[string]string{
		"host":           "localhost",
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	granted := map[string]string{
		"host":           "localhost",
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: All fields should match
	is.Equal(result, true)
}

func Test_host_wildcard_with_other_fields(t *testing.T) {
	is := is.New(t)

	// Given: Host matches via wildcard, other fields match exactly
	requested := map[string]string{
		"host":           "localhost",
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	granted := map[string]string{
		"host":           "*",
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: All fields should match (host via wildcard, others exactly)
	is.Equal(result, true)
}

func Test_schema_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching schema values
	requested := map[string]string{
		"schema": "image",
	}

	granted := map[string]string{
		"schema": "image",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_schema_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with schema field and granted object with "*" wildcard for schema
	requested := map[string]string{
		"schema": "image",
	}

	granted := map[string]string{
		"schema": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The schema field should match via wildcard
	is.Equal(result, true)
}

func Test_schema_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different schema values
	requested := map[string]string{
		"schema": "image",
	}

	granted := map[string]string{
		"schema": "hive",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_schema_matches_but_host_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Schema matches but host doesn't
	requested := map[string]string{
		"schema": "image",
		"host":   "localhost",
	}

	granted := map[string]string{
		"schema": "image",
		"host":   "remotehost",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because host doesn't match
	is.Equal(result, false)
}

func Test_schema_matches_but_container_name_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Schema matches but container_name doesn't
	requested := map[string]string{
		"schema":         "image",
		"container_name": "vendor/confetti-cms/image/container",
	}

	granted := map[string]string{
		"schema":         "image",
		"container_name": "vendor/confetti-cms/different/container",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because container_name doesn't match
	is.Equal(result, false)
}

func Test_all_four_fields_match(t *testing.T) {
	is := is.New(t)

	// Given: All four fields match exactly
	requested := map[string]string{
		"schema":         "image",
		"host":           "localhost",
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	granted := map[string]string{
		"schema":         "image",
		"host":           "localhost",
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: All fields should match
	is.Equal(result, true)
}

func Test_schema_wildcard_with_other_fields(t *testing.T) {
	is := is.New(t)

	// Given: Schema matches via wildcard, other fields match exactly
	requested := map[string]string{
		"schema":         "image",
		"host":           "localhost",
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	granted := map[string]string{
		"schema":         "*",
		"host":           "localhost",
		"container_name": "vendor/confetti-cms/image/container",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: All fields should match (schema via wildcard, others exactly)
	is.Equal(result, true)
}

func Test_action_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching action values
	requested := map[string]string{
		"action": "push",
	}

	granted := map[string]string{
		"action": "push",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_action_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with action field and granted object with "*" wildcard for action
	requested := map[string]string{
		"action": "push",
	}

	granted := map[string]string{
		"action": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The action field should match via wildcard
	is.Equal(result, true)
}

func Test_action_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different action values
	requested := map[string]string{
		"action": "push",
	}

	granted := map[string]string{
		"action": "pull",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_action_matches_but_schema_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Action matches but schema doesn't
	requested := map[string]string{
		"action": "push",
		"schema": "image",
	}

	granted := map[string]string{
		"action": "push",
		"schema": "hive",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because schema doesn't match
	is.Equal(result, false)
}

func Test_source_organization_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching source_organization values
	requested := map[string]string{
		"source_organization": "myorg",
	}

	granted := map[string]string{
		"source_organization": "myorg",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_source_organization_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with source_organization field and granted object with "*" wildcard
	requested := map[string]string{
		"source_organization": "myorg",
	}

	granted := map[string]string{
		"source_organization": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The source_organization field should match via wildcard
	is.Equal(result, true)
}

func Test_source_organization_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different source_organization values
	requested := map[string]string{
		"source_organization": "myorg",
	}

	granted := map[string]string{
		"source_organization": "otherorg",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_source_organization_matches_but_action_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Source organization matches but action doesn't
	requested := map[string]string{
		"source_organization": "myorg",
		"action":              "push",
	}

	granted := map[string]string{
		"source_organization": "myorg",
		"action":              "pull",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because action doesn't match
	is.Equal(result, false)
}

func Test_source_repository_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching source_repository values
	requested := map[string]string{
		"source_repository": "myrepo",
	}

	granted := map[string]string{
		"source_repository": "myrepo",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_source_repository_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with source_repository field and granted object with "*" wildcard
	requested := map[string]string{
		"source_repository": "myrepo",
	}

	granted := map[string]string{
		"source_repository": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The source_repository field should match via wildcard
	is.Equal(result, true)
}

func Test_source_repository_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different source_repository values
	requested := map[string]string{
		"source_repository": "myrepo",
	}

	granted := map[string]string{
		"source_repository": "otherrepo",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_source_repository_matches_but_source_org_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Source repository matches but source organization doesn't
	requested := map[string]string{
		"source_repository":   "myrepo",
		"source_organization": "myorg",
	}

	granted := map[string]string{
		"source_repository":   "myrepo",
		"source_organization": "otherorg",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because source organization doesn't match
	is.Equal(result, false)
}

func Test_umbrella_organization_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching umbrella_organization values
	requested := map[string]string{
		"umbrella_organization": "parentorg",
	}

	granted := map[string]string{
		"umbrella_organization": "parentorg",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_umbrella_organization_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with umbrella_organization field and granted object with "*" wildcard
	requested := map[string]string{
		"umbrella_organization": "parentorg",
	}

	granted := map[string]string{
		"umbrella_organization": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The umbrella_organization field should match via wildcard
	is.Equal(result, true)
}

func Test_umbrella_organization_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different umbrella_organization values
	requested := map[string]string{
		"umbrella_organization": "parentorg",
	}

	granted := map[string]string{
		"umbrella_organization": "differentorg",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_umbrella_organization_matches_but_source_repo_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Umbrella organization matches but source repository doesn't
	requested := map[string]string{
		"umbrella_organization": "parentorg",
		"source_repository":     "myrepo",
	}

	granted := map[string]string{
		"umbrella_organization": "parentorg",
		"source_repository":     "otherrepo",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because source repository doesn't match
	is.Equal(result, false)
}

func Test_umbrella_repository_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching umbrella_repository values
	requested := map[string]string{
		"umbrella_repository": "parentrepo",
	}

	granted := map[string]string{
		"umbrella_repository": "parentrepo",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_umbrella_repository_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with umbrella_repository field and granted object with "*" wildcard
	requested := map[string]string{
		"umbrella_repository": "parentrepo",
	}

	granted := map[string]string{
		"umbrella_repository": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The umbrella_repository field should match via wildcard
	is.Equal(result, true)
}

func Test_umbrella_repository_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different umbrella_repository values
	requested := map[string]string{
		"umbrella_repository": "parentrepo",
	}

	granted := map[string]string{
		"umbrella_repository": "differentrepo",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_umbrella_repository_matches_but_umbrella_org_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Umbrella repository matches but umbrella organization doesn't
	requested := map[string]string{
		"umbrella_repository":   "parentrepo",
		"umbrella_organization": "parentorg",
	}

	granted := map[string]string{
		"umbrella_repository":   "parentrepo",
		"umbrella_organization": "differentorg",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because umbrella organization doesn't match
	is.Equal(result, false)
}

func Test_all_nine_fields_match(t *testing.T) {
	is := is.New(t)

	// Given: All nine fields match exactly
	requested := map[string]string{
		"schema":                "image",
		"host":                  "localhost",
		"container_name":        "vendor/confetti-cms/image/container",
		"target":                "cmd",
		"action":                "push",
		"source_organization":   "myorg",
		"source_repository":     "myrepo",
		"umbrella_organization": "parentorg",
		"umbrella_repository":   "parentrepo",
	}

	granted := map[string]string{
		"schema":                "image",
		"host":                  "localhost",
		"container_name":        "vendor/confetti-cms/image/container",
		"target":                "cmd",
		"action":                "push",
		"source_organization":   "myorg",
		"source_repository":     "myrepo",
		"umbrella_organization": "parentorg",
		"umbrella_repository":   "parentrepo",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: All fields should match
	is.Equal(result, true)
}
