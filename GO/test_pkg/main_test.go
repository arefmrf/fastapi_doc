package test_pkg

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	output := iGenericAdd(1, 2)
	if output != 3 {
		t.Errorf("Sum function was incorrect, got: %d, want: %d.", output, 3)
	}
	// multiple tests
	var tests = []struct {
		name       string
		a, b, want int
	}{
		{"int", 1, 2, 3},
		{"int", 3, 4, 7},
		{"int", 5, 6, 11},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := iGenericAdd(test.a, test.b)
			if test.want != got {
				t.Errorf("Sum function was incorrect, got: %d, want: %d.", got, test.want)
			}
		})
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iGenericAdd(1, 2)
	}
}

// we can use TestMain to have control on before and after
func TestMain(m *testing.M) {
	fmt.Println("before starting . . .")
	m.Run()
	fmt.Println("after ending . . .")
}
