package patterns

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
)

// Overcast is the "Overcast" pattern from Hero Patterns <https://www.heropatterns.com>
type Overcast struct {
	ID            string
	maskID        string
	patternWidth  int
	patternHeight int
}

func NewOvercast() *Overcast {
	return &Overcast{
		ID:            "Overcast",
		maskID:        "Overcast-mask",
		patternWidth:  80,
		patternHeight: 80,
	}
}

func (p *Overcast) DefinePattern(width int, height int, canvas *svg.SVG) {
	canvas.Def()
	canvas.Pattern(p.ID, 0, 0, p.patternWidth, p.patternHeight, "user", "stroke:white;stroke-linecap:square;stroke-width:1")

	canvas.Gstyle("fill:#000")
	canvas.Path("M0 0h80v80H0V0zm20 20v40h40V20H20zm20 35a15 15 0 1 1 0-30 15 15 0 0 1 0 30z")
	canvas.Path("M15 15h50l-5 5H20v40l-5 5V15zm0 50h50V15L80 0v80H0l15-15zm32.07-32.07l3.54-3.54A15 15 0 0 1 29.4 50.6l3.53-3.53a10 10 0 1 0 14.14-14.14zM32.93 47.07a10 10 0 1 1 14.14-14.14L32.93 47.07z")
	canvas.Gend()

	canvas.PatternEnd()

	canvas.Mask(p.maskID, 0, 0, width, height)
	canvas.Rect(0, 0, width, height, fmt.Sprintf("fill:url(#%s)", p.ID))
	canvas.MaskEnd()

	canvas.DefEnd()
}

func (p *Overcast) Style(color string) string {
	return fmt.Sprintf("mask:url(#%s);fill:%s", p.maskID, color)
}
