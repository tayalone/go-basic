package realworld

import (
	"fmt"
	"testing"
)

type fakeRepo struct{}

func (f fakeRepo) QueryLang(i uint) Langeuage {
	return Langeuage{ID: 1, Name: "GO"}
}

func TestHandler_Do(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		id uint
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			fields: fields{
				repo: fakeRepo{},
			},
			args: args{
				id: 1,
			},
			want: fmt.Sprintf("%s language", "GO"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				repo: tt.fields.repo,
			}
			if got := h.Do(tt.args.id); got != tt.want {
				t.Errorf("Handler.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
