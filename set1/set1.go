package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"unicode"
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
	var (
		output       []byte
		maxPlainText string
		maxScore     int
	)

	for key := byte('A'); key <= byte('Z'); key++ {
		output = output[:0]
		for _, element := range s_bytes {
			output = append(output, element^key)
		}
		// check if current key gives a higher english plaintext score
		if ScorePlainText(string(output)) > maxScore {
			maxPlainText = string(output)
			maxScore = ScorePlainText(maxPlainText)
		}
	}
	fmt.Println(maxPlainText)
	return maxPlainText
}

// return 0 is string is not ASCII nor printable
// if it is, base score of 5
// add 1 for each space
func ScorePlainText(s string) int {
	score := 5
	for _, r := range s {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			score = 0
			break
		}
		if r == ' ' {
			score += 1
		}
	}
	return score
}
