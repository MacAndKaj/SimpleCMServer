package utils

// Finds element in array.
// Return: index of found element, -1 if not found.
func Find(a []string, elem string) int {
	for i, val := range a {
		if val == elem {
			return i
		}
	}
	return -1
}
