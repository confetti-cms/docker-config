package dockerconfig

type Requested struct {
	Description          string `json:"description,omitempty"`
	DestinationPath      string `json:"destination_path,omitempty"`
	Scheme               string `json:"scheme,omitempty"`
	Action               string `json:"action,omitempty"`
	SourceOrganization   string `json:"source_organization,omitempty"`
	SourceRepository     string `json:"source_repository,omitempty"`
	UmbrellaOrganization string `json:"umbrella_organization,omitempty"`
	UmbrellaRepository   string `json:"umbrella_repository,omitempty"`
	ContainerName        string `json:"container_name,omitempty"`
	Target               string `json:"target,omitempty"`
}

type Granted struct {
	Description          string `json:"description,omitempty"`
	ExposePath           string `json:"expose_path,omitempty"`
	Scheme               string `json:"scheme,omitempty"`
	Action               string `json:"action,omitempty"`
	SourceOrganization   string `json:"source_organization,omitempty"`
	SourceRepository     string `json:"source_repository,omitempty"`
	UmbrellaOrganization string `json:"umbrella_organization,omitempty"`
	UmbrellaRepository   string `json:"umbrella_repository,omitempty"`
	ContainerName        string `json:"container_name,omitempty"`
	Target               string `json:"target,omitempty"`
}

func GetGranted(requested []Requested) []Granted {
	return []Granted{}
}
