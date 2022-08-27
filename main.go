package dopewars

import (
	"fmt"
	"strconv"
)

func main() {
	var Drug struct {
		Name  string
		Price int
		Stock int
		RaiseWanted int
	}
	var Weapon struct {
		Name     string
		Price    int
		Damage   int
		Accuracy int
		FiringRate int
		Melee bool
		MeleeOnly bool
		MeleeDmg int
		Throwable bool
		ThrowingDamage int
		ThrowingAccuracy int
		MaxStock int
	}
	if weapon.MeleeOnly == true {
		weapon.Damage = weapon.MeleeDmg
		weapon.Melee = true
	}
	if weapon.Throwable == true {
		weapon.ThrowingDamage == 0
		weapon.ThrowingAccuracy == 0
	}
	cash := 10000
	debt := 15000
	var player struct {
		Name            string
		Health          int
		Reputation      int
		Armed           bool
		WantedLevel     int
		CurrentWeapon   []Weapon
		CurrentDistrict string
	}
	var district struct {
		Name     string
		NeighbouringDistricts [2]string
		DrugsAvailable        [5]Drug
		hospital              bool
		bank				  bool
		loanShark             bool
}
	var inventory struct {
		Drugs     [8]Drug
		Weapons   [2]Weapon
		cash	  int
	}
	
	var weed = Drug{"Weed", 50, 0, 2}
	var cocaine = Drug{"Cocaine", 300, 0, 4}
	var heroin = Drug{"Heroin", 200, 0, 6}
	var acid = Drug{"Acid", 40, 0, 0}
	var ketamine = Drug{"Ketamine", 100, 0, 1}
	var amphetamine = Drug{"Amphetamine", 60, 0, 3}
	var meth = Drug{"Meth", 150, 0, 5}
	var morphine = Drug{"Morphine", 80, 0, 5}
	var shrooms = Drug{"Shrooms", 30, 0, 1}	
	
	var knuckle = Weapon{"Knuckle", 0, 3, 100, 1, true, true, 1, false, 0, 0, 1}
	var knife = Weapon{"Knife", 100, 10, 100, 1, false, true, 0, true, 5, 50, 5}
	var baseballBat = Weapon{"Baseball Bat", 200, 20, 100, 1, false, true, 0, false, 0, 0, 1}
	var machete = Weapon{"Machete", 300, 30, 100, 1, false, true, 0, true, 15, 30, 2}
	var pistol = Weapon{"Pistol", 1200, 10, 80, 2, false, false, 0, false, 0, 0, 1}
	var SMG = Weapon{"SMG", 3000, 20, 50, 3, false, false, 0, false, 0, 0, 1}
	var shotgun = Weapon{"Shotgun", 2000, 30, 60, 1, false, false, 0, false, 0, 0, 1}
	var machineGun = Weapon{"Machine Gun", 5000, 40, 30, 4, true, false, 10, false, 0, 0, 1}
	var handgrenade = Weapon{"Handgrenade", 800, 50, 100, 1, false, false, 0, true, 30, 80, 8}

	var manhattan = District{"Manhattan", ["Brooklyn", "Queens"], [5]Drug{weed, cocaine, heroin, meth, ketamine}, true, true, false}
	var brooklyn = District{"Brooklyn", ["Staten Island", "Queens"], [5]Drug{amphetamine, meth, morphine, shrooms, heroin}, false, true, false}
	var queens = District{"Queens", ["Manhattan", "Bronx"], [5]Drug{weed, cocaine, heroin, acid, amphetamine}, true, false, false}
	var statenIsland = District{"Staten Island", ["Manhattan", "Brooklyn"], [5]Drug{weed, amphetamine, shrooms, acid, ketamine}, false, true, false}
	var bronx = District{"Bronx", ["Manhattan", "Queens"], [5]Drug{meth, morphine, heroin, shrooms, acid}, true, false, true}
	
}

func init() {
	fmt.Println("Welcome to Dope Wars!")
	fmt.Println("What is your name?")
	fmt.Scanln(&player.Name)
	fmt.Println("Welcome to the world of Dope Wars, " + player.Name + "!")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("You are a small time drug dealer in the city of New York.")
	fmt.Println("After failing one job after another, you have decided to start a small business. You have a small amount of cash, but you need to make a lot of money.")
	fmt.Println("After one of your drug deals went down, you were left with a debt.")
	fmt.Println("You have $" + strconv.Itoa(debt) + " to pay off.")
	fmt.Println("Press h for help. Press q to quit.")
	fmt.Scanln()
}
func keys() {
	fmt.println("h - help")
	fmt.println("q - quit")
	fmt.println("i - inventory")
	fmt.println("d - district info and the available drugs. Press d again to see the drugs in stock.\n to buy a drug, type the drug number and press enter.")
	fmt.println("t - travel to a district. Press t again to see the districts you can travel to.\n to travel to a district, type the district number and press enter.")
	fmt.println("w - weapon info and the available weapons. Press w again to see the weapons in stock.\n to buy a weapon, type the weapon number and press enter.")
	fmt.println("a - current weapon info. Press a again to see the current weapon stats. Press s to sell the current weapon.")
	fmt.println("f - fight the opponent. For throwable weapons, press j to throw the weapon. Note you will lose the weapon if you do not deal a critical hit\n or if it's a handgrenade.")
	fmt.println("s - sell the drugs. Type the drug number and press enter.")
	fmt.println("o - make a payment or withdraw/borrow money from the bank or loan shark. Type the amount and press enter.")
	fmt.println("r - run away. You might lose some cash or drugs and the wanted level will go down.")
	fmt.println("b - bribe the law enforcement. You will lose some cash and the wanted level will go down.")
	fmt.println("g - visit the bank or loan shark.")
	fmt.println("u - visit the hospital.")
	fmt.println("Press enter to continue.")
	fmt.Scanln()
}
func reputation(){

	switch {
	case player.Reputation > 0 && player.Reputation < 10:
		weaponsAvailable = [2]Weapon{knife, baseballBat}
		//a chance of 20% to multiply the price of up to 2 drugs in the inventory by 1.5
		if rand.Intn(100) < 30 {
			for i := 0; i < len(inventory.Drugs); i++ {
				if inventory.Drugs[i].Price > 0 {
					inventory.Drugs[i].Price = int(float64(inventory.Drugs[i].Price) * 1.5)
				}
			}
		}
	case player.Reputation > 10 && player.Reputation < 25:
		weaponsAvailable = [4]Weapon{knife, baseballBat, machete, pistol}
		//a chance of 40% to multiply the price of up to 3 drugs in the inventory by 1.5
		if rand.Intn(100) < 40 {
			for i := 0; i < len(inventory.Drugs); i++ {
				if inventory.Drugs[i].Price > 0 {
					inventory.Drugs[i].Price = int(float64(inventory.Drugs[i].Price) * 1.5)
				}
			}
	case player.Reputation > 25 && player.Reputation < 50:
		weaponsAvailable = [6]Weapon{knife, baseballBat, machete, pistol, SMG, shotgun}
		//a chance of 60% to multiply the price of up to 4 drugs in the inventory by 1.75
		if rand.Intn(100) < 60 {
			for i := 0; i < len(inventory.Drugs); i++ {
				if inventory.Drugs[i].Price > 0 {
					inventory.Drugs[i].Price = int(float64(inventory.Drugs[i].Price) * 1.75)
				}
			}
		}
	case player.Reputation > 50:
		weaponsAvailable = [8]Weapon{knife, baseballBat, machete, pistol, SMG, shotgun, machineGun, handgrenade}
		//a chance of 80% to multiply the price of up to 5 drugs in the inventory by 2
		if rand.Intn(100) < 80 {
			for i := 0; i < len(inventory.Drugs); i++ {
				if inventory.Drugs[i].Price > 0 {
					inventory.Drugs[i].Price = int(float64(inventory.Drugs[i].Price) * 2)
				}
			}
	}
}
	}

}

func buyDrug() {
	fmt.Println("You have $" + strconv.Itoa(inventory.cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("What drug would you like to buy?")
	fmt.Println(district.DrugsAvailable)
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("How many would you like to buy?")
	fmt.Scanln(&inventory.Drugs[0].Stock)
	fmt.Println("You have $" + strconv.Itoa(inventory.cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("You have bought " + strconv.Itoa(inventory.Drugs[0].Stock) + " " + inventory.Drugs[0].Name + ".")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
}
// sellDrug is a function that allows the player to sell drugs. Each sale will increase the player's reputation, but also increase the wanted level, multiplied by the amount of drugs sold.
func sellDrug() {
	fmt.Println("You have " + strconv.Itoa(inventory.Drugs[0].Stock) + " " + inventory.Drugs[0].Name + " to sell.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	// print the numbered list of drugs in the inventory with their current stock and price per unit
	for i := 0; i < len(inventory.Drugs); i++ {
		if inventory.Drugs[i].Stock > 0 {
			fmt.Println(strconv.Itoa(i+1) + ". " + inventory.Drugs[i].Name + " - " + strconv.Itoa(inventory.Drugs[i].Stock) + " units - $" + strconv.Itoa(inventory.Drugs[i].Price) + " per unit")
		}
	}
	fmt.Println("Which drug would you like to sell?.  Please type the number and press enter.")
	fmt.Scanln(&inventory.Drugs[0].Name)
	fmt.Println("How many would you like to sell?")
	fmt.Scanln(&unitsSell)

	if unitsSell > inventory.Drugs[0].Stock {
		fmt.Println("You don't have that many units to sell.")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	} else {
		inventory.Drugs[0].Stock -= unitsSell
		inventory.cash += unitsSell * inventory.Drugs[0].Price
		player.WantedLevel += inventory.Drugs[0].RaiseWanted * unitsSell
		fmt.Println("You have sold " + strconv.Itoa(unitsSell) + " " + inventory.Drugs[0].Name + ".")
		fmt.Println("You have" + strconv.Itoa(inventory.Drugs[0].Stock) + " " + inventory.Drugs[0].Name + " left.")
		fmt.Println("Your current cash is $" + strconv.Itoa(inventory.cash) + ".")
		fmt.Println("Your reputation has increased to " + strconv.Itoa(player.Reputation) + ".")
		fmt.Println("Your wanted level has increased to " + strconv.Itoa(player.WantedLevel) + ".")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	}
			//If the player has a reputation lower than 25, the reputation will increase by 4 for each 4 units sold.
			if player.Reputation < 25 {
				player.Reputation += 4 * (unitsSell / 4)
			} else {
				//If the player has a reputation higher than 25 and lower than 50, the reputation will increase by 3 for each 5 units sold.
				if player.Reputation > 25 && player.Reputation < 50 {
					player.Reputation += 3 * (unitsSell / 5)
				} else {
					//If the player has a reputation higher than 50, the reputation will increase by 2 for each 6 units sold.
					if player.Reputation > 50 {
						player.Reputation += 2 * (unitsSell / 6)
					}
				}
			}

	fmt.Println("Press enter to continue.")
	fmt.Scanln()
}