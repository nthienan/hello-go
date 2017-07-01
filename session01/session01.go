package main

import "fmt"
import "math"

func main() {
	aNumber := 8128
	fmt.Printf("checkPrimeNumber(%d) return %t \n", aNumber, checkPrimeNumber(aNumber))
	fmt.Printf("checkPerfectNumber(%d) return %t \n", aNumber, checkPerfectNumber(aNumber))
}

func checkPrimeNumber(aNumber int) bool {
	if aNumber < 2 {
		return false
	} else if aNumber > 2 {
		if aNumber%2 == 0 {
			return false
		}
		for i := 3; i <= int(math.Sqrt(float64(aNumber))); i += 2 {
			if aNumber%i == 0 {
				return false
			}
		}
	}
	return true
}

func checkPerfectNumber(aNumber int) bool {
	//divs := divisors(aNumber)
	//var sum int
	//for _, v := range divs {
	//	if v != aNumber {
	//		sum += v
	//	}
	//}

	i, sum := 1, 0
	for i < aNumber {
		if aNumber%i == 0 {
			sum += i
		}
		i++
	}
	if sum == aNumber {
		return true
	}
	return false
}

func divisors(aNumber int) (divs []int) {
	for i := 1; i <= aNumber; i++ {
		if aNumber%i == 0 {
			divs = append(divs, i)
		}
	}
	return divs
}
