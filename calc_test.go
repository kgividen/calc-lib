package calc

import "testing"

func TestAddition_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
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
