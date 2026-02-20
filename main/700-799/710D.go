package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf710D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var a, b, c, d, l, r, t int
	Fscan(in, &a, &b, &c, &d, &l, &r)
	if a < c {
		a, c = c, a
		b, d = d, b
	}
	for b <= r && ((b-d)%c+c)%c != 0 {
		b += a
		t++
		if t > c {
			b = r + 1
		}
	}
	t = a / gcd(a, c) * c
	b += max(0, max(l, d)-b+t-1) / t * t
	if b > r {
		Fprint(out, 0)
	} else {
		Fprint(out, (r-b)/t+1)
	}
}

//func main() { cf710D(bufio.NewReader(os.Stdin), os.Stdout) }
