package detectors

import "os/exec"
import "strings"

func GetDefaultMonitor() string {

	result := ""

	var out strings.Builder

	cmd := exec.Command("xrandr", "--query")
	cmd.Stdout = &out

	err := cmd.Run()

	if err == nil {

		monitors := make([]string, 0)
		lines := strings.Split(strings.TrimSpace(out.String()), "\n")

		for l := 0; l < len(lines); l++ {

			line := lines[l]

			if strings.Contains(line, " connected primary ") {

				name := strings.TrimSpace(line[0:strings.Index(line, " connected ")])

				if name != "" {
					monitors = append(monitors, name)
					result = name
				}

			} else if strings.Contains(line, " connected ") {

				name := strings.TrimSpace(line[0:strings.Index(line, " connected ")])

				if name != "" {
					monitors = append(monitors, name)
				}

			} else if strings.Contains(line, " disconnected ") {
				// Ignore disconnected monitors
			} else {
				// Ignore resolutions
			}

		}

		if len(monitors) > 0 && result == "" {
			result = monitors[0]
		}

	}

	return result

}

