package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 10, 30, 255})

	/* vector.StrokeCircle(screen, 320, 240, 100, 5, color.RGBA{255, 105, 180, 255}, false) */
	/* vector.StrokeCircle(screen, 200, 200, 100, 2, color.RGBA{255, 0, 0, 255}, true)
	vector.StrokeCircle(screen, 200, 200, 75, 2, color.RGBA{255, 0, 0, 255}, true)
	vector.StrokeCircle(screen, 200, 200, 50, 2, color.RGBA{255, 0, 0, 255}, true)
	vector.StrokeCircle(screen, 200, 200, 25, 2, color.RGBA{255, 0, 0, 255}, true)
	vector.StrokeCircle(screen, 200, 200, 12.5, 2, color.RGBA{255, 0, 0, 255}, true) */
	var geoM ebiten.GeoM
	geoM.Translate(300, 200)
	geoM.Rotate(0.5)
	transformedRect := ebiten.NewImage(100, 50)
	vector.StrokeRect(transformedRect, 0, 0, 100, 50, 5, color.RGBA{0, 0, 255, 255}, false)
	/* screen.DrawImage(transformedRect, &ebiten.DrawImageOptions{GeoM: geoM}) */
	/* vector.StrokeRect(screen, 100, 100, 100, 50, 5, color.RGBA{0, 255, 0, 255}, false) */

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
