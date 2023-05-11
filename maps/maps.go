package main

import (
	"fmt"
)

func main() {
	// Mapa(map) é uma coleção não ordenada de pares chave/valor em que cada chave é unica e o valor associado a uma dada chave pode ser recuperado, atualizado ou removido usando, em média, um número constante de comparações de chaves, independente do tamanho da tabela.
	// Definição retirada do excelente livro: https://www.amazon.com.br/Linguagem-Programa%C3%A7%C3%A3o-Go-Alan-Donovan/dp/8575225464
	// Em outras linguagens são normalmente chamados dicionários, hash, hashmap, map, arrays associativos.

	// Sua declaração pode ser feita utilizando var
	var mVar map[string]int

	// Ou através do comando make
	mMake := make(map[string]int)

	// Ou mesmo de forma literal
	mLiteralVazia := map[string]int{}

	// Em todos os exemplos acima, vemos o mapeamento de uma chave string a um valor int.
	fmt.Println("As três maneiras de se definir um mapa, chegam ao mesmo resultado: ", mVar, mMake, mLiteralVazia)

	mLiteralNaoVazia := map[string]int{
		"janeiro":   2,
		"fevereiro": 4,
		"março":     6,
		"abril":     2,
	}
	fmt.Println("Exemplo de mapa já preenchido: ", mLiteralNaoVazia)

	// Elementos de um mapa podem ser acessados à paritr de suas chaves, na notação de índice, similar a arrays
	fmt.Println("O valor da chave janeiro é", mLiteralNaoVazia["janeiro"])
	fmt.Println("O valor da chave março é", mLiteralNaoVazia["março"])

	// Buscar uma chave não existente, retorna o valor zero para o tipo mapeado
	fmt.Println("A chave 'dezembro' não existe, o valor retornado é: ", mLiteralNaoVazia["dezembro"])
	// Caso necessário, podemos testar a existência de uma chave antes de realizar alguma operação
	valor, ok := mLiteralNaoVazia["dezembro"]
	if !ok {
		fmt.Println("A chave não existe!")
	} else {
		fmt.Println("Operação não executa com ", valor)
	}

	// Podemos adicinar uma nova chave ou alterar o valor de uma chave fazendo atribuição a ela:
	// exemplo de uma chave nova
	mLiteralNaoVazia["agosto"] = 42
	// exemplo de atualização de uma chave antiga
	mLiteralNaoVazia["março"] = 10
	fmt.Println("Map após alterações: ", mLiteralNaoVazia)

	// Para deletar uma associação, utilizamos a função delete
	fmt.Println("Map antes de deletar agosto: ", mLiteralNaoVazia)
	delete(mLiteralNaoVazia, "agosto")
	fmt.Println("Map após deletar agosto: ", mLiteralNaoVazia)

	// É seguro realizar operações de delete, len e range em mapa nulo
	delete(mVar, "chave")
	_ = len(mVar)
	for _, v := range mVar {
		fmt.Println("Se houvesse valores, imprimiria: ", v)
	}
	// Mais adiante temos um exemplo de atribuição a um mapa nulo
	// que não é uma operação segura e um panic ocorrerá

	// A iteração dos elementos de um mapa são realizadas utilizando range
	for chave, valor := range mLiteralNaoVazia {
		fmt.Println("A chave ", chave, "está associada ao valor ", valor)
	}

	// Operações de escrita e leitura em mapas em diferentes gorotinas não são seguras
	// Veja um exemplo através deste link: https://go.dev/play/p/8ksDHEguqcJ
	// Uma maneira de contornar este problema é utilizando ferramentas de sincronia como Mutex e canais ou sync.Map como neste exemplo: https://go.dev/play/p/l_pM1LfOTHT


    // Com a adição de generics na versão 1.18 da linguagem Go
    // Aumentamos a capacidade de reutilização de funções de tipos distintos
    // mas que apresentam mesmos comportamentos
    // No exemplo abaixo, a mesma função pode ser utilizada com inteiros
    // ou string pois ambas as coleções são de elementos comparáveis (==)
    fmt.Println("Mapa de ocorrência dos valores", calculaOcorrencias([]int{1, 2, 3, 1, 2, 1, 1, 3, 4}))
    fmt.Println("Mapa de ocorrência dos valores", calculaOcorrencias([]int{1: 3, 4: 6, 99: 0, 1, 1}))

    fmt.Println("Mapa de ocorrência dos valores", calculaOcorrencias([]string{"um", "texto", "partido", "em", "pedaços"}))
    fmt.Println("Mapa de ocorrência dos valores", calculaOcorrencias([]string{"ab", "bc", "aa", "aa", "bc", "aa"}))

	// Tenta atribuir um valor a uma chave no mapa nulo
	// Isso deve gerar um pânico, já que o mapa não foi inicializado
	// corretamente
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ocorreu um pânico:", r)
		}
	}()
	mVar["chave"] = 1

}

// Uma feature interessante e com grande potencial é a utilização de generics
// como visto no exemplo abaixo.
// Já existem outras features experimentais sendo testadas na linguagem
// utilizando generics, veja em https://github.com/golang/exp/
func calculaOcorrencias[T comparable](colecao []T) map[T]int {
    ocorrencias := make(map[T]int, len(colecao))
    for _, item := range colecao{
        ocorrencias[item]++
    }
    return ocorrencias
}
