package helpers

func AbsInt(first, second int) int {
	if first-second < 0 {
		return second - first
	}
	return first - second
}
