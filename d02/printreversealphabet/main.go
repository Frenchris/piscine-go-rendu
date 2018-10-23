package main

import "github.com/01-edu/O1"

func printReverseAlphabet() {
  i := byte('z')
  for i >= 'a'{
    putchar(i)
    i--
  }
}

func main() {
  printReverseAlphabet()
}
