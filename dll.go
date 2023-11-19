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
func Test1(r uintptr, args *uintptr) {
	arg := (*[]uintptr)(unsafe.Pointer(args))
	str := dllLoader.ParsePtrToString(uintptr(unsafe.Pointer((*arg)[0])))
	b := str + "mew"
	r_ := dllLoader.ParseStringToPtr(b)
	defer println("完成")
	re := (*[]*interface{})(unsafe.Pointer(r))
	println(re)
	*(*re)[0] = r_
}

func main() {

}
