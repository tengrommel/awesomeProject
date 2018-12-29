package main

import "fmt"

type SideInCountry interface {
	GetDrivingSideInCountry() string
}

type CountrySpain struct{}

func (c *CountrySpain) GetDrivingSideInCountry() string {
	return "left"
}

type CountryEngland struct{}

func (c *CountryEngland) GetDrivingSideInCountry() string {
	return "left"
}

type Driver struct{}

func (d *Driver) GetSideToDrive(country SideInCountry) string {
	return country.GetDrivingSideInCountry()
}

type Cyclist struct{}

func (c *Cyclist) GetSideToRide(country SideInCountry) string {
	return country.GetDrivingSideInCountry()
}

func main() {
	d := Driver{}
	spain := CountrySpain{}
	fmt.Printf("I'm in Spain so I have to drive on the %s side\n", d.GetSideToDrive(&spain))
	england := CountryEngland{}
	fmt.Printf("I'm in England now so I have to drive on the %s side\n", d.GetSideToDrive(&england))
	c := Cyclist{}
	fmt.Printf("I'm still in England and the side for bikes is: %s\n", c.GetSideToRide(&england))
}
