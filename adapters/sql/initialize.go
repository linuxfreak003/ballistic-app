package sql

import (
	"context"
	"database/sql"
)

// InitializeDatabase ...
func InitializeDatabase(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `CREATE TABLE IF NOT EXISTS rifles (
		id INTEGER PRIMARY KEY,
		name TEXT,
		sight_height FLOAT,
		barrel_twist FLOAT,
		twist_direction_left BOOL,
		zero_range FLOAT
	)`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS loads (
		id INTEGER PRIMARY KEY,
		bullet_caliber FLOAT,
		bullet_weight FLOAT,
		bullet_bc_value FLOAT,
		bullet_bc_drag_func INTEGER,
		bullet_length FLOAT,
		muzzle_velocity FLOAT,
		name TEXT
	)`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS environments (
		id INTEGER PRIMARY KEY,
		temperature FLOAT,
		altitude INTEGER,
		pressure FLOAT,
		humidity FLOAT,
		wind_angle FLOAT,
		wind_speed FLOAT,
		pressure_is_absolute BOOL,
		latitude FLOAT,
		azimuth FLOAT,
		name TEXT
	)`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS scenarios (
		id INTEGER PRIMARY KEY,
		name TEXT,
		environment_id INTEGER,
		rifle_id INTEGER,
		load_id INTEGER,
		FOREIGN KEY(environment_id) REFERENCES environments(id),
		FOREIGN KEY(rifle_id) REFERENCES rifles(id),
		FOREIGN KEY(load_id) REFERENCES loads(id)
	)`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
