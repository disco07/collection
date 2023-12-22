package collections

import (
	"golang.org/x/exp/constraints"
)

type Num interface {
	constraints.Float | constraints.Integer
}

// Number est une structure générique qui contient une valeur de n'importe quel type numérique.
type Number[T Num] struct {
	value T
}

// NewNumber crée une nouvelle instance de Number avec des fonctions d'addition et de soustraction spécifiques.
func NewNumber[T Num](value T) Number[T] {
	return Number[T]{value}
}

// Add additionne la valeur de Number avec une autre valeur et retourne le résultat.
func (n *Number[T]) Add(other T) T {
	return add(n.value, other)
}

// Sub soustrait une valeur de la valeur de Number et retourne le résultat.
func (n *Number[T]) Sub(other T) T {
	return sub(n.value, other)
}

// Mul multiplie une valeur de la valeur de Number et retourne le résultat.
func (n *Number[T]) Mul(other T) T {
	return mul(n.value, other)
}

// Div divise une valeur de la valeur de Number et retourne le résultat.
func (n *Number[T]) Div(other T) T {
	return div(n.value, other)
}

func (n *Number[T]) GreaterThan(other T) bool {
	return n.value > other
}

func (n *Number[T]) LessThan(other T) bool {
	return n.value < other
}

func add[T Num](a, b T) T {
	return a + b
}

func sub[T Num](a, b T) T {
	return a - b
}

func div[T Num](a, b T) T {
	return a / b
}

func mul[T Num](a, b T) T {
	return a * b
}
