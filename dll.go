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
func Test1(args *uintptr) uintptr {
	println(args)
	arg := (*[]uintptr)(unsafe.Pointer(args))
	str := dllLoader.ParsePtrToString(uintptr(unsafe.Pointer((*arg)[0])))
	b := str + "mew"
	r := dllLoader.ParseStringToPtr(b)
	re := make([]*uintptr, 1)
	re[0] = &r
	defer println("完成")
	rep := uintptr(unsafe.Pointer(&re))
	println("内部地址")
	println(rep)
	k := (*[]*uintptr)(unsafe.Pointer(rep))
	println(*(*k)[0])
	return rep
}

func main() {

}
