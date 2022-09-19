package basegame

import (
	"fmt"
)

type District struct {
	Name           string
	DrugsAvailable Drugs
	ID             int
	Properties     DistrictProperties
}

type DistrictProperties struct {
	Hospital     bool
	Bank         bool
	LoanShark    bool
	NeighbourIDs []int
}

// create the Districts and their Properties just like in New York
var manhattan = District{
	Name:           "Manhattan",
	DrugsAvailable: make(Drugs, 6),
	ID:             0,
	Properties: DistrictProperties{
		Hospital:     true,
		Bank:         true,
		LoanShark:    true,
		NeighbourIDs: []int{1, 2, 3},
	},
}
var brooklyn = District{
	Name:           "Brooklyn",
	DrugsAvailable: make(Drugs, 6),
	ID:             1,
	Properties: DistrictProperties{
		Hospital:     true,
		Bank:         false,
		LoanShark:    false,
		NeighbourIDs: []int{0, 2, 4},
	},
}
var queens = District{
	Name:           "Queens",
	DrugsAvailable: make(Drugs, 6),
	ID:             2,
	Properties: DistrictProperties{
		Hospital:     false,
		Bank:         true,
		LoanShark:    false,
		NeighbourIDs: []int{0, 1, 3},
	},
}
var Bronx = District{
	Name:           "Bronx",
	DrugsAvailable: make(Drugs, 6),
	ID:             3,
	Properties: DistrictProperties{
		Hospital:     true,
		Bank:         false,
		LoanShark:    true,
		NeighbourIDs: []int{0, 2},
	},
}
var statenIsland = District{
	Name:           "Staten Island",
	DrugsAvailable: make(Drugs, 6),
	ID:             4,
	Properties: DistrictProperties{
		Hospital:     false,
		Bank:         true,
		LoanShark:    false,
		NeighbourIDs: []int{1},
	},
}

var Districts = []District{manhattan, brooklyn, queens, Bronx, statenIsland}

/*type District interface {
	Name() string
	neighbour_a() []District
	neighbour_b() []District
	//list of up to 5 drugs available in the District. This must be indexable
	DrugsAvailable() Drugs
	Hospital() bool
	Bank() bool
	LoanShark() bool
	starting() bool
	ID() int
	Properties() DistrictProperties
}
*/
// print a numbered list of the Districts the player can travel with their Names to according to their IDs. The length of the list will change depending on the length of the player's current District's NeighbourIDs slice. A switch case will be used, where the cases are currentDistrict.ID and the cases are the IDs of the Districts in the NeighbourIDs slice. The player will be able to travel to any of the Districts in the NeighbourIDs slice.
func (d District) PrintNeighbours() {
	fmt.Println("You can travel to the following Districts:")
	for i, v := range d.Properties.NeighbourIDs {
		switch v {
		case 0:
			fmt.Printf("%d. %s", i+1, manhattan.Name)
		case 1:
			fmt.Printf("%d. %s", i+1, brooklyn.Name)
		case 2:
			fmt.Printf("%d. %s", i+1, queens.Name)
		case 3:
			fmt.Printf("%d. %s", i+1, Bronx.Name)
		case 4:
			fmt.Printf("%d. %s", i+1, statenIsland.Name)
		}
	}
}

func travel(c *Character, d *District) {

	currentDistrict := c.CurrentDistrict
	currentDistrict.PrintNeighbours()
	fmt.Println("Type the number of the District you want to travel to. To cancel, press any other numeric key.")
	var travelChoice int
	fmt.Scan(&travelChoice)
	if travelChoice > len(currentDistrict.Properties.NeighbourIDs) || travelChoice < 1 {
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
				c.CurrentDistrict = Bronx
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
				c.CurrentDistrict = Bronx
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
	fmt.Println("You have arrived at " + c.CurrentDistrict.Name + ".")
}
