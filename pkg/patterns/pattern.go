package patterns

import svg "github.com/ajstarks/svgo"

type Pattern interface {
	DefinePattern(width int, height int, canvas *svg.SVG)
	Style(color string) string
}

var _ Pattern = (*Fleur)(nil)
var _ Pattern = (*FlowerAndStemSwirl)(nil)
var _ Pattern = (*Jigsaw)(nil)
var _ Pattern = (*Jupiter)(nil)
var _ Pattern = (*Overcast)(nil)
var _ Pattern = (*SwirlyStem)(nil)
