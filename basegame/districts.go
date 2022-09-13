package basegame

import (
	"fmt"
)

type District struct {
	name           string
	drugsAvailable Drugs
	ID             int
	properties     districtProperties
}

type districtProperties struct {
	hospital     bool
	bank         bool
	loanShark    bool
	neighbourids []int
}

// create the districts and their properties just like in New York
var manhattan = District{
	name:           "Manhattan",
	drugsAvailable: make(Drugs, 5),
	ID:             0,
	properties: districtProperties{
		hospital:     true,
		bank:         true,
		loanShark:    true,
		neighbourids: []int{1, 2, 3},
	},
}
var brooklyn = District{
	name:           "Brooklyn",
	drugsAvailable: make(Drugs, 5),
	ID:             1,
	properties: districtProperties{
		hospital:     true,
		bank:         false,
		loanShark:    false,
		neighbourids: []int{0, 2, 4},
	},
}
var queens = District{
	name:           "Queens",
	drugsAvailable: make(Drugs, 5),
	ID:             2,
	properties: districtProperties{
		hospital:     false,
		bank:         true,
		loanShark:    false,
		neighbourids: []int{0, 1, 3},
	},
}
var bronx = District{
	name:           "Bronx",
	drugsAvailable: make(Drugs, 5),
	ID:             3,
	properties: districtProperties{
		hospital:     true,
		bank:         false,
		loanShark:    true,
		neighbourids: []int{0, 2},
	},
}
var statenIsland = District{
	name:           "Staten Island",
	drugsAvailable: make(Drugs, 5),
	ID:             4,
	properties: districtProperties{
		hospital:     false,
		bank:         true,
		loanShark:    false,
		neighbourids: []int{1},
	},
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
// print a numbered list of the districts the player can travel with their names to according to their IDs. The length of the list will change depending on the length of the player's current district's neighbourids slice. A switch case will be used, where the cases are currentDistrict.ID and the cases are the IDs of the districts in the neighbourids slice. The player will be able to travel to any of the districts in the neighbourids slice.
func (d District) PrintNeighbours() {
	fmt.Println("You can travel to the following districts:")
	for i, v := range d.properties.neighbourids {
		switch v {
		case 0:
			fmt.Printf("%d. %s", i+1, manhattan.name)
		case 1:
			fmt.Printf("%d. %s", i+1, brooklyn.name)
		case 2:
			fmt.Printf("%d. %s", i+1, queens.name)
		case 3:
			fmt.Printf("%d. %s", i+1, bronx.name)
		case 4:
			fmt.Printf("%d. %s", i+1, statenIsland.name)
		}
	}
}

func travel(c *Character, d *District) {

	currentDistrict := c.CurrentDistrict
	currentDistrict.PrintNeighbours()
	fmt.Println("Type the number of the district you want to travel to. To cancel, press any other numeric key.")
	var travelChoice int
	fmt.Scan(&travelChoice)
	if travelChoice > len(currentDistrict.properties.neighbourids) || travelChoice < 1 {
		fmt.Println("You have cancelled your travel.")
		return
	} else {
		switch c.CurrentDistrict.ID {
		case 0:
			switch travelChoice {
			case 1:
				c.CurrentDistrict = brooklyn
			case 2:
				c.CurrentDistrict = queens
			case 3:
				c.CurrentDistrict = bronx
			}
		case 1:
			switch travelChoice {
			case 1:
				c.CurrentDistrict = manhattan
			case 2:
				c.CurrentDistrict = queens
			case 3:
				c.CurrentDistrict = statenIsland
			}
		case 2:
			switch travelChoice {
			case 1:
				c.CurrentDistrict = manhattan
			case 2:
				c.CurrentDistrict = brooklyn
			case 3:
				c.CurrentDistrict = bronx
			}
		case 3:
			switch travelChoice {
			case 1:
				c.CurrentDistrict = manhattan
			case 2:
				c.CurrentDistrict = queens
			}
		case 4:
			switch travelChoice {
			case 1:
				c.CurrentDistrict = brooklyn
			}
		}
	}

	fmt.Scanln("%s", &travelChoice)
	fmt.Println("You have arrived at " + c.CurrentDistrict.name + ".")
}
