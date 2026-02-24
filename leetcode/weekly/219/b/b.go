package main

// github.com/EndlessCheng/codeforces-go
func minPartitions(n string) int {
	ans := rune(0)
	for _, ch := range n {
		ans = max(ans, ch)
	}
	return int(ans - '0')
}
