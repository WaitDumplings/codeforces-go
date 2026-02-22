package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod85 = 1_000_000_007

func pow85(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod85
		}
		x = x * x % mod85
	}
	return
}

type comb85 struct{ _f, _invF []int }

func newComb85(mx int) *comb85 {
	c := &comb85{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb85) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod85
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow85(c._f[mx], mod85-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod85
	}
}

func (c *comb85) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb85) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb85) c(n, k int) int {
	return c.f(n) * c.invF(k) % mod85 * c.invF(n-k) % mod85
}

func cf285E(in io.Reader, out io.Writer) {
	cm := newComb85(0)
	var n, k, ans int
	Fscan(in, &n, &k)
	f := make([]int, n+1)
	for i := range n/2 + 1 {
		for j := range n/2 + 1 {
			f[i+j] = (f[i+j] + cm.c(n-i, i)*cm.c(n-j, j)) % mod85
		}
	}
	for i := range n + 1 {
		f[i] = f[i] * cm.f(n-i) % mod85
	}
	for i := k; i <= n; i++ {
		sign := 1 - (i-k)%2*2
		ans = (ans + sign*f[i]*cm.c(i, k)) % mod85
	}
	Fprint(out, (ans+mod85)%mod85)
}

//func main() { cf285E(bufio.NewReader(os.Stdin), os.Stdout) }
