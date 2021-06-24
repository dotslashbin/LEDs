package pkgLED

type Light interface {
	TurnOn()
	TurnOff()
	GetStatus() string
}

type LED struct {
	On   bool
	Name string
}

func (led *LED) GetStatus() string {
	if led.On {
		return "LED " + led.Name + " is ON"
	}
	return "LED " + led.Name + " is off"
}

func (led *LED) TurnOn() {
	led.On = true
}

func (led *LED) TurnOff() {
	led.On = false
}
