package curves

func IncreaseEaseIn(current uint8) uint8 {

	var result uint8 = 100;

	if current == 0 {
		result = 1;
	} else if current >= 0 && current < 10 {
		result = current + 1;
	} else if current >= 10 && current < 90 {
		result = current + 10;
	} else if current >= 90 && current <= 100 {
		result = 100;
	}

	return result

}
