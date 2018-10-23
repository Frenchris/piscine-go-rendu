package main

import "github.com/01-edu/O1"

func printalphabet() {
	i := byte('a')
	for i <= 'z' {
		O1.Putchar(i)
		i++
	}
}

func main() {
	printalphabet()
}
