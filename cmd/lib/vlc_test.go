package lib

import "testing"

func Test_prepareText(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},

		{
			name: "no one upper",
			str:  "hi i am sasha",
			want: "hi i am sasha",
		},

		{
			name: "upper middle in word",
			str:  "alexAnder g",
			want: "alex!ander g",
		},

		{
			name: "last upper",
			str:  "hello filiP",
			want: "hello fili!p",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "!ted",
			want: "001000100110100101",
		},

		{
			name: "one char test",
			str:  "tttt",
			want: "1001100110011001",
		},

		{
			name: "only spaces",
			str:  "   ",
			want: "111111",
		},

		{
			name: "only uppers",
			str:  "!t!t!t!t",
			want: "0010001001001000100100100010010010001001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str string
		want string
	}{
		{
			name: "base test",
			str: "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}