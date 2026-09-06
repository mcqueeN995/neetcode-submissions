import "slices"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := []rune(s)
	    slices.Sort(key)
		groups[string(key)] = append(groups[string(key)], s)
	}

	ans := make([][]string, 0, len(groups))
	for _, arr := range groups{
		ans = append(ans, arr)
	}
	return ans
}
