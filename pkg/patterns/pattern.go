package patterns

import (
	"errors"

	svg "github.com/ajstarks/svgo"
)

type Pattern interface {
	Name() string
	DefinePattern(width int, height int, canvas *svg.SVG)
	Style(color string, offsetX int, offsetY int) string
}

var _ Pattern = (*Fleur)(nil)
var _ Pattern = (*FlowerAndStem)(nil)
var _ Pattern = (*Jigsaw)(nil)
var _ Pattern = (*Jupiter)(nil)
var _ Pattern = (*Overcast)(nil)
var _ Pattern = (*SwirlyStem)(nil)

func PatternNames() []string {
	return []string{
		"fleur",
		"flowerAndStem",
		"jigsaw",
		"jupiter",
		"overcast",
		"sCurve",
		"swirlyStem",
	}
}

func PatternNameToPattern(name string) (Pattern, error) {
	if name == "fleur" {
		return NewFleur(), nil
	}
	if name == "flowerAndStem" {
		return NewFlowerAndStem(), nil
	}
	if name == "jigsaw" {
		return NewJigsaw(), nil
	}
	if name == "jupiter" {
		return NewJupiter(), nil
	}
	if name == "overcast" {
		return NewOvercast(), nil
	}
	if name == "sCurve" {
		return NewSCurve(), nil
	}
	if name == "swirlyStem" {
		return NewSwirlyStem(), nil
	}
	return nil, errors.New("Invalid pattern name: " + name)
}
