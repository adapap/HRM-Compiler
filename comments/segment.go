package main	

import (
	"fmt"
)

type Segment struct {
	x float64
	y float64
}
var DIVIDER Segment = Segment{2, 2}

func CharacterSegments(c rune) ([]Segment, error) {
	switch c {
	case 'A':
		return []Segment{
			Segment{-1./2, -1},
			Segment{-1./2, 1./2},
			Segment{-1./4, 1},
			Segment{1./4, 1},
			Segment{1./2, 1./2},
			Segment{1./2, -1},
			DIVIDER,
			Segment{-1./2, 0},
			Segment{1./2, 0},
		}, nil
	case 'B':
		return []Segment{
			Segment{-1./2, -1},
			Segment{-1./2, 1},
			Segment{1./4, 1},
			Segment{1./2, 1./2},
			Segment{1./4, 0},
			Segment{-1./2, 0},
			DIVIDER,
			Segment{1./4, 0},
			Segment{1./2, -1./2},
			Segment{1./4, -1},
			Segment{-1./2, -1},
		}, nil
	case 'C':
		return []Segment{
			Segment{1./2, 2./3},
			Segment{0, 1},
			Segment{-1./2, 2./3},
			Segment{-1./2, -2./3},
			Segment{0, -1},
			Segment{1./2, -2./3},
		}, nil
	case ' ':
		return nil, nil
	default:
		return nil, fmt.Errorf("No data for character '%c'", c)
	}
}
