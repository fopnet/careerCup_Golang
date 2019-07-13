package arrays

func IsIntArrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		for j, _ := range b {
			if a[i] != b[j] {
				return false
			}
		}
	}

	return true

}
