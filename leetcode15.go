package main

import "fmt"

/**
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
*/
type element struct {
	num1 int
	num2 int
	num3 int
}

//从给定的数组里先筛选出所有的三数组合,然后再过滤和等于0的组合
//先对所有元素进行排序,然后对于后续的element判重会轻松很多
//时间复杂度:o(n^3)+
func threeElementSumEqualZero01(array []int) {
	matches := make([]element, 0)
	for i := 0; i < len(array)-3; i++ {
		for j := i + 1; j < len(array)-2; j++ {
			for k := i + 2; k < len(array)-1; k++ {
				if array[i]+array[j]+array[k] == 0 {
					var e element
					e.num1 = array[i]
					e.num2 = array[j]
					e.num3 = array[k]
					//这里再去里面查重一下
					matches = append(matches, e)
				}
			}
		}
	}
	fmt.Println(matches)
}

//先将接收到的参数进行排序,
func threeElementSumEqualZero02(array []int) {
	matches := make([]element, 0)
	for i := 0; i < len(array)-3; i++ {
		for j := i + 1; j < len(array)-2; j++ {
			for k := i + 2; k < len(array)-1; k++ {
				if array[i]+array[j]+array[k] == 0 {
					var e element
					e.num1 = array[i]
					e.num2 = array[j]
					e.num3 = array[k]
					//这里再去里面查重一下
					matches = append(matches, e)
				}
			}
		}
	}
	fmt.Println(matches)
}

func main15() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeElementSumEqualZero01(nums)

}
