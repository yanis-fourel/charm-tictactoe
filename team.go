package main

type Team rune

const (
	Team_None Team = '.'
	Team_O    Team = 'O'
	Team_X    Team = 'X'
)

func (t Team) toRune() rune {
	return rune(t)
}

func (t Team) Name() string {
	switch t {
	case Team_None:
		return "No one"
	case Team_O:
		return "Circle"
	case Team_X:
		return "Cross"
	}
	return "<unknown>"
}
