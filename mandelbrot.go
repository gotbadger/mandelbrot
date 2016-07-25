package main

import (
  "fmt"
  "image"
  "image/color"
  "image/png"
  "math"
  "math/cmplx"
  "bytes"
  "github.com/gotbadger/terminalImage"
  "github.com/gotbadger/palettes"
)

const (
  maxIterations = 30
  step float64 = 0.003
)
//
func mandelbrot(x float64, y float64) uint8 {
  a := complex(x,y)
  var i uint8 = 0
  for z := a; cmplx.Abs(z) < 2 && i < maxIterations; i++ {
      z = z*z + a
  }
  return i
}

func transposeCoordiante(i int, min float64) float64 {
  return min + (step * float64(i))
}

func getRange(min float64, max float64) int {
  return int(math.Floor((max-min)/step))
}

func main() {

  var yMMin, yMMax float64 = -1.2, 1.2
  var xMMin, xMMax float64 = -2.0, 1.0

  // transparent := color.RGBA{0,0,0,0}
  black := color.RGBA{0,0,0,255}
  colours := palettes.Rainbow(maxIterations)

  h := getRange(yMMin, yMMax)
  w := getRange(xMMin, xMMax)

  fmt.Printf("Generating mandelbrot W:%dpx H:%dpx\n", w, h)

  img := image.NewRGBA(image.Rect(0, 0, w, h))

  for x:=0;  x < w; x++ {
    for y:=0; y < h; y++ {

      depth := mandelbrot(transposeCoordiante(x, xMMin),transposeCoordiante(y, yMMin))
      shade := black
      if(depth < maxIterations){
        shade = colours[int(depth)]
      }

      img.Set(x, y, shade)
    }
  }

  buf := new(bytes.Buffer)
  png.Encode(buf, img)
  terminalImage.Print(buf.Bytes())
}
