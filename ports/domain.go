package ports

import "context"

type Domain interface {
	Start(ctx context.Context)
	Greet(name string) string

	CreateEnvironment(*Environment) (*Environment, error)
	ListEnvironments() ([]*Environment, error)

	CreateRifle(*Rifle) (*Rifle, error)
	ListRifles() ([]*Rifle, error)

	CreateLoad(*Load) (*Load, error)
	ListLoads() ([]*Load, error)

	CreateScenario(*Scenario) (*Scenario, error)
	ListScenarios() ([]*Scenario, error)
}

// Environment ...
type Environment struct {
	EnvironmentId      int64   `json:"environment_id"`
	Temperature        float64 `json:"temperature"`          // In Fahrenheit
	Altitude           int64   `json:"altitude"`             // in ft
	Pressure           float64 `json:"pressure"`             // in Hg
	Humidity           float64 `json:"humidity"`             // Himidity percentage
	WindAngle          float64 `json:"wind_angle"`           // Wind direction in degrees (0 to 360)
	WindSpeed          float64 `json:"wind_speed"`           // Wind speed in mph
	PressureIsAbsolute bool    `json:"pressure_is_absolute"` // If true, only Pressure will be used
	Latitude           float64 `json:"latitude"`             // Latitude (deg)
	Azimuth            float64 `json:"azimuth"`              // Azimuth of fire (deg)
	Name               string  `json:"name"`
}

// Rifle contains all the variables of the rifle
type Rifle struct {
	RifleId            int64
	Name               string  `json:"name"`
	SightHeight        float64 `json:"sight_height"`         // Sight height in inches (default 1.5)
	BarrelTwist        float64 `json:"barrel_twist"`         // Barrel twist rate (default 7)
	TwistDirectionLeft bool    `json:"twist_direction_left"` // If twist direction is left
	ZeroRange          float64 `json:"zero_range"`           // Zero range of rifle in yrds (default 100)
}

// Load includes relevant data about the load
type Load struct {
	LoadId         int64   `json:"load_id"`
	Bullet         *Bullet `json:"bullet"`
	MuzzleVelocity float64 `json:"muzzle_velocity"`
	Name           string  `json:"name"`
}

// Bullet ...
type Bullet struct {
	Caliber float64      `json:"caliber"`
	Name    string       `json:"name"`
	Weight  float64      `json:"weight"`
	BC      *Coefficient `json:"bc"`
	Length  float64      `json:"length"`
}

// Coefficient ...
type Coefficient struct {
	Value    float64 `json:"value"`
	Name     string  `json:"name"`
	DragFunc int     `json:"drag_func"`
}

type Scenario struct {
	ScenarioId    int64  `json:"scenario_id"`
	Name          string `json:"name"`
	EnvironmentId int64  `json:"environment_id"`
	RifleId       int64  `json:"rifle_id"`
	LoadId        int64  `json:"load_id"`
}
