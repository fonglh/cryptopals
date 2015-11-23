package set1

import (
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"strings"
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

func SingleByteXorCipher(s string) (string, int) {
	s_bytes, _ := hex.DecodeString(s)
	var (
		output       []byte
		maxPlainText string
		maxScore     int
	//	maxKey       string
	)

	for key := 0; key < 256; key++ {
		output = output[:0]
		for _, element := range s_bytes {
			output = append(output, element^byte(key))
		}
		score := ScorePlainText(string(output))

		// check if current key gives a higher english plaintext score
		if score > maxScore {
			maxPlainText = string(output)
			maxScore = ScorePlainText(maxPlainText)
			//		maxKey = string(key)
		}
	}
	return maxPlainText, maxScore
}

// add 1 point for each space
func ScorePlainText(s string) int {
	score := 0
	for _, r := range s {
		if r > unicode.MaxASCII {
			score = 0
			break
		} else if r == ' ' {
			score += 1
		} /*else if r > unicode.MaxASCII /*|| (!unicode.IsPrint(r) && (r != '\r' || r != '\n' || r != '\t'))*/
	}
	return score
}

// helper function to read file and return lines as array of strings
func ReadStrings(filename string) ([]string, error) {
	fileContents, err := ioutil.ReadFile(filename)
	fileStrings := string(fileContents)
	fileStringsArray := strings.Fields(fileStrings)
	return fileStringsArray, err
}

// Answer is "Now that the party is jumping"
// key is 53, in stringsArray[170]
func SingleCharXor(filename string) string {
	stringsArray, _ := ReadStrings(filename)
	var (
		maxScore     int
		maxPlaintext string
	)

	for _, str := range stringsArray {
		plaintext, score := SingleByteXorCipher(str)
		if score > maxScore {
			maxScore = score
			maxPlaintext = plaintext
		}
	}
	return maxPlaintext
}

// Abstract repeat key byte XOR so both encrypt and decrypt can use it
func ByteRepeatKeyXor(input []byte, key string) (output []byte) {
	var keyPtr int
	for i := range input {
		output = append(output, input[i]^key[keyPtr])
		keyPtr = (keyPtr + 1) % len(key)
	}
	return
}

// Set 1 Challenge 5
func RepeatKeyXor(plaintext, key string) string {
	inputBytes := []byte(plaintext)
	output := ByteRepeatKeyXor(inputBytes, key)
	return hex.EncodeToString(output)
}

// Inverse of Set 1 Challenge 5
func DecryptRepeatKeyXor(ciphertext, key string) string {
	ciphertextBytes, _ := hex.DecodeString(ciphertext)
	output := ByteRepeatKeyXor(ciphertextBytes, key)
	return string(output)
}

// Calculate Hamming Distance (number of differing bits) between 2 strings of equal length
func HammingDistance(s1, s2 string) (result int) {
	for i := range s1 {
		diffBits := s1[i] ^ s2[i]
		var mask byte
		// when byte overflows, it becomes 0
		for mask = 1; mask > 0; mask = mask << 1 {
			if (mask & diffBits) != 0 {
				result++
			}
		}
	}

	return
}
