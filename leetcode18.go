package main

/***
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。
请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]]
（若两个四元组元素一一对应，则认为两个四元组重复）：
0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

示例 1：
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

示例 2：
输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

提示：
1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109

**/

type Sum4Element struct {
	numA int
	numB int
	numC int
	numD int
}

//步骤一:先对array进行排序,后续形成的四元组就是有序的,从而避免了结果集里面的排序
//步骤二:在不讲究效率的情况下就是进行4次for循环遍历array里的元素,尝试所有的四次组合看看谁更适合,后续需要对四元组进行排序或者比较
//优化步骤二:a,b,c,d为target的四元素下标,a最小,d最大,同理:array[a]<=array[b]<=array[c]<=array[d]
//先选择一对指针a和d;a是array的最左边界,d为array的最右边界,计算subAD=target-(array[a]+array[d]),
//然后在array[a+1]:array[d-1]去匹配余下的2个元素,直到b==c,注意此处也使用2个指针,
//如果subAD<0,且array[c]+array[d]+subAD<0,则d--;如果subAD<0,且array[c]+array[d]+subAD>0,则c++
//如果subAD>0,且array[c]+array[d]+subAD<0,则c++;如果subAD<0,且array[c]+array[d]+subAD>0,则d--
//按照这种情形
func getTargetSumElement(array []int, target int) {
	//步骤一:先对array进行排序,后续形成的四元组就是有序的,从而避免了结果集里面的排序
	QuickSort_InPlace(array, 0, len(array), 0)
	if len(array) < 4 {
		return
	}
	elements := make([]Sum4Element, 0)
	//1,2,3,4,5,6,7,8,9
	var aIndex int = 0
	for len(array)-(aIndex+1) >= 4 {
		var dIndex int = len(array) - 1
		for dIndex-aIndex >= 2 {
			var bIndex int = aIndex + 1
			var cIndex int = dIndex - 1
			sumAD := array[aIndex] + array[dIndex]
			extraAD := target - sumAD
			for cIndex-bIndex >= 1 {
				sumBC := array[bIndex] + array[cIndex]
				if sumBC > extraAD {
					//由于array已经经过排序,说明可以换一个小一点的c试试
					cIndex--
				} else if sumBC < extraAD {
					//由于array已经经过排序,说明可以换一个大一点的b试试
					bIndex++
				} else {
					//answer,并且ad确定后,bc的答案可能不止一个
					sum4Element := Sum4Element{numA: array[aIndex], numB: array[bIndex], numC: array[cIndex], numD: array[dIndex]}
					elements = append(elements, sum4Element)
					bIndex++ //这里还是沿用上一轮循环的cIndex,因为,b可能和和此刻d后面的元素组成可能
				}
			}
			dIndex--
		}
		aIndex++ //当a和有可能的数字匹配一遍以后a就接着后走,直到array里最后四个元素
	}
	println(elements)

}
func main18() {
	nums := []int{1, 0, -1, 0, -2, 2}
	//-2,-1,0,0,1,2
	getTargetSumElement(nums, 0)

}
