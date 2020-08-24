package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	svg "github.com/ajstarks/svgo"
	"github.com/cheshire137/go-brocade/pkg/models"
	"github.com/cheshire137/go-brocade/pkg/patterns"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	options, err := models.ParseOptions()
	if err != nil {
		fmt.Println("Could not parse options: " + err.Error())
		os.Exit(1)
		return
	}

	if options.ListPatterns {
		fmt.Print("Patterns: ")
		patternNames := patterns.PatternNames()
		fmt.Printf("%s\n", strings.Join(patternNames, ", "))
		os.Exit(0)
		return
	}

	if len(options.OutPath) < 1 {
		fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
		return
	}

	outFile, err := os.Create(options.OutPath)
	if err != nil {
		fmt.Println("Could not create SVG file: " + err.Error())
		os.Exit(1)
		return
	}

	fmt.Println("Using options:")
	fmt.Println(options.String())

	canvas := svg.New(outFile)
	canvas.Start(options.Width, options.Height)

	canvas.Rect(0, 0, options.Width, options.Height, fmt.Sprintf("fill:%s", options.Background))

	for _, config := range options.PatternConfigs {
		pattern := config.Pattern
		pattern.DefinePattern(options.Width, options.Height, canvas)
		style := pattern.Style(config.Color, config.Xoffset, config.Yoffset)
		canvas.Rect(0, 0, options.Width, options.Height, style)
	}

	canvas.End()

	fmt.Printf("Wrote %s\n", options.OutPath)
}
