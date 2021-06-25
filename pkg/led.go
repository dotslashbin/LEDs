package pkgLED

type Light interface {
	TurnOn()
	TurnOff()
	GetStatus() string
}

type LED struct {
	Name  string
	State string
}

func (led *LED) GetStatus() string {
	if led.State == "on" {
		return "LED " + led.Name + " is ON"
	}
	return "LED " + led.Name + " is off"
}

func (led *LED) TurnOn() {
	led.State = "on"
}

func (led *LED) TurnOff() {
	led.State = "off"
}
