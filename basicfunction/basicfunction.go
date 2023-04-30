package basicfunction

import (
	"errors"
	"fmt"
	"reflect"
)



//Add function with parameters
func Add(a, b int) {
	println(a + b)
}

//Subtract function with return value
func Subtract(a, b int) int {
	return a - b
}

//Learn function with multiple return value
func Learn() (string, string, string) {
	return "Golang", "Basic", "Functions"
}

//Divide function with naming return value
func Divide(a, b float32) (ans float32, err error) {
	if b == 0 {
		err = errors.New("cannot divide by 0")
	} else {
		ans = a / b
	}
	return
}

//Sum function with variadic parameters
func Sum(number ...float32) float32 {
	fmt.Println(reflect.TypeOf(number))
	var total float32
	for _, value := range number {
		total += value
	}
	return total
}

//PassByValue function
func PassByValue(a int) {
	a = 10
}

//PassByReference function
func PassByReference(a *int) {
	*a = 10
}

//Defer function
func Defer() {
	defer print("Learners")
	defer print("all ")
	print("Hello ")
}
