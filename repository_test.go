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

func TestRepository_no_request_does_not_match_none_granted(t *testing.T) {
	// Given
	is, dbManager := setupTestDB(t)

	requested := []Requested{}

	result, err := dbManager.FindGranted(requested)

	// Then
	is.NoErr(err)
	is.Equal(len(result), 0)
}
