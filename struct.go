package main

import "fmt"

type Book struct {
	name string
	id string
	author string
}

func main() {
    var book1 Book

    book1.name = "hello kity"
    book1.id = "1"
    book1.author = "chenglong"

    fmt.Println("book1.author",book1.author)
}