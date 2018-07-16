package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "this is a string"
	fmt.Printf("%t\n", strings.HasPrefix(str, "th"))
	fmt.Printf("%d\n", strings.Index(str, "is"))

	new_str := strings.Replace(str, "is", "IS", -1)
	fmt.Println(new_str)

	s := "hey you"
	var p *string = &s
	*p = "haha"
	fmt.Printf("%p", p)
	fmt.Printf("%s, %s", *p, s)

	x := min(1,2,4,0,-1)
	fmt.Println(x)

	slice := []int{9,3,5,8,1}
	x = min(slice...)
	fmt.Println(x)

	f()
}

// 变长参数
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v:= range s {
		if v < min {
			min = v
		}
	}
	return min
}

// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
