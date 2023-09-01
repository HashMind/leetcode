package main

import (
	"errors"
	"fmt"
	"strings"
)

func main2() {

	rsult := strings.Split("a*bcd*ef*", "*")
	fmt.Println(rsult)
}

// 正则解析器,匹配:*.
// *不能出现在第一位,*不能跟在*后面
func regex(target, pattern string) error {
	//aa*.*aa
	//pattern的解析:通配符出现的位置
	//1:slice(pattern,“*”),先扫描*号,然后切割,*号前面的模式,最终可能匹配出一个变长字符串
	//然后就是从pattern的数组里从前往后匹配,遇到.则代表可以是任意字符,最后一个字符
	patterns := strings.Split(pattern, "*") //数组的最后一个是不会包含*号的
	//先对pattens的格式校验一遍,防止后续的匹配到最后才发现在做无用功
	//patterns里合规的元素需要遵循: (“字母”或“点”至少出现1次,其后“星号”出现0次或1次),前面这条约束至少要出现1次
	for i := range patterns {
		pattern := patterns[i]
		if len(pattern) == 0 {
			errors.New("invalid regex pattern")
		}
	}
	//

	//targetRunes := []rune(target)
	for i := range patterns {
		pattern := patterns[i]
		patternRunes := []rune(pattern)
		for u := range patternRunes {
			character := patternRunes[u]

			switch character {
			case '.':

			}
		}
	}
	return nil
}
