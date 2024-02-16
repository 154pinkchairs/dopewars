package main

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/154pinkchairs/dopewars2d/basegame"
	"github.com/154pinkchairs/dopewars2d/core"
	"github.com/yohamta/furex/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/rs/zerolog/log"
)

type Game struct {
	init      bool
	gameUI    *furex.View
	Character basegame.Character
	CG        core.Game
	// must implement ebiten.Game interface
	ebiten.Game
	UI GameUI
}

type button struct {
	img, imgHover *ebiten.Image
	x, y, w, h    *int
}

type buttonPath struct {
	base, hover string
}

// nolint:gochecknoglobals
var (
	bg *ebiten.Image

	buttonPaths = []buttonPath{
		{"assets/newgame.png", "assets/newgame_hoover.png"},
		{"assets/loadsave.png", "assets/loadsave_hoover.png"},
		{"assets/donate.png", "assets/donate_hoover.png"},
		{"assets/issues.png", "assets/issues_hoover.png"},
		{"assets/quit.png", "assets/quit_hoover.png"},
	}

	buttonImages []*ebiten.Image
)

func initButtonImages() (err error) {
	// create a slice of images to hold the button images
	buttonImages = make([]*ebiten.Image, len(buttonPaths)*2)
	const loadErr = "failed to load button image:"

	// interpolate the button paths into image variable names
	for i, path := range buttonPaths {
		baseImg, _, err := ebitenutil.NewImageFromFile(path.base)
		if err != nil {
			return fmt.Errorf("%s %w", loadErr, err)
		}
		hoverImg, _, err := ebitenutil.NewImageFromFile(path.hover)
		if err != nil {
			return fmt.Errorf("%s %w", loadErr, err)
		}

		buttonImages[i*2] = baseImg
		buttonImages[i*2+1] = hoverImg
	}

	return nil
}

func initBackground() (err error) {
	bg, _, err = ebitenutil.NewImageFromFile("assets/menu_bg.png")
	if err != nil {
		return fmt.Errorf("failed to load background image: %w", err)
	}
	return nil
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
	}
	if g.gameUI == nil {
		return g.setupUI()
	}
	g.gameUI.UpdateWithSize(ebiten.WindowSize())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if bg == nil {
		if err := initBackground(); err != nil {
			log.Fatal().Err(err).Msg("Failed to load background image")
		}
	}

	screen.DrawImage(bg, nil)
	g.UI.DrawMenu(*ebiten.NewImageFromImage(bg))

	buttonNames := []string{"newgame", "loadsave", "donate", "issues", "quit"}
	buttonPositions := []struct{ x, y, w, h int }{
		{340, 150, 200, 50},
		{340, 200, 200, 50},
		{340, 250, 200, 50},
		{340, 300, 200, 50},
		{340, 350, 280, 50},
	}

	if err := initButtonImages(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize button images")
	}

	buttons := make([]button, len(buttonNames))
	for i := range buttonNames {
		buttons[i] = button{
			img:      buttonImages[i*2],
			imgHover: buttonImages[i*2+1],
			x:        &buttonPositions[i].x,
			y:        &buttonPositions[i].y,
			w:        &buttonPositions[i].w,
			h:        &buttonPositions[i].h,
		}
	}
	for _, btn := range buttons {
		pos := &ebiten.DrawImageOptions{}
		pos.GeoM.Translate(float64(*btn.x), float64(*btn.y))
		screen.DrawImage(btn.img, pos)
		if mouseOverButton(*btn.x, *btn.y, *btn.w, *btn.h) {
			screen.DrawImage(btn.imgHover, pos)
		}
	}
}

func (g *Game) StartGame(c *basegame.Character, cg *core.Game) error {
	c.InitDefault()
	err := core.NewGame(c, cg)
	if err != nil {
		return err
	}
	g.CG.HasStarted = true
	return nil
}

// create a function that checks if the mouse is over a button
func mouseOverButton(x, y, width, height int) bool {
	// get the mouse position
	mouseX, mouseY := ebiten.CursorPosition()
	// check if the mouse is within the button's x and y bounds
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
		log.Fatal().Err(err).Msg("Failed to run game")
	}
}
