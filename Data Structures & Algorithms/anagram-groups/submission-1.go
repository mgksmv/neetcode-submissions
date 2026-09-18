import (
	"maps"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	anagramArrays := make(map[[26]byte][]string)

	for i := 0; i < len(strs); i++ {
		freqs := [26]byte{}
		for _, l := range strs[i] {
			freqs[l - 'a']++
		}
		anagramArrays[freqs] = append(anagramArrays[freqs], strs[i])
	}

	return slices.Collect(maps.Values(anagramArrays))
}
