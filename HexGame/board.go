package hexgame

import (
	"bytes"
	"image"
	_ "image/png"	
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	boardHRatio = 0.85
	boardTopOffeset = 0.1 
	boardHeight = 7
	boardWidth = 13
	hexImgHeight = 30
	hexLongWidth = 90 
	hexShortWidth = 40
)

type tile struct {
	x, y int
	background *ebiten.Image
	foreground *ebiten.Image
}

type coord struct {
	x, y int
}

type board struct {
	tiles map[coord]*tile
	boardImage *ebiten.Image
}

func populateBoard(width int, height int) map[coord]*tile{
	tempBoard := make(map[coord]*tile)
	for x := 0; x < width; x++ {
		for y := 0; y < height*2; y += 2{
			offset := x % 2
			tempBoard[coord{x,y+offset}] = newTile(x, y+offset)
		}
	}
	return tempBoard
}

func (t *tile) draw(image *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.x*(hexLongWidth+hexShortWidth)), float64(t.y*hexImgHeight))
	image.DrawImage(t.background, op)
	
}

func newTile(x, y int) *tile {
	bgimg, _ , err := image.Decode(bytes.NewReader(blueHex_png))
	if err != nil{
		log.Fatal(err)
	}
	return &tile{
		x: x,
		y: y,
		background: ebiten.NewImageFromImage(bgimg) ,
		
	}
}

func newBoard() *board {
	return &board{ tiles : populateBoard(boardWidth, boardHeight),
	}
}

func (b *board) draw() {
	if b.boardImage == nil {
		b.boardImage = ebiten.NewImage(GameWidth,GameHeight*boardHRatio)
	}
	b.boardImage.Fill(backgroundColor)
	for _, t := range b.tiles {
		t.draw(b.boardImage)
	}
	//draw background
	//draw foreground
}
