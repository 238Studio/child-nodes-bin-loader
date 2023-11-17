package main

import "C"
import (
	"fmt"
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
func Test1(args *uintptr) {
	println(args)
	arg := (*[]uintptr)(unsafe.Pointer(args))
	fmt.Printf("args:%+v", arg)
	str := dllLoader.ParsePtrToString(uintptr(unsafe.Pointer((*arg)[0])))
	println(str)
	b := str + "mew"
	re := make([]uintptr, 1)
	re[0] = (uintptr)(unsafe.Pointer(&b))
	println(b)
}

func main() {

}
