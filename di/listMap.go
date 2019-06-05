package main

import "fmt"

type alarm struct {
	id int
	name string
}

func (a *alarm)add()  {
	a.name += "name"
}

func main() {
	//a := alarm{}
	//a.add()
	//a.add()
	//fmt.Println(a)
	//alarmMap := make(map[string]alarm)
	//alarmMap["ok"] = alarm{1,"d"}
	//alarmMap["ok"].add()
	alarmMap := make(map[string]*alarm)
	if _, ok := alarmMap["ok"]; !ok{
		alarmMap["ok"] = &alarm{}
	}
	alarmMap["ok"].add()
	fmt.Println(alarmMap)

	alarmList := make([]alarm, 0)
	alarmList = append(alarmList, alarm{2, "p"})
	for i :=range alarmList {
		alarmList[i].add()
	}
	fmt.Println(alarmList)
}
