package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var wailsContext *context.Context

var filestring = ""
var filestrings []string
var secondInstanceArgs []string

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

	wailsContext = &ctx

	argsWithoutProg := os.Args[1:]
	// send callback to frontend when it's loaded with customURL
	runtime.EventsOn(ctx, "frontEndLoaded", func(args ...interface{}) {
		if filestring != "" {
			println("startup file", filestring)
			runtime.EventsEmit(ctx, "fileOpened", filestring)
		}
		if len(filestrings) != 0 {
			println("startup files", filestrings)
			runtime.EventsEmit(ctx, "filesOpened", filestrings)
		}
		if len(argsWithoutProg) != 0 {
			println("startup arguments", strings.Join(argsWithoutProg, ","))
			runtime.EventsEmit(ctx, "launchArgs", argsWithoutProg)
		}
	})
}

func (a *App) onFileOpen(filePath string) {
	filestring = filePath
	go runtime.EventsEmit(*wailsContext, "fileOpened", filePath)

	println("user opened associated file", filePath)
}

func (a *App) onFilesOpen(filePaths []string) {
	filestrings = filePaths

	println("user opened associated files", strings.Join(filePaths, ","))
	go runtime.EventsEmit(*wailsContext, "filesOpened", filePaths)
}

func (a *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	secondInstanceArgs = secondInstanceData.Args

	println("user opened second instance", strings.Join(secondInstanceData.Args, ","))
	runtime.WindowUnminimise(*wailsContext)
	runtime.Show(*wailsContext)
	go runtime.EventsEmit(*wailsContext, "launchArgs", secondInstanceArgs)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
