package gordle

// hint describes the validity of a char in a word.
type hint byte

const (
	absentChar hint = iota
	wrongPosition
	correctPosition
)

func (h hint) String() string {
	switch h {
		case absentChar:
			return "Ⅹ"
		case wrongPosition: 
			return "🟡"
		case correctPosition:
			return "✅"
		default:
			return "⚠️"
		}
	}