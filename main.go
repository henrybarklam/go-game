package main

import (
	"embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var assets embed.FS

type Game struct{}

func (game *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWIdth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
