package main

import (
	"fmt"
	"reflect"
)

func typeGo() {
	var x bool          // boolean mặc đinh là false
	var Float32 float32 // float32|64|int mặc định là 0
	var Float64 float64
	var StringT string
	var StringP *string
	g := "a"                                                    // rune (int32) mặc định là 0
	fmt.Println(StringT == "", 'a', StringP, reflect.TypeOf(g)) // in ra chuỗi rỗng
	fmt.Println(Float32, Float64, x)
	if iftest := 10; iftest > 5 {
		fmt.Println("x lớn hơn 5")
	}

}
func main() {
	typeGo()
	fmt.Println("Hello, World!")
}
