package lib

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		binaryString string
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "base test",
			args: args{
				binaryString: "001000100110100101",
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.binaryString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name  string
		bnChs BinaryChunks
		want  HexChunks
	}{
		{
			name:  "base test",
			bnChs: BinaryChunks{"0101111", "10000000"},
			want:  HexChunks{"2F", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bnChs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryChunks.ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name string
		str string
		want HexChunks
	}{
		{
			name: "base test",
			str: "20 30 3C 18",
			want: HexChunks{"20", "30", "3C", "18"},
		},

		{
			name: "only one hex",
			str: "3C",
			want: HexChunks{"3C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
