package main

import "github.com/01-edu/O1"

func show(a byte, b byte, c byte) {
	if a < b && b < c {
		O1.Putchar(a)
		O1.Putchar(b)
		O1.Putchar(c)
		if !(a == '7' && b == '8' && c == '9') {
			O1.Putchar(',')
			O1.Putchar(' ')
		}
	}
}

func printCombn() {
	a := byte('0')
	b := byte('0')
	c := byte('0')
	for a <= '7' {
		for b <= '8' {
			for c <= '9' {
				show(a, b, c)
				c++
			}
			c = '0'
			b++
		}
		b = '0'
		a++
	}
}
