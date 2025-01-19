package actions

import "brightnessctl/controllers"
import "fmt"

func Increase(name string, typ controllers.Type) bool {

	var result bool = false

	if typ == controllers.TypeBacklight {

		controller := controllers.NewBacklight(name)
		result = controller.Increase()

		if result == true {

			percentage := controller.Status()

			if percentage != "" {
				fmt.Println("Increased brightness to " + percentage + "%")
			} else {
				fmt.Println("Increased brightness to ???%")
			}

		}

	} else if typ == controllers.TypeXrandr {

		controller := controllers.NewXrandr(name)
		result = controller.Increase()

		if result == true {

			percentage := controller.Status()

			if percentage != "" {
				fmt.Println("Increased brightness to " + percentage + "%")
			} else {
				fmt.Println("Increased brightness to ???%")
			}

		}

	}

	return result

}
