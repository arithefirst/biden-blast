package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// bidenBlast creates a new projectile at the player's pos
func (g *Game) bidenBlast(player *Player) {
	g.Projs = append(g.Projs, Proj{
		X:      player.XPos + float64(player.Image.Bounds().Dx()/2) + 5,
		Y:      player.YPos,
		Show:   true,
		Delete: false,
	})
}

// drawProj draws a single projectile
func (g *Game) drawProj(screen *ebiten.Image, proj *Proj) {
	if proj.Show {
		geoM := &ebiten.GeoM{}
		geoM.Translate(proj.X, proj.Y)

		playerImg := ebiten.NewImage(10, 20)
		playerImg.Fill(color.RGBA{
			R: 255, G: 255,
			B: 255, A: 255,
		})

		screen.DrawImage(playerImg, &ebiten.DrawImageOptions{GeoM: *geoM})
	}
}

// updateProj updates all projectile positions
func (g *Game) updateProjs(proj *Proj) {
	if proj.Y < -20 {
		// Remove projectile when it goes offscreen
		proj.Delete = true
	} else {
		for i := range g.Enemies {
			// Check if the projectile is colliding with an enemy
			if proj.X+10 >= g.Enemies[i].X && proj.X <= g.Enemies[i].X+64 &&
				proj.Y+20 >= g.Enemies[i].Y && proj.Y <= g.Enemies[i].Y+64 &&
				g.Enemies[i].Show && proj.Show {

				// Remove projectile
				proj.Delete = true
				proj.Show = false

				// Remove enemy
				g.Enemies[i].Show = false
				g.Score += 100
			}
		}
		// Otherwise advance the projectile
		proj.Y -= 4
	}
}
