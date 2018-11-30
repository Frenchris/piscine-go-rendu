package main

import "github.com/01-edu/O1"

func putStr(str string) {
	i := 0
	for i < len(str) {
		O1.Putchar(str[i])
		i++
	}
}
