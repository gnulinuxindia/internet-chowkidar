package main

import (
	"log"
	"os"
	"os/signal"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	utils "github.com/gnulinuxindia/internet-chowkidar/clientutils"
)

func main() {
	// Create Fyne app
	a := app.NewWithID("watch.inet.gui")

	// Apply custom theme
	a.Settings().SetTheme(&CustomTheme{})

	// Get default paths
	conf, data := utils.DefaultStoragePath()

	db, err := utils.FindDatabase(data)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return
	}
	defer db.Close()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() {
		<-done
		db.Close()
		os.Exit(0)
	}()

	// Check if config exists
	_, err = utils.FindConfigData(conf)
	needsSetup := err != nil

	// If setup is needed, show GUI setup wizard
	if needsSetup {
		log.Println("No configuration found, starting setup wizard...")
		err := runSetupWizard(a, conf, 0)
		if err != nil {
			log.Fatalf("Setup failed: %v", err)
		}
	}

	desk, ok := a.(desktop.App)
	// Setup systray if supported
	if !needsSetup && ok {
		setupSystray(desk, conf, db)
	}

	// Run app (keeps running for systray)
	a.Run()
}
