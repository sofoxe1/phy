package main

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

func main(){
	render()
}
func square( x float64, y float64, img *draw2dimg.GraphicContext, size float64){
	img.MoveTo(x,y)
	img.LineTo(x+size,0)
	img.LineTo(x+size,y+size)
	img.LineTo(0,y+size)
	img.LineTo(0,0)
	img.FillStroke()

}

func render(){
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
	gc := draw2dimg.NewGraphicContext(dest)
	
	// Set some properties
	co:= color.RGBA{0xff, 0x05, 0x11, 0xff}
	gc.SetFillColor(co)
	gc.SetStrokeColor(co)
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.MoveTo(10, 10) // should always be called first for a new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()
	square(0,0,gc,10)

	// Save to file
	draw2dimg.SaveToPngFile("../dynamic/img.png", dest)

}