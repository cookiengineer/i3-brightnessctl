package detectors

import "os/exec"
import "strings"

func GetMonitors() []string {

	result := make([]string, 0)

	var out strings.Builder

	cmd := exec.Command("xrandr", "--query")
	cmd.Stdout = &out

	err := cmd.Run()

	if err == nil {

		lines := strings.Split(strings.TrimSpace(out.String()), "\n")

		for l := 0; l < len(lines); l++ {

			line := lines[l]

			if strings.Contains(line, " connected ") {

				name := strings.TrimSpace(line[0:strings.Index(line, " connected ")])

				if name != "" {
					result = append(result, name)
				}

			} else if strings.Contains(line, " disconnected ") {
				// Ignore disconnected monitors
			} else {
				// Ignore resolutions
			}

		}

	}

	return result

}
