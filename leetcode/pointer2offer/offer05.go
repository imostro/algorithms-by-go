package main

import "strings"

/*
剑指 Offer 05. 替换空格

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
*/

// api
func replaceSpaceByApi(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// 遍历替换
func replaceSpaceByByte(s string) string {
	data := []byte(s)
	var newStrB []byte

	for _, ch := range data {
		if ch == ' ' {
			newStrB = append(newStrB, '%', '2', '0')
		} else {
			newStrB = append(newStrB, ch)
		}
	}

	return string(newStrB)
}
