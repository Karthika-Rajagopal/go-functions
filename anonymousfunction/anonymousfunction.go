package anonymousfunction

import "math"


//Calculate function as parameter
func Calculate(a, b int, c func(int, int) int) {
	r := c(a, b)
	println(r)
}

//GetFunction return function from function
func GetFunction() func(string) string {
	return func(s string) string {
		return s
	}
}

//StateFul function
func StateFul() func() int64 {
	a := 0.0
	return func() int64 {
		a += 1
		return int64(math.Pow(a, 2))
	}
}
