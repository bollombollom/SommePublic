package sommepublic

import "testing"

func TestSomme(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positifs", 2.5, 3.5, 6.0},
		{"négatifs", -1.5, -2.5, -4.0},
		{"mixte", 5.0, -2.0, 3.0},
		{"zéros", 0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Somme(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Somme(%f, %f) = %f, expected %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
