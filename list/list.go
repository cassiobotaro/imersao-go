package main

import (
	"container/list"
	"fmt"
	"strconv"
)

// Para ilustrar como o conhecimento e utilização de uma estrtura de dados no momento correto pode contribuir com a eficiência do seu código vamos ao seguinte código.
// Criaremos uma coleção de elementos (strings no exemplo) com dois elementos, em seguida faremos múltiplas inserções de novos elementos no início da coleção.
// Imagine algo assim: ["passaro", "cachorro"] e vamos adicionar "1", em seguida "2".
// Como a inserção é feita no início da coleção, teriamos ao final ["1", "2", "passaro", "cachorro"]

// Abaixo temos a demonstração do que foi explicado, utilizando slices e listas.

// Lista é uma estrutura de dados que armazena uma sequência de elementos, onde cada elemento é composto por um valor e uma referência para o próximo elemento da lista. Diferentemente de um array, a lista não tem um tamanho fixo, podendo ser aumentada ou diminuída durante a execução do programa. As listas podem ser implementadas de várias maneiras, sendo duas das mais comuns a lista encadeada (linked list) e a lista duplamente encadeada (doubly linked list). As listas são amplamente utilizadas em algoritmos de processamento de dados, como ordenação e pesquisa, devido à sua flexibilidade e eficiência.

// Go possui uma implementação de lista duplamente encadeada em um pacote chamado container na própria linguagem.
// Para uma melhor vizualização desta estrtura de dados, acesse: https://visualgo.net/en/list

// Para garantir que nossas implementações são equivalentes, você pode executar este arquivo.

// Com a intenção de comparar as duas implementações, vamos invocar as duas funções com o número de interações de 10.000 vezes.
// Para isto podemos utilizar o comando "go test"

// A lista (list) é mais eficiente que o slice ([]T) em cenários em que há a necessidade frequente de inserções ou remoções de elementos na parte inicial (ou em qualquer posição intermediária) da sequência de dados.
// Isso ocorre porque as listas são implementadas como uma estrutura encadeada, em que cada elemento é armazenado em um nó que contém um ponteiro para o próximo elemento (e, em caso de lista duplamente encadeada, um ponteiro para o elemento anterior).
// Dessa forma, a inserção ou remoção de elementos na parte inicial da lista requer apenas a atualização dos ponteiros dos nós adjacentes, o que é uma operação muito rápida.
// Já o slice é implementado como um array com tamanho fixo, e, portanto, inserir ou remover elementos no início requer a realocação de todos os elementos do array para frente ou para trás, o que pode ser uma operação muito custosa, especialmente para grandes conjuntos de dados.
// No código abaixo , a lista é mais eficiente porque há a necessidade de inserir 20 elementos na parte inicial da sequência em cada iteração, o que é uma operação muito mais rápida com uma lista do que com um slice.

// Ainda no pacote container, podemos encontrar implementação de mais duas estruturas de dados, Heap e Ring.
// Recomendo a leitura da documentação sobre as estruturas.
// https://pkg.go.dev/container/heap
// https://pkg.go.dev/container/ring


func usaContainerList(numIteracoes int) {

	for i := 0; i < numIteracoes; i++ {
		// Cria uma nova lista
		valores := list.New()
		// Adiciona 2 elementos na lista.
		valores.PushBack("bird")
		valores.PushBack("cat")
		// Adiciona 20 elementos no início da lista.
		for i := 0; i < 20; i++ {
			// Adiciona os numeros convertidos em string.
			valores.PushFront(strconv.Itoa(i))
		}
		imprimeListaComoSlide(valores)
	}

}

func usaSlice(numIteracoes int) {

	for i := 0; i < numIteracoes; i++ {
		// Cria um slice vazio
		valores := []string{}
		// Adiciona 2 elementos no slice.
		valores = append(valores, "bird")
		valores = append(valores, "cat")
		// Adiciona 20 elementos no início do slice.
		for i := 0; i < 20; i++ {
			// Cria uma nova fatia e coloca a string em seu início.
			tempSlice := []string{}
			tempSlice = append(tempSlice, strconv.Itoa(i))
			// Agora adicione todas as strings após a primeira
			valores = append(tempSlice, valores...)
		}
		fmt.Println(valores)
	}

}

// Para facilitar a vizualização dos elementos da lista, esta função foi criada
func imprimeListaComoSlide(l *list.List) {
	// Cria um slice vazio
	slice := make([]interface{}, l.Len())

	// Para cada elemento da lista, adicione ao slice
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		slice[i] = e.Value
		i++
	}
	// Imprime o slice final
	fmt.Println(slice)
}

func main() {
	// Podemos invocar as duas funções para garantir que se comportam de maneira idêntica
	usaSlice(5)
	usaContainerList(5)
}
