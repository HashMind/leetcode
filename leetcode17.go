package main

import (
	"fmt"
	"strings"
)

type PhoneTyper struct {
	keys []string
}
type TargetTyper struct {
	keys   []string
	values []string
}

func letterCombinations(digits string) {
	phoneNums = make(map[string]PhoneTyper, 0)
	phoneNums["2"] = PhoneTyper{keys: []string{"a", "b", "c"}}
	phoneNums["3"] = PhoneTyper{keys: []string{"d", "e", "f"}}
	phoneNums["4"] = PhoneTyper{keys: []string{"g", "h", "i"}}
	phoneNums["5"] = PhoneTyper{keys: []string{"j", "k", "l"}}
	phoneNums["6"] = PhoneTyper{keys: []string{"m", "n", "o"}}
	phoneNums["7"] = PhoneTyper{keys: []string{"p", "q", "r", "s"}}
	phoneNums["8"] = PhoneTyper{keys: []string{"t", "u", "v"}}
	phoneNums["9"] = PhoneTyper{keys: []string{"w", "x", "y", "z"}}
	numbers := strings.Split(digits, "")

	result := make([]string, 0)

	output := getLetter(numbers, result)
	fmt.Println(output)

}

var phoneNums = make(map[string]PhoneTyper, 0)

func getLetter(keys []string, input []string) (output []string) {
	//从keys里取key然后从key-values取出匹配的字符元素出来,然后values里的每一个元素末尾都拼接上(笛卡尔积)
	//keys:=keys[1:],直到没有按键
	typer := phoneNums[keys[0]]
	if len(input) == 0 {
		output = append(output, typer.keys...)
	} else {
		for i := range input {
			valuePrefix := input[i]
			for j := range typer.keys {
				valueSuffix := typer.keys[j]
				value := string(valuePrefix) + string(valueSuffix)
				output = append(output, value)
			}

		}
	}
	if len(keys) > 1 {
		output = getLetter(keys[1:], output)
	}
	return output
}

func main17() {
	letterCombinations("22")
}
