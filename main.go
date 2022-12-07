package main

import (
	"flag"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/154pinkchairs/dopewars2d/helpers"

	"github.com/154pinkchairs/dopewars2d/basegame"
	"github.com/154pinkchairs/dopewars2d/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/components"
	"golang.org/x/exp/shiny/screen"
)

type Game struct {
 init         bool
	gameUI    *furex.View
	screen    screen.Screen
	Character basegame.Character
	CG        core.Game
	//must implement ebiten.Game interface
	ebiten.Game
}

var bg *ebiten.Image

var (
	Bg                *ebiten.Image
	Loadsave          *ebiten.Image
	Newgameimg        *ebiten.Image
	Donate            *ebiten.Image
	Issues            *ebiten.Image
	Quitimg           *ebiten.Image
	Loadsave_hoover   *ebiten.Image
	Newgameimg_hoover *ebiten.Image
	Donate_hoover     *ebiten.Image
	Issues_hoover     *ebiten.Image
	Quitimg_hoover    *ebiten.Image
	InitDone          chan bool
	wg                sync.WaitGroup
)

func Init() error {
	var err error
	wg.Add(1)
	Bg, _, _ = ebitenutil.NewImageFromFile("assets/menu_bg.png")
	Newgameimg, _, _ = ebitenutil.NewImageFromFile("assets/newgame.png")
	Loadsave, _, _ = ebitenutil.NewImageFromFile("assets/loadsave.png")
	Donate, _, _ = ebitenutil.NewImageFromFile("assets/donate.png")
	Issues, _, _ = ebitenutil.NewImageFromFile("assets/issues.png")
	Quitimg, _, _ = ebitenutil.NewImageFromFile("assets/quit.png")
	Loadsave_hoover, _, _ = ebitenutil.NewImageFromFile("assets/loadsave_hoover.png")
	Newgameimg_hoover, _, _ = ebitenutil.NewImageFromFile("assets/newgame_hoover.png")
	Donate_hoover, _, _ = ebitenutil.NewImageFromFile("assets/donate_hoover.png")
	Issues_hoover, _, _ = ebitenutil.NewImageFromFile("assets/issues_hoover.png")
	Quitimg_hoover, _, err = ebitenutil.NewImageFromFile("assets/quit_hoover.png")
	if err != nil {
		log.Fatal(err)
	}
	InitDone <- true
	wg.Wait()
	wg.Done()
	return nil
}

func ClearScreen() error {
	Bg = nil
	Loadsave = nil
	Newgameimg = nil
	Donate = nil
	Issues = nil
	Quitimg = nil
	Loadsave_hoover = nil
	Newgameimg_hoover = nil
	Donate_hoover = nil
	Issues_hoover = nil
	Quitimg_hoover = nil
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go Init()
	wg.Wait()
	wg.Done()

	<-InitDone

	wg.Add(1)
	go g.setupUI()
	go g.gameUI.Draw(screen)
	wg.Wait()
	wg.Done()

	screen.DrawImage(Bg, nil)
	pos1 := &ebiten.DrawImageOptions{}
	pos1.GeoM.Translate(340, 150)
	screen.DrawImage(Newgameimg, pos1)
	if mouseOverButton(340, 150, 200, 50) {
		screen.DrawImage(Newgameimg_hoover, pos1)
	}
	// add a handler for the new game button using MouseleftButtonHandler from furex

	pos2 := &ebiten.DrawImageOptions{}
	pos2.GeoM.Translate(340, 200)
	screen.DrawImage(Loadsave, pos2)
	if mouseOverButton(340, 200, 200, 50) {
		screen.DrawImage(Loadsave_hoover, pos2)
	}

	pos3 := &ebiten.DrawImageOptions{}
	pos3.GeoM.Translate(340, 250)
	screen.DrawImage(Donate, pos3)
	if mouseOverButton(340, 250, 200, 50) {
		screen.DrawImage(Donate_hoover, pos3)
	}

	pos4 := &ebiten.DrawImageOptions{}
	pos4.GeoM.Translate(340, 300)
	screen.DrawImage(Issues, pos4)
	if mouseOverButton(340, 300, 200, 50) {
		screen.DrawImage(Issues_hoover, pos4)
	}

	pos5 := &ebiten.DrawImageOptions{}
	pos5.GeoM.Translate(340, 350)
	screen.DrawImage(Quitimg, pos5)
	if mouseOverButton(340, 350, 280, 50) {
		screen.DrawImage(Quitimg_hoover, pos5)
	}
}

// TODO: convert the functions called here to gorooutines
func (g *Game) StartGame(c *basegame.Character, cg *core.Game) error {
	//run the game
	ClearScreen()
	c.InitDefault()
	core.NewGame(c, cg)
	g.CG.HasStarted = true
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
				ClearScreen()
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
			Handler:      &components.Button{Text: "", OnClick: func() { helpers.Openbrowser("https://www.liberapay.com/") }}, //TODO: setup donations
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
			Handler:      &components.Button{Text: "", OnClick: func() { helpers.Openbrowser("https://github.com/154pinkchairs/dopewars/issues") }},
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
	if g.CG.HasStarted {
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

func resflag() (int, int) {
	//get the resolution in intxint format, parsing the string and converting it to 2 ints
	res := strings.Split(flag.Arg(0), "x")
	resx, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err)
	}
	resy, err := strconv.Atoi(res[1])
	if err != nil {
		panic(err)
	}
	return resx, resy
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 540
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Dopewars 2D")
	help := flag.Bool("help", false, "show help")
	flag.Parse()
	if *help {
		log.Println("Usage: dopewars [options] [resolution]")
		log.Println("Options:")
		flag.PrintDefaults()
		return
	}
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	if *debug {
		log.Println("Debug mode enabled")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.SetPrefix("DEBUG: ")
	} else if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

//Clear the screen when "New Game" is pressed
