package player

import (
	"math/rand"
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
