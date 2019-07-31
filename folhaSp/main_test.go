package main

import (
	"testing"
)

func TestTargetStar(t *testing.T) {
	planets := []string{"SIRIUS", "LALANDE", "PROCION", "ALPHA CENTAURI", "BARNARD"}

	expected := "PROCION"
	engineers := 75
	targetStart := targetStart(engineers, planets)
	if targetStart != expected {
		t.Errorf("targetStar was %s, but it was expeceted %s", targetStart, expected)
	}
}
func TestFibo1(t *testing.T) {

	n := 1
	expected := 1
	result := fibo(n)
	if expected != result {
		t.Errorf("Fibo(%d) was %d, but it was expeceted %d", n, result, expected)
	}
}

func TestFibo2(t *testing.T) {

	n := 2
	expected := 1
	result := fibo(n)
	if expected != result {
		t.Errorf("Fibo(%d) was %d, but it was expeceted %d", n, result, expected)
	}
}

func TestFibo3(t *testing.T) {

	n := 3
	expected := 2
	result := fibo(n)
	if expected != result {
		t.Errorf("Fibo(%d) was %d, but it was expeceted %d", n, result, expected)
	}
}

func TestFibo4(t *testing.T) {

	n := 4
	expected := 3
	result := fibo(n)
	if expected != result {
		t.Errorf("Fibo(%d) was %d, but it was expeceted %d", n, result, expected)
	}
}

func TestFibo5(t *testing.T) {

	n := 5
	expected := 5
	result := fibo(n)
	if expected != result {
		t.Errorf("Fibo(%d) was %d, but it was expeceted %d", n, result, expected)
	}
}

func TestFibo6(t *testing.T) {

	n := 6
	expected := 8
	result := fibo(n)
	if expected != result {
		t.Errorf("Fibo(%d) was %d, but it was expeceted %d", n, result, expected)
	}
}
