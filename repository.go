package dockerconfig

import (
	"database/sql"
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
	CREATE TABLE IF NOT EXISTS granted_permissions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
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
	);`

	_, err := dm.db.Exec(query)
	return err
}

func (dm *DbManager) SaveGranted(granted Granted) error {
	query := `
	INSERT INTO granted_permissions (
		description, expose_path, scheme, action, source_organization,
		source_repository, umbrella_organization, umbrella_repository,
		container_name, target
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

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
	)

	return err
}

func (dm *DbManager) GetAllGranted() ([]Granted, error) {
	query := `SELECT description, expose_path, scheme, action, source_organization,
		source_repository, umbrella_organization, umbrella_repository,
		container_name, target FROM granted_permissions`

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

func GetGranted(dbManager *DbManager, requested []Requested) []Granted {
	granted, err := dbManager.GetAllGranted()
	if err != nil {
		return []Granted{}
	}
	return granted
}
