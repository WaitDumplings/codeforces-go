package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2085C(in io.Reader, out io.Writer) {
	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)
		if x == y {
			Fprintln(out, -1)
		} else {
			x = max(x, y)
			mask := 1<<bits.Len(uint(x)) - 1
			Fprintln(out, mask^x+1)
		}
	}
}

//func main() { cf2085C(bufio.NewReader(os.Stdin), os.Stdout) }
