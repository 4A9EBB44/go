package main

import "fmt"

func main() {
	i, j := 24, 235
	fmt.Println(&i, &j)

	p := &i
	fmt.Println("p ", p)
	fmt.Println("*p ", *p)
	fmt.Println("&p ", &p)
	fmt.Printf("Hello, World!\n")
	for ; i < 29; i++ {
		fmt.Println(i)
	}
}
