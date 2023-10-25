package evenorodd

import "testing"


func TestOddNumber(t *testing.T) {
    input := 5
    want := "Odd"
    output := EvenOrOdd(input)
    if output != want {
        t.Fatalf("Le fail")
    }
}

func TestEvenNumber(t *testing.T) {
    input := 8
    want := "Even"
    output := EvenOrOdd(input)
    if output != want {
        t.Fatalf("Le fail")
    }
}