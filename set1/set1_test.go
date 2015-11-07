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
