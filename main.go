package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/teonet-go/teonet"
)

const (
	appName    = "Teonet fortune golang GUI application"
	appShort   = "teofortunegui"
	appVersion = "0.0.1"
	teoFortune = "8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF"
)

func main() {

	// Teonet application logo
	teonet.Logo(appName, appVersion)

	// Connect to Teonet and Teofortune server
	teo, err := newTeofortune(appShort, teoFortune)
	if err != nil {
		teo.Log().Debug.Println("can't connect to Teonet, error:", err)
		return
	}

	// Create and run gui interface
	teo.newGui()
}

// teofortune contains teonet data and holds methods to start gui, process 
// teonet connection and teofortune api
type teofortune struct {
	addr           string            // Teofortune address
	client         *teonet.APIClient // Teofortune api client
	*teonet.Teonet                   // Teonet object
}

// newTeofortune connects to Teonet and Teofortune server(peer)
func newTeofortune(appShort, teoFortune string) (teo *teofortune, err error) {

	teo = new(teofortune)
	teo.addr = teoFortune

	// Start Teonet client
	teo.Teonet, err = teonet.New(appShort)
	if err != nil {
		err = fmt.Errorf("can't init Teonet, error: " + err.Error())
		return
	}

	// Connect to Teonet
	err = teo.Connect()
	if err != nil {
		// time.Sleep(1 * time.Second)
		// goto connect
		err = fmt.Errorf("can't connect to Teonet, error: " + err.Error())
		return
	}

	// Connect to teoFortune server(peer)
	if err = teo.ConnectTo(teo.addr); err != nil {
		err = fmt.Errorf("can't connect to 'fortune', error: %s" + err.Error())
		return
	}

	// Connet to fortune api
	if teo.client, err = teo.NewAPIClient(teo.addr); err != nil {
		err = fmt.Errorf("can't connect to 'fortune' api, error: %s", err.Error())
		return
	}

	return
}

// newGui creates and run gui interface
func (teo *teofortune) newGui() {
	a := app.New()
	w := a.NewWindow("Teofortune")

	label := widget.NewLabel("Fortune message from Teofortune server:")
	fmsg, _ := teo.fortune()
	message := widget.NewLabel(fmsg)
	w.SetContent(container.NewVBox(
		label,
		message,
		widget.NewButton("Show next", func() {
			fmsg, _ := teo.fortune()
			message.SetText(fmsg)
		}),
	))

	w.Resize(fyne.Size{Width: 600, Height: 600})
	w.ShowAndRun()
}

// fortune gets fortune messsage from teofortune microservice
func (teo *teofortune) fortune() (msg string, err error) {

	// Get fortune message from teofortune microservice
	id, err := teo.client.SendTo("fortb", nil)
	if err != nil {
		return
	}
	data, err := teo.WaitFrom(teo.addr, uint32(id))
	if err != nil {
		return
	}

	msg = string(data)
	return
}
