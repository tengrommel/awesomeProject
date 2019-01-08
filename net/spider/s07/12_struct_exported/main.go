package main

import (
	"./types"
	"fmt"
	"strings"
)

func changeAthleteName1(p types.Player) {
	p.Name = "Anderson Silva-Spider"
}

func changeAthleteName2(p *types.Player) {
	p.Name = "Anderson Silva-Spider"
	p.Country = strings.ToUpper(p.Country)
}

func main() {
	player1 := types.Player{"Anderson Silva", "MMA", 43, types.Info{"Brazil", "Black"}}
	// info1 := athletes.Info{Country: "Brazil", HairColor: "Black"}
	// player1 := athletes.Player{Name: "Anderson Silva", Sport: "MMA", Age: 43, Info: info1}
	fmt.Println("(1) player1:", player1)

	changeAthleteName1(player1)
	fmt.Println("(2) player1:", player1)

	changeAthleteName2(&player1)
	fmt.Println("(3) player1:", player1)

	fmt.Println("(4) player1:", *player1.ToLowercase())
}
