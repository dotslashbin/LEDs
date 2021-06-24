package controller

import (
	"fmt"
	"net/http"

	pkgLED "github.com/dotslashbin/LEDs/pkg"
)

type Controller struct {
	light pkgLED.LED
}

func (ctrl *Controller) Init(light pkgLED.LED) {
	ctrl.light = light
}

func (ctrl *Controller) TurnLedOn(w http.ResponseWriter, r *http.Request) {
	ctrl.light.TurnOn()
	fmt.Println(ctrl.light.GetStatus())
}

func (ctrl *Controller) TurnLedOff(w http.ResponseWriter, r *http.Request) {
	ctrl.light.TurnOff()
	fmt.Println(ctrl.light.GetStatus())
}
