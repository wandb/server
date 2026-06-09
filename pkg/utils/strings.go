package utils

func RemoveDuplicates(strings []string) []string {
	// Create a map to track seen strings.
	seen := make(map[string]struct{})
	var result []string

	// Iterate over the input slice.
	for _, str := range strings {
		// If we haven't seen the string before, add it to the result slice.
		if _, ok := seen[str]; !ok {
			result = append(result, str)
			seen[str] = struct{}{} // Mark string as seen.
		}
	}

	return result
}
