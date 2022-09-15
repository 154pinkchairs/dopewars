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
)

type Game struct{ txtRenderer *etxt.Renderer }

func (self *Game) Layout(int, int) (int, int) { return 960, 540 }
func (self *Game) Update() error              { return nil }
func (self *Game) Draw(screen *ebiten.Image) {
	millis := time.Now().UnixMilli() // (you should usually avoid using time)
	blue := (millis / 16) % 512
	if blue >= 256 {
		blue = 511 - blue
	}
	changingColor := color.RGBA{0, 255, uint8(blue), 255}

	// set relevant text renderer properties and draw
	self.txtRenderer.SetTarget(screen)
	self.txtRenderer.SetColor(changingColor)
	self.txtRenderer.Draw("Welcome to Dope Wars 2D!", 480, 200)
}

type Character struct {
	Name                    string
	Health                  int
	Reputation, WantedLevel int
	cash                    int
	bank                    int
	debt                    int
	CurrentDistrict         District
	drugs                   Drugs
	weapons                 WeaponUnits
	weaponsAvailable        []Weapon
}

//leave this here for debugging only
/*func Player(c *Character) {
	c.cash = 10000
	c.debt = 15000
	fmt.Println("Welcome to Dope Wars!")
	fmt.Println("What is your name?")
	fmt.Scanln(&c.Name)
	fmt.Println("Welcome to the world of Dope Wars, " + c.Name + "!")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("You are a small time drug dealer in the city of New York.\n After failing one job after another, you have decided to start a small business.")
	//place the player in a random district from the districts slice
	c.CurrentDistrict = districts[rand.Intn(len(districts))]
	fmt.Println("You have decided to start your business in " + c.CurrentDistrict.name + ".")
	//randomize the district's availability of drugs
	fmt.Println("You have a small amount of cash, but you need to make a lot of money.\nAfter one of your drug deals went down, you were left with a debt.")
	fmt.Println("You have $" + strconv.Itoa(c.debt) + " to pay off.")
	fmt.Println("Press h for help. Press q to quit.")
	var key string
	fmt.Scanln("%s", &key)
	if key == "h" {
		fmt.Println("h - help")
		fmt.Println("q - quit")
		fmt.Println("i - Player")
		fmt.Println("d - district info and the available drugs. Press d again to see the drugs in stock.\n to buy a drug, type the drug number and press enter.")
		fmt.Println("t - travel to a district. Press t again to see the districts you can travel to.\n to travel to a district, type the district number and press enter.")
		fmt.Println("w - weapon info and the available weapons. Press w again to see the weapons in stock.\n to buy a weapon, type the weapon number and press enter.")
		fmt.Println("a - current weapon info. Press a again to see the current weapon stats. Press s to sell the current weapon.")
		fmt.Println("f - fight the opponent. For throwable weapons, press j to throw the weapon. Note you will lose the weapon if you do not deal a critical hit\n or if it's a handgrenade.")
		fmt.Println("s - sell the drugs. Type the drug number and press enter.")
		fmt.Println("o - make a payment or withdraw/borrow money from the bank or loan shark. Type the amount and press enter.")
		fmt.Println("r - run away. You might lose some cash or drugs and the wanted level will go down.")
		fmt.Println("b - bribe the law enforcement. You will lose some cash and the wanted level will go down.")
		fmt.Println("g - visit the bank or loan shark.")
		fmt.Println("u - visit the hospital.")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	} else if key == "q" {
		os.Exit(0)
	}
}
*/

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
	Bank                    int
	Debt                    int
	CurrentDistrict         District
	Drugs                   Drugs
	Weapons                 WeaponUnits
	WeaponsAvailable        []Weapon
}

//parse the ../savegame.json file and pass the data to the character struct. If any fields are missing, use the default values
func Loadsave(c *Character) {
	//open the file
	file, err := os.Open("../savegame.json")
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
	c.cash = save.Cash
	c.bank = save.Bank
	c.debt = save.Debt
	c.CurrentDistrict = save.CurrentDistrict
	c.drugs = save.Drugs
	c.weapons = save.Weapons
	c.weaponsAvailable = save.WeaponsAvailable
}

func NewGame() {
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

	ebiten.SetWindowSize(400, 400)
	err = ebiten.RunGame(&Game{txtRenderer})
	if err != nil {
		log.Fatal(err)
	}
}
