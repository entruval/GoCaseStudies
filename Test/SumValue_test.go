package sumvalue

import "testing"

func TestSumValue(t *testing.T) {
	a := 100
	b := 100

	value := SumValue(a, b)
	if value != 200 {
		t.Error("Expected 200, got ", value)
	}
}

func BenchmarkSumValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumValue(100, 100)
	}
}
