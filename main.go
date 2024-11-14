package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var (
	moveSpeed     float64 = 2
	GscreenWidth  int     = 800
	GscreenHeight int     = 600
)

func NewGame() *Game {
	image, _, err := ebitenutil.NewImageFromFile("sprites/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Game{
		player: &Player{
			Image: image,
			XPos:  0,
			YPos:  0,
		},
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	geoM := &ebiten.GeoM{}
	geoM.Translate(g.player.XPos, g.player.YPos)

	screen.DrawImage(g.player.Image, &ebiten.DrawImageOptions{
		GeoM: *geoM,
	})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GscreenWidth, GscreenHeight
}

func main() {
	ebiten.SetWindowTitle("Space invaders")
	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
