package main
import (
       "fmt"
       "math"
)

func Sqrt(x, test float64) float64 {
     z := test
     for {
     if (z*z - x) / (2*z) > 0.000001 {
          z -= (z*z - x) / (2*z)
     } else if (z*z - x) / (2*z) < -0.000001 {
          z -= (z*z - x) / (2*z)
     } else {
     return z
     }
     }
}

func main() {
     fmt.Printf("true sqrt      : %g\n", math.Sqrt(2))
     fmt.Printf("my sqrt init %g : %g\n", 2.0, Sqrt(2, 2))
     fmt.Printf("my sqrt init %g : %g\n", 1.0, Sqrt(2, 1))
     fmt.Printf("my sqrt init %g : %g\n", 4.0, Sqrt(2, 4))
     fmt.Printf("my sqrt init %g : %g\n", 8.0, Sqrt(2, 0.5))
}