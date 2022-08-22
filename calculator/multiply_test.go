package calculator

import (
	testing "testing"

	"github.com/golang/mock/gomock"
)

func TestMultiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				a: 10,
				b: 20,
			},
			want: 200,
		},
		{
			name: "case 2",
			args: args{
				a: 30,
				b: 40,
			},
			want: 1200,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum = func(a int, b int) int {
				return tt.args.b
			}

			if got := Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Log(got)
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
