package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var bg *ebiten.Image
var loadsave *ebiten.Image
var newgameimg *ebiten.Image
var donate *ebiten.Image
var issues *ebiten.Image
var quitimg *ebiten.Image

func init() {
	var err error
	var err2 error
	var err3 error
	var err4 error
	var err5 error
	var err6 error
	bg, _, err2 = ebitenutil.NewImageFromFile("assets/menu_bg.png")
	if err2 != nil {
		log.Fatal(err2)
	}
	newgameimg, _, err3 = ebitenutil.NewImageFromFile("assets/newgame.png")
	if err3 != nil {
		log.Fatal(err3)
	}
	loadsave, _, err4 = ebitenutil.NewImageFromFile("assets/loadsave.png")
	if err4 != nil {
		log.Fatal(err4)
	}
	donate, _, err5 = ebitenutil.NewImageFromFile("assets/donate.png")
	if err5 != nil {
		log.Fatal(err5)
	}
	issues, _, err6 = ebitenutil.NewImageFromFile("assets/issues.png")
	if err6 != nil {
		log.Fatal(err6)
	}
	quitimg, _, err = ebitenutil.NewImageFromFile("assets/quit.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)
	pos1 := &ebiten.DrawImageOptions{}
	pos1.GeoM.Translate(340, 150)
	screen.DrawImage(newgameimg, pos1)
	pos2 := &ebiten.DrawImageOptions{}
	pos2.GeoM.Translate(340, 200)
	screen.DrawImage(loadsave, pos2)
	pos3 := &ebiten.DrawImageOptions{}
	pos3.GeoM.Translate(340, 250)
	screen.DrawImage(donate, pos3)
	pos4 := &ebiten.DrawImageOptions{}
	pos4.GeoM.Translate(340, 300)
	screen.DrawImage(issues, pos4)
	pos5 := &ebiten.DrawImageOptions{}
	pos5.GeoM.Translate(340, 350)
	screen.DrawImage(quitimg, pos5)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 540
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Dopewars 2D")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

/*func execute() {
	//create a character
	//create a city
	//create a district
	//create a district
	main()
	//basegame.Player.Character()
}
*/
