package ports

import "context"

type RifleRepository interface {
	CreateRifle(context.Context, *Rifle) (*Rifle, error)
	ListRifles(context.Context) ([]*Rifle, error)
	GetRifle(context.Context, int64) (*Rifle, error)
}
