package dockerconfig

import "reflect"

// SyncMatcher compares a requested object with a granted object and returns true if they match
// Supports wildcard "*" matching for string values in the granted object
// Parameters:
//   - requested: The object containing the requested permissions/values
//   - granted: The object containing the granted permissions/values (may contain "*" wildcards)
func SyncMatcher(requested, granted map[string]string) bool {
	// If objects are exactly equal, return true
	if reflect.DeepEqual(requested, granted) {
		return true
	}

	// Check if both container_name and target match (either exact or wildcard)
	containerNameMatches := false
	targetMatches := false

	// Check container_name matching
	if grantedContainerName, exists := granted["container_name"]; exists {
		if grantedContainerName == "*" {
			// Wildcard matches any non-empty container_name
			if requestedContainerName, exists := requested["container_name"]; exists && requestedContainerName != "" {
				containerNameMatches = true
			}
		} else {
			// Exact match required
			if requestedContainerName, exists := requested["container_name"]; exists && requestedContainerName == grantedContainerName {
				containerNameMatches = true
			}
		}
	}

	// Check target matching
	if grantedTarget, exists := granted["target"]; exists {
		if grantedTarget == "*" {
			// Wildcard matches any non-empty target
			if requestedTarget, exists := requested["target"]; exists && requestedTarget != "" {
				targetMatches = true
			}
		} else {
			// Exact match required
			if requestedTarget, exists := requested["target"]; exists && requestedTarget == grantedTarget {
				targetMatches = true
			}
		}
	}

	// Both fields must match for overall match
	return containerNameMatches && targetMatches

	return false
}
