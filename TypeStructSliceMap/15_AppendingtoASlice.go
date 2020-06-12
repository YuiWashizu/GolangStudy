package main
import "fmt"

func main() {
	var s[]int
	printSlice(s)

	//append works on nil slice
	s = append(s, 0)
	printSlice(s)

	//The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	//We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(x []int) {
	fmt.Printf("len = %d, cap = %d %v\n", len(x), cap(x), x)
}
