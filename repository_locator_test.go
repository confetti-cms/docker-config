package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func TestRepositoryLocator_fill_requested_with_empty_requested(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd/image/container?environment_name=local&environment_stage=development&target=cmd&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms"
	requested := Requested{}

	// When
	result, err := FillRequestedByLocator(locator, requested)

	// Then
	is := is.New(t)
	is.NoErr(err)
	is.Equal(result.Host, "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd")
	is.Equal(result.SourceOrganization, "confetti-sites")
	is.Equal(result.SourceRepository, "confetti-cms")
	is.Equal(result.UmbrellaOrganization, "confetti-sites")
	is.Equal(result.UmbrellaRepository, "confetti-cms")
	is.Equal(result.ContainerName, "image/container")
	is.Equal(result.Target, "cmd")
	is.Equal(result.Description, "")
	is.Equal(result.DestinationPath, "")
	is.Equal(result.RequestScheme, "")
	is.Equal(result.RequestAction, "")
	is.Equal(result.RequestSourceOrganization, "")
	is.Equal(result.RequestSourceRepository, "")
	is.Equal(result.RequestUmbrellaOrganization, "")
	is.Equal(result.RequestUmbrellaRepository, "")
	is.Equal(result.RequestContainerName, "")
	is.Equal(result.RequestTarget, "")
}

func TestRepositoryLocator_fill_requested_with_filled_requested(t *testing.T) {
	// Given
	locator := "//confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-all_up/image/container?environment_name=local&environment_stage=development&umbrella_organization=confetti-sites&umbrella_repository=confetti-cms&target=all_up"
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
	is.Equal(result.SourceOrganization, "confetti-sites")
	is.Equal(result.SourceRepository, "confetti-cms")
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
