package main

import "github.com/01-edu/O1"

func isNegative(n int) {
	if n >= 0 {
		O1.Putchar(byte('P'))
	} else {
		O1.Putchar(byte('N'))
	}
}
