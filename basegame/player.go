package basegame

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Character struct {
	Name                    string
	Health                  int
	Reputation, WantedLevel int
	Cash                    int
	Days                    int
	Bank                    int
	Debt                    int
	CurrentDistrict         District
	drugs                   Drugs
	Weapons                 WeaponUnits
	weaponsAvailable        []Weapon
}

func InitPlayer() *Character {
	c := &Character{}
	c.Name = "Heisenberg"
	c.Cash = 10000
	c.Debt = 15000
	c.Reputation = 0
	c.Days = 0
	c.WantedLevel = 0
	c.CurrentDistrict = Bronx

	return c
}

// using the following doc: https://docs.rocketnine.space/code.rocketnine.space/tslocum/messeji/
// create a new text field at the bottom of the screen
// window will be 1152x648
// text field will be 1152x200
// set the font.Face to assets/VT323_Regular.17.ttf
// set the font size to 32
// var textField = messeji.NewTextField(image.Rect(0, 448, 1152, 648), font.Face("assets/VT323_Regular.ttf"))
type Save struct {
	Name                    string
	Health                  int
	Reputation, WantedLevel int
	Cash                    int
	Days                    int
	Bank                    int
	Debt                    int
	CurrentDistrict         District
	Drugs                   Drugs
	Weapons                 WeaponUnits
	WeaponsAvailable        []Weapon
}

// parse the ../savegame.json file and pass the data to the character struct. If any fields are missing, use the default values
func Loadsave() *Character {
	c := &Character{}
	// open the file
	file, err := os.Open("savegame.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// parse the json
	var save Save
	err = json.Unmarshal(data, &save)
	if err != nil {
		log.Fatal(err)
	}

	// set the character struct values
	c.Name = save.Name
	c.Health = save.Health
	c.Reputation = save.Reputation
	c.WantedLevel = save.WantedLevel
	c.Cash = save.Cash
	c.Bank = save.Bank
	c.Debt = save.Debt
	c.Days = save.Days
	c.CurrentDistrict = save.CurrentDistrict
	c.drugs = save.Drugs
	c.Weapons = save.Weapons
	c.weaponsAvailable = save.WeaponsAvailable

	return c
}
