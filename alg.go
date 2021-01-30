package cli

const notFound = -1

func find(slice []string, what string) int {
	for i, s := range slice {
		if what == s {
			return i
		}
	}
	return -1
}
