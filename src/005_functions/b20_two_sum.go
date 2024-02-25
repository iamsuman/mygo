package main

import "fmt"

func main() {
	nums := []int{2, 7, 7, 7, 9, 11, 15}
	target := 16

	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	res := map[int]int{
		target - nums[0]: 0,
	}
	r := make([]int, 2)

	for i := 1; i < len(nums); i++ {
		v, ok := res[nums[i]]
		if ok {
			// r = append(r, res[nums[i]], i)
			// r = append(r, v, i)
			r[0] = v
			r[1] = i
			break
		}
		x := target - nums[i]
		res[x] = i

	}
	return r
}
