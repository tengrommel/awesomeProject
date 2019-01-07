package main

import "fmt"

func main() {
	employees := map[string]map[string]string{
		"BT": {
			"firstName": "Blake",
			"lastName":  "Travis",
		},
		"PC": {
			"firstName": "Parker",
			"lastName":  "Cooper",
		},
		"DC": {
			"firstName": "Dakota",
			"lastName":  "Carrington",
		},
	}
	if emp, ok := employees["PC"]; ok {
		fmt.Println(emp["firstName"], emp["lastName"])
	}
	for initials, emp := range employees {
		fmt.Println(initials, emp["firstName"], emp["lastName"])
	}
}
