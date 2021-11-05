package main

import "fmt"

func main() {
	var s string
	s = "as df"
	s = "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

//给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。
//
//单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
//
//
//
//示例 1：
//
//输入：s = "Hello World"
//输出：5
//示例 2：
//
//输入：s = "   fly me   to   the moon  "
//输出：4
//示例 3：
//
//输入：s = "luffy is still joyboy"
//输出：6
//
//
//提示：
//
//1 <= s.length <= 104
//s 仅有英文字母和空格 ' ' 组成
//s 中至少存在一个单词

func lengthOfLastWord(s string) int {
	return lengthOfLastWord_B(s)
}

func lengthOfLastWord_B(s string) int {
	var (
		sl = len(s)
		l  int
	)

	for i := sl - 1; i >= 0; i-- {
		if s[i] == 32 {
			if l > 0 {
				return l
			}
		} else {
			l++
		}
	}

	return l
}

func lengthOfLastWord_A(s string) int {
	var (
		l int
		r bool
	)
	for _, v := range s {
		if v == 32 {
			r = true
		} else {
			if r {
				l = 0
				r = !r
			}
			l++
		}
	}

	return l
}
