package main

import (
	"fmt"
)

func main() {
	var a, b string
	a = "1010"
	b = "1011"
	fmt.Println(addBinary(a, b))
}

//67. 二进制求和
//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
//输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//提示：
//
//每个字符串仅由字符 '0' 或 '1' 组成。
//1 <= a.length, b.length <= 10^4
//字符串如果不是 "0" ，就都不含前导零。

func addBinary(a string, b string) string {
	return addBinary_B(a, b)
}

func addBinary_B(a string, b string) string {
	var (
		la            = len(a)
		lb            = len(b)
		r             = make([]uint8, 0)
		rb      uint8 = 0
		reverse       = func(s []uint8) []uint8 {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			return s
		}
	)

	for la > 0 || lb > 0 || rb > 0 {
		var (
			cnt uint8
			rr  uint8
		)

		if la > 0 {
			rr += a[la-1]
			cnt++
		}

		if lb > 0 {
			rr += b[lb-1]
			cnt++
		}

		rr = rr - 48*cnt + rb
		if rr > 1 {
			rb = rr / 2
			rr = rr % 2
		} else {
			rb = 0
		}

		r = append(r, rr+48) // 逆序存储，结果反序处理
		//r = append([]uint8{rr + 48}, r...)
		la--
		lb--
	}

	return string(reverse(r))
}

//addBinary_A 耗内存多
func addBinary_A(a string, b string) string {
	var (
		add = func(rs ...uint8) (uint8, uint8) {
			var (
				l = len(rs)
				r uint8
				p uint8
			)

			for _, v := range rs {
				r += v
			}
			r = r - 48*uint8(l)

			p = r / 2
			r = r % 2
			return r + 48, p + 48
		}

		la       = len(a)
		lb       = len(b)
		r        = make([]uint8, 0)
		rb uint8 = 48
	)

	for la > 0 || lb > 0 || rb > 48 {
		var (
			rs []uint8 = make([]uint8, 0)
			rr uint8
		)

		if la > 0 {
			rs = append(rs, a[la-1])
		}

		if lb > 0 {
			rs = append(rs, b[lb-1])
		}

		rs = append(rs, rb)

		rr, rb = add(rs...)
		r = append([]uint8{rr}, r...)
		la--
		lb--
	}

	return string(r)
}
