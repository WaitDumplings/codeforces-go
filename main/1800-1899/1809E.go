package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1809E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, a, b int
	Fscan(in, &n, &a, &b)
	v := make([]int, n)
	for i := range v {
		Fscan(in, &v[i])
	}

	ans := make([][]int, a+1)
	for i := range ans {
		ans[i] = make([]int, b+1)
	}
	for i := range a + b + 1 {
		d := 0
		lb := max(0, i-b)
		rb := min(a, i)
		l, r := lb, rb
		for _, vj := range v {
			d -= vj
			l = min(max(l, lb-d), rb-d)
			r = min(max(r, lb-d), rb-d)
		}
		for j := lb; j <= rb; j++ {
			ans[j][i-j] = min(max(j, l), r) + d
		}
	}

	for _, r := range ans {
		for _, v := range r {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1809E(bufio.NewReader(os.Stdin), os.Stdout) }
