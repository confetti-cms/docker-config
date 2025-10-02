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

	// Check if all specified fields in granted match (either exact or wildcard)
	// Only fields present in the granted object need to match

	// Check container_name matching
	if grantedContainerName, exists := granted["container_name"]; exists {
		if grantedContainerName == "*" {
			// Wildcard matches any non-empty container_name
			if requestedContainerName, exists := requested["container_name"]; !exists || requestedContainerName == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedContainerName, exists := requested["container_name"]; !exists || requestedContainerName != grantedContainerName {
				return false
			}
		}
	}

	// Check target matching
	if grantedTarget, exists := granted["target"]; exists {
		if grantedTarget == "*" {
			// Wildcard matches any non-empty target
			if requestedTarget, exists := requested["target"]; !exists || requestedTarget == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedTarget, exists := requested["target"]; !exists || requestedTarget != grantedTarget {
				return false
			}
		}
	}

	// Check host matching
	if grantedHost, exists := granted["host"]; exists {
		if grantedHost == "*" {
			// Wildcard matches any non-empty host
			if requestedHost, exists := requested["host"]; !exists || requestedHost == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedHost, exists := requested["host"]; !exists || requestedHost != grantedHost {
				return false
			}
		}
	}

	// Check schema matching
	if grantedSchema, exists := granted["schema"]; exists {
		if grantedSchema == "*" {
			// Wildcard matches any non-empty schema
			if requestedSchema, exists := requested["schema"]; !exists || requestedSchema == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedSchema, exists := requested["schema"]; !exists || requestedSchema != grantedSchema {
				return false
			}
		}
	}

	// Check action matching
	if grantedAction, exists := granted["action"]; exists {
		if grantedAction == "*" {
			// Wildcard matches any non-empty action
			if requestedAction, exists := requested["action"]; !exists || requestedAction == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedAction, exists := requested["action"]; !exists || requestedAction != grantedAction {
				return false
			}
		}
	}

	// Check source_organization matching
	if grantedSourceOrg, exists := granted["source_organization"]; exists {
		if grantedSourceOrg == "*" {
			// Wildcard matches any non-empty source_organization
			if requestedSourceOrg, exists := requested["source_organization"]; !exists || requestedSourceOrg == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedSourceOrg, exists := requested["source_organization"]; !exists || requestedSourceOrg != grantedSourceOrg {
				return false
			}
		}
	}

	// Check source_repository matching
	if grantedSourceRepo, exists := granted["source_repository"]; exists {
		if grantedSourceRepo == "*" {
			// Wildcard matches any non-empty source_repository
			if requestedSourceRepo, exists := requested["source_repository"]; !exists || requestedSourceRepo == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedSourceRepo, exists := requested["source_repository"]; !exists || requestedSourceRepo != grantedSourceRepo {
				return false
			}
		}
	}

	// Check umbrella_organization matching
	if grantedUmbrellaOrg, exists := granted["umbrella_organization"]; exists {
		if grantedUmbrellaOrg == "*" {
			// Wildcard matches any non-empty umbrella_organization
			if requestedUmbrellaOrg, exists := requested["umbrella_organization"]; !exists || requestedUmbrellaOrg == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedUmbrellaOrg, exists := requested["umbrella_organization"]; !exists || requestedUmbrellaOrg != grantedUmbrellaOrg {
				return false
			}
		}
	}

	// Check umbrella_repository matching
	if grantedUmbrellaRepo, exists := granted["umbrella_repository"]; exists {
		if grantedUmbrellaRepo == "*" {
			// Wildcard matches any non-empty umbrella_repository
			if requestedUmbrellaRepo, exists := requested["umbrella_repository"]; !exists || requestedUmbrellaRepo == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedUmbrellaRepo, exists := requested["umbrella_repository"]; !exists || requestedUmbrellaRepo != grantedUmbrellaRepo {
				return false
			}
		}
	}

	// All specified fields match
	return true

	return false
}
