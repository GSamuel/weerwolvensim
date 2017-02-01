package character

type Type int

const (
	WEERWOLF Type = iota
	CUPIDO
	GENEZER
	HEKS
	BURGER
)

func (this Type) String() string {
	switch this {
	case WEERWOLF:
		return "weerwolf"
	case HEKS:
		return "heks"
	case CUPIDO:
		return "cupido"
	case GENEZER:
		return "genezer"
	case BURGER:
		return "burger"
	default:
		return "unknown"
	}
}

func New(card Type) Character {
	switch card {
	case WEERWOLF:
		return NewWeerwolf()
	case CUPIDO:
		return NewCupido()
	case GENEZER:
		return NewGenezer()
	case HEKS:
		return NewHeks()
	case BURGER:
		return NewBurger()
	default:
		panic("help!")
		return nil
	}
}

func NewPool(types []Type) *Pool {

	characters := []Character{}

	for _, t := range types {
		characters = append(characters, New(t))
	}

	return &Pool{characters: characters}
}
