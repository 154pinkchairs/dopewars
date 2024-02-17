package main

import (
	"flag"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/154pinkchairs/dopewars2d/basegame"
	"github.com/154pinkchairs/dopewars2d/core"
	"github.com/154pinkchairs/dopewars2d/debug"

	"github.com/rs/zerolog/log"
)

func resflag() (int, int) {
	// get the resolution in intxint format, parsing the string and converting it to 2 ints
	res := strings.Split(flag.Arg(0), "x")
	resx, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err)
	}
	resy, err := strconv.Atoi(res[1])
	if err != nil {
		log.Error().Err(err).Msg("non-numeric argument provided")
	}
	return resx, resy
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Dopewars 2D")
	logger := debug.SetupLogger()
	player := basegame.InitPlayer()
	gh := core.NewGameHandler(logger, player)
	scr := ebiten.NewImage(960, 540)
	gh.DrawMenu(scr)
}
