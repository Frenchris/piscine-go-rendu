package main

import "O1"

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
