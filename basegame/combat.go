package basegame

import (
	"fmt"
	"strconv"
)

type FightFlight interface {
	MeleeAttack()
	RangedAttack()
	Throw()
	Dodge()
	Bribe()
	Submit()
	Run()
}

type Weapon struct {
	Name             string
	Price            int
	Damage, Accuracy int
	FiringRate       float32
	Melee            bool
	MeleeOnly        bool
	MeleeDmg         int
	Throwable        bool
	ThrowingDamage   int
	ThrowingAccuracy int
	MaxStock         int
	Default          bool
}

// change WeaponUnits to a map of Weapon to int
type WeaponUnits map[Weapon]int

type Enemy struct {
	Name         string
	HP           int
	Weapon       Weapon
	Weapon2      Weapon
	Speed        float64
	Perseverance float32
	likelihood   float32
	canSteal     bool
	canArrest    bool
}

var PolicePed = Enemy{"Police Officer on foot", 100, pistol, baton, 1.0, 1.0, 0.5, false, true}                //TODO: likelihood will depend on WantedLevel. If the c has a very high reputation, it will be possible to visit the police station and bribe the police to decrease the WantedLevel and the likelihood of encounter.
var PoliceCar = Enemy{"Police Officer in a car", 100, pistol, baton, 2.5, 1.5, 0.33, false, true}              //TODO: add cars for the player
var PoliceHeli = Enemy{"Police Officer in a helicopter", 100, machineGun, shotgun, 5.0, 3.0, 0.1, false, true} //TODO: fights escalation
var Crook = Enemy{"Crook", 100, Knuckle, Knife, 1.2, 1.0, 0.4, true, false}                                    //print only that he's armed with a Knife
var Junkie = Enemy{"Junkie", 70, Knuckle, Knife, 1.0, 1.75, 0.6, true, false}                                  //TODO: random junkie health, changes depending on which drug the player is selling the most
var Gangster = Enemy{"Gangster", 100, pistol, machete, 1.5, 2.0, 0.08, true, false}

// TODO: weapon randomization
var Gangster2 = Enemy{"Gangster", 100, SMG, baseballBat, 1.3, 1.9, 0.1, true, false}
var Gangster3 = Enemy{"Gangster", 100, shotgun, Knuckle, 2.0, 2.1, 0.07, true, false}
var GangLeader = Enemy{"Gang Leader", 100, machineGun, pistol, 1.7, 2.5, 0.05, true, false}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var Knuckle = Weapon{"Knuckle", 0, 3, 100, 1.0, true, true, 1, false, 0, 0, 1, true}
var Knife = Weapon{"Knife", 100, 10, 100, 1.0, false, true, 0, true, 5, 50, 5, false}
var baton = Weapon{"Police Baton", 160, 22, 100, 0.75, false, true, 0, false, 0, 0, 1, false}
var baseballBat = Weapon{"Baseball Bat", 200, 20, 100, 0.67, false, true, 0, false, 0, 0, 1, false}
var machete = Weapon{"Machete", 300, 30, 100, 1, false, true, 0, true, 15, 30, 2, false}
var pistol = Weapon{"Pistol", 1200, 10, 80, 2, false, false, 0, false, 0, 0, 1, false}
var SMG = Weapon{"SMG", 3000, 20, 50, 3, false, false, 0, false, 0, 0, 1, false}
var shotgun = Weapon{"Shotgun", 2000, 30, 60, 1, false, false, 0, false, 0, 0, 1, false}
var machineGun = Weapon{"Machine Gun", 5000, 40, 30, 4, true, false, 10, false, 0, 0, 1, false}
var handgrenade = Weapon{"Handgrenade", 800, 50, 100, 1, false, false, 0, true, 30, 80, 8, false}

func unlockWeapons(c *Character) {
	//unlock weapons based on the Character's Reputation
	if c.Reputation >= 0 {
		c.weaponsAvailable = append(c.weaponsAvailable, Knife)
	}
	if c.Reputation >= 1 {
		c.weaponsAvailable = append(c.weaponsAvailable, baseballBat)
	}
	if c.Reputation >= 2 {
		c.weaponsAvailable = append(c.weaponsAvailable, machete)
	}
	if c.Reputation >= 5 {
		c.weaponsAvailable = append(c.weaponsAvailable, pistol)
	}
	if c.Reputation >= 8 {
		c.weaponsAvailable = append(c.weaponsAvailable, SMG)
	}
	if c.Reputation >= 12 {
		c.weaponsAvailable = append(c.weaponsAvailable, shotgun)
	}
	if c.Reputation >= 25 {
		c.weaponsAvailable = append(c.weaponsAvailable, machineGun)
	}
	if c.Reputation >= 7 {
		c.weaponsAvailable = append(c.weaponsAvailable, handgrenade)
	}
}

func buyWeapon(c *Character, w *Weapon, wu *WeaponUnits) {
	//using a map of Weapon to int give c one unit of Knuckle
	(*wu)[Knuckle] = 1
	fmt.Println("You have $" + strconv.Itoa(c.Cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("What weapon would you like to buy?")
	//print a numbered list of weapons available to the c, based on their Reputation and amount of Cash, writable to a weaponChoice variable
	var weaponChoice int
	var maxObtainable int
	for i := 0; i < len(c.weaponsAvailable); i++ {
		fmt.Println(strconv.Itoa(i+1) + ". " + c.weaponsAvailable[i].Name + " $" + strconv.Itoa(c.weaponsAvailable[i].Price))
	}
	//prompt the c to select the number of the weapon to buy, using the weapon's number in the list as the index
	fmt.Println("Type the number of the weapon you would like to buy and press enter.")
	fmt.Scanln(&weaponChoice)
	//if weapon.MaxStock > 1, prompt the c to provide the quantity. Read the weapon quantity owned.
	//If the c owns at least 1 unit of a weapon, subtract the quantity owned and set the maxObtainable variable.
	minObtainable := 1
	//max obtainable is the minimum of the max stock and the c's Cash modulo divided by the weapon's price
	maxObtainable = min(c.weaponsAvailable[weaponChoice].MaxStock, c.Cash/c.weaponsAvailable[weaponChoice].Price)
	fmt.Printf("Please provide the quantity you wish to purchase (%d - %d):", minObtainable, maxObtainable)
	var weaponQuantity int
	fmt.Scanln(&weaponQuantity)
	switch {
	case weaponQuantity < minObtainable:
		//terminate the function if the c tries to buy less than 1 unit
		return
	case weaponQuantity > maxObtainable:
		fmt.Println("You cannot afford or carry that many.")
		fmt.Println("Press space to continue. To abort purchase, press c.")
		var abort string
		fmt.Scanln(&abort)
		if abort == "c" {
			return
		}
	default:
		//if the c has enough Cash, subtract the cost of the weapon from the c's Cash and add the weapon to the c's inventory
		if c.Cash >= weaponQuantity*c.weaponsAvailable[weaponChoice].Price {
			c.Cash -= weaponQuantity * c.weaponsAvailable[weaponChoice].Price
			//give c the weapon(s) they bought, by modifying the c's weapons Name and count
			(*wu)[c.weaponsAvailable[weaponChoice]] = weaponQuantity
			fmt.Println("You have purchased " + strconv.Itoa(weaponQuantity) + " " + c.weaponsAvailable[weaponChoice].Name + ".")
			fmt.Println("You have $" + strconv.Itoa(c.Cash) + " left.")
			fmt.Println("Press enter to continue.")
			fmt.Scanln()
		} else {
			fmt.Println("You cannot afford that many.")
			fmt.Println("Press enter to continue.")
			fmt.Scanln()
		}

	}
}