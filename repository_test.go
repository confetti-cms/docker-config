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

func TestRepository_matching(t *testing.T) {
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
			requested:     Requested{Action: "read", RequestAction: "read"},
			granted:       Granted{Action: "read", GrandAction: "read"},
			expectedCount: 1,
		},
		{
			name:          "requested action does not match grand action",
			requested:     Requested{Action: "read", RequestAction: "read"},
			granted:       Granted{Action: "read", GrandAction: "write"},
			expectedCount: 0,
		},
		{
			name:          "request action does not match granted action",
			requested:     Requested{Action: "read", RequestAction: "read"},
			granted:       Granted{Action: "write", GrandAction: "read"},
			expectedCount: 0,
		},
		{
			name:          "wildcard in grand action",
			requested:     Requested{Action: "read", RequestAction: "read"},
			granted:       Granted{Action: "read", GrandAction: "*"},
			expectedCount: 1,
		},
		{
			name:          "wildcard in request action",
			requested:     Requested{Action: "read", RequestAction: "*"},
			granted:       Granted{Action: "read", GrandAction: "read"},
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
