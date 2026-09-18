import (
	"maps"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	anagramArrays := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		sortedStr := sortString(strs[i])
		anagramArrays[sortedStr] = append(anagramArrays[sortedStr], strs[i])
	}

	return slices.Collect(maps.Values(anagramArrays))
}

func sortString(str string) string {
	bytes := []byte(str)
	slices.Sort(bytes)

	return string(bytes)
}
