package reverse

const maxInt = 2147483647

func reverse(x int) int {
	new := 0
	original := 0

	if x < 0 {
		original = x
		x *= -1
	}

	for x > 0 {
		remainder := x % 10 // use modulus to get the last digit // 3 // 2 // 1
		new *= 10           // shift the new digit left each time // 0 // 30 // 320
		new += remainder    // adds remainder to the the counter 3 // 32 // 321
		x /= 10             // this drops off the last digit // 12 // 1 // 0
	}

	if original != 0 {
		new *= -1
	}

	if new > maxInt || new < maxInt*-1 {
		return 0
	}

	return new
}
