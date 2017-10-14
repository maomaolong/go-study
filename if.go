package main

import "fmt"

func main() {
    if 1 == 1 {
        fmt.Println("1 == 1")
    }

    switch 10 {
        case 1: fmt.Println("case 1")
        case 10: fmt.Println("case 10")
        case 100: fmt.Println("case 100")
        default: fmt.Println("default")
    }
}