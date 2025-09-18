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
		Skills:         []string{"Dépassement"},
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
	fmt.Printf("Pilote: %s\nEcurie: %s\nNiveau: %d\nStamina: %d/%d\nFocus: %d/%d\nCredits: %d\nInitiative: %d\nExp: %d/%d\nInventaire: %v\nSkills: %v\nItems: Casque=%s, Combinaison=%s, Bottes=%s\n",
		d.Name, d.Team, d.Level, d.CurrStamina, d.MaxStamina, d.Focus, d.MaxFocus, d.SponsorCredits, d.Initiative, d.CurrExp, d.MaxExp, d.PitItems, d.Skills, d.Gear.Helmet, d.Gear.Suit, d.Gear.Boots)
}

func accessPitItems(d *Driver) {
	fmt.Println("Inventaire: ")
	for i, item := range d.PitItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func useEnergyDrink(d *Driver, item string) bool {
	for i, it := range d.PitItems {
		if it == item {
			d.PitItems = append(d.PitItems[:i], d.PitItems[i+1:]...)
			d.CurrStamina += 50
			if d.CurrStamina > d.MaxStamina {
				d.CurrStamina = d.MaxStamina
			}
			fmt.Printf("Red Bull utilisé. Stamina: %d/%d\n", d.CurrStamina, d.MaxStamina)
			fmt.Printf("Objet '%s' retiré de l'inventaire.\n", item)
			return true
		}
	}
	fmt.Println("Pas de Red Bull disponible :(")
	return false
}

func readInput() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func mainMenu(d *Driver) {
	for {
		fmt.Println("\n===========MAIN MENU============\n1. Info sur le pilote \n2. Accéder à l'inventaire\n3. Pit Stop\n4. Garage Mécanique\n5. Lancer la Course d'entrainement\n6. Lancer le Grand Prix\n7. Qui sont les artistes cachés?\n8. Quitter\n=================================")
		choice := readInput()
		switch choice {
		case "1":
			displayInfo(d)
		case "2":
			accessPitItemsMenu(d, false, nil)
		case "3":
			pitShopMenu(d)
		case "4":
			garageMenu(d)
		case "5":
			trainingRace(d)
		case "6":
			grandPrix(d)
		case "7":
			fmt.Println("Les artistes cachés sont ABBA et Steven Spielberg")
		case "8":
			fmt.Println("Merci d'avoir joué(e)!")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func containsSkill(d *Driver, skill string) bool {
	for _, s := range d.Skills {
		if s == skill {
			return true
		}
	}
	return false
}

func accessPitItemsMenu(d *Driver, inRace bool, r *Rival) {
	for {
		accessPitItems(d)
		fmt.Println("Choisissez un objet à utiliser ou écrivez 'back' pour revenir en arrière:")
		input := readInput()
		if input == "back" {
			return
		}
		num, err := strconv.Atoi(input)
		if err == nil && num > 0 && num <= len(d.PitItems) {
			item := d.PitItems[num-1]
			if inRace {
				if item == "Skill Manuel: DRS Boost" || item == "Skill Manuel: One Shot" {
					fmt.Println("Cet objet ne peut pas être utilisé pendant la course.")
					continue
				}
				if item == "Casque de course" || item == "Combinaison de course" || item == "Bottes de course" {
					fmt.Println("Les équipements ne peuvent pas être équipés pendant la course.")
					continue
				}
			}
			if (item == "Casque de course" || item == "Combinaison de course" || item == "Bottes de course") && !inRace {
				equipGear(d, item)
			} else if item == "Red Bull (Gives you wings)" {
				if useEnergyDrink(d, item) {
					if inRace {
						fmt.Printf("Inventaire après utilisation: %v\n", d.PitItems)
					}
				}
			} else if item == "Drapeau Jaune" {
				if inRace && r != nil {
					if useYellowFlag(d, r, item) {
						fmt.Printf("Inventaire après utilisation: %v\n", d.PitItems)
					}
				} else {
					if useYellowFlag(d, nil, item) {
						fmt.Printf("Inventaire après utilisation: %v\n", d.PitItems)
					}
				}
			} else if item == "Skill Manuel: DRS Boost" && !inRace {
				learnSkill(d, "Dépassement agressif")
				if removeFromPitItems(d, item) {
					fmt.Printf("Objet '%s' retiré de l'inventaire.\n", item)
					if inRace {
						fmt.Printf("Inventaire après utilisation: %v\n", d.PitItems)
					}
				}
			} else if item == "Skill Manuel: One Shot" && !inRace {
				learnSkill(d, "One Shot")
				if removeFromPitItems(d, item) {
					fmt.Printf("Objet '%s' retiré de l'inventaire.\n", item)
					if inRace {
						fmt.Printf("Inventaire après utilisation: %v\n", d.PitItems)
					}
				}
			} else {
				fmt.Println("Item n'est pas utilisable ici.")
			}
		} else {
			fmt.Println("Invalide Input")
		}
	}
}

func addToPitItems(d *Driver, item string) {
	if !checkInventoryLimit(d) {
		fmt.Println("Inventaire full!")
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
		fmt.Println("\n=====INVENTAIRE=====\n1. Red Bull (Gives You Wings) (3 credits)\n2. Drapeau jaune (6 credits)\n3. Skill Manuel: DRS Boost (25 credits)\n4. Skill Manuel: One Shot (30 credits)\n5. Carbone (4 credits)\n6. Titane (7 credits)\n7. Caoutchouc (3 credits)\n8. Aerodynamique (1 credit)\n9. Upgrade inventaire (30 credits)\nback. Return")
		input := readInput()
		if input == "back" {
			return
		}
		var cost int
		var item string
		switch input {
		case "1":
			cost = 3
			item = "Red Bull (Gives you wings)"
		case "2":
			cost = 6
			item = "Drapeau Jaune"
		case "3":
			cost = 25
			item = "Skill Manuel: DRS Boost"
		case "4":
			cost = 30
			item = "Skill Manuel: One Shot"
		case "5":
			cost = 4
			item = "Carbone"
		case "6":
			cost = 7
			item = "Titane"
		case "7":
			cost = 3
			item = "Caoutchouc"
		case "8":
			cost = 1
			item = "Aerodynamique"
		case "9":
			upgradePitBox(d)
			continue
		default:
			fmt.Println("Choix invalide")
			continue
		}
		if d.SponsorCredits >= cost {
			d.SponsorCredits -= cost
			addToPitItems(d, item)
			fmt.Printf("Acheter %s\n", item)
		} else {
			fmt.Println("Pas assez de crédits")
		}
	}
}

func isCrashed(d *Driver) bool {
	if d.CurrStamina <= 0 {
		d.CurrStamina = d.MaxStamina / 2
		fmt.Printf("ACCIDENT!!! Rescussiter avec: %d/%d stamina.\n", d.CurrStamina, d.MaxStamina)
		return true
	}
	return false
}

func isRivalCrashed(r *Rival) bool {
	if r.CurrStamina <= 0 {
		r.CurrStamina = r.MaxStamina / 2
		fmt.Printf("%s A CRASHÉ!!! Relancé avec: %d/%d stamina.\n", r.Name, r.CurrStamina, r.MaxStamina)
		return true
	}
	return false
}

func useYellowFlag(d *Driver, r *Rival, item string) bool {
	if d == nil {
		fmt.Println("Erreur: pilote non défini pour retirer l'objet.")
		return false
	}
	if r != nil {
		for i := 0; i < 3; i++ {
			r.CurrStamina -= 10
			fmt.Printf("Drapeau Jaune! Pénalité pour %s: Stamina: %d/%d\n", r.Name, r.CurrStamina, r.MaxStamina)
			time.Sleep(time.Second)
		}
		isRivalCrashed(r)
	} else {
		for i := 0; i < 3; i++ {
			d.CurrStamina -= 10
			fmt.Printf("Penalty! Stamina: %d/%d\n", d.CurrStamina, d.MaxStamina)
			time.Sleep(time.Second)
		}
		isCrashed(d)
	}
	if removeFromPitItems(d, item) {
		fmt.Printf("Objet '%s' retiré de l'inventaire.\n", item)
		return true
	}
	fmt.Println("Erreur: impossible de retirer l'objet de l'inventaire.")
	return false
}

func learnSkill(d *Driver, skill string) {
	for _, s := range d.Skills {
		if s == skill {
			fmt.Println("Skill déjà connu.")
			return
		}
	}
	d.Skills = append(d.Skills, skill)
	fmt.Printf("Appris %s!\n", skill)
}

func driverCreation() Driver {
	var name string
	for {
		fmt.Print("Entrez le nom de votre pilote (lettres seulement): ")
		name = readInput()
		if strings.Trim(name, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ") == "" {
			name = strings.ToLower(name)
			break
		}
		fmt.Println("Prénom invalide.")
	}

	var team string
	var maxStamina, initiative int
	for {
		fmt.Println("Choisissez l'écurie: 1. Ferrari (100 stamina, 10 initiative), 2. Alpine (80 stamina, 12 initiative), 3. Mclaren (120 stamina, 8 initiative)")
		choice := readInput()
		switch choice {
		case "1":
			team = "Ferrari"
			maxStamina = 100
			initiative = 10
		case "2":
			team = "Alpine"
			maxStamina = 80
			initiative = 12
		case "3":
			team = "Mclaren"
			maxStamina = 120
			initiative = 8
		default:
			fmt.Println("Choix invalide.")
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
		fmt.Println("\nGarage:\n1. Casque de course (Aerodynamique + Rubber)\n2. Combinaison de course (2 Carbone + Titane)\n3. Bottes de course(Carbone + caoutchouc)\nback. Return")
		input := readInput()
		if input == "back" {
			return
		}
		if d.SponsorCredits < 5 {
			fmt.Println("Pas assez de crédits.")
			continue
		}
		switch input {
		case "1":
			if removeFromPitItems(d, "Aerodynamique") && removeFromPitItems(d, "Caoutchouc") {
				crafted = true
				item = "Casque de course"
			}
		case "2":
			cf1 := removeFromPitItems(d, "Carbone")
			cf2 := removeFromPitItems(d, "Carbone")
			ta := removeFromPitItems(d, "Titane")
			if cf1 && cf2 && ta {
				crafted = true
				item = "Combinaison de course"
			}
		case "3":
			if removeFromPitItems(d, "Carbone") && removeFromPitItems(d, "Caoutchouc") {
				crafted = true
				item = "Bottes de course"
			}
		default:
			fmt.Println("Choix invalide.")
			continue
		}
		if crafted {
			d.SponsorCredits -= 5
			addToPitItems(d, item)
			fmt.Printf("Création de %s\n", item)
			crafted = false
		} else {
			fmt.Println("Manque de matériaux.")
		}
	}
}

func equipGear(d *Driver, item string) {
	switch item {
	case "Casque de course":
		if d.Gear.Helmet != "" {
			addToPitItems(d, d.Gear.Helmet)
		}
		d.Gear.Helmet = item
		d.MaxStamina += 10
	case "Combinaison de course":
		if d.Gear.Suit != "" {
			addToPitItems(d, d.Gear.Suit)
		}
		d.Gear.Suit = item
		d.MaxStamina += 25
	case "Bottes de course":
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
	fmt.Printf("Équipé %s. Max Stamina est maintenant %d\n", item, d.MaxStamina)
}

func upgradePitBox(d *Driver) {
	if upgradesUsed >= maxUpgrades {
		fmt.Println("Max upgrades atteint.")
		return
	}
	if d.SponsorCredits < 30 {
		fmt.Println("Pas assez de crédits.")
		return
	}
	d.SponsorCredits -= 30
	maxPitItems += 10
	upgradesUsed++
	fmt.Printf("Inventaire Upgraded! Nouveau max: %d\n", maxPitItems)
}

func gainExp(d *Driver, exp int) {
	d.CurrExp += exp
	for d.CurrExp >= d.MaxExp {
		d.CurrExp -= d.MaxExp
		d.Level++
		d.MaxExp += 10
		d.MaxStamina += 5
		d.CurrStamina = d.MaxStamina
		fmt.Printf("Level up à %d! Stamina Max +5\n", d.Level)
	}
}

func driverTurn(d *Driver, r *Rival, weather int) bool {
	for {
		fmt.Println("\nTon Tour:")
		fmt.Println("1. Attaque (Dépassement)")
		if containsSkill(d, "Dépassement agressif") {
			fmt.Println("2. Attaque (Dépassement agressif)")
		}
		if containsSkill(d, "One Shot") {
			fmt.Println("3. Attaque (One Shot)")
		}
		fmt.Println("4. Inventaire")
		choice := readInput()
		switch choice {
		case "1":
			dmg := 8
			if weather == 1 {
				dmg = int(float64(dmg) * 0.8)
			}
			d.CurrStamina -= 5
			r.CurrStamina -= dmg
			fmt.Printf("%s utilise Dépassement sur %s pour %d dégâts!\n", d.Name, r.Name, dmg)
			fmt.Printf("%s Stamina: %d/%d\n", r.Name, r.CurrStamina, r.MaxStamina)
			return true
		case "2":
			if containsSkill(d, "Dépassement agressif") {
				dmg := 12
				if weather == 1 {
					dmg = int(float64(dmg) * 0.8)
				}
				d.CurrStamina -= 10
				r.CurrStamina -= dmg
				fmt.Printf("%s utilise Dépassement agressif sur %s pour %d dégâts!\n", d.Name, r.Name, dmg)
				fmt.Printf("%s Stamina: %d/%d\n", r.Name, r.CurrStamina, r.MaxStamina)
				return true
			} else {
				fmt.Println("Skill non disponible.")
			}
		case "3":
			if containsSkill(d, "One Shot (Seulement pour test)") {
				r.CurrStamina = 0
				fmt.Printf("%s utilise One Shot sur %s, le vainquant instantanément!\n", d.Name, r.Name)
				fmt.Printf("%s Stamina: %d/%d\n", r.Name, r.CurrStamina, r.MaxStamina)
				return true
			} else {
				fmt.Println("Skill non disponible.")
			}
		case "4":
			accessPitItemsMenu(d, true, r)
			return true
		default:
			fmt.Println("Choix invalide")
		}
	}
}

func initRookieRival() Rival {
	return Rival{Name: "Rookie Rival", MaxStamina: 40, CurrStamina: 40, AttackPts: 10, Initiative: 9}
}

func initF1Rival(name string, rank int) Rival {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxStamina := 40 + (10-rank)*2
	attackPts := 8 + (10-rank)/2
	initiative := 8 + rng.Intn(5)
	return Rival{
		Name:        name,
		MaxStamina:  maxStamina,
		CurrStamina: maxStamina,
		AttackPts:   attackPts,
		Initiative:  initiative,
	}
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
	fmt.Printf("%s accélère sur %s pour %d dégâts!\n", r.Name, d.Name, dmg)
	fmt.Printf("%s Stamina: %d/%d\n", d.Name, d.CurrStamina, d.MaxStamina)
	isCrashed(d)
}

func trainingRace(d *Driver) {
	r := initRookieRival()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	weather := 0
	if rng.Intn(100) < 20 {
		weather = 1
		fmt.Println("La pluie est arrivée!")
	} else {
		fmt.Println("Le temps est clair pour la course!")
	}
	turn := 1
	playerFirst := d.Initiative >= r.Initiative
	for d.CurrStamina > 0 && r.CurrStamina > 0 {
		fmt.Printf("\nTour %d\n", turn)
		if playerFirst {
			if !driverTurn(d, &r, weather) {
				continue
			}
			if r.CurrStamina <= 0 {
				fmt.Println("VICTOIRE!")
				gainExp(d, 20)
				break
			}
			rivalPattern(&r, d, turn, weather)
			if d.CurrStamina <= 0 {
				fmt.Println("YOU DIED")
				break
			}
		} else {
			rivalPattern(&r, d, turn, weather)
			if d.CurrStamina <= 0 {
				fmt.Println("YOU DIED")
				break
			}
			if !driverTurn(d, &r, weather) {
				continue
			}
			if r.CurrStamina <= 0 {
				fmt.Println("VICTOIRE!")
				gainExp(d, 20)
				break
			}
		}
		turn++
	}
	fmt.Println("Fin de la Course d'entrainement. Retour au menu principal.")
}

func grandPrix(d *Driver) {
	f1Drivers := []string{
		"Nico Hulkenberg",
		"Isack Hadjar",
		"Kimi antonelli",
		"Alexander Albon",
		"Lewis Hamilton",
		"Charles Leclerc",
		"George Russell",
		"Max Verstappen",
		"Lando Norris",
		"Oscar Piastri",
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	weather := 0
	if rng.Intn(100) < 20 {
		weather = 1
		fmt.Println("La pluie est arrivée!")
	} else {
		fmt.Println("Le temps est clair pour le Grand Prix!")
	}
	totalTurns := 0
	defeated := 0
	for i, driverName := range f1Drivers {
		if totalTurns >= 70 {
			fmt.Println("Vous avez dépassé la limite de 70 tours. Défaite!")
			break
		}
		rival := initF1Rival(driverName, i+1)
		fmt.Printf("\nCombat contre %s (Rival %d/10)\n", rival.Name, i+1)
		playerFirst := d.Initiative >= rival.Initiative
		for d.CurrStamina > 0 && rival.CurrStamina > 0 && totalTurns < 70 {
			fmt.Printf("\nTour %d (Total: %d/70)\n", totalTurns+1, totalTurns+1)
			if playerFirst {
				if !driverTurn(d, &rival, weather) {
					continue
				}
				if rival.CurrStamina <= 0 {
					fmt.Printf("%s a été vaincu!\n", rival.Name)
					defeated++
					gainExp(d, 20)
					break
				}
				rivalPattern(&rival, d, totalTurns+1, weather)
				if d.CurrStamina <= 0 {
					fmt.Println("YOU DIED! Défaite dans le Grand Prix!")
					break
				}
			} else {
				rivalPattern(&rival, d, totalTurns+1, weather)
				if d.CurrStamina <= 0 {
					fmt.Println("YOU DIED! Défaite dans le Grand Prix!")
					break
				}
				if !driverTurn(d, &rival, weather) {
					continue
				}
				if rival.CurrStamina <= 0 {
					fmt.Printf("%s a été vaincu!\n", rival.Name)
					defeated++
					gainExp(d, 20)
					break
				}
			}
			totalTurns++
		}
		if d.CurrStamina <= 0 || totalTurns >= 70 {
			break
		}
	}
	if defeated == len(f1Drivers) && d.CurrStamina > 0 && totalTurns < 70 {
		fmt.Println("You are the World Champion!")
		gainExp(d, 50)
	} else if totalTurns >= 70 {
		fmt.Println("Limite de tours atteinte. Vous n'êtes pas champion du monde.")
	} else if d.CurrStamina <= 0 {
		fmt.Println("Vous avez crashé. Défaite dans le Grand Prix!")
	}
	fmt.Println("Fin du Grand Prix. Retour au menu principal.")
}

func main() {
	driver := driverCreation()
	mainMenu(&driver)
}
