package main
import "fmt"

type Vertex struct {
     X int
     Y int
     str string
}

func main() {
     fmt.Println(Vertex{1, 2, "ok"})
}