package ports

import "context"

type EnvironmentRepository interface {
	CreateEnvironment(context.Context, *Environment) (*Environment, error)
	ListEnvironments(context.Context) ([]*Environment, error)
	GetEnvironment(context.Context, int64) (*Environment, error)
}
