package main

import (
	"flag"
	"fmt"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/cheshire137/go-brocade/pkg/patterns"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	var width int
	flag.IntVar(&width, "w", 1063,
		"Width of SVG image to produce; defaults to 8.5\" at 125px per inch")

	var height int
	flag.IntVar(&height, "h", 1375,
		"Height of SVG image to produce; defaults to 11\" at 125px per inch")

	var outPath string
	flag.StringVar(&outPath, "out", "",
		"Name of SVG file to create, e.g., my-image.svg")

	var color string
	flag.StringVar(&color, "color", "",
		"Color to apply to the pattern, e.g., #ff00ff; defaults to a randomly chosen color")

	flag.Parse()
	if len(outPath) < 1 {
		fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
		return
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		fmt.Println("Could not create SVG file: " + err.Error())
		os.Exit(1)
		return
	}

	canvas := svg.New(outFile)
	canvas.Start(width, height)

	pattern := patterns.NewFlowerAndStemSwirl()
	pattern.DefinePattern(width, height, canvas)

	if len(color) < 1 {
		color = colorful.FastHappyColor().Hex()
	}
	fmt.Printf("Using color %s\n", color)
	canvas.Rect(0, 0, width, height, pattern.Style(color))

	canvas.End()

	fmt.Printf("Wrote %s\n", outPath)
}
