package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func (g *Game) bidenBlast(player *Player) {
	g.Projs = append(g.Projs, Proj{
		X: player.XPos + float64(player.Image.Bounds().Dx()/2),
		Y: player.YPos,
	})
}

// drawProj draws a single projectile
func (g *Game) drawProj(screen *ebiten.Image, proj *Proj) {
	geoM := &ebiten.GeoM{}
	geoM.Translate(proj.X, proj.Y)

	playerImg := ebiten.NewImage(10, 20)
	playerImg.Fill(color.RGBA{
		R: 255, G: 255,
		B: 255, A: 255,
	})

	screen.DrawImage(playerImg, &ebiten.DrawImageOptions{GeoM: *geoM})
}

// updateProj updates all projectile positions
func (g *Game) updateProjs(proj *Proj) {
	if proj.Y < -20 {
		proj.Delete = true
	} else {
		for i := range g.Enemies {
			if proj.X >= g.Enemies[i].X && proj.X <= g.Enemies[i].X+64 &&
				proj.Y >= g.Enemies[i].Y && proj.Y <= g.Enemies[i].Y+64 &&
				g.Enemies[i].Show {
				proj.Delete = true
				g.Enemies[i].Show = false
				g.Score += 100
			}
		}
		proj.Y -= 4
	}
}
