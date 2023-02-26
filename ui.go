package main

import (
	"image/color"

	"github.com/154pinkchairs/dopewars2d/core"
	"github.com/yohamta/furex/v2"
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
}

func (gu *GameUI) NewGameBtn() (*furex.View, error) {
	g := &Game{}
	cg := core.Game{}
	if err := g.StartGame(&g.Character, &cg); err != nil {
		g.Close()
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
		Handler: &Button{Color: color.RGBA{255, 255, 255, 0}, OnClick: func() { //FIXME: correct types
			g.Update()
			cg := core.Game{}
			g.StartGame(&g.Character, &cg)
			//g.Clear() //TODO: clear the screen
		},
			Sprite: "assets/newgame.png",
		},
		Direction:    0,
		Wrap:         0,
		Justify:      0,
		AlignItems:   0,
		AlignContent: 0,
		Grow:         0,
		Shrink:       0,
	}), nil
}

/*
func (g *Game) setupUI() error {
	newGameBtn := func() *furex.View {
		return (&furex.View{

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
	//if core.NewGame function has started, then .RemoveAll is called on the gameUI
	//NOTE: got segfault here
	if g.CG.HasStarted {
		g.gameUI.RemoveAll()
	}
}
*/
