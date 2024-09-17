package two_sum

// import (
// 	"fmt"
// 	"sort"
// )

// /* solução ainda não funciona para slices com números repetidos que fazem a soma acontecer */
// func init() {
	// fmt.Println(twoSum([]int {3, 3}, 6))
// }

// func twoSum(nums []int, target int) []int {
// 	numsCopy := copySlice(nums)
// 	sortSlice(nums)
// 	sumNums := findSumNums(nums, target)

// 	return []int {indexOf(sumNums[0], numsCopy), indexOf(sumNums[1], numsCopy)}
// }

// func copySlice(nums []int) []int {
// 	copyOfNums := make([]int, len(nums))
// 	copy(copyOfNums, nums)
// 	return copyOfNums
// }

// func sortSlice(nums []int) {
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})
// }

// func findSumNums(nums []int, target int) []int {
// 	idxStart := 0
// 	idxEnd := len(nums) - 1

// 	for range nums {
// 		if(nums[idxStart] + nums[idxEnd]) == target {
// 			return []int {nums[idxStart], nums[idxEnd]}
// 		} else if(nums[idxStart] + nums[idxEnd]) > target {
// 			idxEnd--
// 		} else if(nums[idxStart] + nums[idxEnd]) < target {
// 			idxStart++
// 		}
// 	}
// 	return []int {}
// }

// func indexOf(element int, nums []int) int {
// 	for k, v := range nums {
// 		if element == v {
// 			return k
// 		}
// 	}
// 	return -1
// }
