package parse

import "testing"

func TestString(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test string input",
			args{
				"st",
			},
			"st",
		},
		{
			"test []uint input",
			args{
				[]uint8{116, 101, 115, 116},
			},
			"test",
		},
		{
			"test int64 input",
			args{
				116,
			},
			"116",
		},
		{
			"test int input",
			args{
				116,
			},
			"116",
		},
		{
			"test nil input",
			args{
				nil,
			},
			"",
		},
		{
			"test other input",
			args{
				make(map[interface{}]interface{}),
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.in); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
