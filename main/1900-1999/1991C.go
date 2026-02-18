package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1991C(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		ok := true
		for i := range a {
			Fscan(in, &a[i])
			if (a[i]-a[0])%2 != 0 {
				ok = false
			}
		}
		if !ok {
			Fprintln(out, -1)
			continue
		}

		ans := []any{}
		for {
			mn := slices.Min(a)
			mx := slices.Max(a)
			if mn == mx {
				ans = append(ans, mn)
				break
			}
			x := (mn + mx) / 2
			ans = append(ans, x)
			for i, v := range a {
				a[i] = abs(v - x)
			}
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { cf1991C(bufio.NewReader(os.Stdin), os.Stdout) }
