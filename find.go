package main

import (
	"fmt"

	"github.com/duongphuhiep/pokemon-best-team/pokemon"
)

func main() {
	bestTeams := pokemon.Find()
	for _, team := range bestTeams {
		fmt.Println(team)
	}
}
