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

func TestRepository_no_request_does_not_match_any_granted(t *testing.T) {
	is, dbManager := setupTestDB(t)

	// Given
	requested := []Requested{}

	// When
	result := GetGranted(dbManager, requested)

	// Then
	is.Equal(len(result), 0)
}

func TestRepository_GetGranted_with_db_manager_and_wildcard_target(t *testing.T) {
	// Given
	is, dbManager := setupTestDB(t)

	err := dbManager.SaveGranted(makeGranted("http", "read", "org", "repo", "umbrella_org", "umbrella_repo", "container", "*"))
	is.NoErr(err)

	requested := []Requested{
		{
			Target: "target",
		},
	}

	// When
	result := GetGranted(dbManager, requested)

	// Then
	is.Equal(len(result), 1)
}
