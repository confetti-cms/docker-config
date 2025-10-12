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

func TestRepository_request_scheme_match_granted_scheme(t *testing.T) {
	// Given
	is, db := setupTestDB(t)

	r := mockRequested(db, Requested{Scheme: "image", RequestScheme: "image"})
	mockGranted(db, Granted{Scheme: "image", GrandScheme: "image"})

	// When
	result, err := db.FindGranted(r)

	// Then
	is.NoErr(err)
	is.Equal(len(result), 1)
}

func TestRepository_requested_scheme_does_not_match_grand_scheme(t *testing.T) {
	// Given
	is, db := setupTestDB(t)

	r := mockRequested(db, Requested{Scheme: "image", RequestScheme: "image"})
	mockGranted(db, Granted{Scheme: "image", GrandScheme: "json"})

	// When
	result, err := db.FindGranted(r)

	// Then
	is.NoErr(err)
	is.Equal(len(result), 0)
}

func TestRepository_request_scheme_does_not_match_granted_scheme(t *testing.T) {
	// Given
	is, db := setupTestDB(t)

	r := mockRequested(db, Requested{Scheme: "image", RequestScheme: "image"})
	mockGranted(db, Granted{Scheme: "json", GrandScheme: "image"})

	// When
	result, err := db.FindGranted(r)

	// Then
	is.NoErr(err)
	is.Equal(len(result), 0)
}

func mockGranted(db *DbManager, granted Granted) []Granted {
	// Mock implementation to insert granted entry into the database
	err := db.SaveGranted(granted)
	if err != nil {
		panic(err)
	}
	return []Granted{granted}
}

func TestRepository_requested_scheme_match_grant_scheme_with_a_wildcard(t *testing.T) {
	// Given
	is, db := setupTestDB(t)

	r := mockRequested(db, Requested{Scheme: "image", RequestScheme: "image"})
	mockGranted(db, Granted{Scheme: "image", GrandScheme: "*"})

	// When
	result, err := db.FindGranted(r)

	// Then
	is.NoErr(err)
	is.Equal(len(result), 1)
}

func mockRequested(db *DbManager, requested Requested) []Requested {
	// Mock implementation to insert requested entry into the database
	return []Requested{requested}
}

// confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/confetti-cms/image/pkg/confetti-cms/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms
