package main

import (
	"fmt"
	"github.com/apspan/protogen/commons/proto/demo"
)

func main() {
	x := demo.Demo{
		Date:        "",
		Value:       123,
		SecondValue: 0,
		ThirdValue:  0,
		FourthValue: "Hello, 世界, Απόστολος",
	}
	fmt.Println(x.FourthValue)
}
