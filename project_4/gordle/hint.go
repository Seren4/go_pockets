package gordle

import "strings"

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

// feedback is a list of hints, one per character  of the word.
type feedback []hint

// String implements the Stringer interface for a slice of hints.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

// Equal determines equality of two feedbacks
func (fb feedback) Equal (other feedback) bool {
	if len(fb) != len(other) {
		return false
	}
	for i, v := range(fb) {
		if v != other[i] {
			return false
		}
	}
	return true
}