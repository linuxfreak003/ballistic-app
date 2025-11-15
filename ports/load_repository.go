package ports

import "context"

type LoadRepository interface {
	CreateLoad(context.Context, *Load) (*Load, error)
	ListLoads(context.Context) ([]*Load, error)
	GetLoad(context.Context, int64) (*Load, error)
}
