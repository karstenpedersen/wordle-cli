package utils

func Contains[T comparable](slice []T, element T) bool {
	for _, r := range slice {
		if r == element {
			return true
		}
	}
	return false
}
