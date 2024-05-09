package main

import "fmt"

//Retorno nomeado retorna uma variável sum e uma sub
func calc(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2 //repare que não declaramos a variável com :=, ela já é declarada no retorno
	sub = n1 - n2 // o mesmo que sum
	return        // ao passar o return dessa maneira, ele retorna os duas variáveis nomeadas, sum e sub
	//o mesmo que
	// return sum, sub
}

func main() {
	sum, sub := calc(30, 20)
	fmt.Println(sum, sub)
}
