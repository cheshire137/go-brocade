package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/cheshire137/go-brocade/pkg/patterns"
	"github.com/lucasb-eyer/go-colorful"
)

func parseColors(colorsStr string, totalColors int) ([]string, error) {
	var colors []string
	if len(colorsStr) > 0 {
		colors = strings.Split(colorsStr, ",")
	} else {
		colors = []string{}
	}
	if len(colors) < totalColors {
		palette, err := colorful.HappyPalette(totalColors - len(colors))
		if err != nil {
			return nil, err
		}
		for _, color := range palette {
			colors = append(colors, color.Hex())
		}
	}
	return colors, nil
}

func parseOffsets(offsetsStr string, totalOffsets int) ([]int, error) {
	var offsets []int
	if len(offsetsStr) > 0 {
		offsetStrs := strings.Split(offsetsStr, ",")
		for _, offsetStr := range offsetStrs {
			offset, err := strconv.Atoi(offsetStr)
			if err == nil {
				offsets = append(offsets, offset)
			} else {
				offsets = append(offsets, 0)
			}
		}
	}
	for i := 0; i <= totalOffsets-len(offsets); i++ {
		offsets = append(offsets, 0)
	}
	return offsets, nil
}

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
		"Comma-separated string of colors to apply to the pattern, e.g., #ff00ff,#999999.\n"+
			"Defaults to randomly chosen colors. The first color will be used for the\n"+
			"background color.")

	var xOffsetsStr string
	flag.StringVar(&xOffsetsStr, "xoffsets", "",
		"Comma-separated string of X-axis offset values, in pixels, for each pattern.\n"+
			"If omitted, will default to 0px.")

	var yOffsetsStr string
	flag.StringVar(&yOffsetsStr, "yoffsets", "",
		"Comma-separated string of Y-axis offset values, in pixels, for each pattern.\n"+
			"If omitted, will default to 0px.")

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
	totalPatterns := 5

	colors, err := parseColors(colorsStr, totalPatterns)
	if err != nil {
		fmt.Println("Could not generate colors: " + err.Error())
		os.Exit(1)
	}
	fmt.Printf("Using colors: %s\n", strings.Join(colors, ", "))

	xOffsets, err := parseOffsets(xOffsetsStr, totalPatterns)
	if err != nil {
		fmt.Println("Could not parse X-offsets: " + err.Error())
		os.Exit(1)
	}
	fmt.Printf("Using colors: %s\n", strings.Join(colors, ", "))

	yOffsets, err := parseOffsets(yOffsetsStr, totalPatterns)
	if err != nil {
		fmt.Println("Could not parse Y-offsets: " + err.Error())
		os.Exit(1)
	}
	fmt.Printf("Using colors: %s\n", strings.Join(colors, ", "))

	canvas.Rect(0, 0, width, height, fmt.Sprintf("fill:%s", colors[0]))

	allPatterns := []patterns.Pattern{
		patterns.NewSwirlyStem(),
		patterns.NewFlowerAndStemSwirl(),
		patterns.NewFleur(),
	}

	for i, pattern := range allPatterns {
		color := colors[(i+1)%len(colors)]
		pattern.DefinePattern(width, height, canvas)
		canvas.Rect(0, 0, width, height, pattern.Style(color, xOffsets[i], yOffsets[i]))
	}

	canvas.End()

	fmt.Printf("Wrote %s\n", outPath)
}
