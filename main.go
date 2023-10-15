package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	width       = 600
	height      = 600
	fieldWidth  = 60
	fieldHeight = 60
	ratioW      = width / fieldWidth
	ratioH      = height / fieldHeight
)

type GameState int

const (
	edit GameState = iota
	run
)

var (
	stateToDrawer  = make([]func(*ebiten.Image, *Game), 2)
	stateToUpdater = make([]func(*Game) error, 2)
	colors         = [2]Color{black, white}
)

type Game struct {
	state       GameState
	pixels      []uint8
	field       [][]uint8
	fieldBuffer [][]uint8

	fullscreen bool
}

func (g *Game) Update() error {
	return stateToUpdater[g.state](g)
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			xf, yf := screenCoordToField(x, y)
			p := (y*width + x) * 4
			setColor(g.pixels, p, colors[g.field[yf][xf]])
		}
	}
	drawGrid(g)
	screen.WritePixels(g.pixels)
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("FPS: %.2f\nTPS: %.2f", ebiten.ActualFPS(), ebiten.ActualTPS()),
		0,
		height-30,
	)
	stateToDrawer[g.state](screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func initGame() *Game {
	game := &Game{
		pixels:      make([]uint8, width*height*4),
		field:       make([][]uint8, fieldHeight),
		fieldBuffer: make([][]uint8, fieldHeight),
	}

	for i := 0; i < fieldHeight; i++ {
		game.field[i] = make([]uint8, fieldWidth)
		game.fieldBuffer[i] = make([]uint8, fieldWidth)
	}

	stateToDrawer[edit] = editDrawer
	stateToDrawer[run] = runDrawer

	stateToUpdater[edit] = editUpdater
	stateToUpdater[run] = runUpdater

	return game
}

func main() {
	game := initGame()

	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Game of life!")
	ebiten.SetFullscreen(game.fullscreen)
	ebiten.SetVsyncEnabled(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
