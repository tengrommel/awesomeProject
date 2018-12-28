package main

func main() {
	ten := 10
	if ten == 20 {
		println("This shouldn't be printed as 10 isn't equal to 20")
	} else {
		println("Ten is not equals to 20")
	}
	if "a" == "b" || 10 == 10 || true == false {
		println("10 is equal to 10")
	} else if 11 == 11 && "go" == "go" {
		println("This isn't print because previous condition was satisfied")
	} else {
		println("In case no condition is satisfied, print this")
	}
}
