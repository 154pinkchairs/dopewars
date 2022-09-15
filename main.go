package main

import (
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var bg *ebiten.Image
var loadsave *ebiten.Image
var newgameimg *ebiten.Image
var donate *ebiten.Image
var issues *ebiten.Image
var quitimg *ebiten.Image
var loadsave_hoover *ebiten.Image
var newgameimg_hoover *ebiten.Image
var donate_hoover *ebiten.Image
var issues_hoover *ebiten.Image
var quitimg_hoover *ebiten.Image

func init() {
	var err error
	var err2 error
	var err3 error
	var err4 error
	var err5 error
	var err6 error
	var err7 error
	var err8 error
	var err9 error
	var err10 error
	var err11 error
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
	loadsave_hoover, _, err7 = ebitenutil.NewImageFromFile("assets/loadsave_hoover.png")
	if err7 != nil {
		log.Fatal(err7)
	}
	newgameimg_hoover, _, err8 = ebitenutil.NewImageFromFile("assets/newgame_hoover.png")
	if err8 != nil {
		log.Fatal(err8)
	}
	donate_hoover, _, err9 = ebitenutil.NewImageFromFile("assets/donate_hoover.png")
	if err9 != nil {
		log.Fatal(err9)
	}
	issues_hoover, _, err10 = ebitenutil.NewImageFromFile("assets/issues_hoover.png")
	if err10 != nil {
		log.Fatal(err10)
	}
	quitimg_hoover, _, err11 = ebitenutil.NewImageFromFile("assets/quit_hoover.png")
	if err11 != nil {
		log.Fatal(err11)
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
	if mouseOverButton(340, 150, 200, 50) {
		screen.DrawImage(newgameimg_hoover, pos1)
	}

	pos2 := &ebiten.DrawImageOptions{}
	pos2.GeoM.Translate(340, 200)
	screen.DrawImage(loadsave, pos2)
	if mouseOverButton(340, 200, 200, 50) {
		screen.DrawImage(loadsave_hoover, pos2)
	}

	pos3 := &ebiten.DrawImageOptions{}
	pos3.GeoM.Translate(340, 250)
	screen.DrawImage(donate, pos3)
	if mouseOverButton(340, 250, 200, 50) {
		screen.DrawImage(donate_hoover, pos3)
	}

	pos4 := &ebiten.DrawImageOptions{}
	pos4.GeoM.Translate(340, 300)
	screen.DrawImage(issues, pos4)
	if mouseOverButton(340, 300, 200, 50) {
		screen.DrawImage(issues_hoover, pos4)
	}

	pos5 := &ebiten.DrawImageOptions{}
	pos5.GeoM.Translate(340, 350)
	screen.DrawImage(quitimg, pos5)
	if mouseOverButton(340, 350, 280, 50) {
		//dump the old image and draw the hoover image
		//screen.Clear()
		screen.DrawImage(quitimg_hoover, pos5)
	}
}

// create a function that checks if the mouse is over a button
func mouseOverButton(x, y, width, height int) bool {
	//get the mouse position
	mouseX, mouseY := ebiten.CursorPosition()
	//check if the mouse is within the button's x and y bounds
	if mouseX >= x && mouseX <= x+width {
		if mouseY >= y && mouseY <= y+height {
			return true
		}
	}
	return false
}

func (g *Game) Activate() error {
	//check if the left mouse button is pressed
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		//check if the mouse is over the button
		if mouseOverButton(340, 150, 280, 50) {
			//do something
		}
		if mouseOverButton(340, 200, 280, 50) {
			//do something
		}
		if mouseOverButton(340, 250, 280, 50) {
			//do something
		}
		if mouseOverButton(340, 300, 280, 50) {
			//open the issues page in the browser
			os.Open("https://github.com/154pinkchairs/dopewars/issues")
		}
		if mouseOverButton(340, 350, 280, 50) {
			//exit the game
			os.Exit(0)
		}
	}
	return nil
}

// // create the districts and their properties just like

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 540
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowResizable(true)
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
