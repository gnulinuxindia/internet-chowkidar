package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/getlantern/systray/example/icon"
	"github.com/martinlindhe/notify"
	"go.mills.io/bitcask/v2"

	"github.com/getlantern/systray"
	"github.com/gnulinuxindia/internet-chowkidar/cmd/chowkidar/utils"
	"github.com/urfave/cli/v2"
)

func createSystray(configPath string, db *bitcask.Bitcask) error {
	config, err := utils.FindConfigData(configPath)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Internet Chowkidar")
	systray.SetTooltip("Internet Chowkidar")
	mRun := systray.AddMenuItem("Run manual check", "Force-run a internet-chowkidar check")
	mFrequency := systray.AddMenuItem("Frequency", "Change frequency of the checks")
	menuNames := [6]string{"Hourly", "4 times a day", "2 times a day", "Daily", "Weekly", "Custom"}
	// 0 is for custom frequency
	frequencies := [6]int{1, 6, 12, 24, 24 * 7, 0}
	var menuItems [6]*systray.MenuItem
	alreadySelected := false
	for i, name := range menuNames {
		checked := false
		if frequencies[i] == config.TestFrequency {
			checked = true
			alreadySelected = true
		} else if i == 5 && alreadySelected == false {
			checked = true
		}
		menuItems[i] = mFrequency.AddSubMenuItemCheckbox(name, "", checked)
	}

	freqCh := make(chan *systray.MenuItem)
	for _, fi := range menuItems {
		go func(f *systray.MenuItem) {
			for range f.ClickedCh { // infinite stream
				freqCh <- f
			}
		}(fi)
	}

	mStop := systray.AddMenuItem("Stop automated checks", "Stops internet chowkidar checks")
	mSetup := systray.AddMenuItem("Re-run Setup", "Setup internet-chowkidar again")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() error {
		for {
			select {
			case <-mQuit.ClickedCh:
				os.Exit(0)
			case <-mRun.ClickedCh:
				notify.Notify("Internet Chowkidar", "Running manual check", "Internet Chowkidar is running a manual check based on your request", "")
				fetchAndRun(config, db)
				notify.Notify("Internet Chowkidar", "Finished manual check", "Internet Chowkidar has finished a manual check based on your request", "")
			case <-mSetup.ClickedCh:
				notify.Notify("Internet Chowkidar", "Not implemented yet", "Run chowkidar setup manually", "")
				log.Println("Not IMplemented  Yet")
			case <-mStop.ClickedCh:
				if stopSync {
					stopSync = false
					notify.Notify("Internet Chowkidar", "Restarted sync", "Internet chowkidar has restarted syncing, click to disable again", "")
					mStop.SetTitle("Stop automated checks")
					mStop.SetTooltip("Stops internet chowkidar checks")
				} else {
					stopSync = true
					notify.Notify("Internet Chowkidar", "Stopped sync", "Internet chowkidar has stopped syncing, click to enable again", "")
					mStop.SetTitle("Restart automated checks")
					mStop.SetTooltip("Restart internet chowkidar checks")
				}
			case selected := <-freqCh:
				for i, item := range menuItems {
					if item == selected {
						config.TestFrequency = frequencies[i]
						data, err := json.Marshal(config)
						if err != nil {
							return cli.Exit("Unable to marshal into json: "+err.Error(), 1)
						}

						err = os.WriteFile(configPath, data, 0644)
						if err != nil {
							return cli.Exit("Unable to write config: "+err.Error(), 1)
						}
						updateConf = true
						notify.Notify("Internet Chowkidar", "Check frequency has been changed", "Internet chowkidar has changed testing frequency", "")
					} else {
						item.Uncheck()
					}
				}
			}
		}
		return nil
	}()
	return nil
}
