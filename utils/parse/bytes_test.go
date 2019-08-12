package parse

import (
	"reflect"
	"testing"
)

func TestBytes(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"in.(type) == []byte",
			args{
				[]byte("test data"),
			},
			[]byte("test data"),
		},
		{
			"in.(type) == string",
			args{
				"test data",
			},
			[]byte("test data"),
		},
		{
			"in.(type) == nil",
			args{
				nil,
			},
			nil,
		},
		{
			"in.(type) == int",
			args{
				int(1),
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
