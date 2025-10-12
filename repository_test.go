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

func TestRepository_no_requested_entries(t *testing.T) {
	// Given
	is, dbManager := setupTestDB(t)
	requested := []Requested{}

	// When
	result, err := dbManager.FindGranted(requested)

	// Then
	is.NoErr(err)
	is.Equal(len(result), 0)
}

func TestRepository_scheme_matching(t *testing.T) {
	tests := []struct {
		name          string
		requested     Requested
		granted       Granted
		expectedCount int
	}{
		{
			name:          "exact scheme match",
			requested:     Requested{Scheme: "image", RequestScheme: "image"},
			granted:       Granted{Scheme: "image", GrandScheme: "image"},
			expectedCount: 1,
		},
		{
			name:          "requested scheme does not match grand scheme",
			requested:     Requested{Scheme: "image", RequestScheme: "image"},
			granted:       Granted{Scheme: "image", GrandScheme: "json"},
			expectedCount: 0,
		},
		{
			name:          "request scheme does not match granted scheme",
			requested:     Requested{Scheme: "image", RequestScheme: "image"},
			granted:       Granted{Scheme: "json", GrandScheme: "image"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand scheme",
			requested:     Requested{Scheme: "image", RequestScheme: "image"},
			granted:       Granted{Scheme: "image", GrandScheme: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request scheme",
			requested:     Requested{Scheme: "image", RequestScheme: "*"},
			granted:       Granted{Scheme: "image", GrandScheme: "image"},
			expectedCount: 1,
		},
		{
			name:          "exact action match",
			requested:     Requested{Action: "image", RequestAction: "image"},
			granted:       Granted{Action: "image", GrandAction: "image"},
			expectedCount: 1,
		},
		{
			name:          "requested action does not match grand action",
			requested:     Requested{Action: "image", RequestAction: "image"},
			granted:       Granted{Action: "image", GrandAction: "json"},
			expectedCount: 0,
		},
		{
			name:          "request action does not match granted action",
			requested:     Requested{Action: "image", RequestAction: "image"},
			granted:       Granted{Action: "json", GrandAction: "image"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand action",
			requested:     Requested{Action: "image", RequestAction: "image"},
			granted:       Granted{Action: "image", GrandAction: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request action",
			requested:     Requested{Action: "image", RequestAction: "*"},
			granted:       Granted{Action: "image", GrandAction: "image"},
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

func mockGranted(db *DbManager, granted Granted) []Granted {
	// Mock implementation to insert granted entry into the database
	err := db.SaveGranted(granted)
	if err != nil {
		panic(err)
	}
	return []Granted{granted}
}

func mockRequested(db *DbManager, requested Requested) []Requested {
	// Mock implementation to insert requested entry into the database
	return []Requested{requested}
}

// confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/confetti-cms/image/pkg/confetti-cms/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms
