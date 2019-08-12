package parse

import (
	"encoding/json"
	"testing"
)

func TestFloat(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test json.Number input",
			args{
				json.Number("1"),
			},
			1.0,
		},
		{
			"test string input",
			args{
				"1",
			},
			1.0,
		},
		{
			"test nil string",
			args{
				"",
			},
			0.0,
		},
		{
			"test \"a\" string",
			args{
				"a",
			},
			0.0,
		},
		{
			"test int input",
			args{
				1,
			},
			1.0,
		},
		{
			"test int32 input",
			args{
				int32(1),
			},
			1.0,
		},
		{
			"test int64 input",
			args{
				int64(1),
			},
			1.0,
		},
		{
			"test uint input",
			args{
				uint(1),
			},
			1.0,
		},
		{
			"test uint32 input",
			args{
				uint32(1),
			},
			1.0,
		},
		{
			"test uint64 input",
			args{
				uint64(1),
			},
			1.0,
		},
		{
			"test float64 input",
			args{
				float64(1),
			},
			1.0,
		},
		{
			"test nil input",
			args{
				nil,
			},
			0,
		},
		{
			"test other type input",
			args{
				make(map[interface{}]interface{}, 10),
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float(tt.args.in); got != tt.want {
				t.Errorf("Float() = %v, want %v", got, tt.want)
			}
		})
	}
}
