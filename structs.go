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
	Show bool
}

// GameState stores data on if the game runs,
// and weather it is won or lost
type GameState struct {
	GameEnd bool
	GameWon bool
}

// Proj stores bullet pos
type Proj struct {
	X, Y   float64
	Delete bool
	Show   bool
}

// Game Stores all required game data
type Game struct {
	player      *Player
	Enemies     []Enemy
	DeadEnemies int
	ShowDebug   bool
	GameState   GameState
	Projs       []Proj
	Score       uint
	TTS         int // Number of frames until we let the player shoot again
}
