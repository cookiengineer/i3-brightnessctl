package actions

import "brightnessctl/controllers"

func Decrease(name string, typ controllers.Type) bool {

	var result bool = false

	if typ == controllers.TypeBacklight {

		controller := controllers.NewBacklight(name)
		result = controller.Decrease()

	} else if typ == controllers.TypeXrandr {

		controller := controllers.NewXrandr(name)
		result = controller.Decrease()

	}

	return result

}
