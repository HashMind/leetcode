package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type Sex string

const (
	Male   Sex = "male"
	Female     = "female"
)

type student struct {
	name *string
	age  *int
}

func getStudent(s student) {
	*s.name = "hashmind"
	fmt.Printf("%p\n", &s)
	fmt.Println(s)
}

var letter = []rune("abcdefghijklmnopqrstuvwxyz")

func RandomLowercaseStr(length int) string {
	bytes := make([]rune, length)
	for k, _ := range bytes {
		bytes[k] = letter[rand.Intn(len(letter))]
	}
	return string(bytes)
}
func main4() {
	str := RandomLowercaseStr(7)
	fmt.Println(str)
	fmt.Println("------------")

	data := []int{9, 3, 2, 9, 1, 4, 5, 8, 7, 9}
	//InsertSort(data)
	//ShellInsertSort(data, 3)
	//shell_sort2(data, 3)
	//HeapSort(data)
	data2 := []int{19, 32, 200, 9, 1001, 4001, 500, 80, 79, 93}
	sortedData := MergeSort(data2)
	//RadixSort(data2)
	fmt.Println(data2)
	fmt.Println(sortedData)
	//select_sort(data)
	//BubbleSort(data)
	//QuickSort_InPlace(data, 0, len(data), 0)
	fmt.Println(data)
	//fmt.Println(data)
	//fmt.Println("---------")
	//BubbleSort(data[:])
	//fmt.Println("---------")
	//fmt.Println(data)
}
func BubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		var temp int = 0
		for j := i; j < len(data); j++ {
			if data[i] > data[j] {
				temp = data[i]
				data[i] = data[j]
				data[j] = temp
				fmt.Println(data)
			}
		}
	}
	return data
}

// 快速排序:冒泡排序的升级版,在原数组的基础上进行排序,在满足快速排序的同时,也尽可能的占用少的空间
func QuickSort_InPlace(data []int, leftIndex, rightIndex, compareIndex int) {
	if len(data) < 2 {
		return
	}
	basic := data[compareIndex]
	leftMatchIndex := leftIndex
	rightMatchIndex := rightIndex
	for i := rightIndex - 1; i >= 1; i-- {
		if i == leftMatchIndex {
			if len(data) == 2 {
				if data[0] > data[1] {
					data[0] = data[0] ^ data[leftIndex]
					data[leftIndex] = data[0] ^ data[leftIndex]
					data[0] = data[0] ^ data[leftIndex]
					return
				}
				return
			}
			data[0] = data[0] ^ data[leftIndex]
			data[leftIndex] = data[0] ^ data[leftIndex]
			data[0] = data[0] ^ data[leftIndex]
			leftData := data[:leftIndex]
			rightData := data[leftIndex+1:]
			if len(leftData) > 1 {
				QuickSort_InPlace(leftData, 1, len(leftData), 0) //左侧接着递归
			}
			if len(rightData) > 1 {
				QuickSort_InPlace(rightData, 1, len(rightData), 0) //右侧接着递归
			}
			break
		}
		if data[i] <= basic {
			rightMatchIndex = i
			break
		}
	}
	for i := leftIndex + 1; i <= len(data)-1; i++ {
		if i == rightMatchIndex {
			break
		} //但是右侧可能还有比目标数大的数存在
		if data[i] >= basic {
			leftMatchIndex = i
			break
		}
	}
	if rightIndex == rightMatchIndex { //说明没有比基数小的数,例如:15324
		if leftIndex == leftMatchIndex {
			return //说明是一个相等的数组:5,5,5,5,5
		} else { //例如:1,4,2,3,5
			rightData := data[1:]
			if len(rightData) > 1 {
				QuickSort_InPlace(rightData, 0, len(rightData), 0) //左侧接着递归
			}
		}
	} else {
		if leftIndex == leftMatchIndex { //说明在左指针在碰撞到右指针前依然没有找到合适的数,例如:5432|5324798
			data[0] = data[0] ^ data[rightMatchIndex]
			data[rightMatchIndex] = data[0] ^ data[rightMatchIndex]
			data[0] = data[0] ^ data[rightMatchIndex]

			leftData := data[:rightMatchIndex]
			if len(leftData) > 1 {
				QuickSort_InPlace(leftData, leftMatchIndex, len(leftData), 0) //左侧接着递归
			}
			if len(data)-1 > rightMatchIndex {
				rightData := data[rightMatchIndex+1:]
				if len(rightData) > 1 {
					QuickSort_InPlace(rightData, 0, len(rightData), 0) //左侧接着递归
				}
			}
		} else {
			//例如:3,4,6,9,8,1,5,2,7
			data[leftMatchIndex] = data[leftMatchIndex] ^ data[rightMatchIndex]
			data[rightMatchIndex] = data[leftMatchIndex] ^ data[rightMatchIndex]
			data[leftMatchIndex] = data[leftMatchIndex] ^ data[rightMatchIndex]
			QuickSort_InPlace(data, leftMatchIndex, rightMatchIndex, 0) //左侧接着递归
		}
	}
}

// 插入排序
// 从数组的第一个元素开始,我们就认为它是有序的,然后剩下的元素依次插入到这个序列的合适的位置
// 就像我们拿到一副牌只占用一个额外的空间,然后就能实现把牌给排好序
// 与冒泡排序的区别在于,每一次循环结束,目标元素左侧序列一定有序,单该元素可能没有处于最终应该出现的位置
// 冒泡排序一轮内循环结束后,外循环索引位置的元素就是目标元素
func InsertSort(data []int) {
	length := len(data)
	for i := 0; i <= length-1; i++ {
		temp := data[i]
		if i == 0 {
			continue
		}
		j := i - 1
		for ; j >= 0; j-- {
			if data[j] > temp {
				data[j+1] = data[j]
			} else {
				break
			}
		}
		data[j+1] = temp
	}
}

// 希尔插入排序
// 将数据先按照一定的步长例如:5,然后a[0]&a[6]&a[11]&a[16]...类似这样的去比较
// 相当于将原属组虚拟为:一个有5列的元素,然后每一列的构成一个新数组再进行比较,一轮循环以后,步长再缩小一半(还有更优秀的步长选择算法)
func ShellInsertSort(data []int, step int) {
	for ; step >= 1; step = step >> 1 {
		for columIndex := 0; columIndex < step; columIndex++ {
			for index := columIndex; index <= len(data)-1; index += step {
				temp := data[index]
				compareIndex := index
				for ; compareIndex-step >= 0; compareIndex -= step {
					if data[compareIndex-step] > temp {
						data[compareIndex] = data[compareIndex-step]
					} else {
						break
					}
				}
				data[compareIndex] = temp
			}
		}
	}
}
func shell_sort2(arr []int, len int) {
	var gap, i, j int
	var temp int
	for gap = len >> 1; gap > 0; gap >>= 1 {
		for i = gap; i < len; i++ {
			temp = arr[i]
			for j = i - gap; j >= 0 && arr[j] > temp; j -= gap {
				arr[j+gap] = arr[j]
			}
			arr[j+gap] = temp
		}
	}
}

// 选择排序（Selection sort）是一种简单直观的排序算法。
// 它的工作原理如下。首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
// 然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
// 以此类推，直到所有元素均排序完毕。
// a,[b,c,d,e,f,g],h
func select_sort(data []int) {
	leftIndex := 0
	rightIndex := len(data) - 1
	for leftIndex <= rightIndex { //相遇的时候说明就结束了
		tMiniIndex := leftIndex
		tMaxIndex := rightIndex
		eMaxIndex := rightIndex
		eMiniIndex := leftIndex
		for i := leftIndex; i <= rightIndex; i++ {
			if data[i] > data[eMaxIndex] {
				eMaxIndex = i
			} else if data[i] < data[eMiniIndex] {
				eMiniIndex = i
			}
		}
		rightIndex--
		leftIndex++
		if tMiniIndex != eMiniIndex {
			temp := data[tMiniIndex]
			data[tMiniIndex] = data[eMiniIndex]
			data[eMiniIndex] = temp
			if eMaxIndex == tMiniIndex {
				if eMiniIndex == tMaxIndex {
					continue
				} else if eMaxIndex == tMiniIndex {
					eMaxIndex = eMiniIndex
				}
			}
		}
		if tMaxIndex != eMaxIndex {
			temp := data[tMaxIndex]
			data[tMaxIndex] = data[eMaxIndex]
			data[eMaxIndex] = temp
		}
		fmt.Println(data)
	}
}

func getNextLeftIndex(nodeIndex int, chaosData []int) *int {
	var i *int
	leftLeafIndex := (nodeIndex+1)*2 - 1
	if leftLeafIndex > len(chaosData)-1 {
		return i
	} else {
		i = new(int)
	}
	*i = leftLeafIndex
	return i
}
func getNextRightIndex(nodeIndex int, chaosData []int) *int {
	var i *int
	rightLeafIndex := (nodeIndex + 1) * 2
	if rightLeafIndex > len(chaosData)-1 {
		return i
	} else {
		i = new(int)
	}
	*i = rightLeafIndex
	return i
}
func getParentIndex(nodeIndex int, chaosData []int) *int {
	//如果根节点的坐标是0开始的,则所有的左侧
	var i *int
	if nodeIndex == 0 {
		return nil
	} else {
		i = new(int)
	}
	mod := math.Mod(float64(nodeIndex), 2)
	if mod == 0 {
		*i = (nodeIndex / 2) - 1
	} else {
		*i = (nodeIndex - 1) / 2
	}
	return i
}
func exchangeRoot(chaosData []int) {
	if len(chaosData) > 1 {
		temp := chaosData[0]
		for i := 0; i < len(chaosData)-1; i++ {
			chaosData[i] = chaosData[i+1]
		}
		chaosData[len(chaosData)-1] = temp
	}
}

//direction:true忽略左孩子
func makeBigRootTree(nodeIndex int, chaosData []int, ignoreLeft, ignoreRight bool) {

	nextLeftIndex := getNextLeftIndex(nodeIndex, chaosData)
	if nextLeftIndex == nil {
		return
	} //如果左孩子为空,则右孩子也一定为空
	//说明nodeIndex已经和左侧的比较过了
	if !ignoreLeft {
		makeBigRootTree(*nextLeftIndex, chaosData, false, false)
	}
	if chaosData[*nextLeftIndex] > chaosData[nodeIndex] {
		temp := chaosData[*nextLeftIndex]
		chaosData[*nextLeftIndex] = chaosData[nodeIndex]
		chaosData[nodeIndex] = temp
		fmt.Println(chaosData)
	}

	nextRightIndex := getNextRightIndex(nodeIndex, chaosData)
	if nextRightIndex == nil {
		return
	}
	if !ignoreRight {
		makeBigRootTree(*nextRightIndex, chaosData, false, false)
	}
	if chaosData[*nextRightIndex] > chaosData[nodeIndex] {
		temp := chaosData[*nextRightIndex]
		chaosData[*nextRightIndex] = chaosData[nodeIndex]
		chaosData[nodeIndex] = temp
		fmt.Println(chaosData)
	}
}
func iteratorHeap(trunkNodeIndex int, chaosData []int, ignoreLeft, ignoreRight bool) {
	makeBigRootTree(trunkNodeIndex, chaosData, ignoreLeft, ignoreRight)
	parentTrunkNodeIndex := getParentIndex(trunkNodeIndex, chaosData)
	//根据trunkNodeIndex的奇偶来决定要忽略左右
	if parentTrunkNodeIndex == nil {
		exchangeRoot(chaosData)
		HeapSort(chaosData[:len(chaosData)-1])
		return
	}
	mod := math.Mod(float64(trunkNodeIndex), 2)
	if mod == 0 {
		ignoreRight = true
		if *parentTrunkNodeIndex == 0 {
			ignoreLeft = false
		}
	} else {
		ignoreLeft = true
		if *parentTrunkNodeIndex == 0 {
			ignoreRight = false
		}
	}

	iteratorHeap(*parentTrunkNodeIndex, chaosData, ignoreLeft, ignoreRight)
}

// 用数组存储二叉堆
// 从第一层开始存储,然后是第二层,第三层,第N层
// 每一层的元素排列都选择固定顺序排序,例如从左往右
func HeapSort(data []int) {
	chaosData := data
	//从这颗二叉堆的最左侧待排序开始,待排序数组求幂,然后向下取整,就能得到这个二叉树的倒数一层的深度L,
	//然后len(data)+1,就是这个二叉树的左子树的最后一层树干中的最右侧树干节点元素
	latestIndex := len(chaosData) - 1
	firstLeftIndex := getParentIndex(latestIndex, chaosData)
	if firstLeftIndex != nil {
		iteratorHeap(*firstLeftIndex, chaosData, false, false)

	}
}

// 桶排序的工作的原理:将数组分到有限数量的桶里。每个桶再个别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排序）。桶排序是鸽巢排序的一种归纳结果。
// 设置一个定量的数组当作空桶子。
// 寻访序列，并且把项目一个一个放到对应的桶子去。
// 对每个不是空的桶子进行排序。
// 从不是空的桶子里把项目再放回原来的序列中。
func BucketSort(data []int) {
	//最简单的思路是构建一个二维数组,列是有序的,
	//但是这样可能会造成内存的浪费,尤其是当一个要排序的数组里拥有很多相同的元素的时候
	//为了节省内存,可以考虑开发一个有序的二叉连表来实现这个需求,map还需要对key进行排序
	bucket := make(map[int]int, 0)
	keySet := make([]int, 0)
	for i := 0; i < len(data)-1; i++ {
		count, ok := bucket[data[i]]
		if ok {
			bucket[data[i]] = count + 1
		} else {
			bucket[data[i]] = 1
			keySet = append(keySet, data[i])
		}
	}
	//再对桶进行排序,下面再次对keySet进行排序即可,可以选择其他算法都可以,本次选择快速排序
	QuickSort_InPlace(keySet, 0, len(keySet), 0)
	//现在keyset就是有序的了

}

// 桶是一种区间桶,例如:大于10的小于20的装在同一个桶里,
func BucketSort2(data []int) {
	//最简单的思路是构建一个二维数组,列是有序的,
	//但是这样可能会造成内存的浪费,尤其是当一个要排序的数组里拥有很多相同的元素的时候
	//为了节省内存,可以考虑开发一个有序的二叉连表来实现这个需求,map还需要对key进行排序
	bucket := make(map[int]int, 0)
	keySet := make([]int, 0)
	for i := 0; i < len(data)-1; i++ {
		count, ok := bucket[data[i]]
		if ok {
			bucket[data[i]] = count + 1
		} else {
			bucket[data[i]] = 1
			keySet = append(keySet, data[i])
		}
	}
	//再对桶进行排序,下面再次对keySet进行排序即可,可以选择其他算法都可以,本次选择快速排序
	QuickSort_InPlace(keySet, 0, len(keySet), 0)
	//现在keyset就是有序的了

}

//2147483647
//1,2,3,4,5,6,7,8,9,10

func RadixSort(data []int) {
	getTargetRadix := func(num int, index int) int {
		numStr := strconv.Itoa(num)
		split := strings.Split(numStr, "") //按照位数来切割,
		if len(split) > index {
			target, _ := strconv.Atoi(split[(len(split)-1)-index])
			return target
		} else {
			return 0
		}
	}
	var maxElement int
	for i := range data {
		if data[i] > maxElement {
			maxElement = data[i]
		}
	}
	//0,0,1,2,4,1,9
	itoa := strconv.Itoa(maxElement)
	maxLength := len(itoa)
	for i := 0; i < maxLength; i++ {
		for j := 0; j <= len(data)-1; j++ {
			//接下来选择插入排序
			//for i := 0; i <= length-1; i++ {
			dataJ := getTargetRadix(data[j], i)
			dataJo := data[j]
			if j == 0 {
				continue
			}
			k := j - 1
			for ; k >= 0; k-- {
				dataK := getTargetRadix(data[k], i)
				dataKo := data[k]
				if dataK > dataJ {
					data[k+1] = dataKo
				} else {
					break
				}
			}
			data[k+1] = dataJo
			//}
		}

	}
}

// 将一个数组先对半劈开,然后再劈开,最后劈到每一个小列表里都只剩下2个元素左右
// 这里的排序只考虑正整数
func MergeSort(data []int) []int {
	compare := func(chaosDataA, chaosDataB []int) []int {
		mergerDataLen := len(chaosDataA) + len(chaosDataB)
		resultMergeData := make([]int, mergerDataLen)
		turnA := new(int)
		turnB := new(int)
		for i := 0; i < mergerDataLen; i++ {
			if turnA == nil && turnB == nil {
				break
			}
			if turnA == nil {
				resultMergeData[i] = chaosDataB[*turnB]
				*turnB++
				continue
			}
			if turnB == nil {
				resultMergeData[i] = chaosDataA[*turnA]
				*turnA++
				continue
			}

			if chaosDataA[*turnA] < chaosDataB[*turnB] {
				resultMergeData[i] = chaosDataA[*turnA]
				*turnA++
				if *turnA > len(chaosDataA)-1 {
					turnA = nil
				}
			} else {
				resultMergeData[i] = chaosDataB[*turnB]
				*turnB++
				if *turnB > len(chaosDataB)-1 {
					turnB = nil
				}
			}
		}
		return resultMergeData
	}

	if len(data) > 2 {
		firstData := data[:len(data)/2]
		lastData := data[len(data)/2:]
		firstSortedData := MergeSort(firstData)
		lastSortedData := MergeSort(lastData)
		return compare(firstSortedData, lastSortedData)
	} else if len(data) == 2 {
		return compare(data[:1], data[1:])
	} else {
		return data
	}

}

/**
1->1->:4->5,
	   3->4,
	   2->6
------------
4->5,
1->1->2->3->4,
6
------------
:4->5,
:3->4,
1->1->2:->6
------------
4->5,
1->1->2->3->4,
6
------------
5,
1->1->2->3->4->4,
6
------------
nil
1->1->2->3->4->4->5,
6
------------
nil
1->1->2->3->4->4->5->6
nil
*/
