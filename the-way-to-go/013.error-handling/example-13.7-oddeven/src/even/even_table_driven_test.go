package even
import (
    "testing"
)

var casesForEven = []struct {
    in int
    out bool
}{
    {10, true},
    {9, false},
    {8, true},
}

var casesForOdd = []struct {
    in int
    out bool
}{
    {10, false},
    {9, true},
    {8, false},
}

func TestEvenBatch(t *testing.T) {
    for i, tt := range casesForEven {
        if s := Even(tt.in); s != tt.out {
            t.Errorf("%d. %d => %t, wanted: %t", i, tt.in, s, tt.out)
        }
    }
}
func TestOddBatch(t *testing.T) {
    for i, tt := range casesForOdd {
        if s := Odd(tt.in); s != tt.out {
            t.Errorf("%d. %d => %t, wanted: %t", i, tt.in, s, tt.out)
        }
    }
}
