package main

import (
	"embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var assets embed.FS

var PlayerSprite = mustLoadImage("assets/player.png")

type Game struct{}

func (game *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWIdth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name) //Walrus operator
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func main() {
	g := &Game{}
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
