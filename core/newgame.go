package core

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"

	"github.com/154pinkchairs/dopewars2d/basegame"
	"github.com/154pinkchairs/dopewars2d/helpers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tinne26/etxt"
)

type Game struct {
	ebiten.Game
	HasStarted bool
}

func NewGame(c *basegame.Character) (error, g *Game) {
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

	// create a new text renderer and configure it
	txtRenderer := etxt.NewStdRenderer()
	glyphsCache := etxt.NewDefaultCache(10 * 1024 * 1024) // 10MB
	txtRenderer.SetCacheHandler(glyphsCache.NewHandler())
	txtRenderer.SetFont(fontLib.GetFont("VT323 Regular"))
	txtRenderer.SetAlign(etxt.YCenter, etxt.XCenter)
	txtRenderer.SetSizePx(72)

	// run the game
	ebiten.SetWindowSize(helpers.GetMaxX(), helpers.GetMaxY())
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Dope Wars 2D")
	Bgnew := ebiten.NewImage(helpers.GetMaxX(), helpers.GetMaxY())
	Bgnew.Fill(color.RGBA{0, 0, 0, 255})
	ebitenutil.DebugPrintAt(Bgnew, "Welcome to Dope Wars. Press enter to continue.", 210, 300)
	//if enter is pressed, print the keybindings menu (q to quit, i to display character info, d for district info, w for weapon info, s to sell drugs, o for bank, r to run if attacked and f to fight, b to bribe the police, u for hospital visit, t for time, r for reputation, h for help)
	ebitenutil.DebugPrintAt(Bgnew, "Press q to quit, i to display character info, d for district info, w for weapon info, s to sell drugs, o for bank, r to run if attacked and f to fight, b to bribe the police, u for hospital visit, t for time, r for reputation, h for help", 210, 330)
	ebitenutil.DebugPrintAt(Bgnew, "Press enter to continue", 210, 360)
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		ebitenutil.DebugPrintAt(Bgnew, "You pressed enter", 210, 390)
		Keys(c)
	}
	if err == nil {
		g.HasStarted = true
	}
	return nil, g
}

func Keys(c *basegame.Character) error {
	Bgnew := ebiten.NewImage(helpers.GetMaxX(), helpers.GetMaxY())
	Bgnew.Fill(color.RGBA{0, 0, 0, 255})
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		//display character info
		ebitenutil.DebugPrintAt(Bgnew, "Name: "+c.Name, 210, 420)
		ebitenutil.DebugPrintAt(Bgnew, "Cash: "+strconv.Itoa(c.Cash), 210, 450)
		ebitenutil.DebugPrintAt(Bgnew, "Debt: "+strconv.Itoa(c.Debt), 210, 480)
		ebitenutil.DebugPrintAt(Bgnew, "Reputation: "+strconv.Itoa(c.Reputation), 210, 510)
		ebitenutil.DebugPrintAt(Bgnew, "Days: "+strconv.Itoa(c.Days), 210, 540)
		ebitenutil.DebugPrintAt(Bgnew, "Wanted Level: "+strconv.Itoa(c.WantedLevel), 210, 570)
		x := c.CurrentDistrict.Name
		ebitenutil.DebugPrintAt(Bgnew, "Current District: "+string(x), 210, 600)
		//enumerate the weapons names and quantities in the character's inventory from basegame.WeaponUnits map
		for k, v := range c.Weapons {
			ebitenutil.DebugPrintAt(Bgnew, k.Name+": "+strconv.Itoa(v), 210, 630)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		ebitenutil.DebugPrintAt(Bgnew, "District Name: "+c.CurrentDistrict.Name, 210, 420)

		ebitenutil.DebugPrintAt(Bgnew, "Hospital: "+strconv.FormatBool(c.CurrentDistrict.Properties.Hospital), 210, 480)
		ebitenutil.DebugPrintAt(Bgnew, "Bank: "+strconv.FormatBool(c.CurrentDistrict.Properties.Bank), 210, 510)
		ebitenutil.DebugPrintAt(Bgnew, "LoanShark: "+strconv.FormatBool(c.CurrentDistrict.Properties.LoanShark), 210, 540)
		for _, v := range c.CurrentDistrict.Properties.NeighbourIDs {
			ebitenutil.DebugPrintAt(Bgnew, "Neighbour Districts: "+fmt.Sprint(v), 210, 570)
		}
		for _, v := range c.CurrentDistrict.DrugsAvailable {
			ebitenutil.DebugPrintAt(Bgnew, v.Name+": "+strconv.Itoa(v.Price), 210, 630)
		}
	}
	//TODO: create a menu to choose the weapon about which to display info
	return nil

}
