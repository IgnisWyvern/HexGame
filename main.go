package main

import (
	"log"

	"github.com/IgnisWyvern/HexGame/HexGame"
	"github.com/hajimehoshi/ebiten/v2"
)

const ScreenMulti = 1.5

func main() {
	game, err := hexgame.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(hexgame.GameWidth*ScreenMulti, hexgame.GameHeight*ScreenMulti)
	ebiten.SetWindowTitle("HexGame")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
