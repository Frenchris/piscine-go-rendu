package main

import "github.com/01-edu/O1"

func printNumbers() {
	i := byte('0')
	for i <= '9'{
		O1.Putchar(i)
		i++
	}
}

func main() {
	printNumbers()
}
