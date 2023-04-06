package ma

import (
	"encoding/hex"
)

//初始置换 IP 表
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

//逆初始置换 IP^-1 表
var inverse_ip_table = []byte{
	40, 8, 48, 16, 56, 24, 64, 32,
	39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30,
	37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28,
	35, 3, 43, 11, 51, 19, 59, 27,
	34, 2, 42, 10, 50, 18, 58, 26,
	33, 1, 41, 9, 49, 17, 57, 25,
}

//选择置换 PC1 表:将64位密钥置换为56位，并将其分为左右两个部分。
var PC1 = []byte{
	57, 49, 41, 33, 25, 17, 9,
	1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27,
	19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15,
	7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 28, 20, 12, 4,
}

//选择置换 PC2 表:将56位密钥的左右两个部分分别置换为48位密钥子块，以便与扩展置换表的结果进行异或运算。
var PC2 = []byte{
	14, 17, 11, 24, 1, 5,
	3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8,
	16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55,
	30, 40, 51, 45, 33, 48,
	44, 49, 39, 56, 34, 53,
	46, 42, 50, 36, 29, 32,
}

//E置换表
var E_table = []byte{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1,
}

//P盒置换表
var P_table = []byte{
	16, 7, 20, 21,
	29, 12, 28, 17,
	1, 15, 23, 26,
	5, 18, 31, 10,
	2, 8, 24, 14,
	32, 27, 3, 9,
	19, 13, 30, 6,
	22, 11, 4, 25,
}

//左移置换表
var LeftShifts = [16]byte{
	1, 1, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 2, 2, 2, 2, 1,
}

//S盒
var SBoxes = [8][4][16]byte{
	{ //S1盒子
		{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
		{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
		{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
		{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
	},
	{ //S2盒子
		{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
		{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
		{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
		{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
	},
	{ //S3盒子
		{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
		{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
		{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
		{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
	},
	{ //S4盒子
		{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
		{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
		{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
		{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
	},
	{ //S5盒子
		{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
		{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
		{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
		{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
	},
	{ //S6盒子
		{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
		{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
		{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
		{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
	},
	{ //S7盒子
		{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
		{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
		{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
		{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
	},
	{ //S8盒子
		{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
		{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
		{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
		{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
	},
}

func IP(block []byte) (right [32]byte, left [32]byte) {
	for i := 0; i < len(right); i++ {
		right[i] = block[ip_table[i]-1]
	}
	for i := 0; i < len(left); i++ {
		left[i] = block[ip_table[i+32]-1]
	}
	return
}

//将字符串切为64位的切片
func slice(data []byte) [][]byte {
	if len(data)%64 != 0 {
		for i := 0; i < len(data)%64; i++ {
			data = append(data, 0)
		}
	}
	i := 0
	var blocks [][]byte
	for i < len(data) {
		blocks = append(blocks, data[i:i+64])
		i = i + 64
	}
	return blocks
}

//将字符串转换为64个字节的二进制字节块
func string_to_binary(str string) []byte {
	var bytekey []byte = []byte(str)
	var binary []byte
	for _, b := range bytekey {
		for i := 0; i < 8; i++ {
			bit := (b >> uint(7-i)) & 1
			binary = append(binary, byte(bit))
		}
	}
	return binary
}

//去除密钥的奇偶校验位
func key_to_noParity(binarykey []byte) []byte {
	var noParityKey []byte
	for i := 0; i < len(binarykey); i++ {
		if (i+1)%8 != 0 {
			noParityKey = append(noParityKey, binarykey[i])
		}
	}
	return noParityKey
}

//获得第n轮的子密钥
func getkey(n int, binkey []byte) [48]byte {
	var key [56]byte    //去除奇偶校验位的56位的密钥
	var subkey [48]byte //子密钥
	//PC1密钥置换
	for i := 0; i < len(PC1); i++ {
		key[i] = binkey[PC1[i]-1]
	}
	//fmt.Println(key)
	for i := 0; i < n; i++ {
		for j := 0; j < int(LeftShifts[i]); j++ {
			subkey[len(PC2)-int(LeftShifts[i])+j] = key[len(PC1)-int(LeftShifts[i])+j]
			subkey[len(PC2)/2-int(LeftShifts[i])+j] = key[len(PC1)/2-int(LeftShifts[i])+j]
		}
		key = rotateLeft(key, int(LeftShifts[i]))

		for j := 0; j < int(LeftShifts[i]); j++ {
			key[len(PC1)/2+j] = subkey[len(PC2)-int(LeftShifts[i])+j]
			key[j] = subkey[len(PC2)/2-int(LeftShifts[i])+j]
		}
	}
	for i := 0; i < len(PC2); i++ {
		subkey[i] = key[PC2[i]-1]
	}
	//fmt.Println(subkey)
	return subkey
}

//左移函数
func rotateLeft(b [56]byte, n int) [56]byte {
	m := n % len(b)
	if m == 0 {
		return b
	}
	first := b[0]
	for i := 0; i < len(b)-1; i++ {
		b[i] = b[i+1]
	}
	b[len(b)-1] = first
	return b
}

func xor(a, b [32]byte) (result [32]byte) {
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}
	return
}

func desTurn(left, right [32]byte, subkey [48]byte) (newLeft, newRight [32]byte) {
	expandedRight := make([]byte, 48)
	for i := 0; i < 48; i++ {
		expandedRight[i] = right[E_table[i]-1] ^ subkey[i]
	}
	temp := make([]byte, 32)
	for i := 0; i < 8; i++ {
		row := (expandedRight[i*6] << 1) | expandedRight[i*6+5]
		col := (expandedRight[i*6+1] << 3) | (expandedRight[i*6+2] << 2) | (expandedRight[i*6+3] << 1) | expandedRight[i*6+4]
		value := SBoxes[i][row][col]
		//SBoxes[8][4][16]
		temp[i*4] = (value & 0x8) >> 3
		temp[i*4+1] = (value & 0x4) >> 2
		temp[i*4+2] = (value & 0x2) >> 1
		temp[i*4+3] = value & 0x1
	}
	var permuted [32]byte
	for i := 0; i < 32; i++ {
		permuted[i] = temp[P_table[i]-1]
	}

	newLeft = right
	newRight = xor(left, permuted)
	return
}

//交换L和R
func exchange_LR(left, right [32]byte) (Left, Right [32]byte) {
	for i := 0; i < len(right); i++ {
		Right[i] = left[i]
	}
	for i := 0; i < len(left); i++ {
		Left[i] = right[i]
	}
	return
}

//逆初始置换
func IIP(left, right [32]byte) (block [64]byte) {
	for i := 0; i < len(block); i++ {
		if inverse_ip_table[i] <= 32 {
			block[i] = right[inverse_ip_table[i]-1]
		} else {
			block[i] = left[inverse_ip_table[i]-32-1]
		}
	}
	return
}

func binaryToString(binary []byte) string {
	var str string
	for i := 0; i < len(binary); i += 8 {
		var val byte
		for j := 0; j < 8; j++ {
			val = (val << 1) + binary[i+j]
		}
		str += string(val)
	}
	return str
}

//加密程序
func encrypt(key, text string) string {

	plaintext := slice(string_to_binary(text))
	bkey := string_to_binary(key)
	presult := ""
	for _, block := range plaintext {
		left, right := IP(block)
		for i := 0; i < 16; i++ {
			subkey := getkey(i, bkey)
			left, right = desTurn(left, right, subkey)

			if i != 15 {
				left, right = exchange_LR(left, right)
			}
		}
		result := IIP(left, right)
		presult += binaryToString(result[:])
	}
	return hex.EncodeToString([]byte(presult))
}

//解密程序
func decrypt(key, cirphertext string) string {
	text, _ := hex.DecodeString(cirphertext)
	plaintext := slice(string_to_binary(string(text)))
	bkey := string_to_binary(key)
	presult := ""
	for _, block := range plaintext {
		left, right := IP(block)
		for i := 15; i >= 0; i-- {
			subkey := getkey(i, bkey)
			left, right = desTurn(left, right, subkey)
			if i != 0 {
				left, right = exchange_LR(left, right)
			}
		}
		result := IIP(left, right)
		presult += binaryToString(result[:])
	}
	return presult

}
