package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var (
	moveSpeed     float64 = 2
	GscreenWidth          = 800
	GscreenHeight         = 600
)

func NewGame() *Game {
	image, _, err := ebitenutil.NewImageFromFile("sprites/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Game{
		player: &Player{
			Image: image,
			XPos:  float64((GscreenWidth - image.Bounds().Dx()) / 2),
			YPos:  float64(GscreenHeight - image.Bounds().Dy()),
		},

		Enemies: []Enemy{
			{X: 0, Y: 10},
			{X: 74, Y: 10},
			{X: 148, Y: 10},
			{X: 222, Y: 10},
			{X: 0, Y: 84},
			{X: 74, Y: 84},
			{X: 148, Y: 84},
			{X: 222, Y: 84},
			{X: 0, Y: 158},
			{X: 74, Y: 158},
			{X: 148, Y: 158},
			{X: 222, Y: 158},
		},
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the enemies
	for i := range g.Enemies {
		g.drawEnemy(screen, g.Enemies[i])
	}

	// Draw the player
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
