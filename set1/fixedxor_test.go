package hextobase64

import "testing"

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
