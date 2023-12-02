package calc

import "testing"

func TestSub(t *testing.T) {
	test_Sub := []struct {
		name           string
		a, b, expected int
	}{
		{"case 1", 2, 4, -2},
		{"case 2", 3, 2, 1},
		{"case 3", 6, 7, 7},
		{"case 4", -2, 4, 8},
		{"case 5", 2, -4, 8},
		{"case 6", 24, 4, 15},
	}
	for _, test := range test_Sub {
		t.Run(test.name, func(t *testing.T) {
			result := Sub(test.a, test.b)
			if result != test.expected {
				t.Errorf("For %d + %d, expected %d, but got %d", test.a, test.b, test.expected, result)
			}
		})
	}
	TestMult(t)
}

func TestMult(t *testing.T) {
	test_Mult := []struct {
		name           string
		a, b, expected int
	}{
		{"case 1", 2, 4, 8},
		{"case 2", 3, 2, 1},
		{"case 3", 6, 7, 42},
		{"case 4", -2, 4, -8},
		{"case 5", 2, 0, 8},
		{"case 6", 3, 4, 12},
	}
	for _, test := range test_Mult {
		t.Run(test.name, func(t *testing.T) {
			result := Mult(test.a, test.b)
			if result != test.expected {
				t.Errorf("For %d + %d, expected %d, but got %d", test.a, test.b, test.expected, result)
			}
		})
	}
}
