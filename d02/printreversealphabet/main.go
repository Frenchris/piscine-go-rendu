package main

import "github.com/01-edu/O1"

func printReverseAlphabet() {
	i := byte('z')
	for i >= 'a'{
		O1.Putchar(i)
		i--
	}
}

func main() {
	printReverseAlphabet()
}
