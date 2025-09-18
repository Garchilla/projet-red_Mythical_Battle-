package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var crafted bool
var item string
var scanner = bufio.NewScanner(os.Stdin)
var maxPitItems = 10
var upgradesUsed = 0

const maxUpgrades = 3

type Driver struct {
	Name           string
	Team           string
	Level          int
	MaxStamina     int
	CurrStamina    int
	PitItems       []string
	Skills         []string
	SponsorCredits int
	Gear           Gear
	Initiative     int
	CurrExp        int
	MaxExp         int
	Focus          int
	MaxFocus       int
}

type Gear struct {
	Helmet string
	Suit   string
	Boots  string
}

type Rival struct {
	Name        string
	MaxStamina  int
	CurrStamina int
	AttackPts   int
	Initiative  int
}

func initDriver(name, team string, level, maxStamina, currStamina, initiative, maxExp, maxFocus int, pitItems []string, credits int) Driver {
	return Driver{
		Name:           name,
		Team:           team,
		Level:          level,
		MaxStamina:     maxStamina,
		CurrStamina:    currStamina,
		PitItems:       pitItems,
		Skills:         []string{"Basic Overtake"},
		SponsorCredits: credits,
		Gear:           Gear{},
		Initiative:     initiative,
		CurrExp:        0,
		MaxExp:         maxExp,
		Focus:          maxFocus,
		MaxFocus:       maxFocus,
	}
}

func displayInfo(d *Driver) {
	fmt.Printf("Name: %s\nTeam: %s\nLevel: %d\nStamina: %d/%d\nFocus: %d/%d\nCredits: %d\nInitiative: %d\nExp: %d/%d\nPit Items: %v\nSkills: %v\nGear: Helmet=%s, Suit=%s, Boots=%s\n",
		d.Name, d.Team, d.Level, d.CurrStamina, d.MaxStamina, d.Focus, d.MaxFocus, d.SponsorCredits, d.Initiative, d.CurrExp, d.MaxExp, d.PitItems, d.Skills, d.Gear.Helmet, d.Gear.Suit, d.Gear.Boots) // Fixed Gear order
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
			d.CurrStamina += 50
			if d.CurrStamina > d.MaxStamina {
				d.CurrStamina = d.MaxStamina
			}
			fmt.Printf("Used Energy Drink. Stamina: %d/%d\n", d.CurrStamina, d.MaxStamina)
			return
		}
	}
	fmt.Println("No Energy Drink Available")
}

func readInput() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func mainMenu(d *Driver) {
	for {
		fmt.Println("\nMenu:\n1. Display Driver Info\n2. Access Pit Items\n3. Pit Shop\n4. Garage Mechanic\n5. Training Race\n6. Quit")
		choice := readInput()
		switch choice {
		case "1":
			displayInfo(d)
		case "2":
			accessPitItemsMenu(d)
		case "3":
			pitShopMenu(d)
		case "4":
			garageMenu(d)
		case "5":
			trainingRace(d)
		case "6":
			fmt.Println("Thanks for playing.")
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
			if strings.HasPrefix(item, "Racing") {
				equipGear(d, item)
			}
			if item == "Energy Drink" {
				useEnergyDrink(d)
			}
			if item == "Yellow Flag" {
				useYellowFlag(d)
			}
			if item == "Skill Manual: DRS Boost" {
				learnSkill(d, "Aggressive Pass")
				removeFromPitItems(d, item)
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
		fmt.Println("\nPit Shop:\n1. Energy Drink (3 credits)\n2. Yellow Flag (6 credits)\n3. Skill Manual: DRS Boost (25 credits)\n4. Carbon Fiber (4 credits)\n5. Titanium Alloy (7 credits)\n6. Rubber Compound (3 credits)\n7. Aero Foil (1 credit)\n8. Pit Box Upgrade (30 credits)\nback. Return")
		input := readInput()
		if input == "back" {
			return
		}
		var cost int
		var item string
		switch input {
		case "1":
			cost = 3
			item = "Energy Drink"
		case "2":
			cost = 6
			item = "Yellow Flag"
		case "3":
			cost = 25
			item = "Skill Manual: DRS Boost"
		case "4":
			cost = 4
			item = "Carbon Fiber"
		case "5":
			cost = 7
			item = "Titanium Alloy"
		case "6":
			cost = 3
			item = "Rubber Compound"
		case "7":
			cost = 1
			item = "Aero Foil"
		case "8":
			upgradePitBox(d)
			continue
		default:
			fmt.Println("Invalid choice.")
			continue
		}
		if d.SponsorCredits >= cost {
			d.SponsorCredits -= cost
			addToPitItems(d, item)
			fmt.Printf("Bought %s\n", item)
		} else {
			fmt.Println("Not enough credits.")
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

func driverCreation() Driver {
	var name string
	for {
		fmt.Print("Enter driver name (letters only): ")
		name = readInput()
		if strings.Trim(name, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ") == "" {
			name = (strings.ToLower(name))
			break
		}
		fmt.Println("Invalid name.")
	}

	var team string
	var maxStamina, initiative int
	for {
		fmt.Println("Choose team: 1. Ferrari (100 stamina, 10 initiative), 2. Mercedes (80 stamina, 12 initiative), 3. Red Bull (120 stamina, 8 initiative)")
		choice := readInput()
		switch choice {
		case "1":
			team = "Ferrari"
			maxStamina = 100
			initiative = 10
		case "2":
			team = "Mercedes"
			maxStamina = 80
			initiative = 12
		case "3":
			team = "Red Bull"
			maxStamina = 120
			initiative = 8
		default:
			fmt.Println("Invalid choice.")
			continue
		}
		break
	}

	currStamina := maxStamina / 2
	return initDriver(name, team, 1, maxStamina, currStamina, initiative, 100, 50, []string{}, 100)
}

func checkInventoryLimit(d *Driver) bool {
	return len(d.PitItems) < maxPitItems
}

func garageMenu(d *Driver) {
	for {
		fmt.Println("\nGarage:\n1. Racing Helmet (Aero Foil + Rubber Compound)\n2. Racing Suit (2 Carbon Fiber + Titanium Alloy)\n3. Racing Boots (Carbon Fiber + Rubber Compound)\nback. Return")
		input := readInput()
		if input == "back" {
			return
		}
		if d.SponsorCredits < 5 {
			fmt.Println("Not enough credits.")
			continue
		}
		switch input {
		case "1":
			if removeFromPitItems(d, "Aero Foil") && removeFromPitItems(d, "Rubber Compound") {
				crafted = true
				item = "Racing Helmet"
			}
		case "2":
			cf1 := removeFromPitItems(d, "Carbon Fiber")
			cf2 := removeFromPitItems(d, "Carbon Fiber")
			ta := removeFromPitItems(d, "Titanium Alloy")
			if cf1 && cf2 && ta {
				crafted = true
				item = "Racing Suit"
			}
		case "3":
			if removeFromPitItems(d, "Carbon Fiber") && removeFromPitItems(d, "Rubber Compound") {
				crafted = true
				item = "Racing Boots"
			}
		default:
			fmt.Println("Invalid choice.")
			continue
		}
		if crafted {
			d.SponsorCredits -= 5
			addToPitItems(d, item)
			fmt.Printf("Crafted %s\n", item)
			crafted = false
		} else {
			fmt.Println("Missing materials.")
		}
	}
}

func equipGear(d *Driver, item string) {
	switch item {
	case "Racing Helmet":
		if d.Gear.Helmet != "" {
			addToPitItems(d, d.Gear.Helmet)
		}
		d.Gear.Helmet = item
		d.MaxStamina += 10
	case "Racing Suit":
		if d.Gear.Suit != "" {
			addToPitItems(d, d.Gear.Suit)
		}
		d.Gear.Suit = item
		d.MaxStamina += 25
	case "Racing Boots":
		if d.Gear.Boots != "" {
			addToPitItems(d, d.Gear.Boots)
		}
		d.Gear.Boots = item
		d.MaxStamina += 15
	}
	if d.CurrStamina > d.MaxStamina {
		d.CurrStamina = d.MaxStamina
	}
	removeFromPitItems(d, item)
	fmt.Printf("Equipped %s. Max Stamina now %d\n", item, d.MaxStamina)
}

func upgradePitBox(d *Driver) {
	if upgradesUsed >= maxUpgrades {
		fmt.Println("Max upgrades reached.")
		return
	}
	if d.SponsorCredits < 30 {
		fmt.Println("Not enough credits.")
		return
	}
	d.SponsorCredits -= 30
	maxPitItems += 10
	upgradesUsed++
	fmt.Printf("Pit box upgraded! New max: %d\n", maxPitItems)
}

func gainExp(d *Driver, exp int) {
	d.CurrExp += exp
	for d.CurrExp >= d.MaxExp {
		d.CurrExp -= d.MaxExp
		d.Level++
		d.MaxExp += 10
		d.MaxStamina += 5
		d.CurrStamina = d.MaxStamina
		fmt.Printf("Level up to %d! Max Stamina +5\n", d.Level)
	}
}

func driverTurn(d *Driver, r *Rival, weather int) bool {
	for {
		fmt.Println("\nYour Turn:\n1. Attack\n2. Inventory")
		choice := readInput()
		switch choice {
		case "1":
			dmg := 5
			if weather == 1 {
				dmg = int(float64(dmg) * 0.8)
			}
			r.CurrStamina -= dmg
			fmt.Printf("%s uses Basic Overtake on %s for %d damage!\n", d.Name, r.Name, dmg)
			fmt.Printf("%s Stamina: %d/%d\n", r.Name, r.CurrStamina, r.MaxStamina)
			return true
		case "2":
			accessPitItemsMenu(d)
			return true
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func initRookieRival() Rival {
	return Rival{Name: "Rookie Rival", MaxStamina: 40, CurrStamina: 40, AttackPts: 5, Initiative: 9}
}

func rivalPattern(r *Rival, d *Driver, turn int, weather int) {
	dmg := r.AttackPts
	if turn%3 == 0 {
		dmg *= 2
	}
	if weather == 1 {
		dmg = int(float64(dmg) * 0.8)
	}
	d.CurrStamina -= dmg
	fmt.Printf("%s overtakes %s for %d damage!\n", r.Name, d.Name, dmg)
	fmt.Printf("%s Stamina: %d/%d\n", d.Name, d.CurrStamina, d.MaxStamina)
	isCrashed(d)
}

func trainingRace(d *Driver) {
	r := initRookieRival()
	rand.Seed(time.Now().UnixNano())
	weather := rand.Intn(2)
	if weather == 1 {
		fmt.Println("Rainy conditions!")
	}
	turn := 1
	playerFirst := d.Initiative >= r.Initiative
	for d.CurrStamina > 0 && r.CurrStamina > 0 {
		fmt.Printf("\nTurn %d\n", turn)
		if playerFirst {
			if !driverTurn(d, &r, weather) {
				continue
			}
			if r.CurrStamina <= 0 {
				fmt.Println("You win!")
				gainExp(d, 20)
				break
			}
			rivalPattern(&r, d, turn, weather)
			if d.CurrStamina <= 0 {
				fmt.Println("You lose!")
				break
			}
		} else {
			rivalPattern(&r, d, turn, weather)
			if d.CurrStamina <= 0 {
				fmt.Println("You lose!")
				break
			}
			if !driverTurn(d, &r, weather) {
				continue
			}
			if r.CurrStamina <= 0 {
				fmt.Println("You win!")
				gainExp(d, 20)
				break
			}
		}
		turn++
	}
	fmt.Println("Race over. Back to menu.")
}

func main() {
	driver := driverCreation()
	mainMenu(&driver)
}
