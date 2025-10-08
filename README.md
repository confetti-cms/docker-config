# SyncMatcher

A flexible permission matcher that supports wildcard matching for configuration validation.

## Why Should You Care?

Permission matching isn't just for security nerds. If you want to know who can do what, and why your website sometimes says "nope," you're in the right place. They're your best friends when you want the user to:

- Only edit their own blog posts (not someone else's)
- View all values on a page, but only edit a select few
- Invite another user to edit blogs in a specific category
- Allow someone to edit a draft article, but only let the admin publish and further modify the page

## Basic Usage

```go
result := SyncMatcher(requested, granted)
```

## Type Signature

```go
func SyncMatcher(requested map[string]string, granted Granted) bool
```

## Parameters

- **requested**: The object containing the requested permissions/values
- **granted**: The object containing the granted permissions/values (may contain `"*"` wildcards)

## Supported Fields

| Field | Description |
|-------|-------------|
| `schema` | The schema type (e.g., "image", "hive", or other custom values) |
| `host` | The host or endpoint identifier |
| `container_name` | The container identifier or path |
| `target` | The target operation or command |
| `action` | The action being performed (e.g., "push", "pull") |
| `source_organization` | The source organization identifier |
| `source_repository` | The source repository identifier |
| `umbrella_organization` | The umbrella/parent organization identifier |
| `umbrella_repository` | The umbrella/parent repository identifier |

## Wildcard Matching

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

### Exact Matches

| Requested | Granted | Result |
|-----------|---------|--------|
| `{"schema": "image"}` | `{"schema": "image"}` | ✅ `true` |
| `{"host": "localhost"}` | `{"host": "localhost"}` | ✅ `true` |
| `{"container_name": "vendor/confetti-cms/image/container"}` | `{"container_name": "vendor/confetti-cms/image/container"}` | ✅ `true` |

### No Matches

| Requested | Granted | Result |
|-----------|---------|--------|
| `{"schema": "image"}` | `{"schema": "hive"}` | ❌ `false` |
| `{"host": "localhost"}` | `{"host": "remotehost"}` | ❌ `false` |
| `{"container_name": "vendor/confetti-cms/image/container"}` | `{"container_name": "vendor/confetti-cms/different/container"}` | ❌ `false` |

### Wildcard Matches

| Requested | Granted | Result |
|-----------|---------|--------|
| `{"schema": "image"}` | `{"schema": "*"}` | ✅ `true` |
| `{"host": "localhost"}` | `{"host": "*"}` | ✅ `true` |
| `{"container_name": "vendor/confetti-cms/image/container"}` | `{"container_name": "*"}` | ✅ `true` |

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

The matcher uses Go's `reflect.DeepEqual` for exact matching and provides additional wildcard logic for all supported fields. It works directly with `map[string]string` types and checks for non-empty string values when wildcards are used.

All specified fields in the `granted` object must match their corresponding fields in the `requested` object, either exactly or via wildcard. Fields not present in the `granted` object are not required for matching.

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

## Related Documentation

For more information about the use cases and design decisions, see:
- [GitHub Discussion #22](https://github.com/confetti-cms/community/discussions/22)
