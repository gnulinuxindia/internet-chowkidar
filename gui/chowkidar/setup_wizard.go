package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	utils "github.com/gnulinuxindia/internet-chowkidar/clientutils"
	"github.com/tidwall/gjson"
)

func runSetupWizard(a fyne.App, confPath string, clientID int) error {
	w := a.NewWindow("Internet Chowkidar Setup")
	w.Resize(fyne.NewSize(700, 500))
	w.CenterOnScreen()

	vars := utils.Config{}
	currentStep := 0
	totalSteps := 5

	stepLabel := widget.NewLabelWithStyle("Step 1 of 5", fyne.TextAlignCenter, fyne.TextStyle{})

	// Navigation buttons
	backBtn := widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), nil)
	nextBtn := widget.NewButtonWithIcon("Next", theme.NavigateNextIcon(), nil)
	backBtn.Importance = widget.LowImportance
	nextBtn.Importance = widget.HighImportance
	backBtn.Disable()

	statusLabel := widget.NewLabel("")

	// Helper to create card-like container
	makeCard := func(content fyne.CanvasObject) fyne.CanvasObject {
		return container.NewPadded(content)
	}

	// Helper to create section title
	makeTitle := func(text string) fyne.CanvasObject {
		title := canvas.NewText(text, theme.Color(theme.ColorNameForeground))
		title.TextSize = 20
		title.TextStyle = fyne.TextStyle{Bold: true}
		title.Alignment = fyne.TextAlignCenter
		return title
	}

	// Helper to create subtitle
	makeSubtitle := func(text string) fyne.CanvasObject {
		subtitle := widget.NewLabelWithStyle(text, fyne.TextAlignCenter, fyne.TextStyle{})
		return subtitle
	}

	// Step 1: Server configuration
	serverEntry := widget.NewEntry()
	serverEntry.SetPlaceHolder("https://api.inet.watch")
	serverEntry.Text = "https://api.inet.watch"

	step1 := makeCard(container.NewVBox(
		makeTitle("Server Configuration"),
		makeSubtitle("Configure the Internet Chowkidar server to connect to"),
		layout.NewSpacer(),
		widget.NewLabel("Server URL:"),
		serverEntry,
		widget.NewLabel("The default server is https://api.inet.watch"),
		layout.NewSpacer(),
		statusLabel,
	))

	// Step 2: Location
	cityEntry := widget.NewEntry()
	cityEntry.SetPlaceHolder("e.g., Mumbai, London, New York")

	step2 := makeCard(container.NewVBox(
		makeTitle("Location"),
		makeSubtitle("Help us identify your location for ISP reporting"),
		layout.NewSpacer(),
		widget.NewLabel("Your City:"),
		cityEntry,
		widget.NewLabel("We use OpenStreetMap to validate city names."),
		widget.NewLabel("This information helps correlate blocking with geographic regions."),
		layout.NewSpacer(),
	))

	// Step 3: Categories
	var categoryChecks []*widget.Check
	categoryContainer := container.NewVBox()
	categoryScroll := container.NewScroll(categoryContainer)
	categoryScroll.SetMinSize(fyne.NewSize(0, 250))

	selectAllBtn := widget.NewButton("Select All", func() {
		for _, check := range categoryChecks {
			check.Checked = true
			check.Refresh()
		}
	})
	selectAllBtn.Importance = widget.LowImportance

	deselectAllBtn := widget.NewButton("Deselect All", func() {
		for _, check := range categoryChecks {
			check.Checked = false
			check.Refresh()
		}
	})
	deselectAllBtn.Importance = widget.LowImportance

	step3 := makeCard(container.NewVBox(
		makeTitle("Select Categories"),
		makeSubtitle("Choose which types of websites to monitor for blocking"),
		layout.NewSpacer(),
		container.NewHBox(selectAllBtn, deselectAllBtn),
		categoryScroll,
		layout.NewSpacer(),
	))

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
		"Hourly":                         1,
		"4 times a day (every 6 hours)":  6,
		"2 times a day (every 12 hours)": 12,
		"Once a day":                     24,
		"Once a week":                    24 * 7,
	}

	step4 := makeCard(container.NewVBox(
		makeTitle("Check Frequency"),
		makeSubtitle("Choose how often to check for website blocks"),
		layout.NewSpacer(),
		widget.NewLabel("Frequency:"),
		freqSelect,
		layout.NewSpacer(),
		widget.NewCard("", "", widget.NewLabel(
			"More frequent checks provide better data but use more bandwidth and resources.",
		)),
		layout.NewSpacer(),
	))

	// Step 5: Review
	reviewCard := widget.NewCard("Configuration Summary", "", container.NewVBox())

	step5 := makeCard(container.NewVBox(
		makeTitle("Review & Complete"),
		makeSubtitle("Review your settings before completing setup\n You have to restart application to start using Internet Chowkidar"),
		layout.NewSpacer(),
		reviewCard,
		layout.NewSpacer(),
		statusLabel,
	))

	steps := []fyne.CanvasObject{step1, step2, step3, step4, step5}

	// Main content area
	contentArea := container.NewStack(steps[0])

	// Footer with buttons
	footer := container.NewBorder(
		nil, nil,
		backBtn,
		nextBtn,
		stepLabel,
		layout.NewSpacer(),
	)

	// Main layout
	content := container.NewBorder(
		nil,
		container.NewPadded(footer),
		nil, nil,
		contentArea,
	)

	// Update step display
	updateStep := func() {
		contentArea.Objects = []fyne.CanvasObject{steps[currentStep]}
		contentArea.Refresh()

		stepLabel.SetText(fmt.Sprintf("Step %d of %d", currentStep+1, totalSteps))

		backBtn.Enable()
		if currentStep == 0 {
			backBtn.Disable()
		}

		if currentStep == len(steps)-1 {
			nextBtn.SetText("Finish")
			nextBtn.Icon = theme.ConfirmIcon()
		} else {
			nextBtn.SetText("Next")
			nextBtn.Icon = theme.NavigateNextIcon()
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

			statusLabel.SetText("⏳ Validating server...")
			if !utils.ValidateServer(serverURL) {
				statusLabel.SetText("❌ Invalid server URL")
				dialog.ShowError(fmt.Errorf("cannot connect to server"), w)
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
				statusLabel.SetText("❌ Invalid response from server")
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
				if catName == "all" {
					check.Checked = true // Select only ALL category by default
				}
				categoryChecks = append(categoryChecks, check)
				categoryContainer.Add(check)
			}

			categoryContainer.Refresh()
			statusLabel.SetText("✅ Server validated successfully")

		case 1: // City validation
			if cityEntry.Text == "" {
				dialog.ShowError(fmt.Errorf("please enter your city name"), w)
				return
			}

			statusLabel.SetText("⏳ Validating city...")
			city, lat, lon, valid := utils.ValidateCity(cityEntry.Text)
			if !valid {
				statusLabel.SetText("❌ City not found")
				dialog.ShowError(fmt.Errorf("could not find city, please check spelling"), w)
				return
			}

			vars.City = city
			vars.Latitude = lat
			vars.Longitude = lon
			statusLabel.SetText(fmt.Sprintf("✅ City validated: %s", city))

		case 2: // Categories
			vars.CheckCategories = []string{}
			for _, check := range categoryChecks {
				if check.Checked {
					vars.CheckCategories = append(vars.CheckCategories, check.Text)
				}
			}

			if len(vars.CheckCategories) == 0 {
				dialog.ShowInformation("No categories selected",
					"No categories were selected. All categories will be monitored by default.", w)
				vars.CheckCategories = []string{"all"}
			}

		case 3: // Frequency
			vars.TestFrequency = freqValues[freqSelect.Selected]

			// Build review summary
			categoriesStr := ""
			if len(vars.CheckCategories) == 1 && vars.CheckCategories[0] == "all" {
				categoriesStr = "All categories"
			} else if len(vars.CheckCategories) > 4 {
				categoriesStr = fmt.Sprintf("%d categories selected", len(vars.CheckCategories))
			} else {
				categoriesStr = fmt.Sprintf("%v", vars.CheckCategories)
			}

			reviewContent := container.NewVBox(
				widget.NewLabelWithStyle("Server:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel(vars.Server),
				widget.NewLabel(""),
				widget.NewLabelWithStyle("Location:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel(fmt.Sprintf("%s (%.4f, %.4f)", vars.City, vars.Latitude, vars.Longitude)),
				widget.NewLabel(""),
				widget.NewLabelWithStyle("Categories:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel(categoriesStr),
				widget.NewLabel(""),
				widget.NewLabelWithStyle("Frequency:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel(freqSelect.Selected),
			)

			reviewCard.SetContent(reviewContent)

		case 4: // Final step - complete setup
			statusLabel.SetText("⏳ Fetching ISP information...")

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
			statusLabel.SetText("⏳ Registering with server...")

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
				dialog.ShowError(fmt.Errorf("failed to prepare registration data: %v", err), w)
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
			if clientID == 0 {
				vars.ClientID = rand.IntN(999999999)
			} else {
				vars.ClientID = clientID
			}

			// Save config
			configData, err := json.Marshal(vars)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to save configuration: %v", err), w)
				return
			}

			err = os.WriteFile(confPath, configData, 0644)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to write configuration file: %v", err), w)
				return
			}

			statusLabel.SetText("✅ Setup complete!")
			dialog.ShowInformation("Setup Complete",
				"Internet Chowkidar is now configured and will begin monitoring.\n\nCheck the system tray for options.", w)
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
