package lib

import (
	"strings"
	"unicode"
)

func Encode(str string) string {
	textWhithOutUpper := prepareText(str)
	
	binaryString := encodeBin(textWhithOutUpper)

	chunks := splitByChunks(binaryString)

	res := chunks.ToHex()

	return res.ToString()
}

func Decode(encodeText string) string {

	hChunks := NewHexChunks(encodeText)

	bChunks := hChunks.ToBinary()

	bString := bChunks.Join()

	dTree := getEncodingTable().DecodingTree() 

	res := dTree.Decode(bString)

	return exportText(res)
}

func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch){
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func exportText(str string) string {
	var buf strings.Builder

	var isCapital bool

	for _, ch := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false 
			continue
		}
		if ch == '!'{
			isCapital = true
			continue
		}else{
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func encodeBin (str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("uncknown symbol: " + string(ch))
	}

	return res
}

func getEncodingTable() EncodingTable {
	return EncodingTable{
		' ': "11",

		't': "1001",

		'n': "10000",

		's': "0101",

		'r': "01000",

		'd': "00101",

		'!': "001000",

		'c': "000101",

		'm': "000011",

		'g': "0000100",

		'b': "0000010",

		'v': "00000001",

		'k': "0000000001",

		'q': "0000000000001",

		'e': "101",

		'o': "10001",

		'a': "011",
		
		'i': "01001",

		'h': "0011",

		'l': "001001",

		'u': "00011",

		'f': "000100",

		'p': "0000101",

		'w': "0000011",

		'y': "0000001",

		'j': "0000000001",

		'x': "0000000000001",

		'z': "00000000000000",
	}
}



