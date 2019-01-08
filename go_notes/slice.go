package main

import "fmt"

func main() {
	data := [...]int {0,1,2,3,4,5,6,7,8,9}
	// [low: high: max], len = high - low, cap = max - low
	s := data[2:4] // 省略max
	fmt.Println(len(s), cap(s))
	s = data[5:] // len=5, cap=5, 省略high和max
	s = data[:3] // 省略low和max, len=3, cap = 10
	// slice的读写实际操作的是底层的数组
	s[0] += 100
	s[2] += 200
	fmt.Println(s)
	fmt.Println(data)
}
