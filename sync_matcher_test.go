package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func Test_container_name_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different container_name values
	requested := map[string]string{
		"container_name": "image",
	}

	granted := map[string]string{
		"container_name": "video",
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
		"container_name": "image",
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
		"container_name": "image",
		"target":         "cmd",
	}

	granted := map[string]string{
		"container_name": "image",
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
		"container_name": "image",
		"target":         "cmd",
	}

	granted := map[string]string{
		"container_name": "image",
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
		"container_name": "image",
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
		"container_name": "image",
	}

	granted := map[string]string{
		"host":           "localhost",
		"container_name": "video",
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
		"container_name": "image",
		"target":         "cmd",
	}

	granted := map[string]string{
		"host":           "localhost",
		"container_name": "image",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: All fields should match
	is.Equal(result, true)
}

func Test_host_container_target_combined_match(t *testing.T) {
	is := is.New(t)

	// Given: Testing combination of host, container_name, and target fields
	requested := map[string]string{
		"host":           "localhost",
		"container_name": "image",
		"target":         "cmd",
	}

	granted := map[string]string{
		"host":           "localhost",
		"container_name": "image",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: All three fields should match together
	is.Equal(result, true)
}

func Test_host_wildcard_with_other_fields(t *testing.T) {
	is := is.New(t)

	// Given: Host matches via wildcard, other fields match exactly
	requested := map[string]string{
		"host":           "localhost",
		"container_name": "image",
		"target":         "cmd",
	}

	granted := map[string]string{
		"host":           "*",
		"container_name": "image",
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
		"container_name": "image",
	}

	granted := map[string]string{
		"schema":         "image",
		"container_name": "video",
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
		"container_name": "image",
		"target":         "cmd",
	}

	granted := map[string]string{
		"schema":         "image",
		"host":           "localhost",
		"container_name": "image",
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
		"container_name": "image",
		"target":         "cmd",
	}

	granted := map[string]string{
		"schema":         "*",
		"host":           "localhost",
		"container_name": "image",
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
		"action": "read",
	}

	granted := map[string]string{
		"action": "read",
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
		"action": "read",
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
		"action": "read",
	}

	granted := map[string]string{
		"action": "write",
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
		"action": "read",
		"schema": "image",
	}

	granted := map[string]string{
		"action": "read",
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
		"action":              "read",
	}

	granted := map[string]string{
		"source_organization": "myorg",
		"action":              "write",
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

func Test_environment_name_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching environment_name values
	requested := map[string]string{
		"environment_name": "production",
	}

	granted := map[string]string{
		"environment_name": "production",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_environment_name_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with environment_name field and granted object with "*" wildcard
	requested := map[string]string{
		"environment_name": "production",
	}

	granted := map[string]string{
		"environment_name": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The environment_name field should match via wildcard
	is.Equal(result, true)
}

func Test_environment_name_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different environment_name values
	requested := map[string]string{
		"environment_name": "production",
	}

	granted := map[string]string{
		"environment_name": "staging",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_environment_name_wildcard_but_empty_requested(t *testing.T) {
	is := is.New(t)

	// Given: Granted object with wildcard but requested has empty environment_name
	requested := map[string]string{
		"environment_name": "",
	}

	granted := map[string]string{
		"environment_name": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires non-empty value
	is.Equal(result, false)
}

func Test_environment_name_wildcard_but_missing_requested(t *testing.T) {
	is := is.New(t)

	// Given: Granted object with wildcard but requested is missing environment_name
	requested := map[string]string{}

	granted := map[string]string{
		"environment_name": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires non-empty value
	is.Equal(result, false)
}

func Test_environment_stage_exact_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with matching environment_stage values
	requested := map[string]string{
		"environment_stage": "stable",
	}

	granted := map[string]string{
		"environment_stage": "stable",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should match
	is.Equal(result, true)
}

func Test_environment_stage_wildcard_match(t *testing.T) {
	is := is.New(t)

	// Given: Requested object with environment_stage field and granted object with "*" wildcard
	requested := map[string]string{
		"environment_stage": "stable",
	}

	granted := map[string]string{
		"environment_stage": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The environment_stage field should match via wildcard
	is.Equal(result, true)
}

func Test_environment_stage_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two JSON objects with different environment_stage values
	requested := map[string]string{
		"environment_stage": "stable",
	}

	granted := map[string]string{
		"environment_stage": "beta",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: The objects should not match
	is.Equal(result, false)
}

func Test_environment_stage_wildcard_but_empty_requested(t *testing.T) {
	is := is.New(t)

	// Given: Granted object with wildcard but requested has empty environment_stage
	requested := map[string]string{
		"environment_stage": "",
	}

	granted := map[string]string{
		"environment_stage": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires non-empty value
	is.Equal(result, false)
}

func Test_environment_stage_wildcard_but_missing_requested(t *testing.T) {
	is := is.New(t)

	// Given: Granted object with wildcard but requested is missing environment_stage
	requested := map[string]string{}

	granted := map[string]string{
		"environment_stage": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires non-empty value
	is.Equal(result, false)
}

func Test_environment_fields_together(t *testing.T) {
	is := is.New(t)

	// Given: Both environment fields match exactly
	requested := map[string]string{
		"environment_name":  "production",
		"environment_stage": "stable",
	}

	granted := map[string]string{
		"environment_name":  "production",
		"environment_stage": "stable",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Both fields should match
	is.Equal(result, true)
}

func Test_environment_fields_mixed_wildcards(t *testing.T) {
	is := is.New(t)

	// Given: One environment field matches via wildcard, other exactly
	requested := map[string]string{
		"environment_name":  "production",
		"environment_stage": "stable",
	}

	granted := map[string]string{
		"environment_name":  "*",
		"environment_stage": "stable",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Both fields should match (name via wildcard, stage exactly)
	is.Equal(result, true)
}

func Test_environment_fields_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Environment name matches but stage doesn't
	requested := map[string]string{
		"environment_name":  "production",
		"environment_stage": "stable",
	}

	granted := map[string]string{
		"environment_name":  "production",
		"environment_stage": "beta",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because stage doesn't match
	is.Equal(result, false)
}

func Test_environment_fields_with_other_mismatches(t *testing.T) {
	is := is.New(t)

	// Given: Environment fields match but other fields don't
	requested := map[string]string{
		"environment_name":  "production",
		"environment_stage": "stable",
		"target":            "cmd",
	}

	granted := map[string]string{
		"environment_name":  "production",
		"environment_stage": "stable",
		"target":            "web",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because target doesn't match
	is.Equal(result, false)
}

func Test_exact_equality_match(t *testing.T) {
	is := is.New(t)

	// Given: Two identical objects
	requested := map[string]string{
		"container_name":    "image",
		"target":            "cmd",
		"environment_name":  "production",
		"environment_stage": "stable",
	}

	granted := map[string]string{
		"container_name":    "image",
		"target":            "cmd",
		"environment_name":  "production",
		"environment_stage": "stable",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match exactly (early return path)
	is.Equal(result, true)
}

func Test_exact_equality_no_match(t *testing.T) {
	is := is.New(t)

	// Given: Two different objects
	requested := map[string]string{
		"container_name":    "image",
		"target":            "cmd",
		"environment_name":  "production",
		"environment_stage": "stable",
	}

	granted := map[string]string{
		"container_name":    "video",
		"target":            "web",
		"environment_name":  "staging",
		"environment_stage": "beta",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match (different objects)
	is.Equal(result, false)
}

func Test_granted_missing_fields(t *testing.T) {
	is := is.New(t)

	// Given: Granted object missing fields (only subset matching)
	requested := map[string]string{
		"container_name":    "image",
		"target":            "cmd",
		"environment_name":  "production",
		"environment_stage": "stable",
	}

	granted := map[string]string{
		"target": "cmd",
		// only specifying target, other fields not restricted
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match because only specified fields need to match
	is.Equal(result, true)
}

func Test_mixed_scenario_wildcard_action(t *testing.T) {
	is := is.New(t)

	// Given: Testing wildcard action with other exact matches
	requested := map[string]string{
		"action": "read",
		"schema": "docker",
		"host":   "localhost",
	}

	granted := map[string]string{
		"action": "*",
		"schema": "docker",
		"host":   "localhost",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match with wildcard action and exact other fields
	is.Equal(result, true)
}

func Test_mixed_scenario_wildcard_source_repository(t *testing.T) {
	is := is.New(t)

	// Given: Testing wildcard source_repository with other exact matches
	requested := map[string]string{
		"source_organization":   "myorg",
		"source_repository":     "myrepo",
		"umbrella_organization": "parentorg",
	}

	granted := map[string]string{
		"source_organization":   "myorg",
		"source_repository":     "*",
		"umbrella_organization": "parentorg",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match with wildcard source_repository and exact other fields
	is.Equal(result, true)
}

func Test_mixed_scenario_wildcard_environment_stage(t *testing.T) {
	is := is.New(t)

	// Given: Testing wildcard environment_stage with other exact matches
	requested := map[string]string{
		"environment_name":  "production",
		"environment_stage": "stable",
		"target":            "cmd",
	}

	granted := map[string]string{
		"environment_name":  "production",
		"environment_stage": "*",
		"target":            "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match with wildcard environment_stage and exact other fields
	is.Equal(result, true)
}

func Test_empty_maps(t *testing.T) {
	is := is.New(t)

	// Given: Empty maps
	requested := map[string]string{}
	granted := map[string]string{}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match because empty maps are equal
	is.Equal(result, true)
}

func Test_empty_requested_with_granted_fields(t *testing.T) {
	is := is.New(t)

	// Given: Empty requested but granted has fields
	requested := map[string]string{}
	granted := map[string]string{
		"container_name": "image",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because requested doesn't have required fields
	is.Equal(result, false)
}

func Test_single_field_container_name_wildcard_non_empty(t *testing.T) {
	is := is.New(t)

	// Given: Single field with wildcard and non-empty value
	requested := map[string]string{
		"container_name": "image",
	}

	granted := map[string]string{
		"container_name": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match because wildcard accepts non-empty value
	is.Equal(result, true)
}

func Test_single_field_target_wildcard_non_empty(t *testing.T) {
	is := is.New(t)

	// Given: Single field with wildcard and non-empty value
	requested := map[string]string{
		"target": "cmd",
	}

	granted := map[string]string{
		"target": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match because wildcard accepts non-empty value
	is.Equal(result, true)
}

func Test_wildcard_requires_non_empty_but_empty_provided(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but empty string in requested
	requested := map[string]string{
		"container_name": "",
		"target":         "cmd",
	}

	granted := map[string]string{
		"container_name": "*",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires non-empty value
	is.Equal(result, false)
}

func Test_exact_match_with_empty_string(t *testing.T) {
	is := is.New(t)

	// Given: Exact match with empty strings
	requested := map[string]string{
		"container_name": "",
		"target":         "cmd",
	}

	granted := map[string]string{
		"container_name": "",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match because empty strings match exactly
	is.Equal(result, true)
}

func Test_wildcard_vs_exact_mismatch(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but different non-empty value in requested
	requested := map[string]string{
		"container_name": "different_image",
		"target":         "cmd",
	}

	granted := map[string]string{
		"container_name": "*",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should match because wildcard accepts any non-empty value
	is.Equal(result, true)
}

func Test_wildcard_missing_field_container_name(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"target": "cmd",
		// container_name is missing
	}

	granted := map[string]string{
		"container_name": "*",
		"target":         "cmd",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_target(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// target is missing
	}

	granted := map[string]string{
		"container_name": "image",
		"target":         "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_host(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// host is missing
	}

	granted := map[string]string{
		"container_name": "image",
		"host":           "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_schema(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// schema is missing
	}

	granted := map[string]string{
		"container_name": "image",
		"schema":         "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_action(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// action is missing
	}

	granted := map[string]string{
		"container_name": "image",
		"action":         "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_source_organization(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// source_organization is missing
	}

	granted := map[string]string{
		"container_name":      "image",
		"source_organization": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_source_repository(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// source_repository is missing
	}

	granted := map[string]string{
		"container_name":    "image",
		"source_repository": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_umbrella_organization(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// umbrella_organization is missing
	}

	granted := map[string]string{
		"container_name":        "image",
		"umbrella_organization": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_umbrella_repository(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// umbrella_repository is missing
	}

	granted := map[string]string{
		"container_name":      "image",
		"umbrella_repository": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_environment_name(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// environment_name is missing
	}

	granted := map[string]string{
		"container_name":   "image",
		"environment_name": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}

func Test_wildcard_missing_field_environment_stage(t *testing.T) {
	is := is.New(t)

	// Given: Wildcard in granted but field missing in requested
	requested := map[string]string{
		"container_name": "image",
		// environment_stage is missing
	}

	granted := map[string]string{
		"container_name":    "image",
		"environment_stage": "*",
	}

	// When: SyncMatcher compares the objects
	result := SyncMatcher(requested, granted)

	// Then: Should not match because wildcard requires the field to exist and be non-empty
	is.Equal(result, false)
}
