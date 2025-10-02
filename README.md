# SyncMatcher

A flexible JSON object matcher that supports wildcard matching for permission and configuration validation.

## Overview

`SyncMatcher` compares two JSON objects and returns `true` if they match, with support for wildcard `"*"` matching across multiple fields including schema, host, container details, actions, and organizational information.

## Features

- **Exact matching**: Objects must be identical to match
- **Wildcard support**: Supports `"*"` wildcard across all supported fields
- **Multi-field matching**: Supports schema, host, container_name, target, action, and organizational fields
- **Type-safe comparison**: Works specifically with `map[string]string` types for key-value permission matching
- **Field-specific logic**: Each field type has its own matching rules

## Usage

```go
result := SyncMatcher(requested, granted)
```

## Type Signature

```go
func SyncMatcher(requested, granted map[string]string) bool
```

## Parameters

- **requested**: The object containing the requested permissions/values
- **granted**: The object containing the granted permissions/values (may contain `"*"` wildcards)

## Supported Fields

The `SyncMatcher` supports matching across the following fields:

- **schema**: The schema type (e.g., "image", "hive", or other custom values)
- **host**: The host or endpoint identifier
- **container_name**: The container identifier or path
- **target**: The target operation or command
- **action**: The action being performed (e.g., "push", "pull")
- **source_organization**: The source organization identifier
- **source_repository**: The source repository identifier
- **umbrella_organization**: The umbrella/parent organization identifier
- **umbrella_repository**: The umbrella/parent repository identifier

## Wildcard Matching Rules

All supported fields support wildcard `"*"` matching:

### Schema Matching
```go
// Example: Grant access to any schema type
granted := map[string]string{
    "schema": "*",
}
requested := map[string]string{
    "schema": "image",
}
// Result: true (wildcard matches any non-empty schema)
```

### Host Matching
```go
// Example: Grant access to any host
granted := map[string]string{
    "host": "*",
}
requested := map[string]string{
    "host": "localhost",
}
// Result: true (wildcard matches any non-empty host)
```

### Container Name Matching
```go
// Example: Grant access to any container
granted := map[string]string{
    "container_name": "*",
}
requested := map[string]string{
    "container_name": "vendor/confetti-cms/image/container",
}
// Result: true (wildcard matches any non-empty container name)
```

### Target Matching
```go
// Example: Grant access to any target
granted := map[string]string{
    "target": "*",
}
requested := map[string]string{
    "target": "cmd",
}
// Result: true (wildcard matches any non-empty target)
```

### Multiple Field Wildcards
You can use wildcards across multiple fields simultaneously:

```go
// Example: Grant access with flexible schema and host but specific container
granted := map[string]string{
    "schema": "*",
    "host": "*",
    "container_name": "vendor/confetti-cms/image/container",
}
requested := map[string]string{
    "schema": "image",
    "host": "localhost",
    "container_name": "vendor/confetti-cms/image/container",
}
// Result: true (schema and host wildcards match, container matches exactly)
```

## Test Examples

The following test scenarios demonstrate the matching behavior across all supported fields:

### Exact Match Tests
```go
// Schema exact match
requested := map[string]string{"schema": "image"}
granted := map[string]string{"schema": "image"}
// Result: true

// Host exact match
requested := map[string]string{"host": "localhost"}
granted := map[string]string{"host": "localhost"}
// Result: true

// Container name exact match
requested := map[string]string{"container_name": "vendor/confetti-cms/image/container"}
granted := map[string]string{"container_name": "vendor/confetti-cms/image/container"}
// Result: true

// Target exact match
requested := map[string]string{"target": "cmd"}
granted := map[string]string{"target": "cmd"}
// Result: true

// Action exact match
requested := map[string]string{"action": "push"}
granted := map[string]string{"action": "push"}
// Result: true
```

### No Match Tests
```go
// Schema no match
requested := map[string]string{"schema": "image"}
granted := map[string]string{"schema": "hive"}
// Result: false

// Host no match
requested := map[string]string{"host": "localhost"}
granted := map[string]string{"host": "remotehost"}
// Result: false

// Container name no match
requested := map[string]string{"container_name": "vendor/confetti-cms/image/container"}
granted := map[string]string{"container_name": "vendor/confetti-cms/different/container"}
// Result: false

// Target no match
requested := map[string]string{"target": "cmd"}
granted := map[string]string{"target": "different"}
// Result: false
```

### Wildcard Match Tests
```go
// Schema wildcard match
requested := map[string]string{"schema": "image"}
granted := map[string]string{"schema": "*"}
// Result: true

// Host wildcard match
requested := map[string]string{"host": "localhost"}
granted := map[string]string{"host": "*"}
// Result: true

// Container name wildcard match
requested := map[string]string{"container_name": "vendor/confetti-cms/image/container"}
granted := map[string]string{"container_name": "*"}
// Result: true

// Target wildcard match
requested := map[string]string{"target": "cmd"}
granted := map[string]string{"target": "*"}
// Result: true
```

### Complex Multi-Field Examples
```go
// Multiple field combination with wildcards
requested := map[string]string{
    "schema": "image",
    "host": "localhost",
    "container_name": "vendor/confetti-cms/image/container",
    "target": "cmd",
    "action": "push",
}
granted := map[string]string{
    "schema": "*",
    "host": "localhost",
    "container_name": "vendor/confetti-cms/image/container",
    "target": "*",
    "action": "push",
}
// Result: true (schema and target match via wildcards, others match exactly)

// Organizational field matching
requested := map[string]string{
    "source_organization": "myorg",
    "source_repository": "myrepo",
    "umbrella_organization": "parentorg",
}
granted := map[string]string{
    "source_organization": "myorg",
    "source_repository": "*",
    "umbrella_organization": "parentorg",
}
// Result: true (source_repository matches via wildcard)
```

## Implementation Details

The matcher uses Go's `reflect.DeepEqual` for exact matching and provides additional wildcard logic for all supported fields. It works directly with `map[string]string` types and checks for non-empty string values when wildcards are used across any of the nine supported fields (schema, host, container_name, target, action, source_organization, source_repository, umbrella_organization, umbrella_repository).

All specified fields in the `granted` object must match their corresponding fields in the `requested` object, either exactly or via wildcard. Fields not present in the `granted` object are not required for matching.

## Related Documentation

For more information about the use cases and design decisions, see:
- [GitHub Discussion #22](https://github.com/confetti-cms/community/discussions/22)

## Running Tests

```bash
go test -v
```

This will run all test scenarios including:
- Exact matches, non-matches, and wildcard scenarios for all nine supported fields
- Complex multi-field combinations and edge cases
- Integration tests with various field combinations
- Comprehensive coverage of schema, host, container, action, and organizational field matching

The test suite includes 41 test cases covering all supported fields and their various matching scenarios.
