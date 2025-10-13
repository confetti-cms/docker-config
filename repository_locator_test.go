package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func TestRepositoryLocator_fill_requested_with_empty_requested(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
	is.Equal(result.SourceOrganization, "different-org")
	is.Equal(result.SourceRepository, "different-repo")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "cmd")
	is.Equal(result.Description, "")
	is.Equal(result.DestinationPath, "")
	is.Equal(result.RequestScheme, "")
	is.Equal(result.RequestAction, "*")
	is.Equal(result.RequestSourceOrganization, "different-org")
	is.Equal(result.RequestSourceRepository, "different-repo")
	is.Equal(result.RequestUmbrellaOrganization, "confetti-sites")
	is.Equal(result.RequestUmbrellaRepository, "confetti-cms")
	is.Equal(result.RequestContainerName, "image/container")
	is.Equal(result.RequestTarget, "cmd")
}

func TestRepositoryLocator_fill_requested_with_missing_RequestAction(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{
		RequestScheme:               "docker",
		RequestAction:               "", // Only this field is missing
		RequestSourceOrganization:   "provided-org",
		RequestSourceRepository:     "provided-repo",
		RequestUmbrellaOrganization: "provided-umbrella-org",
		RequestUmbrellaRepository:   "provided-umbrella-repo",
		RequestContainerName:        "provided-container",
		RequestTarget:               "provided-target",
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
	is.Equal(result.SourceOrganization, "different-org")
	is.Equal(result.SourceRepository, "different-repo")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "cmd")
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "*")                                   // Should get default value
	is.Equal(result.RequestSourceOrganization, "provided-org")            // Should keep provided value
	is.Equal(result.RequestSourceRepository, "provided-repo")             // Should keep provided value
	is.Equal(result.RequestUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.RequestUmbrellaRepository, "provided-umbrella-repo")  // Should keep provided value
	is.Equal(result.RequestContainerName, "provided-container")           // Should keep provided value
	is.Equal(result.RequestTarget, "provided-target")                     // Should keep provided value
}

func TestRepositoryLocator_fill_requested_with_missing_RequestSourceOrganization(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{
		RequestScheme:               "docker",
		RequestAction:               "pull",
		RequestSourceOrganization:   "", // Only this field is missing
		RequestSourceRepository:     "provided-repo",
		RequestUmbrellaOrganization: "provided-umbrella-org",
		RequestUmbrellaRepository:   "provided-umbrella-repo",
		RequestContainerName:        "provided-container",
		RequestTarget:               "provided-target",
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
	is.Equal(result.SourceOrganization, "different-org")
	is.Equal(result.SourceRepository, "different-repo")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "cmd")
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "different-org")           // Should get default value
	is.Equal(result.RequestSourceRepository, "provided-repo")             // Should keep provided value
	is.Equal(result.RequestUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.RequestUmbrellaRepository, "provided-umbrella-repo")  // Should keep provided value
	is.Equal(result.RequestContainerName, "provided-container")           // Should keep provided value
	is.Equal(result.RequestTarget, "provided-target")                     // Should keep provided value
}

func TestRepositoryLocator_fill_requested_with_missing_RequestSourceRepository(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{
		RequestScheme:               "docker",
		RequestAction:               "pull",
		RequestSourceOrganization:   "provided-org",
		RequestSourceRepository:     "", // Only this field is missing
		RequestUmbrellaOrganization: "provided-umbrella-org",
		RequestUmbrellaRepository:   "provided-umbrella-repo",
		RequestContainerName:        "provided-container",
		RequestTarget:               "provided-target",
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
	is.Equal(result.SourceOrganization, "different-org")
	is.Equal(result.SourceRepository, "different-repo")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "cmd")
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "provided-org")            // Should keep provided value
	is.Equal(result.RequestSourceRepository, "different-repo")            // Should get default value
	is.Equal(result.RequestUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.RequestUmbrellaRepository, "provided-umbrella-repo")  // Should keep provided value
	is.Equal(result.RequestContainerName, "provided-container")           // Should keep provided value
	is.Equal(result.RequestTarget, "provided-target")                     // Should keep provided value
}

func TestRepositoryLocator_fill_requested_with_missing_RequestUmbrellaOrganization(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{
		RequestScheme:               "docker",
		RequestAction:               "pull",
		RequestSourceOrganization:   "provided-org",
		RequestSourceRepository:     "provided-repo",
		RequestUmbrellaOrganization: "", // Only this field is missing
		RequestUmbrellaRepository:   "provided-umbrella-repo",
		RequestContainerName:        "provided-container",
		RequestTarget:               "provided-target",
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
	is.Equal(result.SourceOrganization, "different-org")
	is.Equal(result.SourceRepository, "different-repo")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "cmd")
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "provided-org")           // Should keep provided value
	is.Equal(result.RequestSourceRepository, "provided-repo")            // Should keep provided value
	is.Equal(result.RequestUmbrellaOrganization, "confetti-sites")       // Should get default value
	is.Equal(result.RequestUmbrellaRepository, "provided-umbrella-repo") // Should keep provided value
	is.Equal(result.RequestContainerName, "provided-container")          // Should keep provided value
	is.Equal(result.RequestTarget, "provided-target")                    // Should keep provided value
}

func TestRepositoryLocator_fill_requested_with_missing_RequestUmbrellaRepository(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{
		RequestScheme:               "docker",
		RequestAction:               "pull",
		RequestSourceOrganization:   "provided-org",
		RequestSourceRepository:     "provided-repo",
		RequestUmbrellaOrganization: "provided-umbrella-org",
		RequestUmbrellaRepository:   "", // Only this field is missing
		RequestContainerName:        "provided-container",
		RequestTarget:               "provided-target",
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
	is.Equal(result.SourceOrganization, "different-org")
	is.Equal(result.SourceRepository, "different-repo")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "cmd")
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "provided-org")            // Should keep provided value
	is.Equal(result.RequestSourceRepository, "provided-repo")             // Should keep provided value
	is.Equal(result.RequestUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.RequestUmbrellaRepository, "confetti-cms")            // Should get default value
	is.Equal(result.RequestContainerName, "provided-container")           // Should keep provided value
	is.Equal(result.RequestTarget, "provided-target")                     // Should keep provided value
}

func TestRepositoryLocator_fill_requested_with_missing_RequestContainerName(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{
		RequestScheme:               "docker",
		RequestAction:               "pull",
		RequestSourceOrganization:   "provided-org",
		RequestSourceRepository:     "provided-repo",
		RequestUmbrellaOrganization: "provided-umbrella-org",
		RequestUmbrellaRepository:   "provided-umbrella-repo",
		RequestContainerName:        "", // Only this field is missing
		RequestTarget:               "provided-target",
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
	is.Equal(result.SourceOrganization, "different-org")
	is.Equal(result.SourceRepository, "different-repo")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "cmd")
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "provided-org")            // Should keep provided value
	is.Equal(result.RequestSourceRepository, "provided-repo")             // Should keep provided value
	is.Equal(result.RequestUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.RequestUmbrellaRepository, "provided-umbrella-repo")  // Should keep provided value
	is.Equal(result.RequestContainerName, "image/container")              // Should get default value
	is.Equal(result.RequestTarget, "provided-target")                     // Should keep provided value
}

func TestRepositoryLocator_fill_requested_with_missing_RequestTarget(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{
		RequestScheme:               "docker",
		RequestAction:               "pull",
		RequestSourceOrganization:   "provided-org",
		RequestSourceRepository:     "provided-repo",
		RequestUmbrellaOrganization: "provided-umbrella-org",
		RequestUmbrellaRepository:   "provided-umbrella-repo",
		RequestContainerName:        "provided-container",
		RequestTarget:               "", // Only this field is missing
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
	is.Equal(result.SourceOrganization, "different-org")
	is.Equal(result.SourceRepository, "different-repo")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "cmd")
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "provided-org")            // Should keep provided value
	is.Equal(result.RequestSourceRepository, "provided-repo")             // Should keep provided value
	is.Equal(result.RequestUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.RequestUmbrellaRepository, "provided-umbrella-repo")  // Should keep provided value
	is.Equal(result.RequestContainerName, "provided-container")           // Should keep provided value
	is.Equal(result.RequestTarget, "cmd")                                 // Should get default value
}

func TestRepositoryLocator_fill_requested_with_filled_requested(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-all_up/image/container?environment_name=local&environment_stage=development&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=custom-source-org&source_repository=custom-source-repo&target=all_up"
	requested := Requested{
		Description:                 "My description",
		DestinationPath:             "/my/path",
		RequestScheme:               "docker",
		RequestAction:               "pull",
		RequestSourceOrganization:   "my-org",
		RequestSourceRepository:     "my-repo",
		RequestUmbrellaOrganization: "my-umbrella-org",
		RequestUmbrellaRepository:   "my-umbrella-repo",
		RequestContainerName:        "my-container",
		RequestTarget:               "all_up",
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-all_up")
	is.Equal(result.SourceOrganization, "custom-source-org")
	is.Equal(result.SourceRepository, "custom-source-repo")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "all_up")
	is.Equal(result.Description, "My description")
	is.Equal(result.DestinationPath, "/my/path")
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "my-org")
	is.Equal(result.RequestSourceRepository, "my-repo")
	is.Equal(result.RequestUmbrellaOrganization, "my-umbrella-org")
	is.Equal(result.RequestUmbrellaRepository, "my-umbrella-repo")
	is.Equal(result.RequestContainerName, "my-container")
	is.Equal(result.RequestTarget, "all_up")
}
