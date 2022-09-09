package main

import (
	"fmt"
	"os"
	p "player"
	"strconv"
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

//create a manhattan struct. The drugs can be the same for each, they should be updated upon drugs_available() call
//var manhattan = dist{ districtProperties{"Manhattan", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 0}
//var brooklyn = dist{ districtProperties{"Brooklyn", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 1}
//var queens = dist{ districtProperties{"Queens", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 2}
//var bronx = dist{ districtProperties{"Bronx", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 3}
//var statenIsland = dist{ districtProperties{"Staten Island", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 4}
//create a manhattan array

func main() {
	fmt.Println("Welcome to the city of New York.")
	err := os.Remove("save.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	p.Player.cash = 10000
	p.Player.debt = 15000
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
