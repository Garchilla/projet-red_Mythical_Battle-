package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Driver struct {
	name        string
	team        string
	level       int
	MaxStamina  int
	currstamina int
	PitItems    []string
}

func initDriver(name, team string, level, maxStamina, currStamina int, pitItems []string) Driver {
	return Driver{
		name:        name,
		team:        team,
		level:       level,
		MaxStamina:  maxStamina,
		currstamina: currStamina,
		PitItems:    pitItems,
	}
}

func displayInfo(d *Driver) {
	fmt.Printf("Name: %s\nTeam: %s\nLevel: %d\nStamina: %d/%d\nPit Items: %v\n", d.name, d.team, d.level, d.currstamina, d.MaxStamina, d.PitItems)
}

func accessPitItems(d *Driver) {
	fmt.Println("Pit Items: ")
	for i, item := range d.PitItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func useEnergyDrink(d *Driver) {
	for i, item := range d.PitItems {
		if item == "Energy Drink" {
			d.PitItems = append(d.PitItems[:i], d.PitItems[i+1:]...)
			d.currstamina += 50
			if d.currstamina > d.MaxStamina {
				d.currstamina = d.MaxStamina
			}
			fmt.Printf("Used Energy Drink. Stamina: %d/%d\n", d.currstamina, d.MaxStamina)
			return
		}
	}
	fmt.Println("No Energy Drink Available")
}

var scanner = bufio.NewScanner(os.Stdin)

func readInput() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func mainMenu(d *Driver) {
	for {
		fmt.Println("\nMenu:\n1. Dislay Driver Info\n2. Access Pit Items\n3. Quit")
		choice := readInput()
		switch choice {
		case "1":
			displayInfo(d)
		case "2":
			accessPitItemsMenu(d)
		case "3":
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}

func accessPitItemsMenu(d *Driver) {
	for {
		accessPitItems(d)
		fmt.Println("Select an item number to use or type 'back' to return:")
		input := readInput()
		if input == "back" {
			return
		}
		num, err := strconv.Atoi(input)
		if err == nil && num > 0 && num <= len(d.PitItems) {
			item := d.PitItems[num-1]
			if item == "Energy Drink" {
				useEnergyDrink(d)
			} else {
				fmt.Println("Item not usable")
			}
		} else {
			fmt.Println("Invalid Input")
		}
	}
}

func main() {
	driver := initDriver("YourName", "Ferrari", 1, 100, 40, []string{"Energy Drink", "Energy Drink", "Energy Drink"})
	mainMenu(&driver)
}
