package hexgame

import (
	"github.com/hajimehoshi/ebiten/v2"
//" image/color"
)
	

const(
	GameWidth = 960 
	GameHeight = 540 
)

type Game struct {
	board *board
}

func (g *Game) Update() error {
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	midground := ebiten.NewImage(GameWidth,60)
	midground.Fill(black)
	g.board.draw()
	screen.DrawImage(g.board.boardImage,nil)
}
func (g *Game) Layout(outsideWidth, ousideHeight int) (screenWidth, screenHeight int) {
	return GameWidth, GameHeight
}

func NewGame() (*Game, error) {
	return &Game{
		newBoard(),
	}, nil
}
