[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-named-return-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-named-return-functions/blob/main/README.md)  
go version 1.22.1

## Funções com retornos nomeados
Ele não é muito utilizado, mas é muito "bonito" de se escrever uma funçao com ele, a diferença é que você pode declarar variáveis com o tipo do retorno para usa-las sem precisar declara-las, e usar o return de maneira simples, retornando os retornos nomeados em sua ordem, no nosso caso ```(sum int, sub int)```, veja o exemplo:  
```go
package main

import "fmt"

//Retorno nomeado retorna uma variável sum e uma sub
func calc(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2 //repare que não declaramos a variável com :=, ela já é declarada no retorno
	sub = n1 - n2 // o mesmo que sum
	return        // ao passar o return dessa maneira, 
  //              ele retorna os duas variáveis nomeadas, sum e sub
	//sendo o mesmo que
	// return sum, sub
}

func main() {
	sum, sub := calc(30, 20)
	fmt.Println(sum, sub) // output 50 10
}
```