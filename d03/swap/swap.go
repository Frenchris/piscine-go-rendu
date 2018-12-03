package main

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
