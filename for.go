package main

import "fmt"

func main() {
    var b int = 5
    var a int = 0

    numbers := [6]int{1,2,3,4}

    for i := 0; i < 5; i++ {
        fmt.Printf("i = %d\n",i)
    }

    for a < b {
        fmt.Printf("a = %d\n",a)
        a++
    }

    for i,x:= range numbers {
        fmt.Printf("i = %d x = %d\n",i,x)
    }
}