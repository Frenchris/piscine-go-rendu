package main

import "github.com/01-edu/O1"

func putNbr(n int64) {
	if n == -9223372036854775808 {
		O1.Putchar('-')
		O1.Putchar('9')
		n = 223372036854775808
	}
	if n < 0 {
		n = n * -1
		O1.Putchar('-')
	}
	if n < 10 {
		O1.Putchar(byte(n + 48))
	} else {
		putNbr(n / 10)
		putNbr(n % 10)
	}
}
