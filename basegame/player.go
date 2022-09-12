package basegame

import (
	"fmt"
	"os"
	"strconv"
)

type Character struct {
	Name                    string
	Health                  int
	Reputation, WantedLevel int
	cash                    int
	debt                    int
	CurrentDistrict         District
	drugs                   Drugs
	weapons                 WeaponUnits
	weaponsAvailable        Weapons
}

//export character

/*
	func NewPlayer(c *character)Set(name string, reputation int, CurrentDistrict city.District) {
		return &character{
			Name:        name,
			Health:      100,
			Reputation:  0,
			WantedLevel: 0,
			cash:        5000,
			debt:        10000,
		}
	}
*/
/*func ModPlayer(reputation int, wantedLevel int, cash int, debt int) *character {
	return &character{
		Reputation:  reputation,
		WantedLevel: wantedLevel,
		cash:        cash,
		debt:        debt,
	}
}
*/

func (c *Character) init() {
	c.cash = 10000
	c.debt = 15000
	fmt.Println("Welcome to Dope Wars!")
	fmt.Println("What is your name?")
	fmt.Scanln(&c.Name)
	fmt.Println("Welcome to the world of Dope Wars, " + c.Name + "!")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("You are a small time drug dealer in the city of New York.\n After failing one job after another, you have decided to start a small business.")
	fmt.Println("You have a small amount of cash, but you need to make a lot of money.\nAfter one of your drug deals went down, you were left with a debt.")
	fmt.Println("You have $" + strconv.Itoa(c.debt) + " to pay off.")
	fmt.Println("Press h for help. Press q to quit.")
	var key string
	fmt.Scanln("%s", &key)
	if key == "h" {
		fmt.Println("h - help")
		fmt.Println("q - quit")
		fmt.Println("i - Player")
		fmt.Println("d - district info and the available drugs. Press d again to see the drugs in stock.\n to buy a drug, type the drug number and press enter.")
		fmt.Println("t - travel to a district. Press t again to see the districts you can travel to.\n to travel to a district, type the district number and press enter.")
		fmt.Println("w - weapon info and the available weapons. Press w again to see the weapons in stock.\n to buy a weapon, type the weapon number and press enter.")
		fmt.Println("a - current weapon info. Press a again to see the current weapon stats. Press s to sell the current weapon.")
		fmt.Println("f - fight the opponent. For throwable weapons, press j to throw the weapon. Note you will lose the weapon if you do not deal a critical hit\n or if it's a handgrenade.")
		fmt.Println("s - sell the drugs. Type the drug number and press enter.")
		fmt.Println("o - make a payment or withdraw/borrow money from the bank or loan shark. Type the amount and press enter.")
		fmt.Println("r - run away. You might lose some cash or drugs and the wanted level will go down.")
		fmt.Println("b - bribe the law enforcement. You will lose some cash and the wanted level will go down.")
		fmt.Println("g - visit the bank or loan shark.")
		fmt.Println("u - visit the hospital.")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	} else if key == "q" {
		os.Exit(0)
	}
}
