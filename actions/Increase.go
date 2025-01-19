package actions

import "brightnessctl/controllers"

func Increase(name string, typ controllers.Type) bool {

	var result bool = false

	if typ == controllers.TypeBacklight {

		controller := controllers.NewBacklight(name)
		result = controller.Increase()

	} else if typ == controllers.TypeXrandr {

		controller := controllers.NewXrandr(name)
		result = controller.Increase()

	}

	return result

}
