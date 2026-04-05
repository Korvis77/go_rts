package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

// Read what Claude was talking about and then do a lot of studying of theory.
// Time to learn about how to decode an image from the image file's byte slice

// var (
// 	ebitenImage *ebiten.Image
// )

// func init() {
// 	// Decode an image from the image file's byte slice.
// 	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	origEbitenImage := ebiten.NewImageFromImage(img)

// 	s := origEbitenImage.Bounds().Size()
// 	ebitenImage = ebiten.NewImage(s.X, s.Y)

// 	op := &ebiten.DrawImageOptions{}
// 	op.ColorScale.ScaleAlpha(1)
// 	ebitenImage.DrawImage(origEbitenImage, op)
// }

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	unitTemplate2x2 := Build2x2UnitTemplate()

	CountNodesUseTemplateInterface(unitTemplate2x2)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}

/* vector length, this is the hypotenuse,
   hypotenuse (vector length a.k.a. offset from x and y)
   cosine for x, sine for y,
   arctan(yoffset/xoffset)
   rotation (passed into the function)

   hypotenuse * cos/sin (arctan*(yoffset/xoffset) + rotation)
                         ^^^ solve once, becomes theta
*/
