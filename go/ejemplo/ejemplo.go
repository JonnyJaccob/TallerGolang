package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main() {
    x := 2.0
    n := 3.0
    lim := 10.0

    result := pow(x, n, lim)
    fmt.Printf("math.Pow(%.2f, %.2f) = %.2f\n", x, n, result)
}

	