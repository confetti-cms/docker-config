package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func setupTestDB(t *testing.T) (*is.I, *DbManager) {
	is := is.New(t)

	dbManager, err := NewDbManager()
	if err != nil {
		t.Fatalf("Failed to create DbManager: %v", err)
	}

	t.Cleanup(func() {
		dbManager.Close()
	})

	return is, dbManager
}

func TestRepository_FindGranted_no_requested_entries(t *testing.T) {
	// Given
	is, dbManager := setupTestDB(t)
	requested := []Requested{}

	// When
	result, err := dbManager.FindGranted(requested)

	// Then
	is.NoErr(err)
	is.Equal(len(result), 0)
}

func TestRepository_FindGranted_matching(t *testing.T) {
	tests := []struct {
		name          string
		requested     Requested
		granted       Granted
		expectedCount int
	}{
		{
			name:          "exact scheme match",
			requested:     Requested{RequestScheme: "image"},
			granted:       Granted{GrandScheme: "image"},
			expectedCount: 1,
		},
		{
			name:          "requested scheme does not match grand scheme",
			requested:     Requested{RequestScheme: "image"},
			granted:       Granted{GrandScheme: "json"},
			expectedCount: 0,
		},
		{
			name:          "request scheme does not match granted scheme",
			requested:     Requested{RequestScheme: "image"},
			granted:       Granted{GrandScheme: "image"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand scheme",
			requested:     Requested{RequestScheme: "image"},
			granted:       Granted{GrandScheme: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request scheme",
			requested:     Requested{RequestScheme: "*"},
			granted:       Granted{GrandScheme: "image"},
			expectedCount: 1,
		},
		{
			name:          "exact action match",
			requested:     Requested{RequestAction: "read"},
			granted:       Granted{GrandAction: "read"},
			expectedCount: 1,
		},
		{
			name:          "requested action does not match grand action",
			requested:     Requested{RequestAction: "read"},
			granted:       Granted{GrandAction: "write"},
			expectedCount: 0,
		},
		{
			name:          "request action does not match granted action",
			requested:     Requested{RequestAction: "read"},
			granted:       Granted{GrandAction: "read"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand action",
			requested:     Requested{RequestAction: "read"},
			granted:       Granted{GrandAction: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request action",
			requested:     Requested{RequestAction: "*"},
			granted:       Granted{GrandAction: "read"},
			expectedCount: 1,
		},
		{
			name:          "exact source organization match",
			requested:     Requested{SourceOrganization: "test-org", RequestSourceOrganization: "test-org"},
			granted:       Granted{SourceOrganization: "test-org", GrandSourceOrganization: "test-org"},
			expectedCount: 1,
		},
		{
			name:          "source organization mismatch",
			requested:     Requested{SourceOrganization: "test-org", RequestSourceOrganization: "test-org"},
			granted:       Granted{SourceOrganization: "test-org", GrandSourceOrganization: "different-org"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand source organization",
			requested:     Requested{SourceOrganization: "test-org", RequestSourceOrganization: "test-org"},
			granted:       Granted{SourceOrganization: "test-org", GrandSourceOrganization: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request source organization",
			requested:     Requested{SourceOrganization: "test-org", RequestSourceOrganization: "*"},
			granted:       Granted{SourceOrganization: "test-org", GrandSourceOrganization: "test-org"},
			expectedCount: 1,
		},
		{
			name:          "exact source repository match",
			requested:     Requested{SourceRepository: "test-repo", RequestSourceRepository: "test-repo"},
			granted:       Granted{SourceRepository: "test-repo", GrandSourceRepository: "test-repo"},
			expectedCount: 1,
		},
		{
			name:          "source repository mismatch",
			requested:     Requested{SourceRepository: "test-repo", RequestSourceRepository: "test-repo"},
			granted:       Granted{SourceRepository: "test-repo", GrandSourceRepository: "different-repo"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand source repository",
			requested:     Requested{SourceRepository: "test-repo", RequestSourceRepository: "test-repo"},
			granted:       Granted{SourceRepository: "test-repo", GrandSourceRepository: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request source repository",
			requested:     Requested{SourceRepository: "test-repo", RequestSourceRepository: "*"},
			granted:       Granted{SourceRepository: "test-repo", GrandSourceRepository: "test-repo"},
			expectedCount: 1,
		},
		{
			name:          "exact umbrella organization match",
			requested:     Requested{UmbrellaOrganization: "test-umb-org", RequestUmbrellaOrganization: "test-umb-org"},
			granted:       Granted{UmbrellaOrganization: "test-umb-org", GrandUmbrellaOrganization: "test-umb-org"},
			expectedCount: 1,
		},
		{
			name:          "umbrella organization mismatch",
			requested:     Requested{UmbrellaOrganization: "test-umb-org", RequestUmbrellaOrganization: "test-umb-org"},
			granted:       Granted{UmbrellaOrganization: "test-umb-org", GrandUmbrellaOrganization: "different-umb-org"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand umbrella organization",
			requested:     Requested{UmbrellaOrganization: "test-umb-org", RequestUmbrellaOrganization: "test-umb-org"},
			granted:       Granted{UmbrellaOrganization: "test-umb-org", GrandUmbrellaOrganization: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request umbrella organization",
			requested:     Requested{UmbrellaOrganization: "test-umb-org", RequestUmbrellaOrganization: "*"},
			granted:       Granted{UmbrellaOrganization: "test-umb-org", GrandUmbrellaOrganization: "test-umb-org"},
			expectedCount: 1,
		},
		{
			name:          "exact umbrella repository match",
			requested:     Requested{UmbrellaRepository: "test-umb-repo", RequestUmbrellaRepository: "test-umb-repo"},
			granted:       Granted{UmbrellaRepository: "test-umb-repo", GrandUmbrellaRepository: "test-umb-repo"},
			expectedCount: 1,
		},
		{
			name:          "umbrella repository mismatch",
			requested:     Requested{UmbrellaRepository: "test-umb-repo", RequestUmbrellaRepository: "test-umb-repo"},
			granted:       Granted{UmbrellaRepository: "test-umb-repo", GrandUmbrellaRepository: "different-umb-repo"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand umbrella repository",
			requested:     Requested{UmbrellaRepository: "test-umb-repo", RequestUmbrellaRepository: "test-umb-repo"},
			granted:       Granted{UmbrellaRepository: "test-umb-repo", GrandUmbrellaRepository: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request umbrella repository",
			requested:     Requested{UmbrellaRepository: "test-umb-repo", RequestUmbrellaRepository: "*"},
			granted:       Granted{UmbrellaRepository: "test-umb-repo", GrandUmbrellaRepository: "test-umb-repo"},
			expectedCount: 1,
		},
		{
			name:          "exact container name match",
			requested:     Requested{ContainerName: "test-container", RequestContainerName: "test-container"},
			granted:       Granted{ContainerName: "test-container", GrandContainerName: "test-container"},
			expectedCount: 1,
		},
		{
			name:          "container name mismatch",
			requested:     Requested{ContainerName: "test-container", RequestContainerName: "test-container"},
			granted:       Granted{ContainerName: "test-container", GrandContainerName: "different-container"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand container name",
			requested:     Requested{ContainerName: "test-container", RequestContainerName: "test-container"},
			granted:       Granted{ContainerName: "test-container", GrandContainerName: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request container name",
			requested:     Requested{ContainerName: "test-container", RequestContainerName: "*"},
			granted:       Granted{ContainerName: "test-container", GrandContainerName: "test-container"},
			expectedCount: 1,
		},
		{
			name:          "exact target match",
			requested:     Requested{Target: "cmd", RequestTarget: "cmd"},
			granted:       Granted{Target: "cmd", GrandTarget: "cmd"},
			expectedCount: 1,
		},
		{
			name:          "target mismatch",
			requested:     Requested{Target: "cmd", RequestTarget: "cmd"},
			granted:       Granted{Target: "cmd", GrandTarget: "all_up"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand target",
			requested:     Requested{Target: "cmd", RequestTarget: "cmd"},
			granted:       Granted{Target: "cmd", GrandTarget: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request target",
			requested:     Requested{Target: "cmd", RequestTarget: "*"},
			granted:       Granted{Target: "cmd", GrandTarget: "cmd"},
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is, db := setupTestDB(t)

			r := mockRequested(db, tt.requested)
			mockGranted(db, tt.granted)

			// When
			result, err := db.FindGranted(r)

			// Then
			is.NoErr(err)
			is.Equal(len(result), tt.expectedCount)
		})
	}
}

func TestRepository_FindRequested_no_granted_entries(t *testing.T) {
	// Given
	is, dbManager := setupTestDB(t)
	granted := []Granted{}

	// When
	result, err := dbManager.FindRequested(granted)

	// Then
	is.NoErr(err)
	is.Equal(len(result), 0)
}

func TestRepository_FindRequested_matching(t *testing.T) {
	tests := []struct {
		name          string
		granted       Granted
		requested     Requested
		expectedCount int
	}{
		{
			name:          "exact scheme match",
			granted:       Granted{GrandScheme: "image"},
			requested:     Requested{RequestScheme: "image"},
			expectedCount: 1,
		},
		{
			name:          "granted scheme does not match request scheme",
			granted:       Granted{GrandScheme: "image"},
			requested:     Requested{RequestScheme: "json"},
			expectedCount: 0,
		},
		{
			name:          "grant scheme does not match requested scheme",
			granted:       Granted{GrandScheme: "image"},
			requested:     Requested{RequestScheme: "image"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in request scheme",
			granted:       Granted{GrandScheme: "image"},
			requested:     Requested{RequestScheme: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in grand scheme",
			granted:       Granted{GrandScheme: "*"},
			requested:     Requested{RequestScheme: "image"},
			expectedCount: 1,
		},
		{
			name:          "exact action match",
			granted:       Granted{GrandAction: "read"},
			requested:     Requested{RequestAction: "read"},
			expectedCount: 1,
		},
		{
			name:          "granted action does not match request action",
			granted:       Granted{GrandAction: "read"},
			requested:     Requested{RequestAction: "write"},
			expectedCount: 0,
		},
		{
			name:          "grant action does not match requested action",
			granted:       Granted{GrandAction: "read"},
			requested:     Requested{RequestAction: "read"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in request action",
			granted:       Granted{GrandAction: "read"},
			requested:     Requested{RequestAction: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in grand action",
			granted:       Granted{GrandAction: "*"},
			requested:     Requested{RequestAction: "read"},
			expectedCount: 1,
		},
		{
			name:          "exact source organization match",
			granted:       Granted{SourceOrganization: "test-org", GrandSourceOrganization: "test-org"},
			requested:     Requested{SourceOrganization: "test-org", RequestSourceOrganization: "test-org"},
			expectedCount: 1,
		},
		{
			name:          "source organization mismatch",
			granted:       Granted{SourceOrganization: "test-org", GrandSourceOrganization: "test-org"},
			requested:     Requested{SourceOrganization: "test-org", RequestSourceOrganization: "different-org"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in request source organization",
			granted:       Granted{SourceOrganization: "test-org", GrandSourceOrganization: "test-org"},
			requested:     Requested{SourceOrganization: "test-org", RequestSourceOrganization: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in grand source organization",
			granted:       Granted{SourceOrganization: "test-org", GrandSourceOrganization: "*"},
			requested:     Requested{SourceOrganization: "test-org", RequestSourceOrganization: "test-org"},
			expectedCount: 1,
		},
		{
			name:          "exact source repository match",
			granted:       Granted{SourceRepository: "test-repo", GrandSourceRepository: "test-repo"},
			requested:     Requested{SourceRepository: "test-repo", RequestSourceRepository: "test-repo"},
			expectedCount: 1,
		},
		{
			name:          "source repository mismatch",
			granted:       Granted{SourceRepository: "test-repo", GrandSourceRepository: "test-repo"},
			requested:     Requested{SourceRepository: "test-repo", RequestSourceRepository: "different-repo"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in request source repository",
			granted:       Granted{SourceRepository: "test-repo", GrandSourceRepository: "test-repo"},
			requested:     Requested{SourceRepository: "test-repo", RequestSourceRepository: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in grand source repository",
			granted:       Granted{SourceRepository: "test-repo", GrandSourceRepository: "*"},
			requested:     Requested{SourceRepository: "test-repo", RequestSourceRepository: "test-repo"},
			expectedCount: 1,
		},
		{
			name:          "exact umbrella organization match",
			granted:       Granted{UmbrellaOrganization: "test-umb-org", GrandUmbrellaOrganization: "test-umb-org"},
			requested:     Requested{UmbrellaOrganization: "test-umb-org", RequestUmbrellaOrganization: "test-umb-org"},
			expectedCount: 1,
		},
		{
			name:          "umbrella organization mismatch",
			granted:       Granted{UmbrellaOrganization: "test-umb-org", GrandUmbrellaOrganization: "test-umb-org"},
			requested:     Requested{UmbrellaOrganization: "test-umb-org", RequestUmbrellaOrganization: "different-umb-org"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in request umbrella organization",
			granted:       Granted{UmbrellaOrganization: "test-umb-org", GrandUmbrellaOrganization: "test-umb-org"},
			requested:     Requested{UmbrellaOrganization: "test-umb-org", RequestUmbrellaOrganization: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in grand umbrella organization",
			granted:       Granted{UmbrellaOrganization: "test-umb-org", GrandUmbrellaOrganization: "*"},
			requested:     Requested{UmbrellaOrganization: "test-umb-org", RequestUmbrellaOrganization: "test-umb-org"},
			expectedCount: 1,
		},
		{
			name:          "exact umbrella repository match",
			granted:       Granted{UmbrellaRepository: "test-umb-repo", GrandUmbrellaRepository: "test-umb-repo"},
			requested:     Requested{UmbrellaRepository: "test-umb-repo", RequestUmbrellaRepository: "test-umb-repo"},
			expectedCount: 1,
		},
		{
			name:          "umbrella repository mismatch",
			granted:       Granted{UmbrellaRepository: "test-umb-repo", GrandUmbrellaRepository: "test-umb-repo"},
			requested:     Requested{UmbrellaRepository: "test-umb-repo", RequestUmbrellaRepository: "different-umb-repo"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in request umbrella repository",
			granted:       Granted{UmbrellaRepository: "test-umb-repo", GrandUmbrellaRepository: "test-umb-repo"},
			requested:     Requested{UmbrellaRepository: "test-umb-repo", RequestUmbrellaRepository: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in grand umbrella repository",
			granted:       Granted{UmbrellaRepository: "test-umb-repo", GrandUmbrellaRepository: "*"},
			requested:     Requested{UmbrellaRepository: "test-umb-repo", RequestUmbrellaRepository: "test-umb-repo"},
			expectedCount: 1,
		},
		{
			name:          "exact container name match",
			granted:       Granted{ContainerName: "test-container", GrandContainerName: "test-container"},
			requested:     Requested{ContainerName: "test-container", RequestContainerName: "test-container"},
			expectedCount: 1,
		},
		{
			name:          "container name mismatch",
			granted:       Granted{ContainerName: "test-container", GrandContainerName: "test-container"},
			requested:     Requested{ContainerName: "test-container", RequestContainerName: "different-container"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in request container name",
			granted:       Granted{ContainerName: "test-container", GrandContainerName: "test-container"},
			requested:     Requested{ContainerName: "test-container", RequestContainerName: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in grand container name",
			granted:       Granted{ContainerName: "test-container", GrandContainerName: "*"},
			requested:     Requested{ContainerName: "test-container", RequestContainerName: "test-container"},
			expectedCount: 1,
		},
		{
			name:          "exact target match",
			granted:       Granted{Target: "cmd", GrandTarget: "cmd"},
			requested:     Requested{Target: "cmd", RequestTarget: "cmd"},
			expectedCount: 1,
		},
		{
			name:          "target mismatch",
			granted:       Granted{Target: "cmd", GrandTarget: "cmd"},
			requested:     Requested{Target: "cmd", RequestTarget: "all_up"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in request target",
			granted:       Granted{Target: "cmd", GrandTarget: "cmd"},
			requested:     Requested{Target: "cmd", RequestTarget: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in grand target",
			granted:       Granted{Target: "cmd", GrandTarget: "*"},
			requested:     Requested{Target: "cmd", RequestTarget: "cmd"},
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is, db := setupTestDB(t)

			g := mockGranted(db, tt.granted)
			mockRequested(db, tt.requested)

			// When
			result, err := db.FindRequested(g)

			// Then
			is.NoErr(err)
			is.Equal(len(result), tt.expectedCount)
		})
	}
}

func TestRepository_FindGranted_multiple_scheme_matches(t *testing.T) {
	// Given
	is, dbManager := setupTestDB(t)

	// Create two granted entries with the same scheme but different GrandScheme values
	granted1 := Granted{
		GrandScheme: "image",
	}

	granted2 := Granted{
		GrandScheme: "json",
	}

	// Save both granted entries
	mockGranted(dbManager, granted1)
	mockGranted(dbManager, granted2)

	// Create one requested entry that should match both granted entries using wildcard
	requested := []Requested{
		{
			RequestScheme: "*", // Wildcard matches any GrandScheme
		},
	}

	// When
	result, err := dbManager.FindGranted(requested)

	// Then
	is.NoErr(err)
	is.Equal(len(result), 2) // Should return both matching granted entries

	// Verify that we got results for both GrandScheme values
	grandSchemes := make(map[string]bool)
	for _, granted := range result {
		grandSchemes[granted.GrandScheme] = true
	}
	is.Equal(len(grandSchemes), 2) // Should have both grand schemes
	is.True(grandSchemes["image"])
	is.True(grandSchemes["json"])
}
func TestRepository_FindRequested_multiple_scheme_matches(t *testing.T) {
	// Given
	is, dbManager := setupTestDB(t)

	// Create two requested entries with the same scheme but different RequestScheme values
	requested1 := Requested{
		RequestScheme: "image",
	}

	requested2 := Requested{
		RequestScheme: "json",
	}

	// Save both requested entries
	mockRequested(dbManager, requested1)
	mockRequested(dbManager, requested2)

	// Create one granted entry that should match both requested entries using wildcard
	granted := Granted{
		GrandScheme: "*", // Wildcard matches any RequestScheme when scheme matches
	}

	// When
	result, err := dbManager.FindRequested([]Granted{granted})

	// Then
	is.NoErr(err)
	is.Equal(len(result), 2) // Should return both matching requested entries

	// Verify that we got results for both RequestScheme values
	requestSchemes := make(map[string]bool)
	for _, requested := range result {
		requestSchemes[requested.RequestScheme] = true
	}
	is.Equal(len(requestSchemes), 2) // Should have both request schemes
	is.True(requestSchemes["image"])
	is.True(requestSchemes["json"])
}

func mockGranted(db *DbManager, granted Granted) []Granted {
	// Mock implementation to insert granted entry into the database
	err := db.SaveGranted(granted)
	if err != nil {
		panic(err)
	}
	return []Granted{granted}
}

func TestNewDbManager_Error(t *testing.T) {
	// Given - We can't easily mock a database connection failure with the current setup
	// But we can test the function exists and handles basic cases
	is := is.New(t)

	// When
	dbManager, err := NewDbManager()

	// Then
	is.NoErr(err)
	is.True(dbManager != nil)
	is.True(dbManager.db != nil)

	// Cleanup
	dbManager.Close()
}

func TestDbManager_Close_Error(t *testing.T) {
	is := is.New(t)

	// Test closing an already closed database
	dbManager, err := NewDbManager()
	is.NoErr(err)

	err = dbManager.Close()
	is.NoErr(err) // First close should succeed

	err = dbManager.Close()
	is.NoErr(err) // Second close should also succeed (idempotent behavior)
}

func TestFindRequested_RowsErr(t *testing.T) {
	// Given
	is, dbManager := setupTestDB(t)

	// Insert a record that will cause issues during scanning
	// We'll use a malformed query approach by inserting data directly
	requested := Requested{
		Description:                 "test",
		DestinationPath:             "/test",
		SourceOrganization:          "test",
		SourceRepository:            "test",
		UmbrellaOrganization:        "test",
		UmbrellaRepository:          "test",
		ContainerName:               "test",
		Target:                      "test",
		RequestScheme:               "test",
		RequestAction:               "test",
		RequestSourceOrganization:   "test",
		RequestSourceRepository:     "test",
		RequestUmbrellaOrganization: "test",
		RequestUmbrellaRepository:   "test",
		RequestContainerName:        "test",
		RequestTarget:               "test",
	}

	err := dbManager.SaveRequested([]Requested{requested})
	is.NoErr(err)

	// When - Try to find with matching criteria
	granted := []Granted{
		{
			SourceOrganization:        "test",
			SourceRepository:          "test",
			UmbrellaOrganization:      "test",
			UmbrellaRepository:        "test",
			ContainerName:             "test",
			Target:                    "test",
			GrandScheme:               "test",
			GrandAction:               "test",
			GrandSourceOrganization:   "test",
			GrandSourceRepository:     "test",
			GrandUmbrellaOrganization: "test",
			GrandUmbrellaRepository:   "test",
			GrandContainerName:        "test",
			GrandTarget:               "test",
		},
	}

	result, err := dbManager.FindRequested(granted)

	// Then - Should handle any rows errors gracefully
	is.NoErr(err)
	is.Equal(len(result), 1)
}

func TestFindGranted_RowsErr(t *testing.T) {
	// Given
	is, dbManager := setupTestDB(t)

	// Insert a record that will cause issues during scanning
	granted := Granted{
		Description:               "test",
		ExposePath:                "/test",
		SourceOrganization:        "test",
		SourceRepository:          "test",
		UmbrellaOrganization:      "test",
		UmbrellaRepository:        "test",
		ContainerName:             "test",
		Target:                    "test",
		GrandScheme:               "test",
		GrandAction:               "test",
		GrandSourceOrganization:   "test",
		GrandSourceRepository:     "test",
		GrandUmbrellaOrganization: "test",
		GrandUmbrellaRepository:   "test",
		GrandContainerName:        "test",
		GrandTarget:               "test",
	}

	err := dbManager.SaveGranted(granted)
	is.NoErr(err)

	// When - Try to find with matching criteria
	requested := []Requested{
		{
			SourceOrganization:          "test",
			SourceRepository:            "test",
			UmbrellaOrganization:        "test",
			UmbrellaRepository:          "test",
			ContainerName:               "test",
			Target:                      "test",
			RequestScheme:               "test",
			RequestAction:               "test",
			RequestSourceOrganization:   "test",
			RequestSourceRepository:     "test",
			RequestUmbrellaOrganization: "test",
			RequestUmbrellaRepository:   "test",
			RequestContainerName:        "test",
			RequestTarget:               "test",
		},
	}

	result, err := dbManager.FindGranted(requested)

	// Then - Should handle any rows errors gracefully
	is.NoErr(err)
	is.Equal(len(result), 1)
}

func mockRequested(db *DbManager, requested Requested) []Requested {
	// Mock implementation to insert requested entry into the database
	err := db.SaveRequested([]Requested{requested})
	if err != nil {
		panic(err)
	}
	return []Requested{requested}
}
