package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/cheshire137/go-brocade/pkg/patterns"
)

func main() {
	width := 1000
	height := 1000
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)

	jigsaw := patterns.NewJigsaw()
	jigsaw.DefinePattern(canvas)

	flowerStemSwirl := patterns.NewFlowerAndStemSwirl()
	flowerStemSwirl.DefinePattern(canvas)

	canvas.Circle(500, 500, 500, jigsaw.Fill())
	canvas.Circle(500, 500, 500, flowerStemSwirl.Fill())

	canvas.End()
}
