package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Hextobase64(s string) string {
	raw_bytes, _ := hex.DecodeString(s)
	base64_result := base64.StdEncoding.EncodeToString(raw_bytes)
	return base64_result
}

func FixedXor(s1 string, s2 string) string {
	s1_bytes, _ := hex.DecodeString(s1)
	s2_bytes, _ := hex.DecodeString(s2)
	var xor_output []byte
	for i := 0; i < len(s1_bytes); i++ {
		xor_output = append(xor_output, s1_bytes[i]^s2_bytes[i])
	}
	return hex.EncodeToString(xor_output)
}

func SingleByteXorCipher(s string) string {
	s_bytes, _ := hex.DecodeString(s)
	var output []byte
	for key := byte('A'); key <= byte('Z'); key++ {
		output = output[:0]
		for i := 0; i < len(s_bytes); i++ {
			output = append(output, s_bytes[i]^key)
		}
		fmt.Println(string(output))
	}
	return string(output)
}
