package core

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

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
	wg.Done()
	return nil
}

func ClearScreen() error {
	Bg.Clear()
	Newgameimg.Clear()
	Loadsave.Clear()
	Donate.Clear()
	Issues.Clear()
	Quitimg.Clear()
	Loadsave_hoover.Clear()
	Donate_hoover.Clear()
	Issues_hoover.Clear()
	Quitimg_hoover.Clear()
	return nil
}
