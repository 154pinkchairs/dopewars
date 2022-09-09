package player

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Drug struct {
	Name        string
	Price       int
	RaiseWanted int
	Stock       int
}

type Drugs []Drug

var drugs = make([]Drug, 9)

func declare_drugs() {
	drugs[0].Name = "weed"
	//random number between 30 and 60
	drugs[0].Price = rand.Intn(30) + 30
	drugs[0].RaiseWanted = 1
	drugs[1].Name = "cocaine"
	drugs[1].Price = rand.Intn(200) + 200
	drugs[1].RaiseWanted = 4
	drugs[2].Name = "heroin"
	drugs[2].Price = rand.Intn(120) + 180
	drugs[2].RaiseWanted = 6
	drugs[3].Name = "acid"
	drugs[3].Price = rand.Intn(10) + 40
	drugs[3].RaiseWanted = 0
	drugs[4].Name = "ketamine"
	drugs[4].Price = rand.Intn(30) + 50
	drugs[4].RaiseWanted = 2
	drugs[5].Name = "amphetamine"
	drugs[5].Price = rand.Intn(35) + 45
	drugs[5].RaiseWanted = 3
	drugs[6].Name = "meth"
	drugs[6].Price = rand.Intn(30) + 60
	drugs[6].RaiseWanted = 5
	drugs[7].Name = "morphine"
	drugs[7].Price = rand.Intn(100) + 200
	drugs[7].RaiseWanted = 5
	drugs[8].Name = "shrooms"
	drugs[8].Price = rand.Intn(10) + 30
	drugs[8].RaiseWanted = 1
}

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
