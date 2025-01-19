package detectors

import "slices"
import "sort"

func GetDefaultBacklight() string {

	result := ""

	backlights := GetBacklights()

	if len(backlights) == 1 {

		result = backlights[0]

	} else if len(backlights) > 1 {

		if slices.Contains(backlights, "intel_backlight") {
			result = "intel_backlight"
		} else {

			sort.Strings(backlights)
			result = backlights[0]

		}

	}

	return result

}
