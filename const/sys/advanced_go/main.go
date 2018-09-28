package main

import "errors"

func division(x, y int) (int, error, error) {
	if y == 0 {
		return 0, nil, errors.New("Cannot divide by zero!")
	}
	if x%y != 0 {
		remainder := errors.New("There is a remainder!")
		return x / y, remainder, nil
	} else {
		return x / y, nil, nil
	}
}

func main() {

}
