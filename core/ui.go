package core

import (
	"image"
	"image/color"
	"os"
	"path"

	"github.com/154pinkchairs/dopewars2d/basegame"
	"github.com/154pinkchairs/dopewars2d/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	eu "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/components"
)

type UI interface {
	NewGameBtn() (*furex.View, error)
	LoadSaveBtn() (*furex.View, error)
	DonateBtn() (*furex.View, error)
	IssuesBtn() (*furex.View, error)
	QuitBtn() (*furex.View, error)
}

type GameUI struct {
	NewGBtn     *furex.View
	LoadSaveBtn *furex.View
	DonateBtn   *furex.View
	IssuesBtn   *furex.View
	QuitBtn     *furex.View
	NewGameImg  *furex.View
}

type Button struct {
	OnClick   func()
	Color     color.RGBA
	mouseover bool
	pressed   bool
	Sprite    string
	DrawHan   FXDHandler
}

type FXDHandler interface {
	furex.DrawHandler
}

type FXHandlerImpl struct{}

func (f FXHandlerImpl) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	var b Button
	eu.DrawRect(screen, float64(frame.Min.X), float64(frame.Min.Y), float64(frame.Dx()), float64(frame.Dy()), color.RGBA{0, 0, 0, 255})
	b.DrawHan.HandleDraw(screen, frame)
}

func (gu *GameUI) DrawMenu(ebiten.Image) {
	screen := ebiten.NewImage(960, 540)
	eu.DrawRect(screen, 0, 0, 960, 540, color.RGBA{0, 0, 0, 255})
}

func (gu *GameUI) NewGameBtn() (*furex.View, error) {
	f := FXHandlerImpl{}
	// TODO: sprite creation here (1st return value)
	_, _, err := eu.NewImageFromFile(path.Join("assets", "newgame.png"))
	if err != nil {
		return nil, err
	}

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
		Handler:      f,
		Direction:    0,
		Wrap:         0,
		Justify:      0,
		AlignItems:   0,
		AlignContent: 0,
		Grow:         0,
		Shrink:       0,
	}), nil
}

func (g *GameHandler) setupUI() error {
	newGameBtn := func() *furex.View {
		return (&furex.View{
			Top:          200,
			Left:         340,
			Width:        120,
			Height:       40,
			MarginLeft:   390,
			MarginTop:    5,
			MarginBottom: 5,
			MarginRight:  5,
			Handler: &components.Button{Text: "", OnClick: func() {
				bg.Clear()
				g.g.gameUI.RemoveAll()
				g.NewGame(g.c)
			}},
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
			Handler: &components.Button{
				Text: "", OnClick: func() {
					playerdata := basegame.Loadsave()
					g.NewGame(playerdata)
					bg.Clear()
					for _, img := range buttonImages {
						img.Clear()
					}
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
			Handler:      &components.Button{Text: "", OnClick: func() { helpers.Openbrowser("https://www.liberapay.com/") }}, // TODO: setup donations
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

	g.g.gameUI = (&furex.View{
		Width:        960,
		Height:       540,
		Direction:    furex.Column,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemStart, // place items in the center, one below the other
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
	// if core.NewGame function has started, then .RemoveAll is called on the gameUI
	return nil
}
