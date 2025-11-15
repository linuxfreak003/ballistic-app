package ports

import "context"

type ScenarioRepository interface {
	CreateScenario(context.Context, *Scenario) (*Scenario, error)
	ListScenarios(context.Context) ([]*Scenario, error)
	GetScenario(context.Context, int64) (*Scenario, error)
}
