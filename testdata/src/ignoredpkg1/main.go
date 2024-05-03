// Package ignoredpkg1 tests NilAway's ability to ignore packages that are configured to be ignored.
package ignoredpkg1

var GlobalVar *int
var GlobalVarInt = 0

func main() {
	// Directly de-referencing a nil pointer, but it is OK since this package is ignored.
	print(*GlobalVar)
}

func RetZero() int {
	return 0
}
