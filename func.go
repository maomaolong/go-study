package main

import "fmt"

func main() {
    fmt.Println("max(1,2) = ",max(1,2))
    a,b:= swap("str1","str2")
    fmt.Println("swap = ",a,b)
}

func max(num1 int,num2 int) (int) {
    if num1 > num2 {
        return num1
    }
    return num2
}

func swap(str1,str2 string) (string,string) {
    return str2,str1
}