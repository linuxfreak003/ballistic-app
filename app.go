package main

import (
	"context"

	"github.com/linuxfreak003/ballistic-app/adapters/sql"
	"github.com/linuxfreak003/ballistic-app/domain"
	"github.com/linuxfreak003/ballistic-app/ports"
)

// App struct
type App struct {
	ports.Domain
}

// NewApp creates a new App application struct
func NewApp() *App {
	db, err := sql.NewSQLiteRepository("ballistic.db")
	if err != nil {
		panic(err)
	}
	return &App{
		Domain: domain.New(&domain.DomainConfig{
			EnvironmentRepo: db,
			RifleRepo:       db,
			LoadRepo:        db,
			ScenarioRepo:    db,
		}),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.Start(ctx)
}
