func hasDuplicate(nums []int) bool {
    // count := make(map[int]int)

	// for _, value := range nums{
	// 	count[value]++
	// }

	// for _, value := range count{
	// 	if value > 1 {
	// 		return true
	// 	}
	// }
	// return false
	count := make(map[int]bool)

	for _,  value := range nums{
		if count[value]{
			return true
		}
		count[value] = true
	}
	return false
}
