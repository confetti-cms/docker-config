# SyncMatcher

A flexible JSON object matcher that supports wildcard matching for permission and configuration validation.

## Overview

`SyncMatcher` compares two JSON objects and returns `true` if they match, with support for wildcard `"*"` matching in specific fields.

## Features

- **Exact matching**: Objects must be identical to match
- **Wildcard support**: Supports `"*"` wildcard in `container_name` and `target` fields
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

## Wildcard Matching Rules

### Container Name Matching
The `container_name` field supports wildcard matching:

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
The `target` field supports wildcard matching:

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
You can use wildcards in multiple fields simultaneously:

```go
// Example: Grant access to any container and any target
granted := map[string]string{
    "container_name": "*",
    "target": "*",
}
requested := map[string]string{
    "container_name": "vendor/confetti-cms/image/container",
    "target": "cmd",
}
// Result: true (both wildcards match)
```

## Test Examples

The following test scenarios demonstrate the matching behavior:

### Exact Match Tests
```go
// Container name exact match
requested := map[string]string{"container_name": "vendor/confetti-cms/image/container"}
granted := map[string]string{"container_name": "vendor/confetti-cms/image/container"}
// Result: true

// Target exact match
requested := map[string]string{"target": "cmd"}
granted := map[string]string{"target": "cmd"}
// Result: true
```

### No Match Tests
```go
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
// Container name wildcard match
requested := map[string]string{"container_name": "vendor/confetti-cms/image/container"}
granted := map[string]string{"container_name": "*"}
// Result: true

// Target wildcard match
requested := map[string]string{"target": "cmd"}
granted := map[string]string{"target": "*"}
// Result: true
```

## Implementation Details

The matcher uses Go's `reflect.DeepEqual` for exact matching and provides additional wildcard logic for specific fields. It works directly with `map[string]string` types and checks for non-empty string values when wildcards are used in the `container_name` and `target` fields.

## Related Documentation

For more information about the use cases and design decisions, see:
- [GitHub Discussion #22](https://github.com/confetti-cms/community/discussions/22)

## Running Tests

```bash
go test -v
```

This will run all test scenarios including exact matches, non-matches, and wildcard scenarios for both container names and targets.# docker-config
