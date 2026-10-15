//import "slices"
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range strs {
		letters := []byte(word)
		sort.Slice(letters, func(i, j int) bool {
    return letters[i] < letters[j]
})
		key := string(letters)
		groups[key] = append(groups[key], word)
	}
	var result [][]string
		for _, group := range groups{
			result = append(result, group)
		}
		return result
}
