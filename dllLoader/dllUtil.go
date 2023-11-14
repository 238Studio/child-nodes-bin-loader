package dllLoader

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// ParseStringToPtr 将字符串转换为字符串数组和首指针 用于传参
// 传入：字符串
// 传出：指针
func ParseStringToPtr(str string) uintptr {
	re, _ := syscall.UTF16PtrFromString(str)
	return uintptr(unsafe.Pointer(re))
}

// ParsePtrToString 将字符串首指针转回字符串 用于解析返回值
// 传入：字符串首指针
// 传出：字符串
// todo
func ParsePtrToString(ptr uintptr) string {
	p := (*uint16)(unsafe.Pointer(ptr))
	if p == nil {
		return ""
	}

	n, end, add := 0, unsafe.Pointer(p), unsafe.Sizeof(*p)
	for *(*uint16)(end) != 0 {
		end = unsafe.Add(end, add)
		n++
	}
	return string(utf16.Decode(unsafe.Slice(p, n)))
}
