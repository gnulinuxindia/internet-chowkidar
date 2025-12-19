package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	utils "github.com/gnulinuxindia/internet-chowkidar/clientutils"
)

var (
	updateConf bool
	stopSync   bool
)

func setupSystray(desk desktop.App, configPath, dbPath string) {
	app := desk.(fyne.App)

	// Create menu
	menu := fyne.NewMenu("Internet Chowkidar",
		fyne.NewMenuItem("Run manual check", func() {
			runManualCheck(configPath, dbPath)
		}),
		fyne.NewMenuItem("Re-run Setup", func() {
			rerunSetup(desk, configPath, dbPath)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Quit", func() {
			app.Quit()
		}),
	)

	desk.SetSystemTrayMenu(menu)

	db, err := utils.FindDatabase(dbPath)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return
	}
	defer db.Close()

	// Start background daemon
	go func() {
		err := utils.Run(configPath, db, &updateConf, &stopSync)
		if err != nil {
			log.Printf("Daemon error: %v", err)
		}
	}()
}

func runManualCheck(configPath, dbPath string) {
	log.Println("Running manual check...")

	config, err := utils.FindConfigData(configPath)
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return
	}

	db, err := utils.FindDatabase(dbPath)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return
	}
	defer db.Close()

	err = utils.FetchAndRun(config, db)
	if err != nil {
		log.Printf("Manual check failed: %v", err)
	} else {
		log.Println("Manual check completed successfully")
	}
}

func rerunSetup(desk desktop.App, configPath, dbPath string) {
	log.Println("Re-running setup wizard...")

	// Stop daemon temporarily
	wasRunning := !stopSync
	stopSync = true

	err := runSetupWizard(desk.(fyne.App), configPath, dbPath)
	if err != nil {
		log.Printf("Setup failed: %v", err)
	} else {
		log.Println("Setup completed successfully")
		updateConf = true
	}

	// Resume daemon if it was running
	if wasRunning {
		stopSync = false
	}
}
