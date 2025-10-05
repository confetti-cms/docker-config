package dockerconfig

import (
	"fmt"
)

var internal_name = "confetti-sites-confetti-cms_local_pkg-confetti-cms-image-container_8609-development-cmd"
var environment_name = "local"
var environment_stage = "development"
var target = "cmd"
var umbrella_organization = "confetti-sites"
var umbrella_repository = "confetti-cms"
var container_name = "vendor/confetti-cms/image/container"

var locator = fmt.Sprintf("locator://%s?environment_name=%s&environment_stage=%s&target=%s&umbrella_organization=%s&umbrella_repository=%s&container_name=%s",
	internal_name, environment_name, environment_stage, target, umbrella_organization, umbrella_repository, container_name)
