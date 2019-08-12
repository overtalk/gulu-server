package parse

import (
	"encoding/json"
	"testing"
)

func TestInt(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"int.(type) == json.Number",
			args{
				json.Number("1"),
			},
			1,
		},
		{
			"int.(type) == string",
			args{
				"1",
			},
			1,
		},
		{
			"int == \"\"",
			args{
				"",
			},
			0,
		},
		{
			"int == \"a\"",
			args{
				"a",
			},
			0,
		},
		{
			"int.(type) == int",
			args{
				int(1),
			},
			1,
		},
		{
			"int.(type) == int32",
			args{
				int32(1),
			},
			1,
		},
		{
			"int.(type) == int64",
			args{
				int64(1),
			},
			1,
		},
		{
			"int.(type) == uint",
			args{
				uint(1),
			},
			1,
		},
		{
			"int.(type) == uint32",
			args{
				uint32(1),
			},
			1,
		},
		{
			"int.(type) == uint64",
			args{
				uint64(1),
			},
			1,
		},
		{
			"int.(type) == float64",
			args{
				float64(1),
			},
			1,
		},
		{
			"int.(type) == nil",
			args{
				nil,
			},
			0,
		},
		{
			"int.(type) == other",
			args{
				map[interface{}]interface{}{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.in); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
