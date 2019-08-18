package parser

var bracketsPair = map[rune]rune{LPAREN: RPAREN}

// CheckBrackets returns true if inbound expression has balanced brackets
// Brackets pair calc
func CheckBrackets(expr string) (bool, error) {
	var queue []rune

	for _, ch := range expr {
		if ch == LPAREN {
			queue = append(queue, bracketsPair[ch])
		} else if ch == RPAREN {
			if 0 < len(queue) && queue[len(queue)-1] == ch {
				queue = queue[:len(queue)-1]
			} else {
				return false, nil
			}
		}
	}
	return len(queue) == 0, nil
}
