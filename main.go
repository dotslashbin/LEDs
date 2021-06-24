package main

import (
	"fmt"
	"log"
	"net/http"

	pkgLED "github.com/dotslashbin/LEDs/pkg"
	"github.com/dotslashbin/LEDs/service/controller"
)

func main() {
	fmt.Println("Running...")

	led := pkgLED.LED{}
	led.Name = "Led ONE"

	controller := controller.Controller{}
	controller.Init(led)

	fmt.Println("Starting server ...")

	http.HandleFunc("/on", controller.TurnLedOn)
	http.HandleFunc("/off", controller.TurnLedOff)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
