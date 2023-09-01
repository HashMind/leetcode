package main

import "fmt"

//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//返回容器可以储存的最大水量。
//说明：你不能倾斜容器。

type Container struct {
	leftIndex  int //左边的坐标
	rightIndex int //右侧的坐标
	leftValue  int //左侧的值
	rightValue int //右侧的值
	capacity   int //容量
}

//最笨的办法,将这个二维数组变为一位数组,本质上是求笛卡尔积中的最大值
func containerWithMostWater(arrays []int) {
	cartesian := make([]Container, 0)
	var maxPoint *Container
	for i := 0; i <= len(arrays)-1; i++ {
		var point Container
		point.leftIndex = i
		point.leftValue = arrays[i]
		for j := i + 1; j <= len(arrays)-1; j++ {
			point.rightIndex = j
			point.rightValue = arrays[j]
			if point.leftValue < point.rightValue {
				point.capacity = point.leftValue * (point.rightIndex - point.leftIndex)
			} else {
				point.capacity = point.rightValue * (point.rightIndex - point.leftIndex)
			}
			//point.capacity = point.leftValue * point.rightValue
			cartesian = append(cartesian, point)
			if maxPoint == nil {
				maxPoint = &point
			} else {
				if point.capacity > maxPoint.capacity {
					maxPoint = &point
				}
			}
		}
	}

	fmt.Println(len(cartesian))
	fmt.Println(maxPoint)

}
func main11() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	containerWithMostWater(arr)
}

//1, 8, 6, 2, 5, 4, 8, 3, 7
//1, 3, 8, 2,2,2,2,2,2,2, 5, 4, 8, 3, 9
