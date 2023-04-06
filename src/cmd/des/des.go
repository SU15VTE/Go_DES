package des

import (
	"encoding/hex"
	"fmt"
)

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
func decrypt(key, cirphertext string) {
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
	fmt.Println("Plaintext")
	fmt.Println(presult)

}
func main() {
	key := "SU15VTE!"
	text := "Hello World!"
	ciphertext := encrypt(key, text)
	fmt.Println(ciphertext)
	decrypt(key, "ae037584bfca3946b7af40041d3e7419")
}
