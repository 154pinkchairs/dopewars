package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/154pinkchairs/dopewars2d/basegame"
	"github.com/154pinkchairs/dopewars2d/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/components"
	"golang.org/x/exp/shiny/screen"
)

type Game struct {
	init        bool
	gameUI      *furex.View
	screen      screen.Screen
	Character   basegame.Character
	txtRenderer *etxt.Renderer
	CG          core.Game
	//must implement ebiten.Game interface
	ebiten.Game
}

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
	bg, _, _ = ebitenutil.NewImageFromFile("assets/menu_bg.png")
	newgameimg, _, _ = ebitenutil.NewImageFromFile("assets/newgame.png")
	loadsave, _, _ = ebitenutil.NewImageFromFile("assets/loadsave.png")
	donate, _, _ = ebitenutil.NewImageFromFile("assets/donate.png")
	issues, _, _ = ebitenutil.NewImageFromFile("assets/issues.png")
	quitimg, _, _ = ebitenutil.NewImageFromFile("assets/quit.png")
	loadsave_hoover, _, _ = ebitenutil.NewImageFromFile("assets/loadsave_hoover.png")
	newgameimg_hoover, _, _ = ebitenutil.NewImageFromFile("assets/newgame_hoover.png")
	donate_hoover, _, _ = ebitenutil.NewImageFromFile("assets/donate_hoover.png")
	issues_hoover, _, _ = ebitenutil.NewImageFromFile("assets/issues_hoover.png")
	quitimg_hoover, _, err = ebitenutil.NewImageFromFile("assets/quit_hoover.png")
	if err != nil {
		log.Fatal(err)
	}
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

func clearScreen() error {
	bg.Clear()
	newgameimg.Clear()
	loadsave.Clear()
	donate.Clear()
	issues.Clear()
	quitimg.Clear()
	loadsave_hoover.Clear()
	donate_hoover.Clear()
	issues_hoover.Clear()
	quitimg_hoover.Clear()
	return nil
}

// TODO: convert the functions called here to gorooutines
func (g *Game) StartGame(c *basegame.Character, cg *core.Game) error {
	//run the game
	clearScreen()
	c.InitDefault()
<<<<<<< HEAD
	core.NewGame(c)
=======
	core.NewGame(c, cg)
	g.CG.HasStarted = true
>>>>>>> fac8704 (rm lockfile on exit, log if resdetect fails)
	return nil
}

/*
1. Render a new screen with the following text, using 	"github.com/tinne26/etxt" package in a 540x240 box at the bottom of the 960x540 screen:
"Welcome to Dope Wars. Press enter to continue."
2. Wait for the user to press enter
The text should be white and use the "assets//fonts/VT323_Regular.17.ttf" font in size 32
*/

func (g *Game) setupUI() {
	newGameBtn := func() *furex.View {
		return (&furex.View{
			Left:         340,
			Top:          210,
			Width:        200,
			Height:       40,
			MarginLeft:   360,
			MarginTop:    25,
			MarginRight:  5,
			MarginBottom: 5,
			Position:     0,
			Handler: &components.Button{Text: "", OnClick: func() {
				g.Update()
				g.StartGame(&g.Character, &g.CG)
			},
			},
			Direction:    0,
			Wrap:         0,
			Justify:      0,
			AlignItems:   0,
			AlignContent: 0,
			Grow:         0,
			Shrink:       0,
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
			Handler: &components.Button{Text: "", OnClick: func() {
				basegame.Loadsave(&basegame.Character{})
				//if savegame.json file does not exist, create it
				basegame.NewGame(&basegame.Game{})
				bg.Clear()
				newgameimg.Clear()
				loadsave.Clear()
				donate.Clear()
				issues.Clear()
				quitimg.Clear()
			},
			},
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
			Handler:      &components.Button{Text: "", OnClick: func() { openbrowser("https://www.liberapay.com/") }}, //TODO: setup donations
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
<<<<<<< HEAD
	//if core.NewGame function has started, then .RemoveAll is called on the gameUI
	if g.CG.IsRunning() {
=======
	if g.CG.HasStarted {
>>>>>>> fac8704 (rm lockfile on exit, log if resdetect fails)
		g.gameUI.RemoveAll()
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

//Clear the screen when "New Game" is pressed
