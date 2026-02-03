package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Image *ebiten.Image
	X, Y  float64
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImage, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		log.Fatal(err)
	}
	gladiatorImage, _, err := ebitenutil.NewImageFromFile("assets/images/gladiator.png")
	if err != nil {
		log.Fatal(err)
	}
	skeletonImage, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}
	lionImage, _, err := ebitenutil.NewImageFromFile("assets/images/lion.png")
	if err != nil {
		log.Fatal(err)
	}

	// sprites start at corners
	playerX, playerY := GetGridCellCenter(0, 0)                   // top left
	enemy1X, enemy1Y := GetGridCellCenter(GridCols-1, 0)          // top right
	enemy2X, enemy2Y := GetGridCellCenter(0, GridRows-1)          // bottom left
	enemy3X, enemy3Y := GetGridCellCenter(GridCols-1, GridRows-1) // bottom right

	if err := ebiten.RunGame(&Game{Player: &Sprite{
		Image: playerImage,
		X:     playerX,
		Y:     playerY,
	}, Enemies: []*Sprite{
		{
			Image: gladiatorImage,
			X:     enemy1X,
			Y:     enemy1Y,
		}, {
			Image: skeletonImage,
			X:     enemy2X,
			Y:     enemy2Y,
		}, {
			Image: lionImage,
			X:     enemy3X,
			Y:     enemy3Y,
		},
	}}); err != nil {
		log.Fatal(err)
	}
}
