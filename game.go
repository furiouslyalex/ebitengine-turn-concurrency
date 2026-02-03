package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player           *Sprite
	Enemies          []*Sprite
	PlayerSelected   bool
	prevMousePressed bool
}

func (g *Game) Update() error {
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButton0)

	if mousePressed && !g.prevMousePressed {
		cX, cY := ebiten.CursorPosition()

		// check if cursor is inside player sprite bounds (16x16 pixels)
		// should change this later to grid cell bounds
		if cX >= int(g.Player.X) && cX < int(g.Player.X)+16 &&
			cY >= int(g.Player.Y) && cY < int(g.Player.Y)+16 {
			// first click: select the player
			g.PlayerSelected = true
		} else if g.PlayerSelected {
			// second click: check if click is inside a grid cell
			col, row, inGrid := GetGridCellFromPosition(cX, cY)

			if inGrid {
				// move player to center of the clicked grid cell
				g.Player.X, g.Player.Y = GetGridCellCenter(col, row)

				// we only set false here because if you select outside you should be able to select again without
				// having to click the player first
				g.PlayerSelected = false
			}
		}
	}

	g.prevMousePressed = mousePressed
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// outer space background (darker)
	screen.Fill(color.RGBA{60, 80, 120, 255})

	// grid area background (lighter)
	DrawGridBackground(screen)

	// draw grid
	DrawGrid(screen)

	// draw player
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.Player.X, g.Player.Y)
	screen.DrawImage(g.Player.Image.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), &opts)
	opts.GeoM.Reset()

	// draw enemies
	for _, e := range g.Enemies {
		opts.GeoM.Translate(e.X, e.Y)
		screen.DrawImage(e.Image.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), &opts)
		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
