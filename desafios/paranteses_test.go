package main

import "testing"

func TestExamploParenteses1(t *testing.T) {
	entrada := "()"
	esperado := true

	obtido := ehValido(entrada)
	if obtido != esperado {
		t.Errorf("Esperado %v, mas obtido %v", esperado, obtido)
	}
}

func TestExamploParenteses2(t *testing.T) {
	entrada := "()[]{}"
	esperado := true

	obtido := ehValido(entrada)
	if obtido != esperado {
		t.Errorf("Esperado %v, mas obtido %v", esperado, obtido)
	}
}

func TestExamploParenteses3(t *testing.T) {
	entrada := "(]"
	esperado := false

	obtido := ehValido(entrada)
	if obtido != esperado {
		t.Errorf("Esperado %v, mas obtido %v", esperado, obtido)
	}
}
