package lib

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

type HexChunks []HexChunk

type HexChunk string

type EncodingTable map[rune]string

const hexChunkSeparator = " "

func (hxChs HexChunks) ToString() string {

	switch len(hxChs) {
	case 0:
		return ""
	case 1:
		return string(hxChs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hxChs[0]))

	for _, ch := range hxChs[1:] {
		buf.WriteString(hexChunkSeparator)
		buf.WriteString(string(ch))
	}
	return buf.String()
}

func (bnChs BinaryChunks) ToHex() HexChunks {

	hexSize := len(bnChs)

	hexChunks := make(HexChunks, 0, hexSize)

	for _, chunk := range bnChs {
		hexChunks = append(hexChunks, chunk.ToHex())
	}

	return hexChunks
}

func (bnCh BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bnCh), 2, 8)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}
	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

func splitByChunks(binaryString string) BinaryChunks{
	strLen := utf8.RuneCountInString(binaryString)
	chunksCount := strLen / 8

	if strLen / 8 != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, ch := range binaryString {
		buf.WriteString(string(ch))

		if (i+1) % 8 == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()

		lastChunk += strings.Repeat("0", 8-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}

func NewHexChunks (str string) HexChunks {
	parts := strings.Split(str, hexChunkSeparator)

	res := make(HexChunks, 0, len(parts))

	for _, part := range parts {
		res = append(res, HexChunk(part))
	}

	return res
}