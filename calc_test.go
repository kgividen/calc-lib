package calc

import "testing"

func TestAddition_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: 1},
		{a: 1, b: 2, want: 3},
	}
	for _, tt := range tests {
		t.Run("name", func(t *testing.T) {
			this := &Addition{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 1, b: 0, want: 1},
		{a: 2, b: 1, want: 1},
		{a: 3, b: 1, want: 2},
	}
	for _, tt := range tests {
		t.Run("name", func(t *testing.T) {
			this := &Subtraction{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 1, b: 3, want: 3},
		{a: 2, b: 3, want: 6},
		{a: 3, b: 4, want: 12},
	}
	for _, tt := range tests {
		t.Run("name", func(t *testing.T) {
			this := &Multiplication{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 3, b: 1, want: 3},
		{a: 6, b: 2, want: 3},
		{a: 12, b: 3, want: 4},
	}
	for _, tt := range tests {
		t.Run("name", func(t *testing.T) {
			this := &Division{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
