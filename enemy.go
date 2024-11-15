package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// updateEnemy updates an enemies position
func (g *Game) updateEnemy(enemy *Enemy) {
	enemy.X += 4
	if enemy.X >= float64(GscreenWidth) {
		// when the enemy hits the edge, have it come slightly closer
		enemy.X = 0
		enemy.Y += 40
	}

	// If the enemy touches the player, end the game
	if enemy.Show && enemy.Y+64 >= g.player.YPos &&
		enemy.X+64 >= g.player.XPos {
		g.GameState.GameEnd = true
		g.GameState.GameWon = false
	}
}

// drawEnemy draws a single enemy
func (g *Game) drawEnemy(screen *ebiten.Image, enemy *Enemy) {
	if enemy.Show {
		geoM := &ebiten.GeoM{}
		geoM.Translate(enemy.X, enemy.Y)

		enemyImg := ebiten.NewImage(64, 64)
		enemyImg.Fill(color.RGBA{R: 255})

		screen.DrawImage(enemyImg, &ebiten.DrawImageOptions{GeoM: *geoM})
	}
}
