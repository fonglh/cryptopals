package set1

import "testing"

func TestHextobase64(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
	}

	for _, c := range cases {
		got := Hextobase64(c.in)
		if got != c.want {
			t.Errorf("Hextobase64(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFixedXor(t *testing.T) {
	cases := []struct {
		in1, in2, want string
	}{
		{"1c0111001f010100061a024b53535009181c",
			"686974207468652062756c6c277320657965",
			"746865206b696420646f6e277420706c6179"},
	}

	for _, c := range cases {
		got := FixedXor(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Fixedxor(%q, %q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}

func TestSingleByteXorCipher(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			"Cooking MC's like a pound of bacon"},
	}

	for _, c := range cases {
		got, _ := SingleByteXorCipher(c.in)
		if got != c.want {
			t.Errorf("SingleByteXorCipher(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSingleCharXor(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"4.txt",
			"Now that the party is jumping\n"},
	}

	for _, c := range cases {
		got := SingleCharXor(c.in)
		if got != c.want {
			t.Errorf("SingleCharXor(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRepeatingKeyXor(t *testing.T) {
	cases := []struct {
		plaintext, key, ciphertext string
	}{
		{"Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", "ICE",
			"0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"},
	}

	for _, c := range cases {
		got := RepeatKeyXor(c.plaintext, c.key)
		if got != c.ciphertext {
			t.Errorf("RepeatKeyXor(%q, %q) == %q, want %q", c.plaintext, c.key, got, c.ciphertext)
		}
	}
}

func TestDecryptRepeatingKeyXor(t *testing.T) {
	cases := []struct {
		ciphertext, key, plaintext string
	}{
		{"0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f", "ICE",
			"Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"},
	}

	for _, c := range cases {
		got := DecryptRepeatKeyXor(c.ciphertext, c.key)
		if got != c.plaintext {
			t.Errorf("DecryptRepeatKeyXor(%q, %q) == %q, want %q", c.ciphertext, c.key, got, c.plaintext)
		}
	}
}
