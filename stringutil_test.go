package gotil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello World","dlroW olleH"},
		{"dlrow olleh","hello world"},
		{"",""},
	}
	for _,cas := range cases {
		got := Reverse(cas.in)
		if got != cas.want {
			t.Errorf("Reverse(%q) expected %q but got %q", cas.in, cas.want, got)
		}
	}	
}
