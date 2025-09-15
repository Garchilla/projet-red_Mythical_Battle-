package main

import (
	"fmt"
)

type character struct {
	name      string
	class     string
	level     int
	maxhp     int
	currenthp int
	inventory map[string]int
}

func initCharacter(name, class string, level, maxhp, currenthp int, inventory map[string]int) character {
	return character{
		name:      name,
		class:     class,
		level:     level,
		maxhp:     maxhp,
		currenthp: currenthp,
		inventory: inventory,
	}
}

func main() {
	c1 := initCharacter("Frank", "Elfe", 1, 100, 40, map[string]int{"potions": 3})

	fmt.Printf("Personnage: %+v\n", c1)
}
