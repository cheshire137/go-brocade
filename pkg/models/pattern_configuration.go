package models

import (
	"fmt"

	"github.com/cheshire137/go-brocade/pkg/patterns"
)

type PatternConfiguration struct {
	Xoffset int
	Yoffset int
	Color   string
	Pattern patterns.Pattern
}

func NewPatternConfiguration(color string, xoffset int, yoffset int) *PatternConfiguration {
	return &PatternConfiguration{
		Color:   color,
		Xoffset: xoffset,
		Yoffset: yoffset,
		Pattern: patterns.NewFleur(),
	}
}

func (c *PatternConfiguration) String() string {
	return fmt.Sprintf("%s %d,%d", c.Color, c.Xoffset, c.Yoffset)
}
