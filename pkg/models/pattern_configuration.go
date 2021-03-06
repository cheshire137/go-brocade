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

func NewPatternConfiguration(patternName string, color string, xoffset int, yoffset int) (*PatternConfiguration, error) {
	pattern, err := patterns.PatternNameToPattern(patternName)
	if err != nil {
		return nil, err
	}
	return &PatternConfiguration{
		Color:   color,
		Xoffset: xoffset,
		Yoffset: yoffset,
		Pattern: pattern,
	}, nil
}

func (c *PatternConfiguration) Name() string {
	pattern := c.Pattern
	return pattern.Name()
}

func (c *PatternConfiguration) String() string {
	return fmt.Sprintf("%s %s %d,%d", c.Name(), c.Color, c.Xoffset, c.Yoffset)
}
