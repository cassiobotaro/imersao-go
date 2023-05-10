package main

import "fmt"

func main() {
	// Arrays são uma sequência de tamanho fixo de zero ou mais  elementos de um tipo específico.

	// Seu tamanho é definido durante a criação.
	var arr [5]int
	// No exemplo acima, podemos ver a declaração de um array de inteiros
	// com o tamanho 5. Como o array não foi inicializado, o valor-zero do
	// tipo é colocado. Resultando assim em um array com zeros [0, 0, 0, 0, 0].
	fmt.Println("Array com valores-zero: ", arr)
	fmt.Println("Tamanho do array: ", len(arr))

	// Outra maneira de criar um array é inicializando-o, declarando de forma literal seus valores.
	arrayDeclarado := [5]string{"a", "e", "i", "o", "u"}
	fmt.Println("Array declarado de forma literal: ", arrayDeclarado)
	// Podemos utilizar a função embutida len para ver o tamanho de um array.
	fmt.Println("Tamanho do array: ", len(arrayDeclarado))

	// O valor entre parênteses diz o tamanho do array e o tipo que aparece após
	// os parênteses são os valores esperados a serem armazenados no array.

	//  Se utilizarmos a forma literal de declaração, mas os valores declarados
	// não forem igual ao tamanho do array, valores zero são adicionados.
	outroArray := [10]string{"g", "h", "i"}
	fmt.Println("Array delcarado de forma literal mas sem preencher todos os elementos: ", outroArray)
	fmt.Println("Tamanho do array: ", len(outroArray))
	fmt.Println("Capacidade do array: ", cap(outroArray))
	// No exemplo acima, o array possui tamanho 10, mas somente 3 valores foram declarados de forma literal.
	// Sendo assim, valores zero ("" já que o tipo é string) foram adicionados.

	// Se reticências estiverem no lugar no tamanho do array, seu tamanho é o número de elementos declarados
	arrayRet := [...]int{7, 8, 9}
	fmt.Println("Array definido com reticências: ", arrayRet)
	fmt.Println("Tamanho do array definido com reticências: ", len(arrayRet))
	fmt.Println("Capacidade do array: ", cap(arrayRet))

	// Na forma literal, os elementos podem ser definidos de forma não ordenada e até serem omitidos
	esparsos := [10]int{1: 1, 5: 1, 3: 1}
	// o valor 1 é colocado na posição 1, 3 e 5
	fmt.Println("Array com valores esparsos: ", esparsos)
	fmt.Println("Tamanho do array: ", len(esparsos))
	fmt.Println("Capacidade do array: ", cap(esparsos))

	// Aqui um exemplo utilizando reticências, e definindo o índice 99 com o valor 1.
	// Todos os outros índices são preenchidos com valores-zero do tipo.
	arrayOmite := [...]int{99: 1}
	fmt.Println("Array com índices omitidos: ", arrayOmite)
	// Teremos um array de 100 elementos, mesmo sem declará-los de forma explícita.
	fmt.Println("Tamanho do array: ", len(arrayOmite))
	fmt.Println("Capacidade do array: ", cap(arrayOmite))

	// arrays podem ter mútiplas dimensões.
	jogoDaVelha := [3][3]string{
		{"x", " ", "o"},
		{" ", "o", " "},
		{"x", " ", "o"},
	}
	fmt.Println("Exemplo de um array multidimensional", jogoDaVelha)
	fmt.Println("Tamanho de um array multidimensional", len(jogoDaVelha))
	// No exemplo acima, temos um array de duas dimensões.
	// O array possui 3 linhas e 3 colunas em cada linha.

	// Para acessar o valor de um array, utilizamos um índice numérico
	// que representa sua posição e inicia-se em zero.
	// Por exemplo, dado um array de vogais, temos:
	// ["a", "e", "i", "o", "u"]
	//   0    1    2    3    4
	fmt.Println("A posição 2 do array é: ", arrayDeclarado[2])
	// Ao acessarmos a posição 2 do array temos a letra "i".
	fmt.Println("A posição 0 do array é: ", arrayDeclarado[0])
	// O mesmo acontece quando acessamos a primeira posição do array.
	fmt.Println("Acessando o segundo elemento, da segunda linha: ", jogoDaVelha[1][1])
	// Lembre-se que indíces iniciam em 0, por a segunda linha está no índice 1.
	// Em arrays de múltiplas dimensões, também acessamos suas posições através dos índices.

	// Para modificarmos um valor de um array, basta atribuir um valor a um índice.
	outroArray[3] = "j"
	fmt.Println("outroArray após a modificação de um dos seus índices: ", outroArray)
	// O outro array antes da modificação ["g", "h", "i", "", "", "", "", "", "", ""]
	// Modificamos a posição 3            ["g", "h", "i", "j", "", "", "", "", "", ""]
	// O valor zero "" se torna "j"

	jogoDaVelha[1][0] = "x"
	fmt.Println("Exemplo de um array multidimensional", jogoDaVelha)
	// O mesmo se aplica quando temos mais de uma dimensão.

	// Como um array possui tamanho e capacidade fixo, não é possível
	// deletar um elemento do array.
	// porém podemos atribuir um valor-zero caso necessário.
	fmt.Println("Antes de \"remover\" uma posição: ", outroArray)
	outroArray[3] = ""
	fmt.Println("Depois de \"remover\" uma posição: ", outroArray)
}
