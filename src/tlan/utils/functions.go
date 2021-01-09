package utils

func Filter(arr []string, cond func(string) bool) []string {
	var result []string
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

func Find(arr []string, cond func(string) bool) bool {
	for i := range arr {
		if cond(arr[i]) {
			return true
		}
	}
	return false
}
