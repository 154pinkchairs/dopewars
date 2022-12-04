package main

import (
	"testing"

	"github.com/154pinkchairs/dopewars2d/core"
)

func TestGameState(t *testing.T) {
	g := core.Game{}
	mg := Game{}
	mg.StartGame(&mg.Character, &mg.CG)
	if !g.HasStarted {
		t.Errorf("StartGame() failed")
	}
}
