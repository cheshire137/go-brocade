package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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

	var colorsStr string
	flag.StringVar(&colorsStr, "colors", "",
		"Comma-separated string of colors to apply to the pattern, e.g., #ff00ff,#999999; defaults to randomly chosen colors")

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

	var colors []string
	if len(colorsStr) > 0 {
		colors = strings.Split(colorsStr, ",")
	} else {
		colors = []string{}
	}
	totalColors := 5
	if len(colors) < totalColors {
		palette, err := colorful.HappyPalette(totalColors - len(colors))
		if err != nil {
			fmt.Println("Could not generate colors: " + err.Error())
			os.Exit(1)
			return
		}
		for _, color := range palette {
			colors = append(colors, color.Hex())
		}
	}
	fmt.Printf("Using colors: %s\n", strings.Join(colors, ", "))

	allPatterns := []patterns.Pattern{
		patterns.NewSwirlyStem(),
		patterns.NewFlowerAndStemSwirl(),
	}

	for i, pattern := range allPatterns {
		color := colors[i%totalColors]
		pattern.DefinePattern(width, height, canvas)
		canvas.Rect(0, 0, width, height, pattern.Style(color))
	}

	canvas.End()

	fmt.Printf("Wrote %s\n", outPath)
}
