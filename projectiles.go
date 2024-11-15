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
func (g *Game) drawProj(screen *ebiten.Image, proj Proj) {
	geoM := &ebiten.GeoM{}
	geoM.Translate(proj.X, proj.Y)

	enemyImg := ebiten.NewImage(10, 20)
	enemyImg.Fill(color.RGBA{
		R: 255, G: 255,
		B: 255, A: 255,
	})

	screen.DrawImage(enemyImg, &ebiten.DrawImageOptions{GeoM: *geoM})
}

// updateProj updates all projectile positions
func (g *Game) updateProjs(projs *Proj) {
	if projs.Y < -20 {
		projs.Delete = true
	} else {
		projs.Y -= 4
	}
}
