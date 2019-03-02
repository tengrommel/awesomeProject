package main

import (
	"fmt"
)

type BadSpouseError struct {
	why string
}

func (bse *BadSpouseError) Error() string {
	return bse.why
}

/*
性别枚举
*/
type Gender int

func (g Gender) String() string {
	return [...]string{
		"Male",
		"Female",
		"Bisexual",
	}[g]
}

const (
	// 男的
	Male Gender = iota + 1
	// 女的
	Female
	// 二椅子
	Bisexual
)

type Human struct {
	Name   string
	Age    int
	Height int
	Weight int

	Looking       int
	TargetLooking int

	Rmb       int
	Sex       Gender
	TargetSex Gender
}

func (h *Human) Marry(o *Human) (happiness int, err error) {
	if o.Sex != h.TargetSex {
		panic(BadSpouseError{"性别不符合要求"})
	}
	if o.Looking < 50 {
		err = &BadSpouseError{"颜值过低"}
		return
	}
	happiness = (o.Height * o.Rmb * o.Looking) / (o.Weight * o.Age)
	return
}

func NewHuman(name string, height, weight, age, rmb, looking, targetLooking int, sex, targetSex Gender) *Human {
	h := new(Human)
	h.Name = name
	h.Height, h.Weight, h.Age, h.Rmb, h.Looking, h.TargetLooking, h.Sex, h.TargetSex = height, weight, age, rmb, looking, targetLooking, sex, targetSex
	return h
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	cook := NewHuman("库克", 180, 150, 60, 11000000000, 60, 90, Male, Male)
	vSisteer := NewHuman("你妹", 180, 110, 20, 10000, 99, 80, Female, Male)
	happiness, err := vSisteer.Marry(cook)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("幸福指数", happiness)
	}
	fmt.Println("oo")

	happiness, err = cook.Marry(vSisteer)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("幸福指数", happiness)
	}
	fmt.Println("oo")
}
