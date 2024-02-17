package core

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

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
