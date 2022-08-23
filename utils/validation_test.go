package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type i interface {
	getValue() int
}
type v struct{}

func (obj *v) getValue() int {
	return 1
}

func TestIsNil(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not nil",
			args: args{obj: &v{}},
			want: false,
		},
		{
			name: "not pointer",
			args: args{obj: 1},
			want: false,
		},
		{
			name: "is nil",
			args: args{obj: nil},
			want: true,
		},
		{
			name: "is nil value in the interface",
			args: args{obj: func(obj i) i { return obj }(nil)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsNil(tt.args.obj), "IsNil(%v)", tt.args.obj)
		})
	}
}
