package main

import (
	"fmt"
	"time"
)

func generateParenthesis(n int) {
	for i := 1; i <= n; i++ {
		es := make([]int, 1)
		es[0] = i
		getNextNumber(n-i, es)
	}
	fmt.Println(res)
	//然后开始织布
	for i := range res {
		e := res[i]
		for k := range e {
			for p := 0; p < e[k]; p++ {
				fmt.Print("(")
			}
			for p := 0; p < e[k]; p++ {
				fmt.Print(")")
			}

		}
		fmt.Println("")

	}
}

var res = make([][]int, 0)

func getNextNumber(extra int, prefix []int) {
	for i := 1; i <= extra; i++ {
		temp := append(prefix, i)
		if extra-i == 0 {
			res = append(res, temp)
		} else if extra-i > 0 {
			getNextNumber(extra-i, temp)
		} else {
			break
		}
	}
}

//n等于几,最终可能的结果就是1-n之间的不多余N的数字和等于N
//例如:n等于4
//1,1,1,1=()()()()
//1,1,2=()()(())
//1,2,1=()(())()
//1,3=()((()))

//3,1=((()))()
//2,1,1=(())()()

func main22() {
	format := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(format)
	generateParenthesis(5)
}
