package usertest

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name   string
		bdyear uint
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "Good",
			args: args{
				name:   "Иван",
				bdyear: 1999,
			},
			want: &usr{
				Name:   "Иван",
				BDYear: 1999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.name, tt.args.bdyear); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
