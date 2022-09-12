package basegame

import (
	"fmt"
)

type District struct {
	name           string
	neighbour_a    []District
	neighbour_b    []District
	drugsAvailable Drugs
	hospital       bool
	bank           bool
	loanShark      bool
	ID             int
}

/*type District interface {
	Name() string
	neighbour_a() []District
	neighbour_b() []District
	//list of up to 5 drugs available in the district. This must be indexable
	drugsAvailable() Drugs
	hospital() bool
	bank() bool
	loanShark() bool
	starting() bool
	ID() int
	Properties() districtProperties
}
*/

type dist struct {
	properties districtProperties
	ID         int
}

func travel(c *Character, d *District) {
	//update neighbour_a and neighbour_b in districtProperties for each district
	manhattan.properties.neighbour_a = brooklyn
	manhattan.properties.neighbour_b = queens

	currentDistrict := c.CurrentDistrict
	//read the t keypress
	//the Player can travel to neighbour_a or neighbour_b
	fmt.Println("Where would you like to travel to? Type 1 or 2 and press enter.")
	var travelChoice int
	fmt.Println("1. " + currentDistrict.neighbour_a.Name)
	fmt.Println("2. " + currentDistrict.Properties.neighbour_b.Name)
	fmt.Scanln("%s", &travelChoice)
	//if the Player selects 1, travel to neighbour_a
	if travelChoice == 1 {
		c.CurrentDistrict = currentDistrict.properties.neighbour_a()[0]
	} else {
		//if the Player selects 2, travel to neighbour_b
		c.CurrentDistrict = currentDistrict.neighbour_b()[0]
	}
	fmt.Println("You have arrived at " + c.CurrentDistrict.Name + ".")
}
