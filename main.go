package main

import (
	"fmt"
	"image"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() { 
	nnnarv := Nnnarv{}

	// nnnarv.Init(5, 3, 0, 5)
	// nnnarv.AddPoint(Point{[]float64{1, 1, 0}, []float64{0}})
	// nnnarv.AddPoint(Point{[]float64{1, 1, 3}, []float64{150}})
	// fmt.Println(nnnarv.GetValueOfPoint([]float64{1,1,2},2))

	fmt.Println("")
	nnnarv.Init(5, 785, 0, 256)
	go loadCsvToNnnarv(&nnnarv, "./train.csv")
	pixelgl.Run(run)
}

func run() {
	img := image.NewRGBA(image.Rect(0, 0, 280, 280))
	win, _ := createWindow(img)
	
	for !win.Closed() {
		if (win.Pressed(pixelgl.MouseButtonLeft)) {
			x,y := getCursorPosition(win)
			drawPixel(img, x, y)
		}
		loop(win, img)
	}
	win.Destroy()
}

func loop(win *pixelgl.Window, img *image.RGBA) {
	pic := pixel.PictureDataFromImage(img)
	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(pixel.RGB(0, 0, 0))
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	win.Update()
}