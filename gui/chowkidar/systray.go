package main

//go:generate fyne bundle -o bundled.go resources/systray.png

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	utils "github.com/gnulinuxindia/internet-chowkidar/clientutils"
	"go.mills.io/bitcask/v2"
)

var updateConf = false
var stopSync = false

func setupSystray(desk desktop.App, configPath string, db *bitcask.Bitcask) {
	app := desk.(fyne.App)

	// autoChecks has to be declared outside menu since that's the only way label of item can be changed.
	// The action function of autoChecks is declared after menu
	var autoChecks *fyne.MenuItem
	autoChecks = fyne.NewMenuItem("Stop automated checks", nil)
	// Create menu
	menu := fyne.NewMenu("Internet Chowkidar",
		fyne.NewMenuItem("Run manual check", func() {
			app.SendNotification(fyne.NewNotification("Internet Chowkidar", "Running a manual check"))
			err := runManualCheck(configPath, db)
			if err != nil {
				app.SendNotification(fyne.NewNotification("Internet Chowkidar", "Could not run manual check: "+err.Error()))
			}
			app.SendNotification(fyne.NewNotification("Internet Chowkidar", "Finished Running Manual Check"))
		}),
		autoChecks,
		fyne.NewMenuItem("Re-run Setup", func() {
			rerunSetup(desk, configPath)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Quit", func() {
			app.Quit()
		}),
	)
	autoChecks.Action = func() {
		if stopSync == true {
			autoChecks.Label = "Stop Automated Checks"
			stopSync = false
			app.SendNotification(fyne.NewNotification("Internet Chowkidar", "Restarted all automated checks"))
		} else {
			autoChecks.Label = "Restart Automated Checks"
			stopSync = true
			app.SendNotification(fyne.NewNotification("Internet Chowkidar", "Stopped all automated checks"))
		}
		menu.Refresh()
	}
	desk.SetSystemTrayMenu(menu)
	desk.SetSystemTrayIcon(resourceSystrayPng)

	// Start background daemon
	go func() {
		err := utils.Run(configPath, db, &updateConf, &stopSync)
		if err != nil {
			log.Printf("Daemon error: %v", err)
		}
	}()
}

func runManualCheck(configPath string, db *bitcask.Bitcask) error {
	log.Println("Running manual check...")

	config, err := utils.FindConfigData(configPath)
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return err
	}

	err = utils.FetchAndRun(config, db)
	if err != nil {
		log.Printf("Manual check failed: %v", err)
		return err
	} else {
		log.Println("Manual check completed successfully")
		return nil
	}
}

func rerunSetup(desk desktop.App, configPath string) {
	log.Println("Re-running setup wizard...")
	config, err := utils.FindConfigData(configPath)
	if err != nil {
		log.Printf("Failed to load config: %v", err)
	}

	// Stop daemon temporarily
	initialState := stopSync
	stopSync = true

	err = runSetupWizard(desk.(fyne.App), configPath, config.ClientID)
	if err != nil {
		log.Printf("Setup failed: %v", err)
	} else {
		log.Println("Setup completed successfully")
		updateConf = true
	}

	// Resume daemon if it was running
	stopSync = initialState
}
