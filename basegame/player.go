package basegame

import (
	"encoding/json"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"golang.org/x/exp/shiny/screen"
)

type Game struct {
	txtRenderer *etxt.Renderer
	screen      *screen.Screen
}

func (g *Game) Layout(int, int) (int, int) { return 960, 540 }
func (g *Game) Update() error              { return nil }
func (g *Game) Draw(screen *ebiten.Image) {
	millis := time.Now().UnixMilli() // (you should usually avoid using time)
	blue := (millis / 16) % 512
	Bgnew := ebiten.NewImage(960, 540)
	Bgnew.Fill(color.RGBA{0, 0, 0, 255})
	//draw the bgnew image to the screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(Bgnew, op)
	if blue >= 256 {
		blue = 511 - blue
	}
	changingColor := color.RGBA{0, 255, uint8(blue), 255}

	// set relevant text renderer properties and draw
	g.txtRenderer.SetTarget(screen)
	g.txtRenderer.SetColor(changingColor)
	g.txtRenderer.Draw("Welcome to Dope Wars 2D!", 480, 200)
}

type Character struct {
	Name                    string
	Health                  int
	Reputation, WantedLevel int
	Cash                    int
	Days                    int
	Bank                    int
	Debt                    int
	CurrentDistrict         District
	drugs                   Drugs
	Weapons                 WeaponUnits
	weaponsAvailable        []Weapon
}

func (c *Character) InitDefault() {
	c.Name = "Heisenberg"
	c.Cash = 10000
	c.Debt = 15000
	c.Reputation = 0
	c.Days = 0
	c.WantedLevel = 0
	c.CurrentDistrict = Bronx
}

//using the following doc: https://docs.rocketnine.space/code.rocketnine.space/tslocum/messeji/
//create a new text field at the bottom of the screen
//window will be 1152x648
//text field will be 1152x200
//set the font.Face to assets/VT323_Regular.17.ttf
//set the font size to 32
//var textField = messeji.NewTextField(image.Rect(0, 448, 1152, 648), font.Face("assets/VT323_Regular.ttf"))

type Save struct {
	Name                    string
	Health                  int
	Reputation, WantedLevel int
	Cash                    int
	Days                    int
	Bank                    int
	Debt                    int
	CurrentDistrict         District
	Drugs                   Drugs
	Weapons                 WeaponUnits
	WeaponsAvailable        []Weapon
}

// parse the ../savegame.json file and pass the data to the character struct. If any fields are missing, use the default values
func Loadsave(c *Character) {
	//open the file
	file, err := os.Open("savegame.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//read the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	//parse the json
	var save Save
	err = json.Unmarshal(data, &save)
	if err != nil {
		log.Fatal(err)
	}

	//set the character struct values
	c.Name = save.Name
	c.Health = save.Health
	c.Reputation = save.Reputation
	c.WantedLevel = save.WantedLevel
	c.Cash = save.Cash
	c.Bank = save.Bank
	c.Debt = save.Debt
	c.Days = save.Days
	c.CurrentDistrict = save.CurrentDistrict
	c.drugs = save.Drugs
	c.Weapons = save.Weapons
	c.weaponsAvailable = save.WeaponsAvailable

	//draw a new black background window

}

func NewGame(g *Game) {
	// load font library
	fontLib := etxt.NewFontLibrary()
	_, _, err := fontLib.ParseDirFonts("assets/fonts")
	if err != nil {
		log.Fatalf("Error while loading fonts: %s", err.Error())
	}
	txtRenderer := etxt.NewStdRenderer()
	glyphsCache := etxt.NewDefaultCache(10 * 1024 * 1024) // 10MB
	txtRenderer.SetCacheHandler(glyphsCache.NewHandler())
	txtRenderer.SetFont(fontLib.GetFont("VT323 Regular"))
	txtRenderer.SetAlign(etxt.Bottom, etxt.XCenter)
	txtRenderer.SetSizePx(36)
	txtRenderer.SetColor(color.RGBA{255, 255, 255, 255})
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Dope Wars")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	err = ebiten.RunGame(&Game{
		txtRenderer: txtRenderer,
	})
	GameBG := ebiten.NewImage(960, 540)
	GameBG.Fill(color.RGBA{0, 0, 0, 255})
	//invoke the Draw function
	/* curruntly we get th following error: panic: runtime error: invalid memory address or nil pointer dereference
	[signal SIGSEGV: segmentation violation code=0x1 addr=0x28 pc=0x66b197] */
	g.Layout(960, 540)
	g.Draw(GameBG)
	g.Update()
	ebiten.SetWindowTitle("Dope Wars")
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err != nil {
		log.Fatal(err)
	}
}
