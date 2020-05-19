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
	case 'D':
		return []Segment{
			Segment{-1./2, -1},
			Segment{-1./2, 1},
			Segment{1./3, 1},
			Segment{2./3, 1./2},
			Segment{2./3, -1./2},
			Segment{1./3, -1},
			Segment{-1./2, -1},
		}, nil
	case 'E':
		return []Segment{
			Segment{1./2, -1},
			Segment{-1./2, -1},
			Segment{-1./2, 1},
			Segment{1./2, 1},
			DIVIDER,
			Segment{-1./2, 0},
			Segment{1./2, 0},
		}, nil
	case 'F':
		return []Segment{
			Segment{-1./2, -1},
			Segment{-1./2, 1},
			Segment{1./2, 1},
			DIVIDER,
			Segment{-1./2, 0},
			Segment{1./2, 0},
		}, nil
	case 'G':
		return []Segment{
			Segment{1./2, 2./3},
			Segment{1./3, 1},
			Segment{-1./3, 1},
			Segment{-1./2, 2./3},
			Segment{-1./2, -2./3},
			Segment{-1./3, -1},
			Segment{1./2, -1},
			Segment{1./2, -1./4},
			Segment{1./4, -1./4},
		}, nil
	case 'H':
		return []Segment{
			Segment{-1./2, 1},
			Segment{-1./2, -1},
			DIVIDER,
			Segment{-1./2, 0},
			Segment{1./2, 0},
			DIVIDER,
			Segment{1./2, 1},
			Segment{1./2, -1},
		}, nil
	case 'I':
		return []Segment{
			Segment{0, -1},
			Segment{0, 1},
			DIVIDER,
			Segment{-1./2, 1},
			Segment{1./2, 1},
			DIVIDER,
			Segment{-1./2, -1},
			Segment{1./2, -1},
		}, nil
	case 'J':
		return []Segment{
			Segment{-1./2, -1./2},
			Segment{0, -1},
			Segment{1./4, -1./2},
			Segment{1./4, 1},
			DIVIDER,
			Segment{-1./2, 1},
			Segment{1./2, 1},
		}, nil
	case 'K':
		return []Segment{
			Segment{-1./2, 1},
			Segment{-1./2, -1},
			DIVIDER,
			Segment{1./2, 1},
			Segment{-1./2, 0},
			Segment{1./2, -1},
		}, nil
	case 'L':
		return []Segment{
			Segment{-1./2, 1},
			Segment{-1./2, -1},
			Segment{1./3, -1},
		}, nil
	case 'M':
		return []Segment{
			Segment{-1./2, -1},
			Segment{-1./2, 1},
			Segment{0, 1./3},
			Segment{1./2, 1},
			Segment{1./2, -1},
		}, nil
	case 'N':
		return []Segment{
			Segment{-1./2, -1},
			Segment{-1./2, 1},
			Segment{1./2, -1},
			Segment{1./2, 1},
		}, nil
	case 'O':
		return []Segment{
			Segment{-1./3, 1},
			Segment{1./3, 1},
			Segment{1./2, 1./2},
			Segment{1./2, -1./2},
			Segment{1./3, -1},
			Segment{-1./3, -1},
			Segment{-1./2, -1./2},
			Segment{-1./2, 1./2},
			Segment{-1./3, 1},
		}, nil
	case 'P':
		return []Segment{
			Segment{-1./2, -1},
			Segment{-1./2, 1},
			Segment{1./3, 1},
			Segment{1./2, 1./2},
			Segment{1./3, 0},
			Segment{-1./2, 0},
		}, nil
	case 'Q':
		return []Segment{
			Segment{-1./3, -1},
			Segment{-1./2, -1./2},
			Segment{-1./2, 1./2},
			Segment{-1./3, 1},
			Segment{-1./3, 1},
			Segment{1./3, 1},
			Segment{1./2, 1./2},
			Segment{1./2, -1./2},
			Segment{1./3, -1},
			Segment{1./2, -1},
			Segment{-1./3, -1},
		}, nil
	case 'R':
		return []Segment{
			Segment{-1./2, -1},
			Segment{-1./2, 1},
			Segment{1./3, 1},
			Segment{1./2, 1./2},
			Segment{1./3, 0},
			Segment{-1./2, 0},
			Segment{1./2, -1},
		}, nil
	case 'S':
		return []Segment{
			Segment{1./2, 2./3},
			Segment{0, 1},
			Segment{-1./2, 2./3},
			Segment{1./2, -2./3},
			Segment{0, -1},
			Segment{-1./2, -2./3},
		}, nil
	case 'T':
		return []Segment{
			Segment{-1./2, 1},
			Segment{1./2, 1},
			DIVIDER,
			Segment{0, 1},
			Segment{0, -1},
		}, nil
	case 'U':
		return []Segment{
			Segment{-1./2, 1},
			Segment{-1./2, -2./3},
			Segment{0, -1},
			Segment{1./2, -2./3},
			Segment{1./2, 1},
		}, nil
	case 'V':
		return []Segment{
			Segment{-1./2, 1},
			Segment{0, -1},
			Segment{1./2, 1},
		}, nil
	case 'W':
		return []Segment{
			Segment{-1./2, 1},
			Segment{-1./4, -1},
			Segment{0, 0},
			Segment{1./4, -1},
			Segment{1./2, 1},
		}, nil
	case 'X':
		return []Segment{
			Segment{-1./2, 1},
			Segment{1./2, -1},
			DIVIDER,
			Segment{-1./2, -1},
			Segment{1./2, 1},
		}, nil
	case 'Y':
		return []Segment{
			Segment{0, -1},
			Segment{0, 0},
			DIVIDER,
			Segment{-1./2, 1},
			Segment{0, 0},
			Segment{1./2, 1},
		}, nil
	case 'Z':
		return []Segment{
			Segment{-1./2, 1},
			Segment{1./2, 1},
			Segment{-1./2, -1},
			Segment{1./2, -1},
		}, nil
	case '!':
		return []Segment{
			Segment{0, 1},
			Segment{0, -1./3},
			DIVIDER,
			Segment{0, -1},
		}, nil
	case ' ':
		return nil, nil
	default:
		return nil, fmt.Errorf("No data for character '%c'", c)
	}
}
