package main

import (
	"fmt"
)

func main() {
	// Estruturas (structs) é um tipo de dado agragado que agrupa zero ou mais valores nomeados de tipos quiasquer como uma única entidade.
	// Definição retirada do excelente livro: https://www.amazon.com.br/Linguagem-Programa%C3%A7%C3%A3o-Go-Alan-Donovan/dp/8575225464

	// A declaração de uma estrutura é feita da seguinte maneira:
	type pessoa struct {
		nome  string
		idade int
	}

	// E a criação de uma instância pode ser feita de várias formas

	// Utilizando var
	var adriano pessoa
	// Nesse caso os campos da estrutura são preenchidos com os valores zero de cada campo
	// string vazia para nome e 0 para idade

	// Utilizando new
	pedro := new(pessoa)
	// Assim como o anterior, valores zero são atribuídos a estrutura.
	// Porém o retorno é um pnteiro para a estrutura
	// equivalente var pedro *pessoa

	// De maneira literal
	maria := &pessoa{}

	// Sendo que a maneira literal pode ser feita com campos nomeados ou não
	debora := pessoa{"Debora", 33}
	// O & indica que será criado um ponteiro para estrutura
	ana := &pessoa{nome: "Ana", idade: 30}

	// Os campos individuais podem ser acessados através da notação de ponto
	adriano.nome = "Adriano"
	adriano.idade = 40

	// Os ponteiros campos dos ponteiros das estrturas podem ser acessados da mesma maneira
	pedro.nome = "Pedro"
	pedro.idade = 35

	maria.nome = "Maria"
	maria.idade = 25

	// Algumas maneiras de imprimir uma estrutura:
	fmt.Printf("A representação de uma pessoa: %v\n", adriano)
	fmt.Printf("A representação de uma pessoa e seus campos: %+v\n", *pedro)
	fmt.Printf("A representação na sintaxe Go de uma pessoa: %#v\n", *maria)

	// Esta função apresenta um problema, consegue dizer qual é?
	incrementaIdadeIncorreta := func(p pessoa) {
		p.idade++
	}
	incrementaIdadeIncorreta(debora)
	fmt.Println("Idade da Debora após incremento errado deveria ser 34, mas o encontrado foi: ", debora)
	// Para alterar um campo de uma estrutura devemos ter a referência da estrutura
	incrementaIdade := func(p *pessoa) {
		p.idade++
	}
	fmt.Println("Idade da Debora antes do incremento é: ", debora.idade)
	fmt.Println("Idade da Ana antes do incremento é: ", ana.idade)
	incrementaIdade(ana)
	incrementaIdade(&debora)
	fmt.Println("Idade da Debora após incremento é: ", debora.idade)
	fmt.Println("Idade da Ana após incremento é: ", ana.idade)

	// A comparação entre duas estrturas é feita campo a campo
	outroAdriano := pessoa{"Adriano", 40}
	fmt.Println("Os Adrianos são iguais sem sobrenome e caso tenham mesma idade? ", adriano == outroAdriano)
	// Tenha cuidado ao comparar ponteiros de estruturas
	fmt.Println("Os Pedros são iguais sem sobrenome e caso tenham mesma idade? ", pedro == &pessoa{"Pedro", 35})
	// As estruturas são iguais, mas os ponteiros não

	type empregado struct {
		// A estrutura pessoa está incluída na estrutura de um empregado
		// o campo pessoa é anônimo, ou seja, não possui um nome vinculado a ele
		pessoa
		ID    int
		cargo string
	}

	lider := empregado{
		pessoa: pessoa{nome: "Marcio", idade: 41},
		ID:     1,
		cargo:  "Líder da Squad",
	}

	liderado := empregado{
		pessoa: pessoa{nome: "Jean", idade: 38},
		ID:     2,
		cargo:  "Engenheiro de Software",
	}

	fmt.Printf("A pessoa empregada com cargo de liderança: %+v\n", lider)
	fmt.Printf("A pessoa empregada sem cargo de liderança: %+v\n", liderado)

	// Campos de estrturas incluídas em outras podem ser acessados dieretamente por notação de ponto.
	// No exemplo abaixo temos o equivalente a lider.pessoa.nome
	fmt.Println("O nome da pessoa empregada com cargo de liderança: ", lider.nome)
	fmt.Println("O nome da pessoa empregada sem cargo de liderança: ", liderado.nome)
	// Ainda que a estrutura empregado não possua o campo nome, como ela inclui a estrtura pessoa, podemos acessar o campo diretamente.

	// Estrutura podem ter métodos vinculados a elas.
	// No exemplo abaixo vemos a estrutura Contador, que apresenta o campo n e dois métodos, N e incrementar.
	contador := &Contador{}
	contador.incrementar()
	fmt.Println("Após incremento o contador possui valor: ", contador.N())

	// Métodos vinculados a estrturas também podem ser acessados diretamente em estruturas incluidas
	vendasPorMinuto := &Metrica{
		Descricao: "Vendas por Minuto",
	}
	// Os métodos de estruturas incluídas em outra podem ser acessadas diretamente
	vendasPorMinuto.incrementar()
	fmt.Println("O número de vendas por minuto agora é: ", vendasPorMinuto.N())
}

type Contador struct{ n int }

// O método N apenas retorna o campo n
func (c Contador) N() int { return c.n }

// O vínculo aqui acontece a um ponteiro de contador, pois é necessário modificar o valor do campo
func (c *Contador) incrementar() { c.n++ }

// A estrutura Metrica, inclui um contador e seus comportamentos. Além de uma descrição.
type Metrica struct {
	Contador
	Descricao string
}
