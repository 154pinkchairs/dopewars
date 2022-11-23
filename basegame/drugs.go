package basegame

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Drug struct {
	Name        string
	Price       int
	RaiseWanted int
	Stock       int
}

type Drugs []Drug

var drugs = make([]Drug, 9)

func declareDrugs() {
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
	drugs[9].Name = "crack"
	drugs[9].Price = rand.Intn(75) + 200
	drugs[9].RaiseWanted = 7
}

//TODO: declare a function that is a merger of randomizeDrugs and declareDrugs

// rewrite the next function with less cyclomatic complexity
func (dr *Drug) randomizeAvailability(d *District) {
	// poorer districts have less more crack, meth, and heroin
	// richer districts have more acid, cocaine and shrooms
	// middle class districts have more weed, ketamine, and amphetamine
	//Wealth scale 1-5

	//So to sum 2 things influence the presence of certain drugs in the d.DrugsAvailable array:
	//1. The wealth of the district (d.Properties.Wealth)
	//2. The random number generator
	//weed always present
	d.DrugsAvailable[0] = drugs[0]
	rng := rand.Intn(100)
	if (rng > 80 && d.Properties.Wealth == 5) || (rng < 70 && d.Properties.Wealth == 4) || (rng < 50 && d.Properties.Wealth == 3 || rng < 30 && d.Properties.Wealth < 3) {
		d.DrugsAvailable = append(d.DrugsAvailable, drugs[1])
		d.DrugsAvailable = append(d.DrugsAvailable, drugs[8])
		d.DrugsAvailable = append(d.DrugsAvailable, drugs[4])
	} else if (rng > 80 && d.Properties.Wealth == 1) || (rng < 70 && d.Properties.Wealth == 2) || (rng < 50 && d.Properties.Wealth == 3 || rng < 30 && d.Properties.Wealth > 3) {
		d.DrugsAvailable = append(d.DrugsAvailable, drugs[2])
		d.DrugsAvailable = append(d.DrugsAvailable, drugs[6])
		d.DrugsAvailable = append(d.DrugsAvailable, drugs[9])
	} else if (rng > 80 && d.Properties.Wealth == 3) || (rng < 70 && d.Properties.Wealth == 4) || (rng < 50 && d.Properties.Wealth == 5 || rng < 30 && d.Properties.Wealth < 2) {
		d.DrugsAvailable = append(d.DrugsAvailable, drugs[3])
		d.DrugsAvailable = append(d.DrugsAvailable, drugs[5])
		d.DrugsAvailable = append(d.DrugsAvailable, drugs[7])
	} else {
		//append up 5 random other drugs to the array if it's not fully populated
		for len(d.DrugsAvailable) < 6 {
			rng := rand.Intn(9)
			if d.DrugsAvailable[rng] != drugs[rng] {
				d.DrugsAvailable = append(d.DrugsAvailable, drugs[rng])
			}
		}
	}
}

func (c *Character) buyDrug() {
	fmt.Println("You have $" + strconv.Itoa(c.Cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("What drug would you like to buy?")
	//prints the drugs in the current district
	//if the drug is not available, it will not be printed
	//get the current district

	fmt.Println(c.CurrentDistrict.DrugsAvailable)
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("How many would you like to buy?")
	fmt.Scanln(&c.drugs[0].Stock)
	fmt.Println("You have $" + strconv.Itoa(c.Cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("You have bought " + strconv.Itoa(c.drugs[0].Stock) + " " + c.drugs[0].Name + ".")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
}

// sellDrug is a function that allows the character to sell drugs. Each sale will increase the character's reputation, but also increase the wanted level, multiplied by the amount of d sold.
func (c *Character) sellDrug() {
	fmt.Println("You have " + strconv.Itoa(c.drugs[0].Stock) + " " + c.drugs[0].Name + " to sell.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	// print the numbered list of drugs in the character with their current stock and price per unit
	for i := 0; i < len(c.drugs); i++ {
		if c.drugs[i].Stock > 0 {
			fmt.Println(strconv.Itoa(i+1) + ". " + c.drugs[i].Name + " - " + strconv.Itoa(c.drugs[i].Stock) + " units - $" + strconv.Itoa(c.drugs[i].Price) + " per unit")
		}
	}
	fmt.Println("Which drug would you like to sell?.  Please type the number and press enter.")
	fmt.Scanln(&c.drugs[0].Name)
	fmt.Println("How many would you like to sell?")
	var unitsSell int
	fmt.Scanln("%d", &unitsSell)

	if unitsSell > c.drugs[0].Stock {
		fmt.Println("You don't have that many units to sell.")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	} else {
		c.drugs[0].Stock -= unitsSell
		c.Cash += unitsSell * c.drugs[0].Price
		c.WantedLevel += c.drugs[0].RaiseWanted * unitsSell
		fmt.Println("You have sold " + strconv.Itoa(unitsSell) + " " + c.drugs[0].Name + ".")
		fmt.Println("You have" + strconv.Itoa(c.drugs[0].Stock) + " " + c.drugs[0].Name + " left.")
		fmt.Println("Your current Cash is $" + strconv.Itoa(c.Cash) + ".")
		fmt.Println("Your reputation has increased to " + strconv.Itoa(c.Reputation) + ".")
		fmt.Println("Your wanted level has increased to " + strconv.Itoa(c.WantedLevel) + ".")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	}
	//If the character has a reputation lower than 25, the reputation will increase by 4 for each 4 units sold.
	if c.Reputation < 25 {
		c.Reputation += 4 * (unitsSell / 4)
	} else {
		//If the character has a reputation higher than 25 and lower than 50, the reputation will increase by 3 for each 5 units sold.
		if c.Reputation > 25 && c.Reputation < 50 {
			c.Reputation += 3 * (unitsSell / 5)
		} else {
			//If the character has a reputation higher than 50, the reputation will increase by 2 for each 6 units sold.
			if c.Reputation > 50 {
				c.Reputation += 2 * (unitsSell / 6)
			}
		}
	}

	fmt.Println("Press enter to continue.")
	fmt.Scanln()
}
