package controllers

import "brightnessctl/curves"
import "os/exec"
import "strconv"
import "strings"

type Xrandr struct {
	Name    string  `json:"name"`
	Current float64 `json:"current"` // from 0.0 to 1.0
}

func NewXrandr(name string) Xrandr {

	var xrandr Xrandr

	xrandr.Name = name

	return xrandr

}

func (xrandr *Xrandr) Decrease() bool {

	xrandr.ReadState()

	current := uint16(xrandr.Current * 100)
	next := curves.DecreaseEaseOut(current)

	xrandr.Current = float64(float64(next) / 100)

	return xrandr.WriteState()

}

func (xrandr *Xrandr) Increase() bool {

	xrandr.ReadState()

	current := uint16(xrandr.Current * 100)
	next := curves.IncreaseEaseIn(current)

	xrandr.Current = float64(float64(next) / 100)

	return xrandr.WriteState()

}

func (xrandr *Xrandr) Status() string {

	xrandr.ReadState()

	current := uint16(xrandr.Current * 100)

	return strconv.FormatUint(uint64(current), 10)

}

func (xrandr *Xrandr) ReadState() bool {

	var result bool = false

	var out strings.Builder

	cmd := exec.Command("xrandr", "--verbose")
	cmd.Stdout = &out

	err0 := cmd.Run()

	if err0 == nil {

		lines := strings.Split(strings.TrimSpace(out.String()), "\n")
		active_section := false

		for l := 0; l < len(lines); l++ {

			line := lines[l]

			if strings.Contains(line, xrandr.Name + " connected ") {
				active_section = true
			} else if strings.Contains(line, " connected ") {
				active_section = false
			} else if strings.Contains(line, " disconnected ") {
				active_section = false
			} else if active_section == true {

				if strings.Contains(line, "Brightness:") {

					tmp := strings.TrimSpace(line[strings.Index(line, "Brightness:")+11:])

					num, err := strconv.ParseFloat(tmp, 64)

					if err == nil {

						if num >= 1 {
							num = float64(1.0)
						} else if num <= 0.01 {
							num = float64(0.01)
						}

						xrandr.Current = num
						result = true
						break

					}

				}

			}

		}

	}

	return result

}

func (xrandr *Xrandr) WriteState() bool {

	var result bool = false

	clamped_float := strconv.FormatFloat(xrandr.Current, 'f', 2, 64)

	cmd := exec.Command("xrandr", "--output", xrandr.Name, "--brightness", clamped_float)
	err := cmd.Run()

	if err == nil {
		result = true
	}

	return result

}
