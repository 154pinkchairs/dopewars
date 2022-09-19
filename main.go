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
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tinne26/etxt"
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
	Character basegame.Character
	txtRenderer *etxt.Renderer
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


/*create a new game function to:
1. initialize the Character struct from basegame/player.go with default/pseudorandom values
2. create a new savegame.json file
3. Draw a new black screen and dismiss the menu assets and UI
4. Render a new screen with the following text, using 	"github.com/tinne26/etxt" package in a 540x240 box at the bottom of the 960x540 screen:
"Welcome to Dope Wars. Press enter to continue."
5. Wait for the user to press enter
The text should be white and use the "assets//fonts/VT323_Regular.17.ttf" font in size 32
*/
func (g *Game) Redraw(screen *ebiten.Image) {
	g.Character = basegame.Character{}
	if screen != nil {
		screen.Fill(color.Black)
		g.txtRenderer.SetTarget(screen)
	} else {
		screen = ebiten.NewImage(960, 540)
		screen.Fill(color.Black)
		g.txtRenderer.SetTarget(screen)
	}
	g.txtRenderer.SetColor(color.White)
	g.txtRenderer.Draw("Welcome to Dope Wars. Press enter to continue.", 200, 200)
}

func (g *Game) NewGame(c *basegame.Character) {
	c.Name = "John Doe"
	c.Cash = 10000
	c.Debt = 15000
	c.Reputation = 0
	c.Days = 0
	c.WantedLevel = 0
	c.CurrentDistrict = basegame.Bronx
	//check if this map is not nil
	if c.Weapons != nil {
		c.Weapons = make(map[basegame.Weapon]int)
	c.Weapons[basegame.Knuckle] = 1
	}
		// load font library
	fontLib := etxt.NewFontLibrary()
	_, _, err := fontLib.ParseDirFonts("assets/fonts")
	if err != nil {
		log.Fatalf("Error while loading fonts: %s", err.Error())
	}

	// check that we have the fonts we want
	// (shown for completeness, you don't need this in most cases)

	// create a new text renderer and configure it
	txtRenderer := etxt.NewStdRenderer()
	glyphsCache := etxt.NewDefaultCache(10*1024*1024) // 10MB
	txtRenderer.SetCacheHandler(glyphsCache.NewHandler())
	txtRenderer.SetFont(fontLib.GetFont("VT323 Regular"))
	txtRenderer.SetAlign(etxt.YCenter, etxt.XCenter)
	txtRenderer.SetSizePx(72)

	// run the "game"
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Dope Wars 2D")
	bg.Clear()
	newgameimg.Clear()
	loadsave.Clear()
	donate.Clear()
	issues.Clear()
	quitimg.Clear()
	bgnew := ebiten.NewImage(960, 540)
	bgnew.Fill(color.Black)
	//make bgnew.NewBuffer poit to an image.Point
	x, x1 := bg.NewBuffer(bgnew.Bounds().Size())
	g.screen = x, x1
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

	//save the values to a new savegame.json file
	/*savegame, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("savegame.json", savegame, 0644)*/
	//create a new black screen and dismiss the menu assets and UI using ebiten.NewImage, bg.Clear, newgameimg.Clear, loadsave.Clear, donate.Clear, issues.Clear, quitimg.Clear and ebiten.Fill
	//set the text color to white
	//render a new screen with the following text, using 	"github.com/tinne26/etxt" package in a 540x240 box at the bottom of the 960x540 screen:
	//The text should be white and use the "assets//fonts/VT323_Regular.17.ttf" font in size 32
	ebitenutil.DebugPrintAt(bgnew, "Welcome to Dope Wars. Press enter to continue.", 210, 300)
	//if enter is pressed, print the keybindings menu (q to quit, i to display character info, d for district info, w for weapon info, s to sell drugs, o for bank, r to run if attacked and f to fight, b to bribe the police, u for hospital visit, t for time, r for reputation, h for help)
	ebitenutil.DebugPrintAt(bgnew, "Press q to quit, i to display character info, d for district info, w for weapon info, s to sell drugs, o for bank, r to run if attacked and f to fight, b to bribe the police, u for hospital visit, t for time, r for reputation, h for help", 210, 330)
	ebitenutil.DebugPrintAt(bgnew, "Press enter to continue", 210, 360)
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		ebitenutil.DebugPrintAt(bgnew, "You pressed enter", 210, 390)
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		//display character info
		ebitenutil.DebugPrintAt(bgnew, "Name: "+c.Name, 210, 420)
		ebitenutil.DebugPrintAt(bgnew, "Cash: "+strconv.Itoa(c.Cash), 210, 450)
		ebitenutil.DebugPrintAt(bgnew, "Debt: "+strconv.Itoa(c.Debt), 210, 480)
		ebitenutil.DebugPrintAt(bgnew, "Reputation: "+strconv.Itoa(c.Reputation), 210, 510)
		ebitenutil.DebugPrintAt(bgnew, "Days: "+strconv.Itoa(c.Days), 210, 540)
		ebitenutil.DebugPrintAt(bgnew, "Wanted Level: "+strconv.Itoa(c.WantedLevel), 210, 570)
		x := c.CurrentDistrict.Name
		ebitenutil.DebugPrintAt(bgnew, "Current District: "+string(x), 210, 600)
		//enumerate the weapons names and quantities in the character's inventory from basegame.WeaponUnits map
		for k, v := range c.Weapons {
			ebitenutil.DebugPrintAt(bgnew, k.Name+": "+strconv.Itoa(v), 210, 630)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		ebitenutil.DebugPrintAt(bgnew, "District Name: "+c.CurrentDistrict.Name, 210, 420)
		
			ebitenutil.DebugPrintAt(bgnew, "Hospital: "+strconv.FormatBool(c.CurrentDistrict.Properties.Hospital), 210, 480)
			ebitenutil.DebugPrintAt(bgnew, "Bank: "+strconv.FormatBool(c.CurrentDistrict.Properties.Bank), 210, 510)
			ebitenutil.DebugPrintAt(bgnew, "LoanShark: "+strconv.FormatBool(c.CurrentDistrict.Properties.LoanShark), 210, 540)
			for _, v := range c.CurrentDistrict.Properties.NeighbourIDs {
				ebitenutil.DebugPrintAt(bgnew, "Neighbour Districts: "+string(v), 210, 570)
			}
		for _, v := range c.CurrentDistrict.DrugsAvailable {
			ebitenutil.DebugPrintAt(bgnew, v.Name+": "+strconv.Itoa(v.Price), 210, 630)
		}
	}
	//create a menu to choose the weapon about which to display info		
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
			Handler:      &components.Button{Text: "", OnClick: func() { 
				g.NewGame(&g.Character)}},
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
			Handler:      &components.Button{Text: "", OnClick: func() { 
				basegame.Loadsave(&basegame.Character{})
				basegame.NewGame(&basegame.Game{})
				bg.Clear()
				newgameimg.Clear()
				loadsave.Clear()
				donate.Clear()
				issues.Clear()
				quitimg.Clear()},
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
