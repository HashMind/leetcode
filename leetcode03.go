package main

import "fmt"

func main_03() {
	splice("1x21x3")
}

func splice(src string) {
	srcSlice := []byte(src)
	maxStr := ""
	tempByteSlice := make([]byte, 0)
	for _, b := range srcSlice {
		//每一个临时列表里之前的元素都需要和这个即将添加进去的新元素进行比对是否已经重复
		for k := 0; k < len(tempByteSlice); k++ {
			if tempByteSlice[k] == b {
				//如果位置N的元素和欲添加到临时列表里进行比对的元素发生了重复,
				//则将列表里的元素假装为长度最长的目标构建为字符串留存,
				//然后临时列表[0:n]的元素舍弃,添加新元素构成一个新临时列表,
				if len(tempByteSlice) > len(maxStr) {
					maxStr = string(tempByteSlice)
				}
				//需要从重复元素位置开始构建一个新的不存在重复的临时列表
				tempByteSlice = tempByteSlice[k+1:]
				break
			}
		}
		tempByteSlice = append(tempByteSlice, b)
		maxStr = string(tempByteSlice)
	}
	fmt.Println(maxStr)
}
