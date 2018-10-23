package main

import "github.com/01-edu/O1"

func show(a int, b int) {
	if a < b {
		O1.Putchar(byte(a/10 + 48))
		O1.Putchar(byte(a%10 + 48))
		O1.Putchar(' ')
		O1.Putchar(byte(b/10 + 48))
		O1.Putchar(byte(b%10 + 48))
		if !(a == 98 && b == 99) {
			O1.Putchar(',')
			O1.Putchar(' ')
		}
	}
}

func printComb2() {
	a := int(0)
	b := int(0)
	for a <= 98 {
		b = a + 1
		for b <= 99 {
			show(a, b)
			b++
		}
		a++
	}
}
