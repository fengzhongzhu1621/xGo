package string_utils

import (
	"bytes"
	"fmt"
	"strings"
)

/**
 * 根据分隔符分割字符串
 */
func Head(str, sep string) (head string, tail string) {
	idx := strings.Index(str, sep)
	if idx < 0 {
		return str, ""
	}
	return str[:idx], str[idx+len(sep):]
}

////////////////////////////////////////////////////////
// 字符串拼接

// Deprecated: 低效的字符串拼接
func StringPlus(p []string) string {
	var s string
	l := len(p)
	for i := 0; i < l; i++ {
		s += p[i]
	}
	return s
}

// Deprecated: 低效的字符串拼接
func StringFmt(p []interface{}) string {
	return fmt.Sprint(p...)
}

// Deprecated: 低效的字符串拼接
func StringBuffer(p []string) string {
	var b bytes.Buffer
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

func StringJoin(p []string) string {
	return strings.Join(p, "")
}

func StringBuilder(p []string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

func StringBuilderEx(p []string, cap int) string {
	var b strings.Builder
	l := len(p)
	// 实现分配足够的内容，减少运行时的内存分配
	b.Grow(cap)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}
