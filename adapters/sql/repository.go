package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/linuxfreak003/ballistic-app/ports"

	_ "github.com/lib/pq"           // postgres
	_ "github.com/mattn/go-sqlite3" // sqlite3
)

type repo struct {
	db *sql.DB
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// NewPostgresRepository returns a new postgres repository
func NewPostgresRepository(config *PostgresConfig) (*repo, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.DBName,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := InitializeDatabase(db); err != nil {
		return nil, err
	}

	return &repo{
		db: db,
	}, nil
}

// NewSQLiteRepository ...
func NewSQLiteRepository(filename string) (*repo, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	if err := InitializeDatabase(db); err != nil {
		return nil, err
	}

	return &repo{
		db: db,
	}, nil
}

// CreateLoad ...
func (r *repo) CreateLoad(ctx context.Context, load *ports.Load) (*ports.Load, error) {
	query := `INSERT INTO loads (
		bullet_caliber,
		bullet_weight,
		bullet_bc_value,
		bullet_bc_drag_func,
		bullet_length,
		muzzle_velocity,
		name
	) VALUES (?,?,?,?,?,?,?);`
	res, err := r.db.ExecContext(ctx, query,
		load.Bullet.Caliber,
		load.Bullet.Weight,
		load.Bullet.BC.Value,
		load.Bullet.BC.DragFunc,
		load.Bullet.Length,
		load.MuzzleVelocity,
		load.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	load.LoadId, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}
	return load, nil
}

// ListLoads ...
func (r *repo) ListLoads(ctx context.Context) (loads []*ports.Load, err error) {
	query := `SELECT
		id,
		bullet_caliber,
		bullet_weight,
		bullet_bc_value,
		bullet_bc_drag_func,
		bullet_length,
		muzzle_velocity,
		name
		FROM loads;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		load := &ports.Load{
			Bullet: &ports.Bullet{
				BC: &ports.Coefficient{},
			},
		}
		if err := rows.Scan(
			&load.LoadId,
			&load.Bullet.Caliber,
			&load.Bullet.Weight,
			&load.Bullet.BC.Value,
			&load.Bullet.BC.DragFunc,
			&load.Bullet.Length,
			&load.MuzzleVelocity,
			&load.Name,
		); err != nil {
			return nil, err
		}
		loads = append(loads, load)
	}
	return
}

// GetLoad ...
func (r *repo) GetLoad(ctx context.Context, id int64) (*ports.Load, error) {
	query := `SELECT
		id,
		bullet_caliber,
		bullet_weight,
		bullet_bc_value,
		bullet_bc_drag_func,
		bullet_length,
		muzzle_velocity,
		name
	FROM loads WHERE id = ?;`
	row := r.db.QueryRowContext(ctx, query, id)
	load := &ports.Load{
		Bullet: &ports.Bullet{
			BC: &ports.Coefficient{},
		},
	}
	if err := row.Scan(
		&load.LoadId,
		&load.Bullet.Caliber,
		&load.Bullet.Weight,
		&load.Bullet.BC.Value,
		&load.Bullet.BC.DragFunc,
		&load.Bullet.Length,
		&load.MuzzleVelocity,
		&load.Name,
	); err != nil {
		return nil, fmt.Errorf("could not scan load: %v", err)
	}
	return load, nil
}

// CreateRifle ...
func (r *repo) CreateRifle(ctx context.Context, in *ports.Rifle) (*ports.Rifle, error) {
	query := `INSERT INTO rifles (
		name,
		sight_height,
		barrel_twist,
		twist_direction_left,
		zero_range
	) VALUES (?,?,?,?,?);`
	res, err := r.db.ExecContext(ctx, query,
		in.Name,
		in.SightHeight,
		in.BarrelTwist,
		in.TwistDirectionLeft,
		in.ZeroRange,
	)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	in.RifleId, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}
	return in, nil
}

// ListRifles ...
func (r *repo) ListRifles(ctx context.Context) ([]*ports.Rifle, error) {
	query := `SELECT
		id,
		name,
		sight_height,
		barrel_twist,
		twist_direction_left,
		zero_range
	FROM rifles;`
	rifles := []*ports.Rifle{}
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	for rows.Next() {
		rifle := &ports.Rifle{}
		rows.Scan(
			&rifle.RifleId,
			&rifle.Name,
			&rifle.SightHeight,
			&rifle.BarrelTwist,
			&rifle.TwistDirectionLeft,
			&rifle.ZeroRange,
		)
		rifles = append(rifles, rifle)
	}
	return rifles, nil
}

// GetRifle ...
func (r *repo) GetRifle(ctx context.Context, id int64) (*ports.Rifle, error) {
	query := `SELECT
		id,
		name,
		sight_height,
		barrel_twist,
		twist_direction_left,
		zero_range
	FROM rifles WHERE id = ?;`
	row := r.db.QueryRowContext(ctx, query, id)
	rifle := &ports.Rifle{}
	err := row.Scan(
		&rifle.RifleId,
		&rifle.Name,
		&rifle.SightHeight,
		&rifle.BarrelTwist,
		&rifle.TwistDirectionLeft,
		&rifle.ZeroRange,
	)
	if err != nil {
		return nil, fmt.Errorf("could not scan rifle: %v", err)
	}
	return rifle, nil
}

// CreateEnvironment ...
func (r *repo) CreateEnvironment(ctx context.Context, in *ports.Environment) (*ports.Environment, error) {
	query := `INSERT INTO environments (
		temperature,
		altitude,
		pressure,
		humidity,
		wind_angle,
		wind_speed,
		pressure_is_absolute,
		latitude,
		azimuth,
		name
	) VALUES (?,?,?,?,?,?,?,?,?,?);`
	res, err := r.db.ExecContext(ctx, query,
		in.Temperature,
		in.Altitude,
		in.Pressure,
		in.Humidity,
		in.WindAngle,
		in.WindSpeed,
		in.PressureIsAbsolute,
		in.Latitude,
		in.Azimuth,
		in.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	in.EnvironmentId, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}
	return in, nil
}

// ListEnvironments ...
func (r *repo) ListEnvironments(ctx context.Context) ([]*ports.Environment, error) {
	query := `SELECT
		id,
		temperature,
		altitude,
		pressure,
		humidity,
		wind_angle,
		wind_speed,
		pressure_is_absolute,
		latitude,
		azimuth,
		name
	FROM environments;`
	environments := []*ports.Environment{}
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	for rows.Next() {
		env := &ports.Environment{}
		rows.Scan(
			&env.EnvironmentId,
			&env.Temperature,
			&env.Altitude,
			&env.Pressure,
			&env.Humidity,
			&env.WindAngle,
			&env.WindSpeed,
			&env.PressureIsAbsolute,
			&env.Latitude,
			&env.Azimuth,
			&env.Name,
		)
		environments = append(environments, env)
	}
	return environments, nil
}

// GetEnvironment ...
func (r *repo) GetEnvironment(ctx context.Context, id int64) (*ports.Environment, error) {
	query := `SELECT
		id,
		temperature,
		altitude,
		pressure,
		humidity,
		wind_angle,
		wind_speed,
		pressure_is_absolute,
		latitude,
		azimuth,
		name
	FROM environments WHERE id = ?;`
	row := r.db.QueryRowContext(ctx, query, id)
	env := &ports.Environment{}
	err := row.Scan(
		&env.EnvironmentId,
		&env.Temperature,
		&env.Altitude,
		&env.Pressure,
		&env.Humidity,
		&env.WindAngle,
		&env.WindSpeed,
		&env.PressureIsAbsolute,
		&env.Latitude,
		&env.Azimuth,
		&env.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("could not scan environment: %v", err)
	}
	return env, nil
}

// CreateScenario ...
func (r *repo) CreateScenario(ctx context.Context, in *ports.Scenario) (*ports.Scenario, error) {
	query := `INSERT INTO scenarios (
		name,
		environment_id,
		rifle_id,
		load_id
	) VALUES (?,?,?,?);`
	res, err := r.db.ExecContext(ctx, query,
		in.Name,
		in.EnvironmentId,
		in.RifleId,
		in.LoadId,
	)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	in.ScenarioId, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}
	return in, nil
}

// ListScenarios ...
func (r *repo) ListScenarios(ctx context.Context) ([]*ports.Scenario, error) {
	query := `SELECT
		id,
		name,
		environment_id,
		rifle_id,
		load_id
	FROM scenarios;`
	scenarios := []*ports.Scenario{}
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	for rows.Next() {
		s := &ports.Scenario{}
		rows.Scan(
			&s.ScenarioId,
			&s.Name,
			&s.EnvironmentId,
			&s.RifleId,
			&s.LoadId,
		)
		scenarios = append(scenarios, s)
	}
	return scenarios, nil
}

// GetScenario ...
func (r *repo) GetScenario(ctx context.Context, id int64) (*ports.Scenario, error) {
	query := `SELECT
		id,
		name,
		environment_id,
		rifle_id,
		load_id
	FROM scenarios
	WHERE id == ?;`
	row := r.db.QueryRowContext(ctx, query, id)
	s := &ports.Scenario{}
	if err := row.Scan(
		&s.ScenarioId,
		&s.Name,
		&s.EnvironmentId,
		&s.RifleId,
		&s.LoadId,
	); err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}

	return s, nil
}
