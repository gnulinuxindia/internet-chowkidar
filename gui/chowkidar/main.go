package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	utils "github.com/gnulinuxindia/internet-chowkidar/clientutils"
)

func main() {
	// Create Fyne app
	a := app.NewWithID("watch.inet.gui")

	// Get default paths
	conf, data := utils.DefaultStoragePath()

	// Check if config exists
	_, err := utils.FindConfigData(conf)
	needsSetup := err != nil

	// If setup is needed, show GUI setup wizard
	if needsSetup {
		log.Println("No configuration found, starting setup wizard...")
		err := runSetupWizard(a, conf, data)
		if err != nil {
			log.Fatalf("Setup failed: %v", err)
		}
	} else {
		log.Println("Configuration found, showing welcome screen...")
		err := runSetupDone(a, conf, data)
		if err != nil {
			log.Printf("Failed to show welcome screen: %v", err)
		}
	}

	// Setup systray if supported
	if desk, ok := a.(desktop.App); ok {
		setupSystray(desk, conf, data)
	}

	// Run app (keeps running for systray)
	a.Run()
}
