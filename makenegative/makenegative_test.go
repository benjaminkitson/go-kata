package makenegative

import "testing"

func TestPositive(t *testing.T) {
    input := 5
    want := -5
    output := MakeNegative(input)
    if output != want {
        t.Fatalf("Le fail")
    }
}

func TestNegative(t *testing.T) {
    input := -8
    want := -8
    output := MakeNegative(input)
    if output != want {
        t.Fatalf("Le fail")
    }
}

func TestZero(t *testing.T) {
    input := 0
    want := 0
    output := MakeNegative(input)
    if output != want {
        t.Fatalf("Le fail")
    }
}
