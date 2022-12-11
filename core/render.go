package core

import (
    "bufio"
		"os"
		"log"
		"sync"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/154pinkchairs/dopewars2d/basegame"
)

type MapDrawer interface {
    Draw(screen *ebiten.Image) error
}

type Tile rune

const (
	building_a Tile = 'a'
	building_b Tile = 'b'
	alley_horiz Tile = '_'
	alley_vert Tile = '!'
	horiz_st Tile = '-' //horizontal street
	vert_st Tile = '|' //vertical street
	police_st Tile = 'p'
	hosp Tile = 'h'
	bank Tile = '$'
	traphouse Tile = 't'
	gunshop Tile = 'g'
	loan_shark Tile = 'l'
	grass Tile = ' '
	water Tile = '~'
	dirt Tile = '#'
	metro Tile = 'm'
)

type Dist_map struct {
	Map [][]Tile
	District *basegame.District
}

//create the file reader, then the render function/method/interface that translates the characters into tile files to be rendered by ebitenutil.NewImageFromFile
func renderMap(screen *ebiten.Image, mapDrawer MapDrawer, srcfile string) (error, *ebiten.Image, *Dist_map) {
	//open the file
	file, err := os.Open(srcfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mapLines := make([]string, 0) //create a slice of strings to hold the map lines
	maxRow := 0
	idx := 0
	for scanner.Scan() {
		mapLines = append(mapLines, scanner.Text())
		if len(mapLines[idx]) > maxRow {
			maxRow = len(mapLines[idx])
		}
		idx++
	}

