type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result string 
	for _, str := range strs{
		encoded := strconv.Itoa(len(str)) + "#" + str
		result = result + encoded
	}
	return result
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	for encoded != "" {
	index := strings.Index(encoded, "#")
	length, _ := strconv.Atoi(encoded[:index])
	word := encoded[index+1 : index+1+length]
	result = append(result, word)
	encoded = encoded[index+1+length:]
}
return result
}