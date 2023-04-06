package main

import "fmt"

var ip_table = []byte{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

func IP(plainText string) (right [32]byte, left [32]byte) {
	for i := 0; i < len(right); i++ {
		right[i] = plainText[ip_table[i]-1]
	}
	for i := 0; i < len(left); i++ {
		left[i] = plainText[ip_table[i+32]-1]
	}
	fmt.Println(right)
	fmt.Println(left)
	return
}

func slice(text string) [][]byte {
	var data []byte = []byte(text)
	if len(data)%64 != 0 {
		for i := 0; i < len(data)%64; i++ {
			data = append(data, 0)
		}
	}
	i := 0
	var blocks [][]byte
	for i < len(data) {
		if i-len(data) > 64 {
			blocks = append(blocks, data[i:i+64])
		} else {
			blocks = append(blocks, data[i:i+64])
		}
		i = i + 64
	}
	return blocks
}
func main() {
	text := "This is a sample text for splitting into 64-byte chunks using Go."
	IP(text)
}
