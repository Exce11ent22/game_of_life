package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func editUpdater(g *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.state = run
		ebiten.SetTPS(10)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		xf, yf := screenCoordToField(x, y)
		if g.field[yf][xf] == 1 {
			g.field[yf][xf] = 0
		} else {
			g.field[yf][xf] = 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDelete) {
		for x := 0; x < fieldWidth; x++ {
			for y := 0; y < fieldHeight; y++ {
				g.field[y][x] = 0
			}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.fullscreen = !g.fullscreen
		ebiten.SetFullscreen(g.fullscreen)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		for x := 0; x < fieldWidth; x++ {
			for y := 0; y < fieldHeight; y++ {
				g.field[x][y] = uint8(rand.Intn(2))
			}
		}
	}
	return nil
}

func runUpdater(g *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.state = edit
		ebiten.SetTPS(60)
	}
	for x := 0; x < fieldWidth; x++ {
		for y := 0; y < fieldHeight; y++ {
			n := getNeighbors(x, y, g.field)
			if n < 2 || n > 3 {
				// die
				g.fieldBuffer[y][x] = 0
			} else if n == 3 {
				// alive
				g.fieldBuffer[y][x] = 1
			} else {
				g.fieldBuffer[y][x] = g.field[y][x]
			}
		}
	}
	for i := 0; i < fieldHeight; i++ {
		copy(g.field[i], g.fieldBuffer[i])
	}
	return nil
}
