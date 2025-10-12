package dockerconfig

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type DbManager struct {
	db *sql.DB
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
	// Create the requested table
	query := `
	CREATE TABLE IF NOT EXISTS requested (
		locator TEXT PRIMARY KEY,
		description TEXT,
		expose_path TEXT,
		scheme TEXT,
		action TEXT,
		source_organization TEXT,
		source_repository TEXT,
		umbrella_organization TEXT,
		umbrella_repository TEXT,
		container_name TEXT,
		target TEXT,
		request_scheme TEXT,
		request_action TEXT,
		request_source_organization TEXT,
		request_source_repository TEXT,
		request_umbrella_organization TEXT,
		request_umbrella_repository TEXT,
		request_container_name TEXT,
		request_target TEXT
	);`

	_, err := dm.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create requested table: %w", err)
	}

	// Create the granted table
	query = `
	CREATE TABLE IF NOT EXISTS granted (
		locator TEXT PRIMARY KEY,
		description TEXT,
		expose_path TEXT,
		scheme TEXT,
		action TEXT,
		source_organization TEXT,
		source_repository TEXT,
		umbrella_organization TEXT,
		umbrella_repository TEXT,
		container_name TEXT,
		target TEXT,
		grand_scheme TEXT,
		grand_action TEXT,
		grand_source_organization TEXT,
		grand_source_repository TEXT,
		grand_umbrella_organization TEXT,
		grand_umbrella_repository TEXT,
		grand_container_name TEXT,
		grand_target TEXT
	);`

	_, err = dm.db.Exec(query)

	return err
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
			req.RequestSourceOrg,
			req.RequestSourceRepo,
			req.RequestUmbrellaOrg,
			req.RequestUmbrellaRepo,
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
			req.RequestSourceOrg,
			req.RequestSourceRepo,
			req.RequestUmbrellaOrg,
			req.RequestUmbrellaRepo,
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

func (dm *DbManager) GetAllGranted() ([]Granted, error) {
	query := `SELECT description, expose_path, scheme, action, source_organization,
		source_repository, umbrella_organization, umbrella_repository,
		container_name, target, grand_scheme, grand_action, grand_source_organization,
		grand_source_repository, grand_umbrella_organization, grand_umbrella_repository,
		grand_container_name, grand_target FROM granted`

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
			return nil, err
		}
		granted = append(granted, g)
	}

	return granted, rows.Err()
}

func (dm *DbManager) Close() error {
	return dm.db.Close()
}

// FindGranted finds granted permissions that match the requested permissions using database queries
func (dm *DbManager) FindGranted(requested []Requested) ([]Granted, error) {
	if len(requested) == 0 {
		return []Granted{}, nil
	}

	// Build query to find granted records where both scheme and grand_scheme match
	// We need to handle multiple requested items, so we'll build conditions for each
	var conditions []string
	var args []interface{}

	for _, req := range requested {
		// Each requested item needs both scheme and grand_scheme to match
		// Handle wildcards in either RequestScheme or GrandScheme
		if req.RequestScheme == "*" {
			// If RequestScheme is "*", match any GrandScheme
			conditions = append(conditions, "scheme = ?")
			args = append(args, req.Scheme)
		} else {
			// Check if we need to look up GrandScheme from database to check for wildcard
			// For now, assume we need exact match unless RequestScheme is "*"
			conditions = append(conditions, "(scheme = ? AND (grand_scheme = ? OR grand_scheme = '*'))")
			args = append(args, req.Scheme, req.RequestScheme)
		}
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
