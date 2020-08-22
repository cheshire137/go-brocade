package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/cheshire137/go-brocade/pkg/patterns"
)

func main() {
	width := 2410
	height := 2520
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)

	pattern := &patterns.FlowerAndStemSwirl{}
	pattern.DefinePattern(canvas)

	canvas.Circle(1205, 1260, 1205, "fill:url(#FlowerAndStemSwirl)")

	canvas.End()
}
