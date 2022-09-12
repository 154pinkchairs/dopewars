package combat

import (
	"fmt"
	"strconv"

	p "dopewars/player"
)

type Weapon struct {
	name                         string
	Price                        int
	Damage, Accuracy, FiringRate int
	Melee                        bool
	MeleeOnly                    bool
	MeleeDmg                     int
	Throwable                    bool
	ThrowingDamage               int
	ThrowingAccuracy             int
	MaxStock                     int
	Default                      bool
}

type WeaponUnits struct {
	Name  Weapon
	Count int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var knuckle = Weapon{"Knuckle", 0, 3, 100, 1, true, true, 1, false, 0, 0, 1, true}
var knife = Weapon{"Knife", 100, 10, 100, 1, false, true, 0, true, 5, 50, 5, false}
var baseballBat = Weapon{"Baseball Bat", 200, 20, 100, 1, false, true, 0, false, 0, 0, 1, false}
var machete = Weapon{"Machete", 300, 30, 100, 1, false, true, 0, true, 15, 30, 2, false}
var pistol = Weapon{"Pistol", 1200, 10, 80, 2, false, false, 0, false, 0, 0, 1, false}
var SMG = Weapon{"SMG", 3000, 20, 50, 3, false, false, 0, false, 0, 0, 1, false}
var shotgun = Weapon{"Shotgun", 2000, 30, 60, 1, false, false, 0, false, 0, 0, 1, false}
var machineGun = Weapon{"Machine Gun", 5000, 40, 30, 4, true, false, 10, false, 0, 0, 1, false}
var handgrenade = Weapon{"Handgrenade", 800, 50, 100, 1, false, false, 0, true, 30, 80, 8, false}

func unlockWeapons() {
	//unlock weapons based on the p.Character's Reputation
	if p.Character.Reputation >= 0 {
		p.Character.weaponsAvailable = append(p.Character.weaponsAvailable, knife)
	}
	if p.Character.Reputation >= 1 {
		p.Character.weaponsAvailable = append(p.Character.weaponsAvailable, baseballBat)
	}
	if p.Character.Reputation >= 2 {
		p.Character.weaponsAvailable = append(p.Character.weaponsAvailable, machete)
	}
	if p.Character.Reputation >= 3 {
		p.Character.weaponsAvailable = append(p.Character.weaponsAvailable, pistol)
	}
	if p.Character.Reputation >= 4 {
		p.Character.weaponsAvailable = append(p.Character.weaponsAvailable, SMG)
	}
	if p.Character.Reputation >= 5 {
		p.Character.weaponsAvailable = append(p.Character.weaponsAvailable, shotgun)
	}
	if p.Character.Reputation >= 6 {
		p.Character.weaponsAvailable = append(p.Character.weaponsAvailable, machineGun)
	}
	if p.Character.Reputation >= 7 {
		p.Character.weaponsAvailable = append(p.Character.weaponsAvailable, handgrenade)
	}
}

func buyWeapon() {
	p.Character.weapons = [4]Weapon{knuckle, knuckle, knuckle, knuckle}
	fmt.Println("You have $" + strconv.Itoa(p.Character.cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("What weapon would you like to buy?")
	//print a numbered list of weapons available to the p.Character, based on their Reputation, writable to a weaponChoice variable
	var weaponChoice int
	var maxObtainable int
	for i := 0; i < len(p.Character.weaponsAvailable); i++ {
		if p.Character.weaponsAvailable[i].Price > 0 {
			weaponChoice = append(weaponChoice, strconv.Itoa(i+1)+". "+p.Character.weaponsAvailable[i].Name+" - $"+strconv.Itoa(p.Character.weaponsAvailable[i].Price)+" per unit")
		}
	}
	//prompt the p.Character to select the number of the weapon to buy, using the weapon's number in the list as the index
	fmt.Println("Type the number of the weapon you would like to buy and press enter.")
	fmt.Scanln(&weaponChoice)
	//if weapon.MaxStock > 1, prompt the p.Character to provide the quantity. Read the weapon quantity owned.
	//If the p.Character owns at least 1 unit of a weapon, subtract the quantity owned and set the maxObtainable variable.
	minObtainable := 1
	//max obtainable is the minimum of the max stock and the p.Character's cash modulo divided by the weapon's price
	maxObtainable = min(p.Character.weaponsAvailable[weaponChoice].MaxStock, p.Character.cash/p.Character.weaponsAvailable[weaponChoice].Price)
	fmt.Println("Please provide the quantity you wish to purchase (%d - %d):", minObtainable, maxObtainable)
	var weaponQuantity int
	fmt.Scanln(&weaponQuantity)
	switch {
	case weaponQuantity < minObtainable:
		//terminate the function if the p.Character tries to buy less than 1 unit
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
		//if the p.Character has enough cash, subtract the cost of the weapon from the p.Character's cash and add the weapon to the p.Character's inventory
		if p.Character.cash >= weaponQuantity*p.Character.weaponsAvailable[weaponChoice].Price {
			p.Character.cash -= weaponQuantity * p.Character.weaponsAvailable[weaponChoice].Price
			p.Character.weapons[weaponChoice].Stock += weaponQuantity
			fmt.Println("You have purchased " + strconv.Itoa(weaponQuantity) + " " + p.Character.weaponsAvailable[weaponChoice].Name + ".")
			fmt.Println("You have $" + strconv.Itoa(p.Character.cash) + " left.")
			fmt.Println("Press enter to continue.")
			fmt.Scanln()
		} else {
			fmt.Println("You cannot afford that many.")
			fmt.Println("Press enter to continue.")
			fmt.Scanln()
		}

	}
	//charge the p.Character the price of the weapon
	p.Character.cash -= p.Character.weaponsAvailable[weaponChoice-1].Price
	//add the weapon to the p.Character's p.Character
	p.Character.weapons = append(p.Character.weapons, p.Character.weaponsAvailable[weaponChoice-1])
}
