package two_sum

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println(twoSum([]int {1, 10, 11, 12, 13, 2}, 3))
}

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for idx, val := range nums {
		complemento := target - val
		
		if mapHasKey(hashmap, val) {
			return []int{hashmap[val], idx}
		}
		hashmap[complemento] = idx
	}
	return []int {}
}

func mapHasKey(hashMap map[int]int, key int) bool {
	keys := make([]int, 0, len(hashMap))
	for k := range hashMap {
		keys = append(keys, k)
	}
	_, error := indexOf(keys, key)
	return error == nil
}

func indexOf(keys []int, element int) (int, error) {
	for k, v := range keys {
		if element == v {
			return k, nil
		}
	}
	return 0, errors.New("chave não encontrada")
}