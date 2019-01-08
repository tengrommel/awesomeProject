package main

import "fmt"

type generalInfo struct {
	country, hairColor string
}
type player struct {
	name, sport string
	age         int
	generalInfo
}

func main() {
	var player1 player
	player1.name = "Wayne Gretzky"
	player1.age = 57
	player1.sport = "Hockey"
	player1.country = "Canada"
	player1.hairColor = "Brown"
	fmt.Println(player1)
}
