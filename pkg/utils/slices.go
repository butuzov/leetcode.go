package utils

// Immutable Slice of strings
func CopySliceString(m []string) []string {
	var slice = make([]string, len(m))
	copy(slice, m)
	return slice
}

// Immutable Slice of ints
func CopySliceInt(m []int) []int {
	var slice = make([]int, len(m))
	copy(slice, m)
	return slice
}

// Immutable Slice of bytes
func CopySliceByte(m []byte) []byte {
	var slice = make([]byte, len(m))
	copy(slice, m)
	return slice
}
