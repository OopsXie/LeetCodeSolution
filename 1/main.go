package main

// O(n^2)
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				result := []int{
// 					i, j,
// 				}
// 				return result
// 			}
// 		}
// 	}
// 	return nil
// }

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))

	result := []int{
		-1, -1,
	}
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		result[0] = i
		v, ok := indexMap[temp]

		if ok {
			if v != result[0] {
				result[1] = v
				return result
			}
		}
		indexMap[nums[i]] = i
	}
	return nil
}

func main() {
	array := []int{
		0, 1, 2, 3,
	}

	target := 6

	twoSum(array, target)
}
