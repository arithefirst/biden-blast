package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

// Update runs every tick/frame
func (g *Game) Update() error {
	// Get the player's size
	playerWidth := g.player.Image.Bounds().Dx()

	// Move Left
	if (ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) &&
		g.player.XPos > 0 {
		g.player.XPos -= 2 * moveSpeed
		fmt.Printf("Player Pos: %v %v\n", g.player.XPos, g.player.YPos)
	}

	// Move Right
	if (ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)) &&
		g.player.XPos < float64(GscreenWidth-playerWidth) {
		g.player.XPos += 2 * moveSpeed
		fmt.Printf("Player Pos: %v %v\n", g.player.XPos, g.player.YPos)
	}

	// Update the enemies
	for i := range g.Enemies {
		g.updateEnemy(&g.Enemies[i])
	}

	// Update the projectiles
	for i := range g.Projs {
		g.updateProjs(&g.Projs[i])
	}

	// If the furthest proj exits the screen, delete it.
	if len(g.Projs) != 0 && g.Projs[len(g.Projs)-1].Delete {
		g.Projs = g.Projs[:len(g.Projs)-1]
	}

	// If you are able to shoot a projectile, do so when
	// space is pressed
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.TTS == 0 {
		g.bidenBlast(g.player)
		g.TTS = 20
	} else if g.TTS > 0 {
		g.TTS -= 1
	}

	return nil
}
