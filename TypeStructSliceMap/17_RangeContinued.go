package main
import "fmt"

func main() {
	pow := make([]int, 10) //空の配列(スライス)を定義
	for i:=range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value:=range pow {
		fmt.Printf("%d\n", value)
	}
}
