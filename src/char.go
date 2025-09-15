package main

import (
	"fmt"
	"os"
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

/*func inventoryMenu(c character) {
	for {
		fmt.Println("\n=== Menu Inventaire ===")
		fmt.Println("1. Afficher l'inventaire")
		fmt.Println("2. Retour")
		fmt.Print("Choisissez une option (1-2) : ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Entrée invalide, veuillez entrer un numéro.")
			continue
		}

		switch choice {
		case 1:
			displayInventory(c)
		case 2:
			return
		default:
			fmt.Println("Choix invalide, veuillez choisir une option entre 1 et 2.")
		}
	}
}*/

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

	for {
		fmt.Println("\n=== Menu Inventaire ===")
		fmt.Println("1. Afficher l'inventaire")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Quitter")
		fmt.Print("Choisissez une option (1-3) : ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Entrée Invalide, veuillez entrer un numéro afficher")
			continue
		}

		switch choice {
		case 1:
			displayInfo(c1)
		case 2:
			inventoryMenu(c1)
		case 3:
			fmt.Println("Bye Bye!!!")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, veuillez choisir une option entre 1 et 3")
		}
	}
}
