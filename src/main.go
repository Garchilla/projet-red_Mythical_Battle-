package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Driver struct {
	name        string
	team        string
	level       int
	MaxStamina  int
	currstamina int
	PitItems    []string
	skills      []string
}

func initDriver(name, team string, level, maxStamina, currStamina int, pitItems []string) Driver {
	return Driver{
		name:        name,
		team:        team,
		level:       level,
		MaxStamina:  maxStamina,
		currstamina: currStamina,
		PitItems:    pitItems,
		skills:      []string{"Basic Overtake"},
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
				if item == "Yellow Flag" {
					useYellowFlag(d)
				}
				if item == "Skill Manual: DRS Boost" {
					learnSkill(d, "Aggressive Pass")
					removeFromPitItems(d, item)
				}
			} else {
				fmt.Println("Item not usable")
			}
		} else {
			fmt.Println("Invalid Input")
		}
	}
}

func addToPitItems(d *Driver, item string) {
	if !checkInventoryLimit(d) {
		fmt.Println("Pit box full!")
		return
	}
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
		fmt.Println("\nPit Shop:\n1. Energy Drink (free)\n2. Penalty Card (free)\n3. Skill Manual: DRS Boost (free)\nback. Return")
		input := readInput()
		if input == "back" {
			return
		}
		if input == "1" {
			addToPitItems(d, "Energy Drink")
			fmt.Println("Added Energy Drink")
		}
		if input == "2" {
			addToPitItems(d, "Yellow Flag")
			fmt.Println("Added Yellow Flag")
		}
		if input == "3" {
			addToPitItems(d, "Skill Manual: DRS Boost")
			fmt.Println("Added Skill Manual: DRS Boost")
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

func isCrashed(d *Driver) bool {
	if d.currstamina <= 0 {
		d.currstamina = d.MaxStamina / 2
		fmt.Printf("Crashed out! Revived with %d/%d stamina.\n", d.currstamina, d.MaxStamina)
		return true
	}
	return false
}

func useYellowFlag(d *Driver) {
	for i := 0; i < 3; i++ {
		d.currstamina -= 10
		fmt.Printf("Penalty damage! Stamina: %d/%d\n", d.currstamina, d.MaxStamina)
		time.Sleep(time.Second)
	}
	isCrashed(d)
}

func learnSkill(d *Driver, skill string) {
	for _, s := range d.skills {
		if s == skill {
			fmt.Println("Skill already learned.")
			return
		}
	}
	d.skills = append(d.skills, skill)
	fmt.Printf("Learned %s!\n", skill)
}

func driverCreation() Driver {
	var name string
	for {
		fmt.Print("Enter driver name (letters only): ")
		name = readInput()
		if strings.Trim(name, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ") == "" {
			name = strings.Title(strings.ToLower(name))
			break
		}
		fmt.Println("Invalid name.")
	}

	var team string
	var maxStamina int
	for {
		fmt.Println("Choose team: 1. Ferrari (100 stamina), 2. Mercedes (80 stamina), 3. Red Bull (120 stamina)")
		choice := readInput()
		switch choice {
		case "1":
			team = "Ferrari"
			maxStamina = 100
		case "2":
			team = "Mercedes"
			maxStamina = 80
		case "3":
			team = "Red Bull"
			maxStamina = 120
		default:
			fmt.Println("Invalid choice.")
			continue
		}
		break
	}

	currStamina := maxStamina / 2
	return initDriver(name, team, 1, maxStamina, currStamina, []string{})
}

const maxPitItems = 10

func checkInventoryLimit(d *Driver) bool {
	return len(d.PitItems) < maxPitItems
}

func main() {
	driver := driverCreation()
	mainMenu(&driver)
}
