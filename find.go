package main

import (
	"./pokemon"
	"fmt"
)

func main() {
	bestTeams := pokemon.Find()
	for _, team := range bestTeams {
		fmt.Println(team)
	}
}
