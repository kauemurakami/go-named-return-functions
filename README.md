[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-named-return-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-named-return-functions/blob/main/README.md)  
go version 1.22.1

## Functions with named returns
It is not widely used, but it is very "beautiful" to write a function with, the difference is that you can declare variables with the return type to use them without having to declare them, and use return in a simple way , returning the named returns in their order, in our case ```(sum int, sub int)```, see the example:  
```go
package main

import "fmt"

//Named return returns a sum and a sub variable
func calc(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2 // notice that we do not declare the variable with :=, it is already declared in the return
	sub = n1 - n2 // the same as sum
	return        // when passing return like this, 
  //              it returns the two named variables, sum and sub
	// being the same as
	// return sum, sub
}

func main() {
	sum, sub := calc(30, 20)
	fmt.Println(sum, sub) // output 50 10
}
```