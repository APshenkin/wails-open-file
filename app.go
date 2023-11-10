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
var urlstring = ""
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
		if urlstring != "" {
			println("startup deeplink string", urlstring)
			runtime.EventsEmit(ctx, "customUrlOpened", urlstring)
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

func (a *App) onUrlOpen(url string) {
	urlstring = url
	go runtime.EventsEmit(*wailsContext, "customUrlOpened", urlstring)

	println("user opened deeplink", urlstring)
}

func (a *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	secondInstanceArgs = secondInstanceData.Args

	println("user opened second instance", strings.Join(secondInstanceData.Args, ","))
	println("user opened second from", secondInstanceData.WorkingDirectory)
	runtime.WindowUnminimise(*wailsContext)
	runtime.Show(*wailsContext)
	go runtime.EventsEmit(*wailsContext, "launchArgs", secondInstanceArgs)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type Test struct {
	Foo string `json:"foo"`
}

func (a *App) Second(name string) *Test {
	return &Test{
		Foo: name,
	}
}

func (a *App) Third() Weekday {
	return Sunday
}

type Weekday string

const (
	Sunday    Weekday = "Sunday"
	Monday    Weekday = "Monday"
	Tuesday   Weekday = "Tuesday"
	Wednesday Weekday = "Wednesday"
	Thursday  Weekday = "Thursday"
	Friday    Weekday = "Friday"
	Saturday  Weekday = "Saturday"
)

var AllWeekdays = []Weekday{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

func (w Weekday) TSName() string {
	switch w {
	case Sunday:
		return "SUNDAY"
	case Monday:
		return "MONDAY"
	case Tuesday:
		return "TUESDAY"
	case Wednesday:
		return "WEDNESDAY"
	case Thursday:
		return "THURSDAY"
	case Friday:
		return "FRIDAY"
	case Saturday:
		return "SATURDAY"
	default:
		return "???"
	}
}
