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

		{
			name: "not enough numbers",
			args: args{
				binaryString: "001101",
			},
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
		str  string
		want HexChunks
	}{
		{
			name: "base test",
			str:  "20 30 3C 18",
			want: HexChunks{"20", "30", "3C", "18"},
		},

		{
			name: "only one hex",
			str:  "3C",
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

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hxCh HexChunk
		want BinaryChunk
	}{
		{
			name: "base test",
			hxCh: HexChunk("2F"),
			want: BinaryChunk("00101111"),
		},
		{
			name: "base test with numbers",
			hxCh: HexChunk("80"),
			want: BinaryChunk("10000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hxCh.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexChunk.ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name  string
		hxChs HexChunks
		want  BinaryChunks
	}{
		{
			name:  "base test",
			hxChs: HexChunks{"2F", "80"},
			want:  BinaryChunks{"00101111", "10000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hxChs.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexChunks.ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		{
			name: "base test",
			bcs: BinaryChunks{"01001111", "10000000"},
			want: "0100111110000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("BinaryChunks.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
