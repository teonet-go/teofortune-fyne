package main

import ("github.com/teonet-go/teonet"
"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget")

const (
	appName    = "Teonet fortune golang GUI application"
	appShort   = "teofortunegui"
	appVersion = "0.0.1"
	// echoServer = "dBTgSEHoZ3XXsOqjSkOTINMARqGxHaXIDxl"
	// sendDelay  = 3000
)

func main() {

	// Teonet application logo
	teonet.Logo(appName, appVersion)

	// Connect to Teonet
	newTeonet(appShort)

	// Create and run gui interface
	newGui()

	// select {}
}

// newGui creates and run gui interface
func newGui() {
	a := app.New()
	w := a.NewWindow("Teofortune")

	hello := widget.NewLabel("Hello Teofortune!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}

// newTeonet connects to Teonet and Teofortune server(peer)
func newTeonet(appShort string) (teo *teonet.Teonet, err error) {

	// Start Teonet client
	teo, err = teonet.New(appShort)
	if err != nil {
		panic("can't init Teonet, error: " + err.Error())
	}

	// Connect to Teonet
	err = teo.Connect()
	if err != nil {
		teo.Log().Debug.Println("can't connect to Teonet, error:", err)
		// time.Sleep(1 * time.Second)
		// goto connect
		panic("can't connect to Teonet, error: " + err.Error())
	}

	return
}
