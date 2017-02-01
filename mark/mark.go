package mark

import ()

type Mark int

const (
	Attacked Mark = iota
	Lover
	Protected
	Revived
	Poisoned
	Lynched
)

func (this Mark) String() string {
	switch this {
	case Attacked:
		return "attacked"
	case Lover:
		return "lover"
	case Protected:
		return "protected"
	case Revived:
		return "revived"
	case Poisoned:
		return "poisoned"
	case Lynched:
		return "lynched"
	default:
		return "unknown"
	}
}

type Character int

const (
	Unassigned Character = iota
	Weerwolf
	Heks
	Cupido
	Genezer
	Burger
)

func (this Character) String() string {
	switch this {
	case Unassigned:
		return "unassigned"
	case Weerwolf:
		return "weerwolf"
	case Heks:
		return "heks"
	case Cupido:
		return "cupido"
	case Genezer:
		return "genezer"
	case Burger:
		return "burger"
	default:
		return "unknown"
	}
}
