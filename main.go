package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Welcome to the city of New York.")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 854, 480
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Dopewars 2D")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
func execute() {
	//create a character
	//create a city
	//create a district
	//create a district
	main()
	//basegame.Player.Character()
}
