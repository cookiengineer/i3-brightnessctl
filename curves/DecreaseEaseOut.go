package curves

func DecreaseEaseOut(current uint16) uint16 {

	var result uint16 = 100;

	if current >= 0 && current <= 1 {
		result = 1;
	} else if current > 1 && current <= 10 {
		result = current - 1
	} else if current > 10 && current <= 100 {
		result = current - 10;
	} else {
		result = 100;
	}

	return result

}
