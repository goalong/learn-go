package main

import "fmt"

type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Println("value is", asset.getValue())
}


func main() {
	var o valuable = stockPosition{"goog", 20.1, 4}
	showValue(o)
	o = car{"bmw", "m3", 9990}
	showValue(o)
}
