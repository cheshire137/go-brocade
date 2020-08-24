package patterns

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
)

type YYY struct {
	name          string
	maskID        string
	patternWidth  int
	patternHeight int
}

func NewYYY() *YYY {
	return &YYY{
		name:          "YYY",
		maskID:        "YYY-mask",
		patternWidth:  60,
		patternHeight: 96,
	}
}

func (p *YYY) Name() string {
	return p.name
}

func (p *YYY) DefinePattern(width int, height int, canvas *svg.SVG) {
	canvas.Def()
	canvas.Pattern(p.name, 0, 0, p.patternWidth, p.patternHeight, "user", "stroke:white;stroke-linecap:square;stroke-width:1")

	canvas.Gstyle("fill-rule:evenodd")
	canvas.Gstyle("fill:#000")
	canvas.Path("M36 10a6 6 0 0 1 12 0v12a6 6 0 0 1-6 6 6 6 0 0 0-6 6 6 6 0 0 1-12 0 6 6 0 0 0-6-6 6 6 0 0 1-6-6V10a6 6 0 1 1 12 0 6 6 0 0 0 12 0zm24 78a6 6 0 0 1-6-6 6 6 0 0 0-6-6 6 6 0 0 1-6-6V58a6 6 0 1 1 12 0 6 6 0 0 0 6 6v24zM0 88V64a6 6 0 0 0 6-6 6 6 0 0 1 12 0v12a6 6 0 0 1-6 6 6 6 0 0 0-6 6 6 6 0 0 1-6 6z")
	canvas.Gend()

	canvas.Gend()

	canvas.PatternEnd()

	canvas.Mask(p.maskID, 0, 0, width, height)
	canvas.Rect(0, 0, width, height, fmt.Sprintf("fill:url(#%s)", p.name))
	canvas.MaskEnd()

	canvas.DefEnd()
}

func (p *YYY) Style(color string, offsetX int, offsetY int) string {
	return fmt.Sprintf("mask:url(#%s);fill:%s;transform:translate(%dpx, %dpx)", p.maskID, color, offsetX, offsetY)
}
