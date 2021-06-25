package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	pkgLED "github.com/dotslashbin/LEDs/pkg"
)

type Controller struct {
	light pkgLED.LED
}

type acceptedInput struct {
	Freq string `json:"freq"`
}

var stopchan chan struct{} = make(chan struct{})
var stoppedchan chan struct{} = make(chan struct{})

func (ctrl *Controller) Init(light pkgLED.LED) {
	ctrl.light = light
}

func (ctrl *Controller) TurnLedOn(w http.ResponseWriter, r *http.Request) {
	ctrl.light.TurnOn()
	fmt.Println(ctrl.light.GetStatus())

	close(stopchan)
	<-stoppedchan
}

func (ctrl *Controller) TurnLedOff(w http.ResponseWriter, r *http.Request) {
	ctrl.light.TurnOff()
	fmt.Println(ctrl.light.GetStatus())

	close(stopchan)
	<-stoppedchan
}

func (ctrl *Controller) Blink(w http.ResponseWriter, r *http.Request) {
	var input acceptedInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Fatal("error decoding body ...", err)
	}

	duration, err := time.ParseDuration(input.Freq + "s")
	if err != nil {
		log.Fatal("error in parsing time...", err)
	}

	go func() {

		defer close(stoppedchan)

		for {
			time.Sleep(duration)
			select {
			default:
				fmt.Println(fmt.Sprintf("Blink (%v)", duration))
			case <-stopchan:
				return
			}
		}
	}()
}
