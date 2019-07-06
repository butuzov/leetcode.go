package accumulate

// Accumulate - return NEW slice with mapped data.
func Accumulate(in []string, f func(string) string) []string {
	// Simple slice mapping

	// initialize new collection
	var out = make([]string, len(in))
	for idx := range in {
		out[idx] = f(in[idx])
	}
	return out
}
