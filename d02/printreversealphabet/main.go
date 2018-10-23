package main

import "01"

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
