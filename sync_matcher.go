package dockerconfig

import (
	"net/url"
	"reflect"
	"strings"
)

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

	// Check environment_name matching
	if grantedEnvName, exists := granted["environment_name"]; exists {
		if grantedEnvName == "*" {
			// Wildcard matches any non-empty environment_name
			if requestedEnvName, exists := requested["environment_name"]; !exists || requestedEnvName == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedEnvName, exists := requested["environment_name"]; !exists || requestedEnvName != grantedEnvName {
				return false
			}
		}
	}

	// Check environment_stage matching
	if grantedEnvStage, exists := granted["environment_stage"]; exists {
		if grantedEnvStage == "*" {
			// Wildcard matches any non-empty environment_stage
			if requestedEnvStage, exists := requested["environment_stage"]; !exists || requestedEnvStage == "" {
				return false
			}
		} else {
			// Exact match required
			if requestedEnvStage, exists := requested["environment_stage"]; !exists || requestedEnvStage != grantedEnvStage {
				return false
			}
		}
	}

	// All specified fields match
	return true
}

// FilterMatchingRequests compares multiple requested objects against multiple granted objects
// and returns all requested objects that match at least one granted object
// Parameters:
//   - requested: Slice of objects containing the requested permissions/values
//   - granted: Slice of objects containing the granted permissions/values (may contain "*" wildcards)
//
// Returns:
//   - Slice of requested objects that match at least one granted object
func FilterMatchingRequests(requested []map[string]string, granted []map[string]string) []map[string]string {
	var matched []map[string]string

	// For each requested object, check if it matches any granted object
	for _, req := range requested {
		for _, gran := range granted {
			if SyncMatcher(req, gran) {
				// Add to matched results and break to avoid duplicates
				matched = append(matched, req)
				break
			}
		}
	}

	return matched
}

// ParseLocator parses a locator URL string into a map[string]string
// The locator format is: locator://path?query_parameters
// Parameters:
//   - locator: The locator string to parse
//
// Returns:
//   - map[string]string containing the parsed fields
func ParseLocator(locator string) map[string]string {
	result := make(map[string]string)

	// Parse the locator URL
	if !strings.HasPrefix(locator, "locator://") {
		return result
	}

	// Remove the scheme
	urlPart := strings.TrimPrefix(locator, "locator://")

	// Split into path and query parts
	parts := strings.SplitN(urlPart, "?", 2)
	path := parts[0]

	// Parse the path part first
	// Format: org-repo_env-stage-target (simplified parsing)
	pathParts := strings.Split(path, "_")

	// Basic path parsing
	if len(pathParts) >= 1 {
		// First part contains org-repo information
		orgRepo := pathParts[0]
		if orgRepoParts := strings.Split(orgRepo, "-"); len(orgRepoParts) >= 2 {
			result["umbrella_organization"] = orgRepoParts[0]
			result["umbrella_repository"] = orgRepoParts[1]
		}
	}

	if len(pathParts) >= 2 {
		// Second part contains environment information
		env := pathParts[1]
		if envParts := strings.Split(env, "-"); len(envParts) >= 1 {
			result["environment_name"] = envParts[0]
			if len(envParts) >= 2 {
				result["environment_stage"] = strings.Join(envParts[1:], "-")
			}
		}
	}

	if len(pathParts) >= 3 {
		// Third part contains target information
		target := pathParts[2]
		if targetParts := strings.Split(target, "-"); len(targetParts) >= 1 {
			result["target"] = targetParts[len(targetParts)-1]
		}
	}

	// Parse query parameters if present (these override path parsing)
	if len(parts) > 1 {
		query := parts[1]
		values, err := url.ParseQuery(query)
		if err == nil {
			// Add query parameters to result (overriding path parsing)
			for key, vals := range values {
				if len(vals) > 0 {
					result[key] = vals[0]
				}
			}
		}
	}

	return result
}

// LocatorMatcher compares a locator string with a granted object and returns true if they match
// The locator is parsed into fields and then compared using the existing SyncMatcher logic
// Parameters:
//   - locator: The locator string to parse and match
//   - granted: The object containing the granted permissions/values (may contain "*" wildcards)
//
// Returns:
//   - true if the parsed locator matches the granted object
func LocatorMatcher(locator string, granted map[string]string) bool {
	// Parse the locator into fields
	parsedLocator := ParseLocator(locator)

	// Use existing SyncMatcher to compare parsed locator with granted object
	return SyncMatcher(parsedLocator, granted)
}
