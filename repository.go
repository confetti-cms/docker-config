package dockerconfig

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

type Requested struct {
	Description                 string `json:"description,omitempty"`
	Host                        string `json:"host,omitempty"`
	DestinationPath             string `json:"destination_path,omitempty"`
	Scheme                      string `json:"scheme,omitempty"`
	Action                      string `json:"action,omitempty"`
	SourceOrganization          string `json:"source_organization,omitempty"`
	SourceRepository            string `json:"source_repository,omitempty"`
	UmbrellaOrganization        string `json:"umbrella_organization,omitempty"`
	UmbrellaRepository          string `json:"umbrella_repository,omitempty"`
	ContainerName               string `json:"container_name,omitempty"`
	Target                      string `json:"target,omitempty"`
	RequestScheme               string `json:"request_scheme,omitempty"`
	RequestAction               string `json:"request_action,omitempty"`
	RequestSourceOrganization   string `json:"request_source_organization,omitempty"`
	RequestSourceRepository     string `json:"request_source_repository,omitempty"`
	RequestUmbrellaOrganization string `json:"request_umbrella_organization,omitempty"`
	RequestUmbrellaRepository   string `json:"request_umbrella_repository,omitempty"`
	RequestContainerName        string `json:"request_container_name,omitempty"`
	RequestTarget               string `json:"request_target,omitempty"`
}

type Granted struct {
	Description               string `json:"description,omitempty"`
	Host                      string `json:"host,omitempty"`
	ExposePath                string `json:"expose_path,omitempty"`
	Scheme                    string `json:"scheme,omitempty"`
	Action                    string `json:"action,omitempty"`
	SourceOrganization        string `json:"source_organization,omitempty"`
	SourceRepository          string `json:"source_repository,omitempty"`
	UmbrellaOrganization      string `json:"umbrella_organization,omitempty"`
	UmbrellaRepository        string `json:"umbrella_repository,omitempty"`
	ContainerName             string `json:"container_name,omitempty"`
	Target                    string `json:"target,omitempty"`
	GrandScheme               string `json:"grand_scheme,omitempty"`
	GrandAction               string `json:"grand_action,omitempty"`
	GrandSourceOrganization   string `json:"grand_source_organization,omitempty"`
	GrandSourceRepository     string `json:"grand_source_repository,omitempty"`
	GrandUmbrellaOrganization string `json:"grand_umbrella_organization,omitempty"`
	GrandUmbrellaRepository   string `json:"grand_umbrella_repository,omitempty"`
	GrandContainerName        string `json:"grand_container_name,omitempty"`
	GrandTarget               string `json:"grand_target,omitempty"`
}

func (dm *DbManager) SaveRequested(requested []Requested) error {
	tx, err := dm.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
	INSERT INTO requested (
		locator,
		description,
		expose_path,
		scheme,
		action,
		source_organization,
		source_repository,
		umbrella_organization,
		umbrella_repository,
		container_name,
		target,
		request_scheme,
		request_action,
		request_source_organization,
		request_source_repository,
		request_umbrella_organization,
		request_umbrella_repository,
		request_container_name,
		request_target
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(locator) DO UPDATE SET
		description=excluded.description,
		expose_path=excluded.expose_path,
		scheme=excluded.scheme,
		action=excluded.action,
		source_organization=excluded.source_organization,
		source_repository=excluded.source_repository,
		umbrella_organization=excluded.umbrella_organization,
		umbrella_repository=excluded.umbrella_repository,
		container_name=excluded.container_name,
		target=excluded.target,
		request_scheme=excluded.request_scheme,
		request_action=excluded.request_action,
		request_source_organization=excluded.request_source_organization,
		request_source_repository=excluded.request_source_repository,
		request_umbrella_organization=excluded.request_umbrella_organization,
		request_umbrella_repository=excluded.request_umbrella_repository,
		request_container_name=excluded.request_container_name,
		request_target=excluded.request_target;
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, req := range requested {
		// Compute the locator hash
		data := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
			req.Description,
			req.DestinationPath,
			req.Scheme,
			req.Action,
			req.SourceOrganization,
			req.SourceRepository,
			req.UmbrellaOrganization,
			req.UmbrellaRepository,
			req.ContainerName,
			req.Target,
			req.RequestScheme,
			req.RequestAction,
			req.RequestSourceOrganization,
			req.RequestSourceRepository,
			req.RequestUmbrellaOrganization,
			req.RequestUmbrellaRepository,
			req.RequestContainerName,
			req.RequestTarget,
		)

		locator := hex.EncodeToString([]byte(data))

		_, err := stmt.Exec(
			locator,
			req.Description,
			req.DestinationPath,
			req.Scheme,
			req.Action,
			req.SourceOrganization,
			req.SourceRepository,
			req.UmbrellaOrganization,
			req.UmbrellaRepository,
			req.ContainerName,
			req.Target,
			req.RequestScheme,
			req.RequestAction,
			req.RequestSourceOrganization,
			req.RequestSourceRepository,
			req.RequestUmbrellaOrganization,
			req.RequestUmbrellaRepository,
			req.RequestContainerName,
			req.RequestTarget,
		)
		if err != nil {
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (dm *DbManager) SaveGranted(granted Granted) error {
	// Compute the locator hash
	data := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
		granted.Description,
		granted.ExposePath,
		granted.Scheme,
		granted.Action,
		granted.SourceOrganization,
		granted.SourceRepository,
		granted.UmbrellaOrganization,
		granted.UmbrellaRepository,
		granted.ContainerName,
		granted.Target,
		granted.GrandScheme,
		granted.GrandAction,
		granted.GrandSourceOrganization,
		granted.GrandSourceRepository,
		granted.GrandUmbrellaOrganization,
		granted.GrandUmbrellaRepository,
		granted.GrandContainerName,
		granted.GrandTarget,
	)

	locator := hex.EncodeToString([]byte(data))

	query := `
	INSERT INTO granted (
		locator,
		description,
		expose_path,
		scheme,
		action,
		source_organization,
		source_repository,
		umbrella_organization,
		umbrella_repository,
		container_name,
		target,
		grand_scheme,
		grand_action,
		grand_source_organization,
		grand_source_repository,
		grand_umbrella_organization,
		grand_umbrella_repository,
		grand_container_name,
		grand_target
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	 ON CONFLICT(locator) DO UPDATE SET
	 		description=excluded.description,
			expose_path=excluded.expose_path,
			scheme=excluded.scheme,
			action=excluded.action,
			source_organization=excluded.source_organization,
			source_repository=excluded.source_repository,
			umbrella_organization=excluded.umbrella_organization,
			umbrella_repository=excluded.umbrella_repository,
			container_name=excluded.container_name,
			target=excluded.target,
			grand_scheme=excluded.grand_scheme,
			grand_action=excluded.grand_action,
			grand_source_organization=excluded.grand_source_organization,
			grand_source_repository=excluded.grand_source_repository,
			grand_umbrella_organization=excluded.grand_umbrella_organization,
			grand_umbrella_repository=excluded.grand_umbrella_repository,
			grand_container_name=excluded.grand_container_name,
			grand_target=excluded.grand_target;
	`

	_, err := dm.db.Exec(query,
		locator,
		granted.Description,
		granted.ExposePath,
		granted.Scheme,
		granted.Action,
		granted.SourceOrganization,
		granted.SourceRepository,
		granted.UmbrellaOrganization,
		granted.UmbrellaRepository,
		granted.ContainerName,
		granted.Target,
		granted.GrandScheme,
		granted.GrandAction,
		granted.GrandSourceOrganization,
		granted.GrandSourceRepository,
		granted.GrandUmbrellaOrganization,
		granted.GrandUmbrellaRepository,
		granted.GrandContainerName,
		granted.GrandTarget,
	)

	return err
}

// FindRequested finds requested permissions that match the granted permissions using database queries
func (dm *DbManager) FindRequested(granted []Granted) ([]Requested, error) {
	if len(granted) == 0 {
		return []Requested{}, nil
	}

	// Build query to find requested records where both scheme and action match
	// We need to handle multiple granted items, so we'll build conditions for each
	var conditions []string
	var args []interface{}

	for _, g := range granted {
		// Each granted item needs both scheme and action to match
		// Handle wildcards in either GrandScheme/RequestScheme and GrandAction/RequestAction

		var schemeCondition string
		if g.GrandScheme == "*" {
			// If GrandScheme is "*", match any RequestScheme
			schemeCondition = "scheme = ?"
			args = append(args, g.Scheme)
		} else {
			// Match when scheme equals the granted scheme AND (request_scheme equals the grand_scheme OR request_scheme is "*")
			schemeCondition = "(scheme = ? AND (request_scheme = ? OR request_scheme = '*'))"
			args = append(args, g.Scheme, g.GrandScheme)
		}

		var actionCondition string
		if g.GrandAction == "*" {
			// If GrandAction is "*", match any RequestAction
			actionCondition = "action = ?"
			args = append(args, g.Action)
		} else {
			// Match when action equals the granted action AND (request_action equals the grand_action OR request_action is "*")
			actionCondition = "(action = ? AND (request_action = ? OR request_action = '*'))"
			args = append(args, g.Action, g.GrandAction)
		}

		// Organization and Repository matching conditions
		var sourceOrgCondition string
		if g.GrandSourceOrganization == "*" {
			// If GrandSourceOrg is "*", match any RequestSourceOrganization
			sourceOrgCondition = "source_organization = ?"
			args = append(args, g.SourceOrganization)
		} else {
			// Match when source_organization equals the granted SourceOrganization AND (request_source_organization equals the grand_source_org OR request_source_organization is "*")
			sourceOrgCondition = "(source_organization = ? AND (request_source_organization = ? OR request_source_organization = '*'))"
			args = append(args, g.SourceOrganization, g.GrandSourceOrganization)
		}

		var sourceRepoCondition string
		if g.GrandSourceRepository == "*" {
			// If GrandSourceRepo is "*", match any RequestSourceRepository
			sourceRepoCondition = "source_repository = ?"
			args = append(args, g.SourceRepository)
		} else {
			// Match when source_repository equals the granted SourceRepository AND (request_source_repository equals the grand_source_repo OR request_source_repository is "*")
			sourceRepoCondition = "(source_repository = ? AND (request_source_repository = ? OR request_source_repository = '*'))"
			args = append(args, g.SourceRepository, g.GrandSourceRepository)
		}

		var umbrellaOrgCondition string
		if g.GrandUmbrellaOrganization == "*" {
			// If GrandUmbrellaOrg is "*", match any RequestUmbrellaOrganization
			umbrellaOrgCondition = "umbrella_organization = ?"
			args = append(args, g.UmbrellaOrganization)
		} else {
			// Match when umbrella_organization equals the granted UmbrellaOrganization AND (request_umbrella_organization equals the grand_umbrella_org OR request_umbrella_organization is "*")
			umbrellaOrgCondition = "(umbrella_organization = ? AND (request_umbrella_organization = ? OR request_umbrella_organization = '*'))"
			args = append(args, g.UmbrellaOrganization, g.GrandUmbrellaOrganization)
		}

		var umbrellaRepoCondition string
		if g.GrandUmbrellaRepository == "*" {
			// If GrandUmbrellaRepo is "*", match any RequestUmbrellaRepository
			umbrellaRepoCondition = "umbrella_repository = ?"
			args = append(args, g.UmbrellaRepository)
		} else {
			// Match when umbrella_repository equals the granted UmbrellaRepository AND (request_umbrella_repository equals the grand_umbrella_repo OR request_umbrella_repository is "*")
			umbrellaRepoCondition = "(umbrella_repository = ? AND (request_umbrella_repository = ? OR request_umbrella_repository = '*'))"
			args = append(args, g.UmbrellaRepository, g.GrandUmbrellaRepository)
		}

		var containerNameCondition string
		if g.GrandContainerName == "*" {
			// If GrandContainerName is "*", match any RequestContainerName
			containerNameCondition = "container_name = ?"
			args = append(args, g.ContainerName)
		} else {
			// Match when container_name equals the granted ContainerName AND (request_container_name equals the grand_container_name OR request_container_name is "*")
			containerNameCondition = "(container_name = ? AND (request_container_name = ? OR request_container_name = '*'))"
			args = append(args, g.ContainerName, g.GrandContainerName)
		}

		var targetCondition string
		if g.GrandTarget == "*" {
			// If GrandTarget is "*", match any RequestTarget
			targetCondition = "target = ?"
			args = append(args, g.Target)
		} else {
			// Match when target equals the granted Target AND (request_target equals the grand_target OR request_target is "*")
			targetCondition = "(target = ? AND (request_target = ? OR request_target = '*'))"
			args = append(args, g.Target, g.GrandTarget)
		}

		// Combine all conditions with AND
		conditions = append(conditions, fmt.Sprintf("(%s AND %s AND %s AND %s AND %s AND %s AND %s AND %s)",
			schemeCondition, actionCondition, sourceOrgCondition, sourceRepoCondition, umbrellaOrgCondition, umbrellaRepoCondition, containerNameCondition, targetCondition))
	}

	query := fmt.Sprintf(`SELECT description, expose_path, scheme, action, source_organization,
		source_repository, umbrella_organization, umbrella_repository,
		container_name, target, request_scheme, request_action, request_source_organization,
		request_source_repository, request_umbrella_organization, request_umbrella_repository,
		request_container_name, request_target FROM requested WHERE %s`,
		strings.Join(conditions, " OR "))

	rows, err := dm.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query requested records: %w", err)
	}
	defer rows.Close()

	var requested []Requested
	for rows.Next() {
		var r Requested
		err := rows.Scan(
			&r.Description,
			&r.DestinationPath,
			&r.Scheme,
			&r.Action,
			&r.SourceOrganization,
			&r.SourceRepository,
			&r.UmbrellaOrganization,
			&r.UmbrellaRepository,
			&r.ContainerName,
			&r.Target,
			&r.RequestScheme,
			&r.RequestAction,
			&r.RequestSourceOrganization,
			&r.RequestSourceRepository,
			&r.RequestUmbrellaOrganization,
			&r.RequestUmbrellaRepository,
			&r.RequestContainerName,
			&r.RequestTarget,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan requested record: %w", err)
		}
		requested = append(requested, r)
	}

	return requested, rows.Err()
}

// FindGranted finds granted permissions that match the requested permissions using database queries
func (dm *DbManager) FindGranted(requested []Requested) ([]Granted, error) {
	if len(requested) == 0 {
		return []Granted{}, nil
	}

	// Build query to find granted records where both scheme and action match
	// We need to handle multiple requested items, so we'll build conditions for each
	var conditions []string
	var args []interface{}

	for _, req := range requested {
		// Each requested item needs both scheme and action to match
		// Handle wildcards in either RequestScheme/GrandScheme and RequestAction/GrandAction

		var schemeCondition string
		if req.RequestScheme == "*" {
			// If RequestScheme is "*", match any GrandScheme
			schemeCondition = "scheme = ?"
			args = append(args, req.Scheme)
		} else {
			// Match when scheme equals the requested scheme AND (grand_scheme equals the request_scheme OR grand_scheme is "*")
			schemeCondition = "(scheme = ? AND (grand_scheme = ? OR grand_scheme = '*'))"
			args = append(args, req.Scheme, req.RequestScheme)
		}

		var actionCondition string
		if req.RequestAction == "*" {
			// If RequestAction is "*", match any GrandAction
			actionCondition = "action = ?"
			args = append(args, req.Action)
		} else {
			// Match when action equals the requested action AND (grand_action equals the request_action OR grand_action is "*")
			actionCondition = "(action = ? AND (grand_action = ? OR grand_action = '*'))"
			args = append(args, req.Action, req.RequestAction)
		}

		// Organization and Repository matching conditions
		var sourceOrgCondition string
		if req.RequestSourceOrganization == "*" {
			// If RequestSourceOrg is "*", match any GrandSourceOrganization
			sourceOrgCondition = "source_organization = ?"
			args = append(args, req.SourceOrganization)
		} else {
			// Match when source_organization equals the requested SourceOrganization AND (grand_source_organization equals the request_source_org OR grand_source_organization is "*")
			sourceOrgCondition = "(source_organization = ? AND (grand_source_organization = ? OR grand_source_organization = '*'))"
			args = append(args, req.SourceOrganization, req.RequestSourceOrganization)
		}

		var sourceRepoCondition string
		if req.RequestSourceRepository == "*" {
			// If RequestSourceRepo is "*", match any GrandSourceRepository
			sourceRepoCondition = "source_repository = ?"
			args = append(args, req.SourceRepository)
		} else {
			// Match when source_repository equals the requested SourceRepository AND (grand_source_repository equals the request_source_repo OR grand_source_repository is "*")
			sourceRepoCondition = "(source_repository = ? AND (grand_source_repository = ? OR grand_source_repository = '*'))"
			args = append(args, req.SourceRepository, req.RequestSourceRepository)
		}

		var umbrellaOrgCondition string
		if req.RequestUmbrellaOrganization == "*" {
			// If RequestUmbrellaOrg is "*", match any GrandUmbrellaOrganization
			umbrellaOrgCondition = "umbrella_organization = ?"
			args = append(args, req.UmbrellaOrganization)
		} else {
			// Match when umbrella_organization equals the requested UmbrellaOrganization AND (grand_umbrella_organization equals the request_umbrella_org OR grand_umbrella_organization is "*")
			umbrellaOrgCondition = "(umbrella_organization = ? AND (grand_umbrella_organization = ? OR grand_umbrella_organization = '*'))"
			args = append(args, req.UmbrellaOrganization, req.RequestUmbrellaOrganization)
		}

		var umbrellaRepoCondition string
		if req.RequestUmbrellaRepository == "*" {
			// If RequestUmbrellaRepo is "*", match any GrandUmbrellaRepository
			umbrellaRepoCondition = "umbrella_repository = ?"
			args = append(args, req.UmbrellaRepository)
		} else {
			// Match when umbrella_repository equals the requested UmbrellaRepository AND (grand_umbrella_repository equals the request_umbrella_repo OR grand_umbrella_repository is "*")
			umbrellaRepoCondition = "(umbrella_repository = ? AND (grand_umbrella_repository = ? OR grand_umbrella_repository = '*'))"
			args = append(args, req.UmbrellaRepository, req.RequestUmbrellaRepository)
		}

		var containerNameCondition string
		if req.RequestContainerName == "*" {
			// If RequestContainerName is "*", match any GrandContainerName
			containerNameCondition = "container_name = ?"
			args = append(args, req.ContainerName)
		} else {
			// Match when container_name equals the requested ContainerName AND (grand_container_name equals the request_container_name OR grand_container_name is "*")
			containerNameCondition = "(container_name = ? AND (grand_container_name = ? OR grand_container_name = '*'))"
			args = append(args, req.ContainerName, req.RequestContainerName)
		}

		var targetCondition string
		if req.RequestTarget == "*" {
			// If RequestTarget is "*", match any GrandTarget
			targetCondition = "target = ?"
			args = append(args, req.Target)
		} else {
			// Match when target equals the requested Target AND (grand_target equals the request_target OR grand_target is "*")
			targetCondition = "(target = ? AND (grand_target = ? OR grand_target = '*'))"
			args = append(args, req.Target, req.RequestTarget)
		}

		// Combine all conditions with AND
		conditions = append(conditions, fmt.Sprintf("(%s AND %s AND %s AND %s AND %s AND %s AND %s AND %s)",
			schemeCondition, actionCondition, sourceOrgCondition, sourceRepoCondition, umbrellaOrgCondition, umbrellaRepoCondition, containerNameCondition, targetCondition))
	}

	query := fmt.Sprintf(`SELECT description, expose_path, scheme, action, source_organization,
		source_repository, umbrella_organization, umbrella_repository,
		container_name, target, grand_scheme, grand_action, grand_source_organization,
		grand_source_repository, grand_umbrella_organization, grand_umbrella_repository,
		grand_container_name, grand_target FROM granted WHERE %s`,
		strings.Join(conditions, " OR "))

	rows, err := dm.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query granted records: %w", err)
	}
	defer rows.Close()

	var granted []Granted
	for rows.Next() {
		var g Granted
		err := rows.Scan(
			&g.Description,
			&g.ExposePath,
			&g.Scheme,
			&g.Action,
			&g.SourceOrganization,
			&g.SourceRepository,
			&g.UmbrellaOrganization,
			&g.UmbrellaRepository,
			&g.ContainerName,
			&g.Target,
			&g.GrandScheme,
			&g.GrandAction,
			&g.GrandSourceOrganization,
			&g.GrandSourceRepository,
			&g.GrandUmbrellaOrganization,
			&g.GrandUmbrellaRepository,
			&g.GrandContainerName,
			&g.GrandTarget,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan granted record: %w", err)
		}
		granted = append(granted, g)
	}

	return granted, rows.Err()
}

// FillRequestedByLocator parses a locator string and fills a Requested struct with the extracted values
func FillRequestedByLocator(locator string, requested ...Requested) (Requested, error) {
	// If a pre-filled Requested struct is provided, use it as base
	var result Requested
	if len(requested) > 0 {
		result = requested[0]
	}

	// Parse the URL
	u, err := url.Parse(locator)
	if err != nil {
		return result, fmt.Errorf("invalid locator format: %w", err)
	}

	// Extract host (without leading //)
	result.Host = strings.TrimPrefix(u.Host, "//")

	// Extract container name from path (remove leading /)
	result.ContainerName = strings.TrimPrefix(u.Path, "/")

	// Extract query parameters
	query := u.Query()
	if target := query.Get("target"); target != "" {
		result.Target = target
	}
	if umbrellaOrg := query.Get("umbrella_organization"); umbrellaOrg != "" {
		result.UmbrellaOrganization = umbrellaOrg
		result.SourceOrganization = umbrellaOrg // Based on test expectations
	}
	if umbrellaRepo := query.Get("umbrella_repository"); umbrellaRepo != "" {
		result.UmbrellaRepository = umbrellaRepo
		result.SourceRepository = umbrellaRepo // Based on test expectations
	}

	return result, nil
}
