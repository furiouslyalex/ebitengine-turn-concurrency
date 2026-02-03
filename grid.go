package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	SpriteSize   = 16
	Padding      = 4
	CellSize     = SpriteSize + Padding*2 // 24 pixels per cell
	GridCols     = 20
	GridRows     = 15
)

var (
	GridOffsetX = (ScreenWidth - GridCols*CellSize) / 2
	GridOffsetY = (ScreenHeight - GridRows*CellSize) / 2
)

// converts screen coordinates to grid cell indices
func GetGridCellFromPosition(x, y int) (int, int, bool) {
	if x < GridOffsetX || y < GridOffsetY {
		return -1, -1, false
	}

	col := (x - GridOffsetX) / CellSize
	row := (y - GridOffsetY) / CellSize

	if col < 0 || col >= GridCols || row < 0 || row >= GridRows {
		return -1, -1, false
	}

	return col, row, true
}

// center position of grid cell for sprite position
func GetGridCellCenter(col, row int) (float64, float64) {
	x := float64(GridOffsetX + col*CellSize + Padding)
	y := float64(GridOffsetY + row*CellSize + Padding)
	return x, y
}

// draws grid background
func DrawGridBackground(screen *ebiten.Image) {
	gridBgColor := color.RGBA{150, 200, 240, 255}
	gridWidth := float32(GridCols * CellSize)
	gridHeight := float32(GridRows * CellSize)
	vector.DrawFilledRect(screen, float32(GridOffsetX), float32(GridOffsetY), gridWidth, gridHeight, gridBgColor, false)
}

// draws game grid 16x16 + padding
func DrawGrid(screen *ebiten.Image) {
	gridColor := color.RGBA{80, 120, 180, 255}

	for row := 0; row < GridRows; row++ {
		for col := 0; col < GridCols; col++ {
			x := float32(GridOffsetX + col*CellSize)
			y := float32(GridOffsetY + row*CellSize)

			// Draw cell border
			vector.StrokeRect(screen, x, y, CellSize, CellSize, 1, gridColor, false)
		}
	}
}
