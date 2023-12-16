package main

import "C"
import (
	"log"
	"unsafe"
)

//export Release
func Release() {
	println("已释放")
}

//export Test0
func Test0() {
	println("测试")
}

// Test1 测试go传入传出 和内部操作指针
//
//export Test1
func Test1(re uintptr, args uintptr, call uintptr) uintptr {
	arg := (*[]uintptr)(unsafe.Pointer(args))
	str := *(*string)(unsafe.Pointer((*arg)[0]))
	b := str + "mew"
	r := (*[]uintptr)(unsafe.Pointer(re))
	println("内部")
	println(*(*string)(unsafe.Pointer((*r)[0])) + "mew")
	*(*string)(unsafe.Pointer((*r)[0])) = b

	log.Println(unsafe.Pointer(call))
	callFuncPointer := (*func(string, string, uintptr) uintptr)(unsafe.Pointer(call))
	log.Println(callFuncPointer)
	reThing := (*callFuncPointer)("package1", "function1", args)
	log.Println(reThing)

	return uintptr(unsafe.Pointer(&re))
}

//ex

func main() {

}
