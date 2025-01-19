package actions

import "brightnessctl/controllers"
import "fmt"

func Status(name string, typ controllers.Type) bool {

	var result bool = false

	if typ == controllers.TypeBacklight {

		controller := controllers.NewBacklight(name)
		percentage := controller.Status()

		if percentage != "" {
			fmt.Println(percentage + "%")
			result = true
		} else {
			fmt.Println("???%")
		}

	} else if typ == controllers.TypeXrandr {

		controller := controllers.NewXrandr(name)
		percentage := controller.Status()

		if percentage != "" {
			fmt.Println(percentage + "%")
			result = true
		} else {
			fmt.Println("???%")
		}

	}

	return result

}
