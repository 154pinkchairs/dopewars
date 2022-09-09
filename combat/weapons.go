package combat

import (
	p "dopewars/player"
	"fmt"
	"player"
	"strconv"
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

type weaponUnits struct {
	Name  string
	Count int
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
	//unlock weapons based on the player.Player's reputation
	if player.Player.reputation >= 0 {
		player.Player.weaponsAvailable = append(player.Player.weaponsAvailable, knife)
	}
	if player.Player.reputation >= 1 {
		player.Player.weaponsAvailable = append(player.Player.weaponsAvailable, baseballBat)
	}
	if player.Player.reputation >= 2 {
		player.Player.weaponsAvailable = append(player.Player.weaponsAvailable, machete)
	}
	if player.Player.reputation >= 3 {
		player.Player.weaponsAvailable = append(player.Player.weaponsAvailable, pistol)
	}
	if player.Player.reputation >= 4 {
		player.Player.weaponsAvailable = append(player.Player.weaponsAvailable, SMG)
	}
	if player.Player.eputation >= 5 {
		player.Player.weaponsAvailable = append(player.Player.weaponsAvailable, shotgun)
	}
	if player.Player.reputation >= 6 {
		player.Player.weaponsAvailable = append(player.Player.weaponsAvailable, machineGun)
	}
	if player.Player.reputation >= 7 {
		player.Player.weaponsAvailable = append(player.Player.weaponsAvailable, handgrenade)
	}
}

func buyWeapon() {
	p.Player.weapons = [4]Weapon{knuckle, knuckle, knuckle, knuckle}
	fmt.Println("You have $" + strconv.Itoa(p.Player.cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("What weapon would you like to buy?")
	//print a numbered list of weapons available to the player.Player, based on their reputation, writable to a weaponChoice variable
	var weaponChoice int
	var maxObtainable int
	for i := 0; i < len(p.Player.weaponsAvailable); i++ {
		if p.Player.weaponsAvailable[i].Price > 0 {
			weaponChoice = append(weaponChoice, strconv.Itoa(i+1)+". "+p.Player.weaponsAvailable[i].Name+" - $"+strconv.Itoa(p.Player.weaponsAvailable[i].Price)+" per unit")
		}
	}
	//prompt the player.Player to select the number of the weapon to buy, using the weapon's number in the list as the index
	fmt.Println("Type the number of the weapon you would like to buy and press enter.")
	fmt.Scanln(&weaponChoice)
	//if weapon.MaxStock > 1, prompt the player.Player to provide the quantity. Read the weapon quantity owned.
	//If the player.Player owns at least 1 unit of a weapon, subtract the quantity owned and set the maxObtainable variable.
	minObtainable := 1
	//max obtainable is the minimum of the max stock and the player.Player's cash modulo divided by the weapon's price
	maxObtainable = min(p.Player.weaponsAvailable[weaponChoice].MaxStock, p.Player.cash/p.Player.weaponsAvailable[weaponChoice].Price)
	fmt.Println("Please provide the quantity you wish to purchase (%d - %d):", minObtainable, maxObtainable)
	var weaponQuantity int
	fmt.Scanln(&weaponQuantity)
	switch {
	case weaponQuantity < minObtainable:
		//terminate the function if the player.Player tries to buy less than 1 unit
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
		//if the player.Player has enough cash, subtract the cost of the weapon from the player.Player's cash and add the weapon to the player.Player's inventory
		if p.Player.cash >= weaponQuantity*player.Player.weaponsAvailable[weaponChoice].Price {
			p.Player.cash -= weaponQuantity * player.Player.weaponsAvailable[weaponChoice].Price
			p.Player.weapons[weaponChoice].Stock += weaponQuantity
			fmt.Println("You have purchased " + strconv.Itoa(weaponQuantity) + " " + player.Player.weaponsAvailable[weaponChoice].Name + ".")
			fmt.Println("You have $" + strconv.Itoa(player.Player.cash) + " left.")
			fmt.Println("Press enter to continue.")
			fmt.Scanln()
		} else {
			fmt.Println("You cannot afford that many.")
			fmt.Println("Press enter to continue.")
			fmt.Scanln()
		}

	}
	//charge the player.Player the price of the weapon
	player.Player.cash -= player.Player.weaponsAvailable[weaponChoice-1].Price
	//add the weapon to the player.Player's player.Player
	player.Player.weapons = append(player.Player.weapons, player.Player.weaponsAvailable[weaponChoice-1])
}
