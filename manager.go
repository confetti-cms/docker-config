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

func (dm *DbManager) Close() error {
	return dm.db.Close()
}
