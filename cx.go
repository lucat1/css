package css

const (
	runes = "abcdefghijklmnopqrstuvwxyz"
	l = len(runes)
)

// prefix to be appended before each generated ID
var prefix = ""

// Next returns the next unique string (the shortest possible)
// we start from a single letter and return the whole alphabet,
// until we reach the and, so we append a second letter and so on
func Next(count int) string {
	res := gen(count)
	
	count++
	return prefix + res
}

func gen(n int) string {
	i := n / l // current position / number of letters
	r := n % l // rest of the prev operation
	res := ""

	if i >= 1 {
		res += gen(i)
	}

	res += string(runes[r])
	return res
}