package main

import (
	"log"
	"time"

	"github.com/wailsapp/wails/exp/pkg/application"
	"github.com/wailsapp/wails/exp/pkg/events"
	"github.com/wailsapp/wails/exp/pkg/options"
)

func main() {
	app := application.New(&options.Application{
		Mac: &options.Mac{
			//ActivationPolicy: options.ActivationPolicyAccessory,
		},
	})
	app.On(events.Mac.ApplicationDidFinishLaunching, func() {
		println("WOOOOO!")
	})
	app.On(events.Mac.ApplicationWillTerminate, func() {
		println("TERMINATION!!")
	})
	myWindow := app.NewWindow(&options.Window{
		Title:         "Basic",
		Width:         600,
		Height:        400,
		AlwaysOnTop:   false,
		URL:           "https://google.com",
		DisableResize: false,
	})

	myWindow2 := app.NewWindow(&options.Window{
		Title:         "#2",
		Width:         1024,
		Height:        768,
		AlwaysOnTop:   false,
		URL:           "https://google.com",
		DisableResize: true,
	})

	go func() {
		time.Sleep(5 * time.Second)
		myWindow.SetTitle("Wooooo")
		myWindow.SetAlwaysOnTop(true)
		myWindow2.SetTitle("OMFG")
		myWindow2.NavigateToURL("https://wails.io")
	}()

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
