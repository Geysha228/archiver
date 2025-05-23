package lib

import (
	"reflect"
	"testing"
)

func TestEncodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   EncodingTable
		want DecodingTree
	}{
		{
			name: "base tree test",
			et: EncodingTable{
				'a':"11",
				'b':"1001",
				'z':"0101",
			},
			want: DecodingTree{
				Zero: &DecodingTree{
					One: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: "z",
							},
						},
					},
				},
				One: &DecodingTree{
					Zero: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: "b",
							},
						},
					},
					One: &DecodingTree{
						Value: "a",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodingTable.DecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
