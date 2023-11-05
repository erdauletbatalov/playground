package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}

func longestOnes(nums []int, k int) int {
	var res, max int
	var tempArr []int

	var nulls int

	for i := range nums {
		if nums[i] == 1 {
			res++
		} else {
			nulls++
		}
		if nums[i] != 1 || i == len(nums)-1 {
			if max < res {
				max = res
			}
			res = 0
		}
	}
	if k == 0 {
		return max
	}
	if nulls < k {
		return len(nums)
	}
	res = 0
	max = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			k--
			nums[i] = 1
			tempArr = append(tempArr, i)
			if k == 0 {
				for j := tempArr[0]; j >= 0; j-- {
					if nums[j] == 0 {
						break
					}
					res++
				}
				for j := tempArr[0]; j < len(nums); j++ {
					if nums[j] == 0 {
						break
					}
					res++
				}
				res = res - 1
				if max < res {
					max = res
				}
				res = 0
				nums[tempArr[0]] = 0
				tempArr = tempArr[1:]
				k++

			}
		}
	}
	return max
}
