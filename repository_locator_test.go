package syncer

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
		RequestScheme:             "docker",
		RequestAction:             "", // Only this field is missing
		RequestSourceOrganization: "provided-org",
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "*")                        // Should get default value
	is.Equal(result.RequestSourceOrganization, "provided-org") // Should keep provided value
}

func TestRepositoryLocator_fill_requested_with_missing_RequestSourceOrganization(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{
		RequestScheme:             "docker",
		RequestAction:             "pull",
		RequestSourceOrganization: "", // Only this field is missing
		RequestSourceRepository:   "provided-repo",
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "different-org") // Should get default value
	is.Equal(result.RequestSourceRepository, "provided-repo")   // Should keep provided value
}

func TestRepositoryLocator_fill_requested_with_missing_RequestSourceRepository(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	requested := Requested{
		RequestScheme:             "docker",
		RequestAction:             "pull",
		RequestSourceOrganization: "provided-org",
		RequestSourceRepository:   "", // Only this field is missing
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "provided-org") // Should keep provided value
	is.Equal(result.RequestSourceRepository, "different-repo") // Should get default value
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
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "provided-org")           // Should keep provided value
	is.Equal(result.RequestSourceRepository, "provided-repo")            // Should keep provided value
	is.Equal(result.RequestUmbrellaOrganization, "confetti-sites")       // Should get default value
	is.Equal(result.RequestUmbrellaRepository, "provided-umbrella-repo") // Should keep provided value
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
	}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.RequestScheme, "docker")
	is.Equal(result.RequestAction, "pull")
	is.Equal(result.RequestSourceOrganization, "provided-org")            // Should keep provided value
	is.Equal(result.RequestSourceRepository, "provided-repo")             // Should keep provided value
	is.Equal(result.RequestUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.RequestUmbrellaRepository, "confetti-cms")            // Should get default value
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

func TestRepositoryLocator_fill_granted_with_empty_granted(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	granted := Granted{}

	// When
	result, err := FillGrantedByLocator(locator, granted)

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
	is.Equal(result.ExposePath, "")
	is.Equal(result.GrandScheme, "")
	is.Equal(result.GrandAction, "*")
	is.Equal(result.GrandSourceOrganization, "different-org")
	is.Equal(result.GrandSourceRepository, "different-repo")
	is.Equal(result.GrandUmbrellaOrganization, "confetti-sites")
	is.Equal(result.GrandUmbrellaRepository, "confetti-cms")
	is.Equal(result.GrandContainerName, "image/container")
	is.Equal(result.GrandTarget, "cmd")
}

func TestRepositoryLocator_fill_granted_with_missing_GrandAction(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	granted := Granted{
		GrandScheme:             "docker",
		GrandAction:             "", // Only this field is missing
		GrandSourceOrganization: "provided-org",
	}

	// When
	result, err := FillGrantedByLocator(locator, granted)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.GrandScheme, "docker")
	is.Equal(result.GrandAction, "*")                        // Should get default value
	is.Equal(result.GrandSourceOrganization, "provided-org") // Should keep provided value
}

func TestRepositoryLocator_fill_granted_with_missing_GrandSourceOrganization(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	granted := Granted{
		GrandScheme:             "docker",
		GrandAction:             "pull",
		GrandSourceOrganization: "", // Only this field is missing
		GrandSourceRepository:   "provided-repo",
	}

	// When
	result, err := FillGrantedByLocator(locator, granted)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.GrandScheme, "docker")
	is.Equal(result.GrandAction, "pull")
	is.Equal(result.GrandSourceOrganization, "different-org") // Should get default value
	is.Equal(result.GrandSourceRepository, "provided-repo")   // Should keep provided value
}

func TestRepositoryLocator_fill_granted_with_missing_GrandSourceRepository(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	granted := Granted{
		GrandScheme:             "docker",
		GrandAction:             "pull",
		GrandSourceOrganization: "provided-org",
		GrandSourceRepository:   "", // Only this field is missing
	}

	// When
	result, err := FillGrantedByLocator(locator, granted)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.GrandScheme, "docker")
	is.Equal(result.GrandAction, "pull")
	is.Equal(result.GrandSourceOrganization, "provided-org") // Should keep provided value
	is.Equal(result.GrandSourceRepository, "different-repo") // Should get default value
}

func TestRepositoryLocator_fill_granted_with_missing_GrandUmbrellaOrganization(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	granted := Granted{
		GrandScheme:               "docker",
		GrandAction:               "pull",
		GrandSourceOrganization:   "provided-org",
		GrandSourceRepository:     "provided-repo",
		GrandUmbrellaOrganization: "", // Only this field is missing
		GrandUmbrellaRepository:   "provided-umbrella-repo",
	}

	// When
	result, err := FillGrantedByLocator(locator, granted)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.GrandScheme, "docker")
	is.Equal(result.GrandAction, "pull")
	is.Equal(result.GrandSourceOrganization, "provided-org")           // Should keep provided value
	is.Equal(result.GrandSourceRepository, "provided-repo")            // Should keep provided value
	is.Equal(result.GrandUmbrellaOrganization, "confetti-sites")       // Should get default value
	is.Equal(result.GrandUmbrellaRepository, "provided-umbrella-repo") // Should keep provided value
}

func TestRepositoryLocator_fill_granted_with_missing_GrandUmbrellaRepository(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	granted := Granted{
		GrandScheme:               "docker",
		GrandAction:               "pull",
		GrandSourceOrganization:   "provided-org",
		GrandSourceRepository:     "provided-repo",
		GrandUmbrellaOrganization: "provided-umbrella-org",
		GrandUmbrellaRepository:   "", // Only this field is missing
	}

	// When
	result, err := FillGrantedByLocator(locator, granted)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.GrandScheme, "docker")
	is.Equal(result.GrandAction, "pull")
	is.Equal(result.GrandSourceOrganization, "provided-org")            // Should keep provided value
	is.Equal(result.GrandSourceRepository, "provided-repo")             // Should keep provided value
	is.Equal(result.GrandUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.GrandUmbrellaRepository, "confetti-cms")            // Should get default value
}

func TestRepositoryLocator_fill_granted_with_missing_GrandContainerName(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	granted := Granted{
		GrandScheme:               "docker",
		GrandAction:               "pull",
		GrandSourceOrganization:   "provided-org",
		GrandSourceRepository:     "provided-repo",
		GrandUmbrellaOrganization: "provided-umbrella-org",
		GrandUmbrellaRepository:   "provided-umbrella-repo",
		GrandContainerName:        "", // Only this field is missing
		GrandTarget:               "provided-target",
	}

	// When
	result, err := FillGrantedByLocator(locator, granted)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.GrandScheme, "docker")
	is.Equal(result.GrandAction, "pull")
	is.Equal(result.GrandSourceOrganization, "provided-org")            // Should keep provided value
	is.Equal(result.GrandSourceRepository, "provided-repo")             // Should keep provided value
	is.Equal(result.GrandUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.GrandUmbrellaRepository, "provided-umbrella-repo")  // Should keep provided value
	is.Equal(result.GrandContainerName, "image/container")              // Should get default value
	is.Equal(result.GrandTarget, "provided-target")                     // Should keep provided value
}

func TestRepositoryLocator_fill_granted_with_missing_GrandTarget(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=different-org&source_repository=different-repo"
	granted := Granted{
		GrandScheme:               "docker",
		GrandAction:               "pull",
		GrandSourceOrganization:   "provided-org",
		GrandSourceRepository:     "provided-repo",
		GrandUmbrellaOrganization: "provided-umbrella-org",
		GrandUmbrellaRepository:   "provided-umbrella-repo",
		GrandContainerName:        "provided-container",
		GrandTarget:               "", // Only this field is missing
	}

	// When
	result, err := FillGrantedByLocator(locator, granted)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.GrandScheme, "docker")
	is.Equal(result.GrandAction, "pull")
	is.Equal(result.GrandSourceOrganization, "provided-org")            // Should keep provided value
	is.Equal(result.GrandSourceRepository, "provided-repo")             // Should keep provided value
	is.Equal(result.GrandUmbrellaOrganization, "provided-umbrella-org") // Should keep provided value
	is.Equal(result.GrandUmbrellaRepository, "provided-umbrella-repo")  // Should keep provided value
	is.Equal(result.GrandContainerName, "provided-container")           // Should keep provided value
	is.Equal(result.GrandTarget, "cmd")                                 // Should get default value
}

func TestRepositoryLocator_fill_granted_with_filled_granted(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-all_up/image/container?environment_name=local&environment_stage=development&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&source_organization=custom-source-org&source_repository=custom-source-repo&target=all_up"
	granted := Granted{
		Description:               "My description",
		ExposePath:                "/my/path",
		GrandScheme:               "docker",
		GrandAction:               "pull",
		GrandSourceOrganization:   "my-org",
		GrandSourceRepository:     "my-repo",
		GrandUmbrellaOrganization: "my-umbrella-org",
		GrandUmbrellaRepository:   "my-umbrella-repo",
		GrandContainerName:        "my-container",
		GrandTarget:               "all_up",
	}

	// When
	result, err := FillGrantedByLocator(locator, granted)

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
	is.Equal(result.ExposePath, "/my/path")
	is.Equal(result.GrandScheme, "docker")
	is.Equal(result.GrandAction, "pull")
	is.Equal(result.GrandSourceOrganization, "my-org")
	is.Equal(result.GrandSourceRepository, "my-repo")
	is.Equal(result.GrandUmbrellaOrganization, "my-umbrella-org")
	is.Equal(result.GrandUmbrellaRepository, "my-umbrella-repo")
	is.Equal(result.GrandContainerName, "my-container")
	is.Equal(result.GrandTarget, "all_up")
}
