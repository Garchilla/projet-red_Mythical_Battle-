package projetredmythicalbattle

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
