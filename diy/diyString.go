package diy

import (
	"reflect"
	"unsafe"
)

// Atoi It can detail change string Type to int.
// Limit 1 << 31 - 1 in -1 << 31 value
func Atoi(s string) int {
	var mark, res, sum, tmp int
	for i, _ := range s {
		if res == 0 {
			if mark == 0 && s[i] == ' ' {
				continue
			}
			if s[i] == '-' {
				if mark == 0 {
					mark = -1
					continue
				}
				return 0
			}
			if s[i] == '0' {
				if mark == 0 {
					mark = 1
				}
				continue
			}
			if s[i] == '+' {
				if mark == 0 {
					mark = 1
					continue
				}
				return 0
			}
		}
		if !(s[i] >= '0' && s[i] <= '9') {
			break
		}
		tmp = int(s[i] - '0')
		sum = res*10 + tmp
		if (sum-tmp)/10 != res || sum < 0 {
			if mark != -1 {
				return MaxInt
			}
			return MinInt
		}
		res = sum
	}
	if mark == -1 {
		res *= -1
	}
	if res > MaxInt {
		res = MaxInt
	}
	if res < MinInt {
		res = MinInt
	}
	return res
}

// Index It will return a int value.
// Have two strings will be judged. The value meaning of second string
// is existing in first string, and it first appears index.
// if it is not exist, and will be return -1
func Index(ftr string, str string) int {
	n, m := len(ftr), len(str)
	if n < m {
		return -1
	}
	for i := 0; i < n; i++ {
		for i < n && ftr[i] != str[0] {
			i++
		}
		j, tmp := 0, i
		for j < m && i < n {
			if ftr[i] != str[j] {
				break
			}
			i++
			j++
		}
		if j == m {
			return tmp
		}
		i = tmp
	}
	return -1
}

// StringToBytes BlackMagic .
// It will make string to bytes and return it.
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// BytesToString BlackMagic .
// It will make bytes to string and return it.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
