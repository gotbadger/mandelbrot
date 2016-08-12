package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/gotbadger/palettes"
	"github.com/gotbadger/terminalImage"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
)

func mandelbrot(x float64, y float64, imax int) int {
	a := complex(x, y)
	var i int = 0
	for z := a; cmplx.Abs(z) < 2 && i < imax; i++ {
		z = z*z + a
	}
	return i
}

func transposeCoordiante(i int, min float64, step float64) float64 {
	return min + (step * float64(i))
}

func getRange(min float64, max float64, step float64) int {
	return int(math.Floor((max - min) / step))
}

func main() {

	// comand line options
	step := flag.Float64("step", 0.003, "a pixel is drawn for each step between coordinates")
	maxIterations := flag.Int("i", 30, "max number of iterations / colours")
	yMMin := flag.Float64("y0", -1.2, "from Y")
	yMMax := flag.Float64("y1", 1.2, "to Y")
	xMMin := flag.Float64("x0", -2.0, "from X")
	xMMax := flag.Float64("x1", 1.0, "to X")
	flag.Parse()

	// transparent := color.RGBA{0,0,0,0}
	black := color.RGBA{0, 0, 0, 255}
	colours := palettes.Rainbow(*maxIterations)

	h := getRange(*yMMin, *yMMax, *step)
	w := getRange(*xMMin, *xMMax, *step)
	fmt.Println("maxIterations:", *maxIterations)
	fmt.Printf("Generating mandelbrot W:%dpx H:%dpx\n", w, h)

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {

			depth := mandelbrot(
				transposeCoordiante(x, *xMMin, *step),
				transposeCoordiante(y, *yMMin, *step),
				*maxIterations)
			shade := black
			if depth < *maxIterations {
				shade = colours[depth]
			}

			img.Set(x, y, shade)
		}
	}

	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	terminalImage.Print(buf.Bytes())
}
