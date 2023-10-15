package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var editInfo = "EDIT MODE\n" +
	"click on cell to edit\n" +
	"press ENTER to run\n" +
	"press DEL to clear\n" +
	"press F for fullscreen\n" +
	"press R to fill randomly"

var runInfo = "RUN MODE\n" +
	"press ESC for edit mode"

func editDrawer(screen *ebiten.Image, game *Game) {
	ebitenutil.DebugPrint(screen, editInfo)
}

func runDrawer(screen *ebiten.Image, game *Game) {
	ebitenutil.DebugPrint(screen, runInfo)
}
