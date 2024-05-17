package main

import (
	"context"
	_ "embed"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx	
}

//go:embed frontend/dist/config.toml
var toml string

func (a *App) InitialConfigRead() string {
	// check if config dir exist, if it doesn't create it
	configPath := configdir.LocalConfig("wibo")
	err := configdir.MakePath(configPath)

	if err != nil {
		panic(err)
	}

	// create path string to config
	configFile := filepath.Join(configPath, "config.toml")

	// check if config already exists, otherwise create a new file
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
    // Create the new config file.
    file, err := os.Create(configFile)

    if err != nil {
			panic(err)
    }
		
		defer file.Close()
  
		_, err = file.WriteString(toml)

		if err != nil {
			file.Close()
			panic(err)
		}		
	} 
	println(configPath)
	return configPath
}

func (a *App) MinimiseApp() {
	runtime.WindowMinimise(a.ctx)
}

func (a *App) CloseApp() {
	runtime.Quit(a.ctx)
}

func (a *App) Copy_To_Clipboard(code string) {
	runtime.ClipboardSetText(a.ctx, code)
}
