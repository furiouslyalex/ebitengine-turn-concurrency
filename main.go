package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Image *ebiten.Image
	X, Y  float64
}

type Game struct {
	Player  *Sprite
	Enemies []*Sprite
}

func (g *Game) Update() error {
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// load sprites
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

	if err := ebiten.RunGame(&Game{Player: &Sprite{
		Image: playerImage,
		X:     100,
		Y:     100,
	}, Enemies: []*Sprite{
		{
			Image: gladiatorImage,
			X:     75,
			Y:     75,
		}, {
			Image: skeletonImage,
			X:     150,
			Y:     150,
		}, {
			Image: lionImage,
			X:     50,
			Y:     50,
		},
	}}); err != nil {
		log.Fatal(err)
	}
}
