package main

import "fmt"

func main(){
	var a int
	var b float64
	var c float64
	c = 0
	fmt.Print("How many words do you want to enter: ")
	fmt.Scanln(&a)
	fmt.Println("Enter ",a," words \n")

	for i:=0;i<a;i++{
		fmt.Scanln(&b)
		c = c+b
	}

	fmt.Println("The answer is")
	fmt.Println(c)

}
