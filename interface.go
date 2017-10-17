package main

import("fmt")

type Phone interface {
	call()
}

type iPhone struct {

}

func (phone iPhone) call(){
	fmt.Println("I am iPhone")
}

type Nokia struct {

}

func (phone Nokia) call(){
	fmt.Println("I am Nokia")
}

func main() {
	var phone Phone

	phone = new(iPhone)
	phone.call()

	phone = new(Nokia)
	phone.call()
}