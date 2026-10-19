func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums{
		count[num]++
	}
	var numList []int
	for num := range count{
		numList = append(numList, num)
	}
	sort.Slice(numList, func(i, j int)bool{
		return (count[numList[i]] > count[numList[j]])
	})
	return numList[0:k]
}
