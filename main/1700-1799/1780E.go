package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1780E(in io.Reader, out io.Writer) {
	var T, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		l--
		ans := 0
		for i, j := 1, 0; i <= l; i = j + 1 {
			h := l / i
			j = l / h
			ans += max(min(j, r/(h+2))-i+1, 0)
		}
		Fprintln(out, ans+max(r/2-l, 0))
	}
}

//func main() { cf1780E(os.Stdin, os.Stdout) }
