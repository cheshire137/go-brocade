package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/cheshire137/go-brocade/pkg/patterns"
)

func main() {
	width := 1063  // 8.5", 125px per in
	height := 1375 // 11", 125px per in
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)

	jupiter := patterns.NewJupiter()
	jupiter.DefinePattern(width, height, canvas)

	canvas.Rect(0, 0, width, height, jupiter.Style("#ff00ff"))

	canvas.End()
}
