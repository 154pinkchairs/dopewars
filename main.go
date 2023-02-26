package main

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/154pinkchairs/dopewars2d/basegame"
	"github.com/154pinkchairs/dopewars2d/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/furex/v2"

	"github.com/rs/zerolog/log"
)

type Game struct {
	init      bool
	gameUI    *furex.View
	Character basegame.Character
	CG        core.Game
	//must implement ebiten.Game interface
	ebiten.Game
	UI GameUI
}

var (
	bg *ebiten.Image

	loadsave          *ebiten.Image
	newgameimg        *ebiten.Image
	donate            *ebiten.Image
	issues            *ebiten.Image
	quitimg           *ebiten.Image
	loadsave_hoover   *ebiten.Image
	newgameimg_hoover *ebiten.Image
	donate_hoover     *ebiten.Image
	issues_hoover     *ebiten.Image

	quitimg_hoover *ebiten.Image
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	var err error
	bg, _, err = ebitenutil.NewImageFromFile("assets/menu_bg.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading background image")
	}
	newgameimg, _, err = ebitenutil.NewImageFromFile("assets/newgame.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading new game image")
	}
	loadsave, _, err = ebitenutil.NewImageFromFile("assets/loadsave.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading load/save image")
	}
	donate, _, err = ebitenutil.NewImageFromFile("assets/donate.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading donate image")
	}
	issues, _, err = ebitenutil.NewImageFromFile("assets/issues.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading issues image")
	}
	quitimg, _, err = ebitenutil.NewImageFromFile("assets/quit.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading quit image")
	}
	loadsave_hoover, _, err = ebitenutil.NewImageFromFile("assets/loadsave_hoover.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading load/save hoover image")
	}
	newgameimg_hoover, _, _ = ebitenutil.NewImageFromFile("assets/newgame_hoover.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading new game hoover image")
	}
	donate_hoover, _, err = ebitenutil.NewImageFromFile("assets/donate_hoover.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading donate hoover image")
	}
	issues_hoover, _, err = ebitenutil.NewImageFromFile("assets/issues_hoover.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading issues hoover image")
	}
	quitimg_hoover, _, err = ebitenutil.NewImageFromFile("assets/quit_hoover.png")
	if err != nil {
		log.Panic().AnErr("error", err).Msg("Error loading quit hoover image")
	}
}

func (g *Game) Close(mode string) error {
	switch mode {
	case "fatal":
		log.Fatal().Msg("Fatal error. Quitting.")
		return nil
	case "error":
		log.Error().Msg("Error. Quitting.")
		return errors.New("Error. Quitting.")
	case "graceful":
		log.Info().Msg("Gracefully quitting.")
		return nil
	default:
		log.Info().Msg("Gracefully quitting.")
		return nil
	}
}

func (g *Game) Update() error {
	if !g.init {
		g.init = true
		g.UI.()
	}
	g.gameUI.UpdateWithSize(ebiten.WindowSize())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)
	g.gameUI.DrawMenu(screen)
	pos1 := &ebiten.DrawImageOptions{}
	pos1.GeoM.Translate(340, 150) // f64, f64
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
	err := clearScreen()
	if err != nil {
		return err
	}
	c.InitDefault()
	err = core.NewGame(c, cg)
	if err != nil {
		return err
	}
	g.CG.HasStarted = true
	return nil
}

/*
1. Render a new screen with the following text, using 	"github.com/tinne26/etxt" package in a 540x240 box at the bottom of the 960x540 screen:
"Welcome to Dope Wars. Press enter to continue."
2. Wait for the user to press enter
The text should be white and use the "assets//fonts/VT323_Regular.17.ttf" font in size 32
*/

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
		log.Error().Err(err).Msg("Failed to open browser")
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
		log.Default().Fatal().Err(err).Msg("Failed to run game")
	}
}

//Clear the screen when "New Game" is pressed
