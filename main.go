package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var (
	GscreenWidth  = 800
	GscreenHeight = 600
)

func NewGame() *Game {
	image, _, err := ebitenutil.NewImageFromFile("sprites/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Game{

		GameState: GameState{
			GameEnd: false,
			GameWon: false,
		},

		player: &Player{
			Image: image,
			XPos:  float64((GscreenWidth - image.Bounds().Dx()) / 2),
			YPos:  float64(GscreenHeight - image.Bounds().Dy()),
		},

		Enemies: []Enemy{
			{X: 0, Y: 10, Show: true},
			{X: 74, Y: 10, Show: true},
			{X: 148, Y: 10, Show: true},
			{X: 222, Y: 10, Show: true},
			{X: 0, Y: 84, Show: true},
			{X: 74, Y: 84, Show: true},
			{X: 148, Y: 84, Show: true},
			{X: 222, Y: 84, Show: true},
			{X: 0, Y: 158, Show: true},
			{X: 74, Y: 158, Show: true},
			{X: 148, Y: 158, Show: true},
			{X: 222, Y: 158, Show: true},
		},
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.GameState.GameWon {

	} else {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Score: %d", g.Score), 0, 0)
		// Draw debug data
		if g.ShowDebug {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()), 0, 12)
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Projs: %d", len(g.Projs)), 0, 24)
		}

		// Draw the enemies
		for i := range g.Enemies {
			g.drawEnemy(screen, &g.Enemies[i])
		}

		// Draw the projectiles
		for i := range g.Projs {
			g.drawProj(screen, &g.Projs[i])
		}

		// Draw the player
		geoM := &ebiten.GeoM{}
		geoM.Translate(g.player.XPos, g.player.YPos)
		screen.DrawImage(g.player.Image, &ebiten.DrawImageOptions{
			GeoM: *geoM,
		})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GscreenWidth, GscreenHeight
}

func main() {
	ebiten.SetWindowTitle("Space invaders")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
