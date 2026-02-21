package main

import "math/bits"

// https://space.bilibili.com/206214
func binaryGap(n int) (ans int) {
	n /= n & -n * 2 // 去掉 n 末尾的 100..0
	for n > 0 {
		gap := bits.TrailingZeros(uint(n)) + 1
		ans = max(ans, gap)
		n >>= gap // 去掉 n 末尾的 100..0
	}
	return
}
