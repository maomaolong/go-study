package main

import "fmt"

func main() {
    var varMap map[string] string
    varMap = make(map[string] string)

    varMap["hunan"] = "changsha"
    varMap["hubei"] = "wuhan"

    for key := range varMap {
    	fmt.Println(key,varMap[key])
    }

    city , ok := varMap["beijing"]
    if ok {
    	fmt.Println(city)
    }else{
    	fmt.Println("not beijing")
    }
}