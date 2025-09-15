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

func displayInfo(c character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Printf("Nom : %s\n", c.name)
	fmt.Printf("Classe : %s\n", c.class)
	fmt.Printf("Niveau : %d\n", c.level)
	fmt.Printf("Points de vie maximum : %d\n", c.maxhp)
	fmt.Printf("Points de vie actuels : %d\n", c.currenthp)
	fmt.Println("Inventaire :")
	for item, quantity := range c.inventory {
		fmt.Printf("  - %s : %d\n", item, quantity)
	}
	fmt.Println("================================")
}

func main() {
	c1 := initCharacter("Frank", "Elfe", 1, 100, 40, map[string]int{"potions": 3})
	displayInfo(c1)
}
oui