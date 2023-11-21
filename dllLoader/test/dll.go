package main

import "C"
import (
	"github.com/UniversalRobotDriveTeam/child-nodes-hdex-loader/dllLoader"
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

//export Test1
func Test1(re uintptr, args *uintptr) uintptr {
	arg := (*[]uintptr)(unsafe.Pointer(args))
	str := dllLoader.ParsePtrToString(uintptr(unsafe.Pointer((*arg)[0])))
	b := str + "mew"
	r := (*[]uintptr)(unsafe.Pointer(re))
	println("内部")
	println(*(*string)(unsafe.Pointer((*r)[0])))
	*(*string)(unsafe.Pointer((*r)[0])) = b
	return uintptr(unsafe.Pointer(&re))
}

func main() {

}
