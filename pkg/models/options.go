package models

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/cheshire137/go-brocade/pkg/patterns"
	"github.com/lucasb-eyer/go-colorful"
)

type Options struct {
	Width           int
	Height          int
	OutPath         string
	Background      string
	PatternConfigs  []*PatternConfiguration
	ListPatterns    bool
	RandomCount     int
	InteractiveMode bool
}

func (o *Options) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Dimensions: %dx%d\n", o.Width, o.Height))

	sb.WriteString("Output: ")
	sb.WriteString(o.OutPath)
	sb.WriteString("\n")

	sb.WriteString("Background: ")
	sb.WriteString(o.Background)
	sb.WriteString("\n")

	sb.WriteString("Patterns:")
	for _, config := range o.PatternConfigs {
		sb.WriteString("\n- ")
		sb.WriteString(config.String())
	}

	return sb.String()
}

func ParseOptions() (*Options, error) {
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
		"Comma-separated list of colors to apply to the pattern, e.g., #ff00ff,#999999.\n"+
			"Defaults to randomly chosen colors. The first color will be used for the\n"+
			"background color.")

	var xOffsetsStr string
	flag.StringVar(&xOffsetsStr, "xoffsets", "",
		"Comma-separated list of X-axis offset values, in pixels, for each pattern.\n"+
			"If omitted, will default to 0px.")

	var yOffsetsStr string
	flag.StringVar(&yOffsetsStr, "yoffsets", "",
		"Comma-separated list of Y-axis offset values, in pixels, for each pattern.\n"+
			"If omitted, will default to 0px.")

	var listPatterns bool
	flag.BoolVar(&listPatterns, "list", false,
		"Pass this to list available patterns.")

	var interactiveMode bool
	flag.BoolVar(&interactiveMode, "i", false,
		"Set to true to be prompted for each option. Ignores other options that were passed.")

	var randomCount int
	flag.IntVar(&randomCount, "random", 0,
		"Number of patterns to randomly include. Set to >0 to use, 0 to specify patterns\n"+
			"yourself.")

	var patternsStr string
	flag.StringVar(&patternsStr, "patterns", "flowerAndStem,swirlyStem,fleur",
		"Comma-separated list of pattern names to include in the generated image.\n"+
			"Ignored when -random is >0.")

	flag.Parse()

	patternNames := parsePatternNames(randomCount, patternsStr)
	totalPatterns := len(patternNames)
	if totalPatterns < 1 {
		return nil, errors.New("At least one pattern must be specified")
	}

	colors, err := parseColors(colorsStr, totalPatterns+1)
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

	patternConfigs := make([]*PatternConfiguration, totalPatterns)
	for i, patternName := range patternNames {
		patternConfig, err := NewPatternConfiguration(patternName, colors[i], xOffsets[i], yOffsets[i])
		if err != nil {
			return nil, err
		}
		patternConfigs[i] = patternConfig
	}

	return &Options{
		Width:           width,
		Height:          height,
		OutPath:         outPath,
		PatternConfigs:  patternConfigs,
		Background:      bgColor,
		ListPatterns:    listPatterns,
		InteractiveMode: interactiveMode,
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
	offsets := make([]int, totalOffsets)
	resumeIndex := 0
	if len(offsetsStr) > 0 {
		offsetStrs := strings.Split(offsetsStr, ",")
		for i, offsetStr := range offsetStrs {
			offset, err := strconv.Atoi(offsetStr)
			if i < totalOffsets {
				if err == nil {
					offsets[i] = offset
				} else {
					offsets[i] = 0
				}
				resumeIndex = i
			}
		}
	}
	for i := resumeIndex; i < totalOffsets; i++ {
		offsets[i] = 0
	}
	return offsets, nil
}

func parsePatternNames(randomCount int, patternsStr string) []string {
	if randomCount < 1 {
		return strings.Split(patternsStr, ",")
	}

	allNames := patterns.PatternNames()
	patternNames := make([]string, randomCount)

	for i := 0; i < randomCount; i++ {
		nameIndex := rand.Intn(len(allNames))
		patternNames[i] = allNames[nameIndex]
	}

	return patternNames
}
