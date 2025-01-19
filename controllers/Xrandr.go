package controllers

type Xrandr struct {
	Name string `json:"name"`
	// TODO
}

func NewXrandr(name string) Xrandr {

	var xrandr Xrandr

	xrandr.Name = name

	return xrandr

}

func (xrandr *Xrandr) Decrease() bool {

	var result bool = false

	// TODO

	return result

}

func (xrandr *Xrandr) Increase() bool {

	var result bool = false

	// TODO

	return result

}

func (xrandr *Xrandr) Status() string {

	var result string = ""

	// TODO

	return result

}
