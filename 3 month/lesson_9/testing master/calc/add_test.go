package calc

import "testing"

// func TestAdd(t *testing.T) {
// 	var a, b = 2, 3
// 	output := Add(a, b)

// 	if output != 5 {
// 		t.Errorf("Input Add(%d, %d) got:%d waiting:%d", a, b, output, a+b)
// 	}
// }

func TestAdd(t *testing.T) {

	tests := []struct {
		name           string
		a, b, expected int
	}{
		{"case 1", 2, 3, 5},
		{"case 2", 0, 0, 0},
		{"case 3", -1, 1, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("For %d + %d, expected %d, but got %d", test.a, test.b, test.expected, result)
			}
		})
	}
}
