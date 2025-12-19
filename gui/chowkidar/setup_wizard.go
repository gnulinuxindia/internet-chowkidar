package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	utils "github.com/gnulinuxindia/internet-chowkidar/clientutils"
	"github.com/tidwall/gjson"
)

func runSetupWizard(a fyne.App, confPath, dataPath string) error {
	w := a.NewWindow("Internet Chowkidar Setup")
	w.Resize(fyne.NewSize(600, 400))

	vars := utils.Config{}
	currentStep := 0

	// Step containers
	var steps []fyne.CanvasObject

	// Navigation buttons
	backBtn := widget.NewButton("Back", nil)
	nextBtn := widget.NewButton("Next", nil)
	backBtn.Disable()

	statusLabel := widget.NewLabel("")

	// Step 1: Server configuration
	serverEntry := widget.NewEntry()
	serverEntry.SetPlaceHolder("https://api.inet.watch")
	serverEntry.Text = "https://api.inet.watch"

	step1 := container.NewVBox(
		widget.NewLabelWithStyle("Step 1: Server Configuration", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("Enter the Internet Chowkidar server URL:"),
		serverEntry,
		widget.NewLabel(""),
		statusLabel,
	)

	// Step 2: Location
	cityEntry := widget.NewEntry()
	cityEntry.SetPlaceHolder("Enter your city name")

	step2 := container.NewVBox(
		widget.NewLabelWithStyle("Step 2: Location", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("Enter your city:"),
		cityEntry,
		widget.NewLabel("This helps identify your location for ISP reporting."),
	)

	// Step 3: Categories (will be populated after server validation)
	var categoryChecks []*widget.Check
	categoryContainer := container.NewVBox()
	categoryScroll := container.NewScroll(categoryContainer)
	categoryScroll.SetMinSize(fyne.NewSize(0, 200))

	step3 := container.NewVBox(
		widget.NewLabelWithStyle("Step 3: Select Categories", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("Choose which categories of sites to monitor:"),
		categoryScroll,
	)

	// Step 4: Frequency
	freqSelect := widget.NewSelect([]string{
		"Hourly",
		"4 times a day (every 6 hours)",
		"2 times a day (every 12 hours)",
		"Once a day",
		"Once a week",
	}, func(string) {})
	freqSelect.SetSelected("Once a day")

	freqValues := map[string]int{
		"Hourly":                           1,
		"4 times a day (every 6 hours)":    6,
		"2 times a day (every 12 hours)":   12,
		"Once a day":                       24,
		"Once a week":                      24 * 7,
	}

	step4 := container.NewVBox(
		widget.NewLabelWithStyle("Step 4: Check Frequency", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("How often should Internet Chowkidar check for blocks?"),
		freqSelect,
		widget.NewLabel(""),
		widget.NewLabel("More frequent checks provide better data but use more resources."),
	)

	// Step 5: Review and Complete
	reviewLabel := widget.NewLabel("")

	step5 := container.NewVBox(
		widget.NewLabelWithStyle("Step 5: Review & Complete", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("Review your configuration:"),
		reviewLabel,
		widget.NewLabel(""),
		statusLabel,
	)

	steps = []fyne.CanvasObject{step1, step2, step3, step4, step5}

	// Content container
	content := container.NewBorder(
		nil,
		container.NewHBox(backBtn, nextBtn),
		nil,
		nil,
		steps[0],
	)

	// Update step display
	updateStep := func() {
		content.Objects[0] = steps[currentStep]
		content.Refresh()

		backBtn.Enable()
		if currentStep == 0 {
			backBtn.Disable()
		}

		if currentStep == len(steps)-1 {
			nextBtn.SetText("Finish")
		} else {
			nextBtn.SetText("Next")
		}

		statusLabel.SetText("")
	}

	// Back button handler
	backBtn.OnTapped = func() {
		if currentStep > 0 {
			currentStep--
			updateStep()
		}
	}

	// Next button handler
	nextBtn.OnTapped = func() {
		// Validate current step
		switch currentStep {
		case 0: // Server validation
			serverURL := serverEntry.Text
			if serverURL == "" {
				serverURL = "https://api.inet.watch"
			}

			statusLabel.SetText("Validating server...")
			if !utils.ValidateServer(serverURL) {
				statusLabel.SetText("❌ Invalid server URL")
				dialog.ShowError(fmt.Errorf("invalid server URL"), w)
				return
			}

			vars.Server = serverURL
			statusLabel.SetText("✅ Server validated. Loading categories...")

			// Fetch categories
			categoriesOut, err := utils.GetRequest(vars.Server + "/categories")
			if err != nil {
				statusLabel.SetText("❌ Failed to load categories")
				dialog.ShowError(fmt.Errorf("failed to fetch categories: %v", err), w)
				return
			}

			if !gjson.Valid(categoriesOut) {
				statusLabel.SetText("❌ Invalid categories response")
				dialog.ShowError(fmt.Errorf("invalid categories response"), w)
				return
			}

			// Populate categories
			gjsonArr := gjson.Get(categoriesOut, "#.name").Array()
			categoryContainer.Objects = nil
			categoryChecks = nil

			for _, cat := range gjsonArr {
				catName := cat.String()
				check := widget.NewCheck(catName, nil)
				categoryChecks = append(categoryChecks, check)
				categoryContainer.Add(check)
			}

			categoryContainer.Refresh()
			statusLabel.SetText("✅ Ready to continue")

		case 1: // City validation
			if cityEntry.Text == "" {
				dialog.ShowError(fmt.Errorf("please enter your city"), w)
				return
			}

			statusLabel.SetText("Validating city...")
			city, lat, lon, valid := utils.ValidateCity(cityEntry.Text)
			if !valid {
				statusLabel.SetText("❌ Invalid city")
				dialog.ShowError(fmt.Errorf("invalid city name"), w)
				return
			}

			vars.City = city
			vars.Latitude = lat
			vars.Longitude = lon
			statusLabel.SetText("✅ City validated")

		case 2: // Categories
			vars.CheckCategories = []string{}
			for _, check := range categoryChecks {
				if check.Checked {
					vars.CheckCategories = append(vars.CheckCategories, check.Text)
				}
			}

			if len(vars.CheckCategories) == 0 {
				vars.CheckCategories = []string{"all"}
			}

		case 3: // Frequency
			vars.TestFrequency = freqValues[freqSelect.Selected]

			// Update review
			categoriesStr := ""
			if len(vars.CheckCategories) > 3 {
				categoriesStr = fmt.Sprintf("%d categories selected", len(vars.CheckCategories))
			} else {
				categoriesStr = fmt.Sprintf("%v", vars.CheckCategories)
			}

			reviewLabel.SetText(fmt.Sprintf(
				"Server: %s\nCity: %s\nCategories: %s\nFrequency: %s",
				vars.Server,
				vars.City,
				categoriesStr,
				freqSelect.Selected,
			))

		case 4: // Final step - complete setup
			statusLabel.SetText("Fetching ISP information...")

			ipInfoOut, err := utils.GetRequest("https://ipinfo.io")
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to get ISP info: %v", err), w)
				return
			}

			if !gjson.Valid(ipInfoOut) {
				dialog.ShowError(fmt.Errorf("invalid ISP info response"), w)
				return
			}

			vars.ISP = gjson.Get(ipInfoOut, "org").String()

			statusLabel.SetText("Registering with server...")

			type ISPStruct struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Name      string  `json:"name"`
				City      string  `json:"city"`
			}

			ispStruct := ISPStruct{
				Latitude:  vars.Latitude,
				Longitude: vars.Longitude,
				Name:      vars.ISP,
				City:      vars.City,
			}

			data, err := json.Marshal(ispStruct)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to marshal ISP data: %v", err), w)
				return
			}

			ISPOut, err := utils.PostRequest(vars.Server+"/isps", data, "application/json")
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to register with server: %v", err), w)
				return
			}

			if !gjson.Valid(ISPOut) {
				dialog.ShowError(fmt.Errorf("invalid server response"), w)
				return
			}

			vars.ISPID = int(gjson.Get(ISPOut, "id").Int())
			vars.ClientID = rand.IntN(999999999)

			// Save config
			configData, err := json.Marshal(vars)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to save config: %v", err), w)
				return
			}

			err = os.WriteFile(confPath, configData, 0644)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to write config: %v", err), w)
				return
			}

			statusLabel.SetText("✅ Setup complete!")
			dialog.ShowInformation("Setup Complete", "Internet Chowkidar is now configured and will start monitoring.", w)
			w.Close()
			return
		}

		// Move to next step
		if currentStep < len(steps)-1 {
			currentStep++
			updateStep()
		}
	}

	w.SetContent(content)
	w.Show()

	return nil
}

func runSetupDone(a fyne.App, confPath, dataPath string) error {
	w := a.NewWindow("Internet Chowkidar")
	w.Resize(fyne.NewSize(400, 200))

	w.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("Setup Complete", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("Internet Chowkidar is configured and running."),
		widget.NewLabel("The application is now monitoring in the background."),
		widget.NewLabel(""),
		widget.NewLabel("Check the system tray icon for options."),
		widget.NewLabel(""),
		widget.NewButton("Close", func() {
			w.Close()
		}),
	))

	w.CenterOnScreen()
	w.Show()

	return nil
}