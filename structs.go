package main

import "github.com/hajimehoshi/ebiten/v2"

// Player stores player pos and sprite
type Player struct {
	Image      *ebiten.Image
	XPos, YPos float64
}

// Enemy stores enemy pos
type Enemy struct {
	X, Y float64
}

// Game Stores all required game data
type Game struct {
	player  *Player
	Enemies []Enemy
}
