package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1037F(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, k int
	Fscan(in, &n, &k)
	k--
	a := make([]int, n)
	l := make([]int, n)
	r := make([]int, n)
	st := []int{-1}
	for i := range a {
		Fscan(in, &a[i])
		for len(st) > 1 && a[st[len(st)-1]] < a[i] {
			r[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		l[i] = st[len(st)-1]
		st = append(st, i)
	}
	for _, i := range st[1:] {
		r[i] = n
	}

	ans := 0
	for i, ai := range a {
		x, y := i-l[i]-1, r[i]-i-1
		if x > y {
			x, y = y, x
		}
		u, v, w := x/k, y/k, (x+y)/k
		res := (u*(u+1)/2*k + u + (v-u)*(x+1) - (v+w+1)*(w-v)/2*k + (x+y+1)*(w-v)) % mod
		ans = (ans + res*ai) % mod
	}
	Fprint(out, ans)
}

//func main() { cf1037F(bufio.NewReader(os.Stdin), os.Stdout) }
