package fizzbuzz

import "strconv"

func Call(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "Fizz Buzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

func Range(s int, f int) []string {
	var result []string = make([]string, 0, 0)
	for i := s; i <= f; i++ {
		result = append(result, Call(i))
	}
	return result
}
