package domain

import (
	"context"
	"fmt"
	rt "runtime"

	"github.com/linuxfreak003/ballistic-app/ports"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Domain struct {
	ctx             context.Context
	EnvironmentRepo ports.EnvironmentRepository
	RifleRepo       ports.RifleRepository
	LoadRepo        ports.LoadRepository
	ScenarioRepo    ports.ScenarioRepository
}

type DomainConfig struct {
	EnvironmentRepo ports.EnvironmentRepository
	RifleRepo       ports.RifleRepository
	LoadRepo        ports.LoadRepository
	ScenarioRepo    ports.ScenarioRepository
}

func New(cfg *DomainConfig) *Domain {
	return &Domain{
		EnvironmentRepo: cfg.EnvironmentRepo,
		RifleRepo:       cfg.RifleRepo,
		LoadRepo:        cfg.LoadRepo,
		ScenarioRepo:    cfg.ScenarioRepo,
	}
}

func (d *Domain) Start(ctx context.Context) {
	d.ctx = ctx
	appMenu := menu.NewMenu()
	if rt.GOOS == "darwin" {
		appMenu.Append(menu.AppMenu())
	}
	fileMenu := appMenu.AddSubmenu("File")
	fileMenu.AddText("Save As...", keys.CmdOrCtrl("s"), func(_ *menu.CallbackData) {
		file, err := runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{
			Title: "Save As",
		})
		if err != nil {
			runtime.LogErrorf(ctx, "save file dialog error: %v", err)
		}
		if file != "" {
			runtime.LogInfof(ctx, "File saved to %s", file)
		}
	})
	fileMenu.AddSeparator()
	fileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(ctx)
	})
	if rt.GOOS == "darwin" {
		appMenu.Append(menu.EditMenu())
	}
	runtime.WindowSetTitle(ctx, "Ballistic Calculator")
	runtime.MenuSetApplicationMenu(ctx, appMenu)
}

func (a *Domain) Greet(name string) string {
	runtime.LogPrint(a.ctx, "Hello "+name)
	selection, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Title:         "Ballistic Calculator",
		Message:       "What ballistic coefficient type are you using?",
		Buttons:       []string{"G1", "G7", "Cancel"},
		DefaultButton: "G7",
		CancelButton:  "Cancel",
	})
	if err != nil {
		runtime.LogErrorf(a.ctx, "message dialog error: %v", err)
	}
	switch selection {
	case "G1":
		return fmt.Sprintf("Hello %s, You selected G1", name)
	case "G7":
		return fmt.Sprintf("Hello %s, You selected G7", name)
	}

	selection, err = runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
	})
	if err != nil {
		runtime.LogErrorf(a.ctx, "open dialog error: %v", err)
		return ""
	}
	return fmt.Sprintf("File: %s", selection)
}

func (d *Domain) CreateEnvironment(env *ports.Environment) (*ports.Environment, error) {
	return d.EnvironmentRepo.CreateEnvironment(d.ctx, env)
}

func (d *Domain) ListEnvironments() ([]*ports.Environment, error) {
	return d.EnvironmentRepo.ListEnvironments(d.ctx)
}

func (d *Domain) CreateRifle(rifle *ports.Rifle) (*ports.Rifle, error) {
	return d.RifleRepo.CreateRifle(d.ctx, rifle)
}

func (d *Domain) ListRifles() ([]*ports.Rifle, error) {
	return d.RifleRepo.ListRifles(d.ctx)
}

func (d *Domain) CreateLoad(load *ports.Load) (*ports.Load, error) {
	return d.LoadRepo.CreateLoad(d.ctx, load)
}

func (d *Domain) ListLoads() ([]*ports.Load, error) {
	return d.LoadRepo.ListLoads(d.ctx)
}

func (d *Domain) CreateScenario(scen *ports.Scenario) (*ports.Scenario, error) {
	return d.ScenarioRepo.CreateScenario(d.ctx, scen)
}

func (d *Domain) ListScenarios() ([]*ports.Scenario, error) {
	return d.ScenarioRepo.ListScenarios(d.ctx)
}
