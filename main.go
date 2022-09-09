package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/154pinkchairs/dopewars/drugs"
)

//var weed = Drug{"Weed", 50, 0, 2}
//var cocaine = Drug{"Cocaine", 300, 0, 4}
//var heroin = Drug{"Heroin", 200, 0, 6}
//var acid = Drug{"Acid", 40, 0, 0}
//var ketamine = Drug{"Ketamine", 100, 0, 1}
//var amphetamine = Drug{"Amphetamine", 60, 0, 3}
//var meth = Drug{"Meth", 150, 0, 5}
//var morphine = Drug{"Morphine", 80, 0, 5}
//var shrooms = Drug{"Shrooms", 30, 0, 1}

type districtProperties struct {
	name           string
	neighbour_a    district
	neighbour_b    district
	drugsAvailable drugs.Drugs
	hospital       bool
	bank           bool
	loanShark      bool
}
type district interface {
	Name() string
	neighbour_a() []district
	neighbour_b() []district
	//list of up to 5 drugs available in the district. This must be indexable
	drugsAvailable() Drugs
	hospital() bool
	bank() bool
	loanShark() bool
	starting() bool
	ID() int
	Properties() districtProperties
}

type dist struct {
	properties districtProperties
	ID         int
}

//create a manhattan struct. The drugs can be the same for each, they should be updated upon drugs_available() call
//var manhattan = dist{ districtProperties{"Manhattan", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 0}
//var brooklyn = dist{ districtProperties{"Brooklyn", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 1}
//var queens = dist{ districtProperties{"Queens", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 2}
//var bronx = dist{ districtProperties{"Bronx", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 3}
//var statenIsland = dist{ districtProperties{"Staten Island", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 4}
//create a manhattan array

func drugs_available() {
	//get up to 5 random drugs from the drugs array.
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		randIndex := rand.Intn(len(drugs))
		//check if the drug is already in the array
		//if it is, generate a new random number
		//if it is not, add it to the array
		if drugs[randIndex].Name == Player.CurrentDistrict.drugsAvailable[i].Name {
			randIndex = rand.Intn(len(drugs))
		} else {
			Player.CurrentDistrict.drugsAvailable[i] = drugs[randIndex]
		}
	}
}

func main() {
	fmt.Println("Welcome to the city of New York.")
	err := os.Remove("save.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	Player.cash = 10000
	Player.debt = 15000
	fmt.Println("Welcome to Dope Wars!")
	fmt.Println("What is your name?")
	fmt.Scanln(&Player.Name)
	fmt.Println("Welcome to the world of Dope Wars, " + Player.Name + "!")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("You are a small time drug dealer in the city of New York.\n After failing one job after another, you have decided to start a small business.")
	fmt.Println("You have a small amount of cash, but you need to make a lot of money.\nAfter one of your drug deals went down, you were left with a debt.")
	fmt.Println("You have $" + strconv.Itoa(Player.debt) + " to pay off.")
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

func reputation() {
	switch {
	case Player.Reputation > 0 && Player.Reputation < 10:
		Player.weaponsAvailable = [8]Weapon{knife, baseballBat, knuckle, knuckle, knuckle, knuckle, knuckle, knuckle}
		//a chance of 20% to multiply the price of up to 2 drugs in the Player by 1.5
		if rand.Intn(100) < 30 {
			for i := 0; i < len(Player.drugs); i++ {
				if Player.drugs[i].Price > 0 {
					Player.drugs[i].Price = int(float64(Player.drugs[i].Price) * 1.5)
				}
			}
		}
	case Player.Reputation > 10 && Player.Reputation < 25:
		Player.weaponsAvailable = [8]Weapon{knife, baseballBat, machete, pistol, knuckle, knuckle, knuckle, knuckle}
		//a chance of 40% to multiply the price of up to 3 drugs in the Player by 1.5
		if rand.Intn(100) < 40 {
			for i := 0; i < len(Player.drugs); i++ {
				if Player.drugs[i].Price > 0 {
					Player.drugs[i].Price = int(float64(Player.drugs[i].Price) * 1.5)
				}
			}
		}
	case Player.Reputation > 25 && Player.Reputation < 50:
		Player.weaponsAvailable = [8]Weapon{knife, baseballBat, machete, pistol, SMG, shotgun, knuckle, knuckle}
		//a chance of 60% to multiply the price of up to 4 drugs in the Player by 1.75
		if rand.Intn(100) < 60 {
			for i := 0; i < len(Player.drugs); i++ {
				if Player.drugs[i].Price > 0 {
					Player.drugs[i].Price = int(float64(Player.drugs[i].Price) * 1.75)
				}
			}
		}
	case Player.Reputation > 50:
		Player.weaponsAvailable = [8]Weapon{knife, baseballBat, machete, pistol, SMG, shotgun, machineGun, handgrenade}
		//a chance of 80% to multiply the price of up to 5 drugs in the Player by 2
		if rand.Intn(100) < 80 {
			for i := 0; i < len(Player.drugs); i++ {
				if Player.drugs[i].Price > 0 {
					Player.drugs[i].Price = int(float64(Player.drugs[i].Price) * 2)
				}
			}
		}
	}
}

func buyDrug() {
	fmt.Println("You have $" + strconv.Itoa(Player.cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("What drug would you like to buy?")
	//prints the drugs in the current district
	//if the drug is not available, it will not be printed
	//get the current district

	fmt.Println(district.drugsAvailable)
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("How many would you like to buy?")
	fmt.Scanln(&Player.drugs[0].Stock)
	fmt.Println("You have $" + strconv.Itoa(Player.cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("You have bought " + strconv.Itoa(Player.drugs[0].Stock) + " " + Player.drugs[0].Name + ".")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
}

// sellDrug is a function that allows the Player to sell drugs. Each sale will increase the Player's reputation, but also increase the wanted level, multiplied by the amount of d sold.
func sellDrug() {
	fmt.Println("You have " + strconv.Itoa(Player.drugs[0].Stock) + " " + Player.drugs[0].Name + " to sell.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	// print the numbered list of drugs in the Player with their current stock and price per unit
	for i := 0; i < len(Player.drugs); i++ {
		if Player.drugs[i].Stock > 0 {
			fmt.Println(strconv.Itoa(i+1) + ". " + Player.drugs[i].Name + " - " + strconv.Itoa(Player.drugs[i].Stock) + " units - $" + strconv.Itoa(Player.drugs[i].Price) + " per unit")
		}
	}
	fmt.Println("Which drug would you like to sell?.  Please type the number and press enter.")
	fmt.Scanln(&Player.drugs[0].Name)
	fmt.Println("How many would you like to sell?")
	var unitsSell int
	fmt.Scanln("%d", &unitsSell)

	if unitsSell > Player.drugs[0].Stock {
		fmt.Println("You don't have that many units to sell.")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	} else {
		Player.drugs[0].Stock -= unitsSell
		Player.cash += unitsSell * Player.drugs[0].Price
		Player.WantedLevel += Player.drugs[0].RaiseWanted * unitsSell
		fmt.Println("You have sold " + strconv.Itoa(unitsSell) + " " + Player.drugs[0].Name + ".")
		fmt.Println("You have" + strconv.Itoa(Player.drugs[0].Stock) + " " + Player.drugs[0].Name + " left.")
		fmt.Println("Your current cash is $" + strconv.Itoa(Player.cash) + ".")
		fmt.Println("Your reputation has increased to " + strconv.Itoa(Player.Reputation) + ".")
		fmt.Println("Your wanted level has increased to " + strconv.Itoa(Player.WantedLevel) + ".")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	}
	//If the Player has a reputation lower than 25, the reputation will increase by 4 for each 4 units sold.
	if Player.Reputation < 25 {
		Player.Reputation += 4 * (unitsSell / 4)
	} else {
		//If the Player has a reputation higher than 25 and lower than 50, the reputation will increase by 3 for each 5 units sold.
		if Player.Reputation > 25 && Player.Reputation < 50 {
			Player.Reputation += 3 * (unitsSell / 5)
		} else {
			//If the Player has a reputation higher than 50, the reputation will increase by 2 for each 6 units sold.
			if Player.Reputation > 50 {
				Player.Reputation += 2 * (unitsSell / 6)
			}
		}
	}

	fmt.Println("Press enter to continue.")
	fmt.Scanln()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func travel() {
	//update neighbour_a and neighbour_b in districtProperties for each district
	manhattan.properties.neighbour_a = brooklyn
	manhattan.properties.neighbour_b = queens

	currentDistrict := Player.CurrentDistrict
	//read the t keypress
	//the Player can travel to neighbour_a or neighbour_b
	fmt.Println("Where would you like to travel to? Type 1 or 2 and press enter.")
	var travelChoice int
	fmt.Println("1. " + currentDistrict.Properties.neighbour_a.name)
	fmt.Println("2. " + currentDistrict.Properties.neighbour_b.Name)
	fmt.Scanln("%s", &travelChoice)
	//if the Player selects 1, travel to neighbour_a
	if travelChoice == 1 {
		Player.CurrentDistrict = currentDistrict.properties.neighbour_a()[0]
	} else {
		//if the Player selects 2, travel to neighbour_b
		Player.CurrentDistrict = currentDistrict.neighbour_b()[0]
	}
	fmt.Println("You have arrived at " + Player.District.Name + ".")
}
