package main

import "testing"

func TestExamploDup1(t *testing.T) {
	esperado := true
	nums := []int{1, 2, 3, 1}

	obtido := contemDuplicado(nums)
	if obtido != esperado {
		t.Errorf("Esperado %v, mas obtido %v", esperado, obtido)
	}
}

func TestExamploDup2(t *testing.T) {
	esperado := false
	nums := []int{1, 2, 3, 4}

	obtido := contemDuplicado(nums)
	if obtido != esperado {
		t.Errorf("Esperado %v, mas obtido %v", esperado, obtido)
	}
}

func TestExamploDup3(t *testing.T) {
	esperado := true
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	obtido := contemDuplicado(nums)
	if obtido != esperado {
		t.Errorf("Esperado %v, mas obtido %v", esperado, obtido)
	}
}
