package main

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestGetShiftOverlap(t *testing.T) {
	t.Parallel()

	valorEsperado := 2
	l1 := Line{2, 5}
	l2 := Line{3, 6}

	valor := l1.getShiftOverlap(l2)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestLength(t *testing.T) {
	t.Parallel()

	valorEsperado := 3
	l2 := Line{1, 4}

	valor := l2.Length()

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestGetContainedOverlap(t *testing.T) {
	t.Parallel()

	valorEsperado := 1
	l1 := Line{2, 3}
	l2 := Line{1, 4}

	valor := l1.getContainedOverlap(l2)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestGetContainedOverlapSameLength(t *testing.T) {
	t.Parallel()

	valorEsperado := 3
	l1 := Line{1, 4}
	l2 := Line{1, 4}

	valor := l1.getContainedOverlap(l2)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestNoOverLap(t *testing.T) {
	t.Parallel()

	valorEsperado := 0
	l1 := Line{1, 2}
	l2 := Line{2, 3}

	valor := l1.GetOverlap(l2)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestSortLines(t *testing.T) {
	t.Parallel()

	l1 := Line{3, 6}
	l2 := Line{2, 5}
	l3 := Line{1, 3}
	valorEsperado := []Line{l3, l2, Line{3, 6}}

	lines := Lines{l1, l2, l3}

	lines.sortByStart()

	if !lines.Equals(valorEsperado) {
		t.Errorf(erroPadrao, valorEsperado, lines)
	}
}

func TestLineNotEquals(t *testing.T) {
	t.Parallel()

	l1 := Line{2, 3}
	l2 := Line{1, 3}
	valorEsperado := false

	if l1.Equals(l2) {
		t.Errorf(erroPadrao, valorEsperado, valorEsperado)
	}
}

func TestLineEquals(t *testing.T) {
	t.Parallel()

	l1 := Line{1, 3}
	l2 := Line{1, 3}
	valorEsperado := false

	if !l1.Equals(l2) {
		t.Errorf(erroPadrao, valorEsperado, valorEsperado)
	}
}

func TestAddUniqueLine(t *testing.T) {
	t.Parallel()

	l1 := Line{1, 3}
	l2 := Line{1, 3}
	lines := Lines{}

	added := lines.addIfNotExists(l1)
	if added == false {
		t.Errorf(erroPadrao, true, added)
	}

	added = lines.addIfNotExists(l2)
	if added == true {
		t.Errorf(erroPadrao, false, added)
	}

	if len(lines) != 1 {
		t.Errorf(erroPadrao, 1, len(lines))
	}

}
