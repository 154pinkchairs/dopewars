package dopewars

import (
	"fmt"
	"os"
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
func execute() {
	//create a character
	//create a city
	//create a district
	//create a district
	main()
}
