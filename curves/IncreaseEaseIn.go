package curves

func IncreaseEaseIn(current uint16) uint16 {

	var result uint16 = 100;

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
