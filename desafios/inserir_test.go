package main

import "testing"

func TestExamploInserir1(t *testing.T) {
	alvo := 5
	esperado := 2
	nums := []int{1, 3, 5, 6}

	obtido := pesquisarInserir(nums, alvo)
	if obtido != esperado {
		t.Errorf("Esperado %v, mas obtido %v", esperado, obtido)
	}
}

func TestExamploInserir2(t *testing.T) {
	alvo := 2
	esperado := 1
	nums := []int{1, 3, 5, 6}

	obtido := pesquisarInserir(nums, alvo)
	if obtido != esperado {
		t.Errorf("Esperado %v, mas obtido %v", esperado, obtido)
	}
}

func TestExamploInserir3(t *testing.T) {
	alvo := 7
	esperado := 4
	nums := []int{1, 3, 5, 6}

	obtido := pesquisarInserir(nums, alvo)
	if obtido != esperado {
		t.Errorf("Esperado %v, mas obtido %v", esperado, obtido)
	}
}
