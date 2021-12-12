//Testing
package main

import "testing"

func TestSum(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 11 {
	// 	t.Errorf("error")
	// }
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{6, 6, 12},
		{25, 25, 50},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Hubo un error")
		}
	}
}
