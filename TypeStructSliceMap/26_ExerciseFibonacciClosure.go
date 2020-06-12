package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		ans := a + b
		a, b = b, ans
		return ans
	}
}

func main() {
	f := fibonacci()
	for i:=0; i<10; i++ {
		fmt.Println(f())
	}
}
