package player

type Player struct {
	Name                    string
	Health                  int
	Reputation, WantedLevel int
	cash, debt              int
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:        name,
		Health:      100,
		Reputation:  0,
		WantedLevel: 0,
		cash:        5000,
		debt:        10000,
	}
}

func ModPlayer(reputation int, wantedLevel int, cash int, debt int) *Player {
	return &Player{
		Reputation:  reputation,
		WantedLevel: wantedLevel,
		cash:        cash,
		debt:        debt,
	}
}
