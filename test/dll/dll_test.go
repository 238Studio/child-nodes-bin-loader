package main_test

import (
	"testing"
	"unsafe"

	"github.com/238Studio/child-nodes-bin-loader/api"
	"github.com/238Studio/child-nodes-bin-loader/initial"
)

func TestName(t *testing.T) {
	args := make([]uintptr, 1)
	a := "helloworld"
	args[0] = uintptr(unsafe.Pointer(&a))
	context, err := initial.InitBinPackage()
	name, id, err := api.LoadBinPackage(context, "./test")
	if err != nil {
		println(err.Error())
		return
	}
	re := make([]uintptr, 1)
	var str = "外部"
	println(&str)
	re[0] = (uintptr)(unsafe.Pointer(&str))

	err = api.Execute(context, name, id, "Test1", args, uintptr(unsafe.Pointer(&re)))
	//	re := execute[0]
	println("mew")
	println(*(*string)(unsafe.Pointer(re[0])))
	args1, err := api.GetFunctionArgsTypes(context, name, id, "testFunc")
	println(args1[0])
	args2, err := api.GetFunctionReturnTypes(context, name, id, "testFunc")
	println(args2[1])

	api.ReleasePackage(context, name, id)
}
