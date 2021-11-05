package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(1))
}

//给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//
//由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//
//注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
//
//
//示例 1：
//
//输入：x = 4
//输出：2
//示例 2：
//
//输入：x = 8
//输出：2
//解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
//
//
//提示：
//
//0 <= x <= 2^31 - 1

func mySqrt(x int) int {
	return mySqrt_B(x)
}

func mySqrt_B(x int) int {
	var (
		fn = func(a, b int) int {
			return int(math.Ceil((float64(a) + float64(b)) / 2))
		}
		b = 0
		e = fn(x, b)
		n = e
	)

	for b < e {
		var nn = n * n

		if nn > x {
			e = n
			n = fn(e, b)
		} else if nn == x || (nn < x && (n+1)*(n+1) > x) {
			break
		} else if nn < x {
			b = n
			n = fn(e, b)
		}
	}

	return n
}

func mySqrt_A(x int) int {
	var n int

	for n < x {
		var nn = n * n
		if nn == x || (nn < x && (n+1)*(n+1) > x) {
			break
		}

		n++
	}

	return n
}
