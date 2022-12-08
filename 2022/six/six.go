package six

func Run(input []byte) int {
	return findMarker(input, 14)
}

func findMarker(s []byte, length int) int {

	buffer := make([]rune, length)

	for i, c := range s {
		buffer[i%length] = rune(c)
		if i > 3 && allElementsDifferent(buffer) {
			return i + 1
		}
	}

	return 0
}

func allElementsDifferent(i []rune) bool {

	lookup := map[rune]bool{}

	for _, r := range i {
		if _, ok := lookup[r]; ok {
			return false
		}
		lookup[r] = true
	}
	return true
}

//1W96
