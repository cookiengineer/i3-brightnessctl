package main

import "brightnessctl/actions"
import "brightnessctl/controllers"
import "brightnessctl/detectors"
import "fmt"
import "os"

func showHelp() {

	fmt.Println("Usage: brightnessctl [action]")
	fmt.Println("")
	fmt.Println("| Action       | Description                                                 |")
	fmt.Println("|--------------|-------------------------------------------------------------|")
	fmt.Println("| increase-all | increase brightness of all backlights and monitors          |")
	fmt.Println("| decrease-all | decrease brightness of all backlights and monitors          |")
	fmt.Println("| increase     | increase brightness of current backlight or monitor         |")
	fmt.Println("| decrease     | decrease brightness of current backlight or monitor         |")
	fmt.Println("| status       | print brightness percentage of current backlight or monitor |")
	fmt.Println("|--------------|-------------------------------------------------------------|")

}

func main() {

	var action string
	var backlight string
	var monitor string

	if len(os.Args) == 2 {

		tmp := os.Args[1]

		if tmp == "increase-all" {
			action = "increase-all"
		} else if tmp == "decrease-all" {
			action = "decrease-all"
		} else if tmp == "increase" {

			action = "increase"

			backlight = detectors.GetDefaultBacklight()
			monitor = detectors.GetDefaultMonitor()

		} else if tmp == "decrease" {

			action = "decrease"

			backlight = detectors.GetDefaultBacklight()
			monitor = detectors.GetDefaultMonitor()

		} else if tmp == "status" {

			action = "status"

			backlight = detectors.GetDefaultBacklight()
			monitor = detectors.GetDefaultMonitor()

		}

	}

	if action == "increase-all" {

		result_all := true
		backlights := detectors.GetBacklights()

		if len(backlights) > 0 {

			for b := 0; b < len(backlights); b++ {

				result := actions.Increase(backlights[b], controllers.TypeBacklight)

				if result == false {
					result_all = false
				}

			}

		}

		monitors := detectors.GetMonitors()

		if len(monitors) > 0 {

			for m := 0; m < len(monitors); m++ {

				result := actions.Increase(monitors[m], controllers.TypeXrandr)

				if result == false {
					result_all = false
				}

			}

		}

		if result_all == true {
			os.Exit(0)
		} else {
			os.Exit(1)
		}

	} else if action == "decrease-all" {

		result_all := true
		backlights := detectors.GetBacklights()

		if len(backlights) > 0 {

			for b := 0; b < len(backlights); b++ {

				result := actions.Decrease(backlights[b], controllers.TypeBacklight)

				if result == false {
					result_all = false
				}

			}

		}

		monitors := detectors.GetMonitors()

		if len(monitors) > 0 {

			for m := 0; m < len(monitors); m++ {

				result := actions.Decrease(monitors[m], controllers.TypeXrandr)

				if result == false {
					result_all = false
				}

			}

		}

		if result_all == true {
			os.Exit(0)
		} else {
			os.Exit(1)
		}

	} else if action == "increase" {

		if backlight != "" {

			result := actions.Increase(backlight, controllers.TypeBacklight)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else if monitor != "" {

			result := actions.Increase(monitor, controllers.TypeXrandr)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else {
			fmt.Println("Could not detect default backlight or monitor")
			os.Exit(2)
		}

	} else if action == "decrease" {

		if backlight != "" {

			result := actions.Decrease(backlight, controllers.TypeBacklight)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else if monitor != "" {

			result := actions.Decrease(monitor, controllers.TypeXrandr)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else {
			fmt.Println("Could not detect default backlight or monitor")
			os.Exit(2)
		}

	} else if action == "status" {

		if backlight != "" {

			result := actions.Status(backlight, controllers.TypeBacklight)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else if monitor != "" {

			result := actions.Status(monitor, controllers.TypeXrandr)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else {
			fmt.Println("Could not detect default backlight or monitor")
			os.Exit(2)
		}


	} else {
		showHelp()
		os.Exit(1)
	}

}
