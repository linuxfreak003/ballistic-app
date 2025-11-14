package domain

import "fmt"

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
