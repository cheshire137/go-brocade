package models

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

type Options struct {
	Width          int
	Height         int
	OutPath        string
	Background     string
	PatternConfigs []*PatternConfiguration
}

func (o *Options) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Dimensions: %dx%d\n", o.Width, o.Height))
	sb.WriteString(fmt.Sprintf("Output: %s\n", o.OutPath))
	for _, config := range o.PatternConfigs {
		sb.WriteString(config.String())
	}
	return sb.String()
}

func ParseOptions(totalPatterns int) (*Options, error) {
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

	colors, err := parseColors(colorsStr, totalPatterns)
	if err != nil {
		return nil, err
	}

	xOffsets, err := parseOffsets(xOffsetsStr, totalPatterns)
	if err != nil {
		return nil, err
	}

	yOffsets, err := parseOffsets(yOffsetsStr, totalPatterns)
	if err != nil {
		return nil, err
	}

	bgColor, colors := colors[0], colors[1:]

	patternConfigs := []*PatternConfiguration{}
	for i, color := range colors {
		patternConfig := NewPatternConfiguration(color, xOffsets[i], yOffsets[i])
		patternConfigs = append(patternConfigs, patternConfig)
	}

	return &Options{
		Width:          width,
		Height:         height,
		OutPath:        outPath,
		PatternConfigs: patternConfigs,
		Background:     bgColor,
	}, nil
}

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
