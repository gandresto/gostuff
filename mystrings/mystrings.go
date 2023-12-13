package mystrings

// Reverse reverses a string left to right
// We need to capitalize the first letter of the function.
// If we don't then we won't be able to access
// this function outside of the mystrings packages
func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
