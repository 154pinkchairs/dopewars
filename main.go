package main

import (
	"dopewars/basegame"
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/components"
	"golang.org/x/exp/shiny/screen"
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

type Game struct {
	init   bool
	gameUI *furex.View
	screen screen.Screen
}

func (g *Game) Update() error {
	if !g.init {
		g.init = true
		g.setupUI()
	}
	g.gameUI.UpdateWithSize(ebiten.WindowSize())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)
	g.gameUI.Draw(screen)
	pos1 := &ebiten.DrawImageOptions{}
	pos1.GeoM.Translate(340, 150)
	screen.DrawImage(newgameimg, pos1)
	if mouseOverButton(340, 150, 200, 50) {
		screen.DrawImage(newgameimg_hoover, pos1)
	} 
	// add a handler for the new game button using MouseleftButtonHandler from furex


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
		screen.DrawImage(quitimg_hoover, pos5)
	}
}

func (g *Game) setupUI() {


	newGameBtn := func() *furex.View {
		return (&furex.View{
			Top:          210,
			Left:         340,
			Width:        200,
			Height:       40,
			MarginLeft:   360,
			MarginTop:    25,
			MarginRight:  5,
			MarginBottom: 5,
			Handler:      &components.Button{Text: "", OnClick: func() { basegame.NewGame() }},
		})
	}

	loadSaveBtn := func() *furex.View {
		return (&furex.View{
			Top:          250,
			Left:         340,
			Width:        235,
			Height:       40,
			MarginLeft:   340,
			MarginTop:    5,
			MarginRight:  5,
			MarginBottom: 5,
			Handler:      &components.Button{Text: "", OnClick: func() { basegame.Loadsave(&basegame.Character{}) }},
		})
	}

	donateBtn := func() *furex.View {
		return (&furex.View{
			Top:          290,
			Left:         340,
			Width:        120,
			Height:       40,
			MarginLeft:   390,
			MarginTop:    5,
			MarginRight:  5,
			MarginBottom: 5,
			Handler:      &components.Button{Text: "", OnClick: func() { openbrowser("https://www.liberapay.com/") }},
		})
	}

	issuesBtn := func() *furex.View {
		return (&furex.View{
			Top:          330,
			Left:         340,
			Width:        200,
			Height:       40,
			MarginLeft:   360,
			MarginTop:    5,
			MarginRight:  5,
			MarginBottom: 5,
			Handler:      &components.Button{Text: "", OnClick: func() { openbrowser("https://github.com/154pinkchairs/dopewars/issues") }},
			Wrap:         furex.NoWrap,
		})
	}

	quitBtn := func() *furex.View {
		return (&furex.View{
			Top:          370,
			Left:         300,
			Width:        110,
			Height:       40,
			MarginLeft:   400,
			MarginTop:    5,
			MarginRight:  5,
			MarginBottom: 285,
			Handler:      &components.Button{Text: "", OnClick: func() { os.Exit(0) }},
		})
	}

	g.gameUI = (&furex.View{
		Width:        960,
		Height:       540,
		Direction:    furex.Column,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemStart, //place items in the center, one below the other
		AlignContent: furex.AlignContentCenter,
		Wrap:         furex.NoWrap,
	}).AddChild(
		(&furex.View{
			Width:      640,
			Height:     200,
			Justify:    furex.JustifySpaceBetween,
			AlignItems: furex.AlignItemCenter,
		}).AddChild(
			&furex.View{
				Width:   100,
				Height:  5,
				Handler: &components.Box{Color: color.RGBA{0, 0, 0, 0}},
			},
			&furex.View{
				Width:   200,
				Height:  5,
				Handler: &components.Box{Color: color.RGBA{0, 0, 0, 0}},
			},
			&furex.View{
				Width:   200,
				Height:  5,
				Handler: &components.Box{Color: color.RGBA{0, 0, 0, 0}},
			},
			&furex.View{
				Width:   100,
				Height:  5,
				Handler: &components.Box{Color: color.RGBA{0, 0, 0, 0}},
			},
			&furex.View{
				Width:   100,
				Height:  5,
				Handler: &components.Box{Color: color.RGBA{0, 0, 0, 0}},
			},
		),
	).AddChild(&furex.View{
		Width:      960,
		Height:     60,
		Justify:    furex.JustifyCenter,
		AlignItems: furex.AlignItemEnd,
	}).AddChild(
		newGameBtn(),
		loadSaveBtn(),
		donateBtn(),
		issuesBtn(),
		quitBtn(),
	)
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



/* Add a handle from inpututil interface. Define button boxes (boundaries) for each pos. Button at pos1 calls basegame.run().
Button at pos2 calls basegame.Loadsave() and then basegame.NewGame().
Button at pos3 opens https://liberapay.com/ in browser. Add a inline comment that this has to be changed to the actual donation link.
Button at pos4 opens https://github.com/154pinkchairs/dopewars/issues in the browser.
Button at pos5 calls os.Exit(0) */

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

// // create the districts and their properties just like

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 540
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Dopewars 2D")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
