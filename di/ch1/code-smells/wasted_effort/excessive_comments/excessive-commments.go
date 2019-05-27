package excessive_comments

import "github.com/tengrommel/awesomeProject/di/ch1/defining-depenency-injection/interface"

// Excessive comments
func outputOrderedPeopleA(in []*defining_depenency_injection.Person)  {
	// This code orders people by name
	// In cases where the name is the same,it will order by phone number.
	// The sort algorithm used is a bubble sort
	// WARNING: this sort will change the items of the input array
	for range in {
		// ... sort code removed ...
	}
	outputPeople(in)
}

// Comments replaced with descriptive names
func outputOrderedPeopleB(in []*defining_depenency_injection.Person)  {
	sortPeople(in)
	outputPeople(in)
}

func outputPeople(in []*defining_depenency_injection.Person)  {
	// TODO: implement
}

// any special instructions that MUST be documented relating to the
func sortPeople(in []*defining_depenency_injection.Person)  {
	// TODO: implement
}

// Person data object
type Person struct {
	Name string
	Phone string
}