package mean

import (
	"fmt"
	"testing"
)

func TestMean(t *testing.T) {
	x := []float64{1.0, 2.0, 3.0}
	want := 2.0
	if got1 := Mean(x...); got1 != want {
		t.Errorf("got: Mean(%v) = %v, want: %v", x, got1, want)
	}

	if got2 := Mean(1, 2, 3); got2 != want {
		t.Errorf("got: Mean(%v, %v, %v) = %v, want: %v", 1, 2, 3, got2, want)
	}

	want2 := 0.0
	if got3 := Mean(1.0, 2, -3); got3 != want2 {
		t.Errorf("got: Mean(%v, %v, %v) = %v, want: %v", 1.0, 2, -3, got3, want2)
	}
}

func ExampleMean() {
	nums := []float64{1.0, 2.0, 3.0}
	fmt.Println(Mean(nums...))
	fmt.Println(Mean(1, 2, 3))
	fmt.Println(Mean(1.0, 2, -3))
	// Output:
	// 2
	// 2
	// 0
}

func BenchmarkMean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mean(1, 2, 3)
	}
}
