package actions

import "brightnessctl/controllers"
import "fmt"

func Decrease(name string, typ controllers.Type) bool {

	var result bool = false

	if typ == controllers.TypeBacklight {

		controller := controllers.NewBacklight(name)
		result = controller.Decrease()

		if result == true {

			percentage := controller.Status()

			if percentage != "" {
				fmt.Println("Decreased brightness to " + percentage + "%")
			} else {
				fmt.Println("Decreased brightness to ???%")
			}

		}

	} else if typ == controllers.TypeXrandr {

		controller := controllers.NewXrandr(name)
		result = controller.Decrease()

		if result == true {

			percentage := controller.Status()

			if percentage != "" {
				fmt.Println("Decreased brightness to " + percentage + "%")
			} else {
				fmt.Println("Decreased brightness to ???%")
			}

		}

	}

	return result

}
