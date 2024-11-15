package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

// Update runs every tick/frame
func (g *Game) Update() error {
	// Get the player's size
	playerWidth := g.player.Image.Bounds().Dx()
	playerHeight := g.player.Image.Bounds().Dy()

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

	// Move Up
	if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)) &&
		g.player.YPos > 0 {
		g.player.YPos -= 2 * moveSpeed
		fmt.Printf("Player Pos: %v %v\n", g.player.XPos, g.player.YPos)
	}

	// Move Down
	if (ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS)) &&
		g.player.YPos < float64(GscreenHeight-playerHeight) {
		g.player.YPos += 2 * moveSpeed
		fmt.Printf("Player Pos: %v %v\n", g.player.XPos, g.player.YPos)
	}

	for i := range g.Enemies {
		g.updateEnemy(&g.Enemies[i])
	}

	return nil
}
