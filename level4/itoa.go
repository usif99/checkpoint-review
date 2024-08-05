package piscine

func Itoa(n int) string {
	var result string
	isNegative := false

	if n < 0 {
		isNegative = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string(rune(digit+'0')) + result
		n /= 10
	}

	if isNegative {
		result = "-" + result
	}

	if result == "" {
		return "0"
	}

	return result
}