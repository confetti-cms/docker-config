package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func TestFilterMatchingRequests_SingleMatch_Count(t *testing.T) {
	is := is.New(t)

	// Given: Multiple requested objects and one matching granted object
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "write", "container_name": "image"},
		{"action": "read", "container_name": "video"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "image"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: Only the first requested object should match
	is.Equal(len(result), 1)
}

func TestFilterMatchingRequests_SingleMatch_FirstResult(t *testing.T) {
	is := is.New(t)

	// Given: Multiple requested objects and one matching granted object
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "write", "container_name": "image"},
		{"action": "read", "container_name": "video"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "image"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: The first result should have correct action and container
	is.Equal(result[0]["action"], "read")
	is.Equal(result[0]["container_name"], "image")
}

func TestFilterMatchingRequests_MultipleMatches_Count(t *testing.T) {
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

	// Then: Two requested objects should match
	is.Equal(len(result), 2)
}

func TestFilterMatchingRequests_MultipleMatches_Actions(t *testing.T) {
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

	// Then: Check that we got the expected matches (order may vary)
	matchedActions := make(map[string]bool)
	for _, match := range result {
		if action, exists := match["action"]; exists {
			matchedActions[action] = true
		}
	}

	is.Equal(len(matchedActions), 1)
	is.Equal(matchedActions["read"], true)
}

func TestFilterMatchingRequests_MultipleMatches_Containers(t *testing.T) {
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

	// Then: Check that we got the expected container matches
	matchedContainers := make(map[string]bool)
	for _, match := range result {
		if container, exists := match["container_name"]; exists {
			matchedContainers[container] = true
		}
	}

	is.Equal(len(matchedContainers), 1)
	is.Equal(matchedContainers["image"], true)
}

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

func TestFilterMatchingRequests_NoMatches(t *testing.T) {
	is := is.New(t)

	// Given: Requested objects that don't match any granted objects
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "write", "container_name": "image"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "video"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: No objects should match
	is.Equal(len(result), 0)
}

func TestFilterMatchingRequests_EmptyInputs(t *testing.T) {
	is := is.New(t)

	// Given: Empty slices
	requested := []map[string]string{}
	granted := []map[string]string{}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: Result should be empty
	is.Equal(len(result), 0)
}

func TestFilterMatchingRequests_WildcardMatching(t *testing.T) {
	is := is.New(t)

	// Given: Granted object with wildcard and various requested objects
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "write", "container_name": "image"},
		{"action": "read", "container_name": "video"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "*"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: Only objects with "read" action should match (via wildcard container_name)
	is.Equal(len(result), 2)
}

func TestFilterMatchingRequests_WildcardMatching_ReadActionOnly(t *testing.T) {
	is := is.New(t)

	// Given: Granted object with wildcard for container_name but specific action requirement
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "write", "container_name": "image"},
		{"action": "read", "container_name": "video"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "*"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: Only objects with "read" action should match (via wildcard container_name)
	is.Equal(len(result), 2)
}

func TestFilterMatchingRequests_MultipleGrantedPatterns_Count(t *testing.T) {
	is := is.New(t)

	// Given: Multiple granted patterns and various requested objects
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "write", "container_name": "image"},
		{"action": "read", "target": "cmd"},
		{"action": "write", "target": "all_up"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "*"},
		{"action": "write", "target": "all_up"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: Two objects should match (first via first pattern, fourth via second pattern)
	is.Equal(len(result), 2)
}

func TestFilterMatchingRequests_MultipleGrantedPatterns_Actions(t *testing.T) {
	is := is.New(t)

	// Given: Multiple granted patterns and various requested objects
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "write", "container_name": "image"},
		{"action": "read", "target": "cmd"},
		{"action": "write", "target": "all_up"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "*"},
		{"action": "write", "target": "all_up"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: Check action distribution in matches
	matchedActions := make(map[string]int)
	for _, match := range result {
		if action, exists := match["action"]; exists {
			matchedActions[action]++
		}
	}

	is.Equal(matchedActions["read"], 1)  // One "read" action matches (object 1)
	is.Equal(matchedActions["write"], 1) // One "write" action matches (object 4)
}

func TestFilterMatchingRequests_DuplicateRequests_Count(t *testing.T) {
	is := is.New(t)

	// Given: Duplicate requested objects
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "read", "container_name": "image"}, // Duplicate
		{"action": "write", "container_name": "image"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "image"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: Should return matches for each occurrence (including duplicates)
	is.Equal(len(result), 2) // Both duplicate requests match and are returned
}

func TestFilterMatchingRequests_DuplicateRequests_FirstResult(t *testing.T) {
	is := is.New(t)

	// Given: Duplicate requested objects
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "read", "container_name": "image"}, // Duplicate
		{"action": "write", "container_name": "image"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "image"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: First result should have correct values
	is.Equal(result[0]["action"], "read")
	is.Equal(result[0]["container_name"], "image")
}

func TestFilterMatchingRequests_DuplicateRequests_SecondResult(t *testing.T) {
	is := is.New(t)

	// Given: Duplicate requested objects
	requested := []map[string]string{
		{"action": "read", "container_name": "image"},
		{"action": "read", "container_name": "image"}, // Duplicate
		{"action": "write", "container_name": "image"},
	}

	granted := []map[string]string{
		{"action": "read", "container_name": "image"},
	}

	// When: FilterMatchingRequests processes the objects
	result := FilterMatchingRequests(requested, granted)

	// Then: Second result should have correct values
	is.Equal(result[1]["action"], "read")
	is.Equal(result[1]["container_name"], "image")
}
