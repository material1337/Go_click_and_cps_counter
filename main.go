package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	// Initialize Fyne app
	a := app.New()
	window := a.NewWindow("Click Counter")
	window.Resize(fyne.NewSize(300, 150))

	// Variables
	clicks := 0
	cpsCount := 0
	var cps float64

	// UI elements
	clickLabel := widget.NewLabel("Clicks: 0")
	cpsLabel := widget.NewLabel("CPS: 0.")
	clickButton := widget.NewButton("Click Me", func() {
		clicks++
		cpsCount++
		clickLabel.SetText(fmt.Sprintf("Clicks: %d", clicks))
	})
	resetButton := widget.NewButton("Reset", func() {
		clicks = 0
		cpsCount = 0
		cps = 0
		clickLabel.SetText("Clicks: 0")
		cpsLabel.SetText("CPS: 0.")
	})

	// Layout
	content := container.NewVBox(clickLabel, cpsLabel, clickButton, resetButton)
	window.SetContent(content)

	// CPS calculation
	go func() {
		for {
			startCount := cpsCount
			time.Sleep(time.Second)
			cps = float64(cpsCount - startCount)
			cpsLabel.SetText(fmt.Sprintf("CPS: %.2f", cps))
			cpsCount = 0
		}
	}()

	// Run app
	window.ShowAndRun()
}
