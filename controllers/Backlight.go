package controllers

import "brightnessctl/curves"
import "os"
import "strconv"
import "strings"

type Backlight struct {
	Name    string `json:"name"`
	Current uint8  `json:"current"`
	Min     uint8  `json:"min"`
	Max     uint8  `json:"max"`
}

func NewBacklight(name string) Backlight {

	var backlight Backlight

	backlight.Name = strings.ToLower(name)
	backlight.Current = 0
	backlight.Min = 0
	backlight.Max = 0

	backlight.ReadState()

	return backlight

}

func (backlight *Backlight) Decrease() bool {

	backlight.ReadState()

	delta := uint8(backlight.Max / 100)
	current := uint8(backlight.Current / delta)

	next := curves.DecreaseEaseOut(current)

	backlight.Current = uint8(next * delta)

	return backlight.WriteState()

}

func (backlight *Backlight) Increase() bool {

	backlight.ReadState()

	delta := uint8(backlight.Max / 100)
	current := uint8(backlight.Current / delta)

	next := curves.IncreaseEaseIn(current)

	backlight.Current = uint8(next * delta)

	return backlight.WriteState()

}

func (backlight *Backlight) Status() string {

	backlight.ReadState()

	delta := uint8(backlight.Max / 100)
	current := uint8(backlight.Current / delta)

	return strconv.FormatUint(uint64(current), 10)

}

func (backlight *Backlight) ReadState() bool {

	var result bool = false

	prefix := "/sys/class/backlight/" + backlight.Name

	buf1, err1 := os.ReadFile(prefix + "/brightness")
	buf2, err2 := os.ReadFile(prefix + "/max_brightness")

	if err1 == nil && err2 == nil {

		num1, err11 := strconv.ParseUint(strings.TrimSpace(string(buf1)), 10, 8)
		num2, err21 := strconv.ParseUint(strings.TrimSpace(string(buf2)), 10, 8)

		if err11 == nil && err21 == nil {

			backlight.Current = uint8(num1)
			backlight.Max = uint8(num2)
			result = true

		}

	}

	return result

}

func (backlight *Backlight) WriteState() bool {

	var result bool = false

	prefix := "/sys/class/backlight/" + backlight.Name

	buf := []byte(strconv.FormatUint(uint64(backlight.Current), 10))

	err1 := os.WriteFile(prefix + "/brightness", buf, 0666)

	if err1 == nil {
		result = true
	}

	return result

}
