func isAnagram(s string, t string)bool{
	if len(s) != len(t){
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s{
		count[ch]++
	}
	for _, ch1 := range t{
		count[ch1]--
	}
	for _, value := range count{
		if value != 0{
			return false
		}
	}
	return true
}