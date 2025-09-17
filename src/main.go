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
		fmt.Println("\nMenu:\n1. Display Driver Info\n2. Access Pit Items\n3. Pit Shop\n4. Quit")
		choice := readInput()
		switch choice {
		case "1":
			displayInfo(d)
		case "2":
			accessPitItemsMenu(d)
		case "3":
			pitShopMenu(d)
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
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

func addToPitItems(d *Driver, item string) {
	d.PitItems = append(d.PitItems, item)
}

func removeFromPitItems(d *Driver, item string) bool {
	for i, it := range d.PitItems {
		if it == item {
			d.PitItems = append(d.PitItems[:i], d.PitItems[i+1:]...)
			return true
		}
	}
	return false
}

func pitShopMenu(d *Driver) {
	for {
		fmt.Println("\nPit Shop:\n1. Energy Drink (free)\nback. Return")
		input := readInput()
		if input == "back" {
			return
		}
		if input == "1" {
			addToPitItems(d, "Energy Drink")
			fmt.Println("Added Energy Drink")
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

func isCrashed(d *Driver) bool {
	if d.CurrStamina <= 0 {
		d.CurrStamina = d.MaxStamina / 2
		fmt.Printf("Crashed out! Revived with %d/%d stamina.\n", d.CurrStamina, d.MaxStamina)
		return true
	}
	return false
}

func useYellowFlag(d *Driver) {
	for i := 0; i < 3; i++ {
		d.CurrStamina -= 10
		fmt.Printf("Penalty damage! Stamina: %d/%d\n", d.CurrStamina, d.MaxStamina)
		time.Sleep(time.Second)
	}
	isCrashed(d)
}

// Update accessPitItemsMenu to handle "Yellow Flag"
if item == "Yellow Flag" {
	useYellowFlag(d)
}

// Update pitShopMenu
fmt.Println("\nPit Shop:\n1. Energy Drink (free)\n2. Yellow Flag (free)\nback. Return")
if input == "2" {
	addToPitItems(d, "Yellow Flag")
	fmt.Println("Added Yellow Flag")
}

// Update Driver struct
type Driver struct {
	// ... existing
	Skills []string
}

// Update initDriver
return Driver{ ... , Skills: []string{"Basic Overtake"} }

// New func
func learnSkill(d *Driver, skill string) {
	for _, s := range d.Skills {
		if s == skill {
			fmt.Println("Skill already learned.")
			return
		}
	}
	d.Skills = append(d.Skills, skill)
	fmt.Printf("Learned %s!\n", skill)
}

// Update accessPitItemsMenu for "Skill Manual: DRS Boost"
if item == "Skill Manual: DRS Boost" {
	learnSkill(d, "Aggressive Pass")
	removeFromPitItems(d, item)
}

// Update pitShopMenu
fmt.Println("\nPit Shop:\n1. Energy Drink (free)\n2. Penalty Card (free)\n3. Skill Manual: DRS Boost (free)\nback. Return")
if input == "3" {
	addToPitItems(d, "Skill Manual: DRS Boost")
	fmt.Println("Added Skill Manual: DRS Boost")
}
