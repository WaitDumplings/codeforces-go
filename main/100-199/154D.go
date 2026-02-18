package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf154D(in io.Reader, out io.Writer) {
	var x, y, a, b int
	Fscan(in, &x, &y, &a, &b)
	if x+a <= y && y <= x+b {
		Fprintln(out, "FIRST")
		Fprint(out, y)
		return
	}
	if a*b <= 0 || (a > 0) != (x < y) {
		Fprintln(out, "DRAW")
		return
	}
	z := (y - x) % (a + b)
	if z == 0 {
		Fprintln(out, "SECOND")
	} else if a <= z && z <= b {
		Fprintln(out, "FIRST")
		Fprint(out, x+z)
	} else {
		Fprint(out, "DRAW")
	}
}

//func main() { cf154D(bufio.NewReader(os.Stdin), os.Stdout) }
