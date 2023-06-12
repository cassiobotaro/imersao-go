package main

import (
	"fmt"
	"strings"
)

func main() {
	// Slices (fatias), são sequências de tamanho variável cujos elementos tem o mesmo tipo..

	// Sua declaração é bastante similar aos arrays, mas seu tamanho não é definido.
	var slice []int
	// No exemplo acima, podemos ver a declaração de um slice de números inteiros
	fmt.Println("Slice vazio", slice)
	fmt.Println("Tamanho do slice: ", len(slice))

	// Outra maneira de criar um slice é inicializando-o, declarando de forma literal seus valores.
	vogais := []string{"a", "e", "i", "o", "u"}
	fmt.Println("Slice declarado de forma literal: ", vogais)
	fmt.Println("Tamanho do slice:", len(vogais))

	// Podemos declarar slices a partir de arrays
	meses := [12]string{"jan", "fev", "mar", "abr", "mai", "jun", "jul", "ago", "set", "out", "nov", "dez"}
	primeiroSemestre := meses[:6]
	// Neste exemplo, pegamos uma fatia do array que vai do primeiro mês até o sexto
	// É o equivalente a [0:6], porém podemos omitir o primeiro valor.
	fmt.Println("O primeiro semestre é: ", primeiroSemestre)

	// Podemos utilizar a notação de fatiamento para recuperar á fatias a partir de determinado valor.
	segundoSemestre := meses[6:]
	// Assim como podemos omitir o primeiro valor no fatiamento, também podemos omitir o segundo
	// Quando fazemos isso, é o equivalente a dizer [6: 13]
	fmt.Println("O segundo semestre é: ", segundoSemestre)

	// E tambem podemos criar uma fatia de todo o array
	ano := meses[:]
	// Slices são como referências a arrays
	fmt.Println("O ano tem os meses: ", ano)

	// Vemos que a sintáxe de fatiamento é s[i:j] com 0 <= i <= j <= cap(s)
	// Podendo ser omitidos i ou j

	// Outra maneira de declarar slices é utilizando a função make
	// Arrays declarados dessa maneira, serão preenchidos com valores-zero de seus determinados tipos, se o tamanho for maior que zero.
	sliceMake1 := make([]int, 5) // slice de tamanho 5
	fmt.Println("Os elementos do array são: ", sliceMake1)
	fmt.Println("O slice possui tamanho: ", len(sliceMake1), " e capacidade: ", cap(sliceMake1))

	// Podemos definir uma capacidade (inicial), porém um tamanho menor.
	// Neste exemplo como o tamanho é zero, não teremos elementos.
	// A capacidade é 5, portanto poderemos adicionar 5 elementos antes do array precisar ser redimensionado
	sliceMake2 := make([]int, 0, 5) // slice de tamanho 0, mas capacidade de 5 elementos
	fmt.Println("Os elementos do array são: ", sliceMake2)
	fmt.Println("O slice possui tamanho: ", len(sliceMake2), " e capacidade: ", cap(sliceMake2))

	// Neste exemplo, mesmo que o slice tenha capacidade para 5 elementos, estamos alocando somente 3. Ou seja adicionamos 3 vezes o valor zero do tipo.
	sliceMake3 := make([]int, 3, 5) // slice de tamnho 3, mas 5 de capacidade
	fmt.Println("Os elementos do array são: ", sliceMake3)
	fmt.Println("O slice possui tamanho: ", len(sliceMake3), " e capacidade: ", cap(sliceMake3))

	// O valor zero de um slice é nil.
	var sliceZero []int
	fmt.Println("O valor zero de um slice é nil?: ", sliceZero == nil)

	// Assim como array, podemos ter slices de slices
	jogoDaVelha := [][]string{
		{"x", " ", "o"},
		{"o", "x", " "},
		{"x", " ", "o"},
	}
	fmt.Println(strings.Join([]string{
		strings.Join(jogoDaVelha[0], " "),
		strings.Join(jogoDaVelha[1], " "),
		strings.Join(jogoDaVelha[2], " ")}, "\n"))

	// Para acessar o valor de um slice, utilizamos um índice numérico
	// que representa sua posição e inicia-se em zero.
	// Por exemplo, dado um slice de vogais, temos:
	// ["a", "e", "i", "o", "u"]
	//   0    1    2    3    4
	fmt.Println("A posição 2 do slice é: ", vogais[2])
	// Ao acessarmos a posição 2 do slice temos a letra "i".
	fmt.Println("A posição 0 do slice é: ", vogais[0])
	// O mesmo acontece quando acessamos a primeira posição do slice.
	fmt.Println("Acessando o segundo elemento, da segunda linha: ", jogoDaVelha[1][1])
	// Lembre-se que indíces iniciam em 0, por a segunda linha está no índice 1.
	// Em slices de múltiplas dimensões, também acessamos suas posições através dos índices.

	// Para modificarmos um valor de um slice, basta atribuir um valor a um índice.
	vogais[3] = "j"
	fmt.Println("Vogais: ", vogais)
	vogais[3] = "o" // corrigindo de volta as vogais, antes que alguém reclame.

	// Slides criados à partir de outros compartilham os mesmos valores
	// como exemplo, ano := meses
	// Se eu modificar o valor presente em ano, a fatia meses também será modificada.
	ano[0] = "janeiro"

	fmt.Println("Anos e meses representam o mesmo array = ", ano, "==", meses)

	// depois voltamos pro lugar
	ano[0] = "jan"

	// O mesmo se aplica quando temos mais de uma dimensão.
	jogoDaVelha[1][2] = "o"
	fmt.Println(strings.Join([]string{
		strings.Join(jogoDaVelha[0], " "),
		strings.Join(jogoDaVelha[1], " "),
		strings.Join(jogoDaVelha[2], " ")}, "\n"))

	// Outra maneira de modificar um slice é adicionando elementos a ele.
	fmt.Println("Antes de inserir um elemento:  ", slice, "len= ", len(slice), "cap=", cap(slice))
	slice = append(slice, 42)
	// O slice cresce a medida que novos elementos são inseridos
	fmt.Println("Depois de inserir um elemento:  ", slice, "len= ", len(slice), "cap=", cap(slice))
	// Nós podemos também adicionar vários elementos de uma vez
	slice = append(slice, 1, 3, 5, 7, 9)
	fmt.Println("Depois de inserir vários elementos:  ", slice, "len= ", len(slice), "cap=", cap(slice))

	// Quando percorremos um slice, dois valores são retornados.
	// O primeiro é o índice e o segundo é a cópia do elemento no índice.
	for indice, valor := range slice {
		fmt.Println("O valor no índice ", indice, "é ", valor)
	}

	// Para copiar elementos de um slice em outro, a maneira mais eficiente é utilizando a função copy
	// Primeiro definimos um fatiamento de tamanho 6
	primeiraFatia := []int{1, 2, 3, 6, 7, 8}
	// Depois outro fatiamento com tamnho 3
	segundaFatia := []int{4, 5, 9}
	// Para a junção, criamos um fatiamento de tamanho 9
	fatiaJuncao := make([]int, 9)
	// primeiro copiamos os três elementos da primeira fatia nas três primeiras posições da nossa fatia de junção
	copy(fatiaJuncao[:3], primeiraFatia[:3])
	fmt.Println("Junção: ", fatiaJuncao)
	// Em seguida, copiamos os dois primeiros elementos da segunda fatia, nas posições 3 e 4 da fatia de junção(fatia de 3 a 5)
	copy(fatiaJuncao[3:5], segundaFatia[:2])
	fmt.Println("Junção: ", fatiaJuncao)
	// Agora copiaremos da terceira posição da primeira fatia até o final, sendo o destino á partir do quinto elemento da fatia de junção
	copy(fatiaJuncao[5:], primeiraFatia[3:])
	fmt.Println("Junção: ", fatiaJuncao)
	// Copia do segundo elemento da segunda fatia até o final, na fatia iniciada na oitava posição do
	copy(fatiaJuncao[8:], segundaFatia[2:])
	fmt.Println("Junção: ", fatiaJuncao)
	// A função copy também é util quando estamos copiando slices de tamanhos diferentes, pois o número de elementos copiados é sempre igual ao slice de menor tamanho.
	copy(segundaFatia, primeiraFatia)
	fmt.Println("Após copiar a primeira fatia na segunda: ", segundaFatia)
	// Repare que nenhum problema ocorre, mesmo a segunda fatia sendo menor que a primeira

	// Não se preocupe se pareceu complexo, a operação de cópia é utilizada para evitar iterações desnecessárias e quando performance for necessária.

	// Uma maneira de deletar os elementos de uma lista é criando uma nova lista a partir dos elementos que restaram.
	// !! Cuidado, pois esta operação possui complexidade linear O(n)
	fmt.Println("Antes de remover um elemento:  ", slice, "len= ", len(slice), "cap=", cap(slice))
	indiceRemovido := 2
	slice = append(slice[:indiceRemovido], slice[indiceRemovido+1:]...)
	fmt.Println("Depois de remover um elemento:  ", slice, "len= ", len(slice), "cap=", cap(slice))

	// Com a adição de generics na versão 1.18 da linguagem Go
	// Aumentamos a capacidade de reutilização de funções de tipos distintos
	// mas que apresentam mesmos comportamentos
	// No exemplo abaixo, a mesma função pode ser utilizada com inteiros
	// ou string pois ambas as coleções são de elementos comparáveis (==)
	fmt.Println("2 in 1,2,3: ", In([]int{1, 2, 3}, 2))
	fmt.Println("5 in 1,2,3: ", In([]int{1, 2, 3}, 5))

	fmt.Println("a in a,b,c: ", In([]string{"a", "b", "c"}, "b"))
	fmt.Println("d in a,b,c: ", In([]string{"a", "b", "c"}, "d"))

	// A tentativa de acessar uma posição em um slice
    // que não exista, irá gerar um pânico
    // No array este erro se dá em tempo de compilação

    // Slice tem tamanho 5
    fmt.Println("Slice=", slice, "tamanho=", len(slice))
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ocorreu um pânico:", r)
		}
	}()
    // Ao acessar a posição len(slice) + 1 (6 neste exemplo), ocorrerá um pânico
	fmt.Println(slice[len(slice) + 1])
}

// Uma feature interessante e com grande potencial é a utilização de generics
// como visto no exemplo abaixo.
// Já existem outras features experimentais sendo testadas na linguagem
// utilizando generics, veja em https://github.com/golang/exp/
func In[T comparable](collection []T, n T) bool {
	for _, item := range collection {
		if n == item {
			return true
		}
	}
	return false
}
