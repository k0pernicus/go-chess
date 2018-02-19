package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/miketmoore/go-chess/board"
	"golang.org/x/image/colornames"
)

// var spriteSheetPath = "assets/spritesheet.png"
var spriteSheetPath = "assets/chess-pieces.png"

func run() {
	// Chess board is 8x8
	// top left is white

	// Load sprite sheet graphic
	piecesPic, err := loadPicture(spriteSheetPath)
	if err != nil {
		panic(err)
	}

	pieces := makePieces(piecesPic)

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  "Chess",
		Bounds: pixel.R(0, 0, 400, 400),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Darkgray)

	board := board.Build(50, colornames.Black, colornames.White)
	for !win.Closed() {
		win.Update()
		mat := pixel.IM
		mat = mat.Moved(win.Bounds().Center())
		pieces["white"]["king"].Draw(win, mat)
		for _, square := range board {
			square.Draw(win)
		}
	}
}

func main() {
	pixelgl.Run(run)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func makePieces(pic pixel.Picture) map[string]map[string]*pixel.Sprite {
	var xInc float64 = 62
	var yInc float64 = 60
	return map[string]map[string]*pixel.Sprite{
		"white": map[string]*pixel.Sprite{
			"king":   newSprite(pic, 0, 0, xInc, yInc),
			"queen":  newSprite(pic, xInc, 0, xInc*2, yInc),
			"rook":   newSprite(pic, xInc*2, 0, xInc*3, yInc),
			"knight": newSprite(pic, xInc*3, 0, xInc*4, yInc),
			"bishop": newSprite(pic, xInc*4, 0, xInc*5+5, yInc),
			"pawn":   newSprite(pic, xInc*5+5, 0, xInc*6, yInc),
		},
		"black": map[string]*pixel.Sprite{
			"king":   newSprite(pic, 0, yInc, xInc, yInc*2),
			"queen":  newSprite(pic, xInc, yInc, xInc*2, yInc*2+5),
			"rook":   newSprite(pic, xInc*2, yInc, xInc*3, yInc*2),
			"knight": newSprite(pic, xInc*3, yInc, xInc*4, yInc*3),
			"bishop": newSprite(pic, xInc*4, yInc, xInc*5+5, yInc*4),
			"pawn":   newSprite(pic, xInc*5+5, yInc, xInc*6, yInc*5),
		},
	}
}

func newSprite(pic pixel.Picture, xa, ya, xb, yb float64) *pixel.Sprite {
	return pixel.NewSprite(pic, pixel.Rect{pixel.Vec{xa, ya}, pixel.Vec{xb, yb}})
}
