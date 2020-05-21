package str

import "fmt"

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if string(haystack[i:i+len(needle)]) == needle {
			fmt.Println(string(haystack[i : i+len(needle)]))
			return i
		}
	}
	return -1
}
