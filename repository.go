package dockerconfig

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DbManager struct {
	db *sql.DB
}

func NewDbManager() (*DbManager, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	manager := &DbManager{db: db}
	if err := manager.initDB(); err != nil {
		return nil, err
	}

	return manager, nil
}

func (dm *DbManager) initDB() error {
	query := `
	CREATE TABLE IF NOT EXISTS granted (
		locator_hash TEXT GENERATED ALWAYS AS (
			make_locator_hash(
				description,
				expose_path,
				scheme,
				action,
				source_organization,
				source_repository,
				umbrella_organization,
				umbrella_repository,
				container_name,
				target
			)
		) STORED PRIMARY KEY,
		description TEXT,
		expose_path TEXT,
		scheme TEXT,
		action TEXT,
		source_organization TEXT,
		source_repository TEXT,
		umbrella_organization TEXT,
		umbrella_repository TEXT,
		container_name TEXT,
		target TEXT
		grand_scheme TEXT,
		grand_action TEXT,
		grand_source_organization TEXT,
		grand_source_repository TEXT,
		grand_umbrella_organization TEXT,
		grand_umbrella_repository TEXT,
		grand_container_name TEXT,
		grand_target TEXT
	);`

	_, err := dm.db.Exec(query)

	dm.db.RegisterFunc("make_locator_hash", func(
		description, exposePath, scheme, action,
		sourceOrg, sourceRepo, umbrellaOrg, umbrellaRepo,
		containerName, target any,
	) (string, error) {
		data := fmt.Sprint(
			description, exposePath, scheme, action,
			sourceOrg, sourceRepo, umbrellaOrg, umbrellaRepo,
			containerName, target,
		)
		sum := sha256.Sum256([]byte(data))
		return hex.EncodeToString(sum[:]), nil
	}, true)

	return err
}

func (dm *DbManager) SaveGranted(granted Granted) error {
	query := `
	INSERT INTO granted (
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
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	 ON CONFLICT(locator_hash) DO UPDATE SET
	 		description=excluded.description,
			expose_path=excluded.expose_path,
			scheme=excluded.scheme,
			action=excluded.action,
			source_organization=excluded.source_organization,
			source_repository=excluded.source_repository,
			umbrella_organization=excluded.umbrella_organization,
			umbrella_repository=excluded.umbrella_repository,
			container_name=excluded.container_name,
			target=excluded.target
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

func (dm *DbManager) GetAllGranted() ([]Granted, error) {
	query := `SELECT description, expose_path, scheme, action, source_organization,
		source_repository, umbrella_organization, umbrella_repository,
		container_name, target FROM granted`

	rows, err := dm.db.Query(query)
	if err != nil {
		return nil, err
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
		)
		if err != nil {
			return nil, err
		}
		granted = append(granted, g)
	}

	return granted, rows.Err()
}

func (dm *DbManager) Close() error {
	return dm.db.Close()
}

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
	RequestScheme        string `json:"request_scheme,omitempty"`
	RequestAction        string `json:"request_action,omitempty"`
	RequestSourceOrg     string `json:"request_source_organization,omitempty"`
	RequestSourceRepo    string `json:"request_source_repository,omitempty"`
	RequestUmbrellaOrg   string `json:"request_umbrella_organization,omitempty"`
	RequestUmbrellaRepo  string `json:"request_umbrella_repository,omitempty"`
	RequestContainerName string `json:"request_container_name,omitempty"`
	RequestTarget        string `json:"request_target,omitempty"`
}

type Granted struct {
	Description               string `json:"description,omitempty"`
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

// requestedToMap converts a Requested struct to a map[string]string for matching
func requestedToMap(r Requested) map[string]string {
	result := make(map[string]string)

	if r.Description != "" {
		result["description"] = r.Description
	}
	if r.DestinationPath != "" {
		result["destination_path"] = r.DestinationPath
	}
	if r.Scheme != "" {
		result["scheme"] = r.Scheme
	}
	if r.Action != "" {
		result["action"] = r.Action
	}
	if r.SourceOrganization != "" {
		result["source_organization"] = r.SourceOrganization
	}
	if r.SourceRepository != "" {
		result["source_repository"] = r.SourceRepository
	}
	if r.UmbrellaOrganization != "" {
		result["umbrella_organization"] = r.UmbrellaOrganization
	}
	if r.UmbrellaRepository != "" {
		result["umbrella_repository"] = r.UmbrellaRepository
	}
	if r.ContainerName != "" {
		result["container_name"] = r.ContainerName
	}
	if r.Target != "" {
		result["target"] = r.Target
	}
	if r.RequestScheme != "" {
		result["request_scheme"] = r.RequestScheme
	}
	if r.RequestAction != "" {
		result["request_action"] = r.RequestAction
	}
	if r.RequestSourceOrg != "" {
		result["request_source_organization"] = r.RequestSourceOrg
	}
	if r.RequestSourceRepo != "" {
		result["request_source_repository"] = r.RequestSourceRepo
	}
	if r.RequestUmbrellaOrg != "" {
		result["request_umbrella_organization"] = r.RequestUmbrellaOrg
	}
	if r.RequestUmbrellaRepo != "" {
		result["request_umbrella_repository"] = r.RequestUmbrellaRepo
	}
	if r.RequestContainerName != "" {
		result["request_container_name"] = r.RequestContainerName
	}
	if r.RequestTarget != "" {
		result["request_target"] = r.RequestTarget
	}

	return result
}

// grantedToMap converts a Granted struct to a map[string]string for matching
func grantedToMap(g Granted) map[string]string {
	result := make(map[string]string)

	if g.Description != "" {
		result["description"] = g.Description
	}
	if g.ExposePath != "" {
		result["expose_path"] = g.ExposePath
	}
	if g.Scheme != "" {
		result["scheme"] = g.Scheme
	}
	if g.Action != "" {
		result["action"] = g.Action
	}
	if g.SourceOrganization != "" {
		result["source_organization"] = g.SourceOrganization
	}
	if g.SourceRepository != "" {
		result["source_repository"] = g.SourceRepository
	}
	if g.UmbrellaOrganization != "" {
		result["umbrella_organization"] = g.UmbrellaOrganization
	}
	if g.UmbrellaRepository != "" {
		result["umbrella_repository"] = g.UmbrellaRepository
	}
	if g.ContainerName != "" {
		result["container_name"] = g.ContainerName
	}
	if g.Target != "" {
		result["target"] = g.Target
	}
	if g.GrandScheme != "" {
		result["grand_scheme"] = g.GrandScheme
	}
	if g.GrandAction != "" {
		result["grand_action"] = g.GrandAction
	}
	if g.GrandSourceOrganization != "" {
		result["grand_source_organization"] = g.GrandSourceOrganization
	}
	if g.GrandSourceRepository != "" {
		result["grand_source_repository"] = g.GrandSourceRepository
	}
	if g.GrandUmbrellaOrganization != "" {
		result["grand_umbrella_organization"] = g.GrandUmbrellaOrganization
	}
	if g.GrandUmbrellaRepository != "" {
		result["grand_umbrella_repository"] = g.GrandUmbrellaRepository
	}
	if g.GrandContainerName != "" {
		result["grand_container_name"] = g.GrandContainerName
	}
	if g.GrandTarget != "" {
		result["grand_target"] = g.GrandTarget
	}

	return result
}

func GetGranted(dbManager *DbManager, requested []Requested) []Granted {
	granted, err := dbManager.GetAllGranted()
	if err != nil {
		return []Granted{}
	}

	var matching []Granted

	// Convert requested permissions to maps
	var requestedMaps []map[string]string
	for _, req := range requested {
		requestedMaps = append(requestedMaps, requestedToMap(req))
	}

	// Check each granted permission against all requested permissions
	for _, g := range granted {
		grantedMap := grantedToMap(g)

		// Check if this granted permission matches any of the requested permissions
		for _, reqMap := range requestedMaps {
			// Use reverse logic: check if requested fields are satisfied by granted fields
			match := true
			for reqKey, reqValue := range reqMap {
				if grantedValue, exists := grantedMap[reqKey]; exists {
					// Field exists in granted, check if it matches
					if grantedValue == "*" {
						// Wildcard in granted matches any non-empty value in requested
						if reqValue == "" {
							match = false
							break
						}
					} else if grantedValue != reqValue {
						// Exact match required
						match = false
						break
					}
				}
				// If field doesn't exist in granted, that's OK (granted doesn't restrict it)
			}

			if match {
				matching = append(matching, g)
				break // Found a match, no need to check other requested permissions
			}
		}
	}

	return matching
}

func makeGranted(scheme, action, sourceOrg, sourceRepo, umbrellaOrg, umbrellaRepo, containerName, target string) Granted {
	return Granted{
		Description:          "Auto-seeded permission",
		ExposePath:           "/var/timeline",
		Scheme:               scheme,
		Action:               action,
		SourceOrganization:   sourceOrg,
		SourceRepository:     sourceRepo,
		UmbrellaOrganization: umbrellaOrg,
		UmbrellaRepository:   umbrellaRepo,
		ContainerName:        containerName,
		Target:               target,
	}
}
