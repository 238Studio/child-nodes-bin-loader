package main_test

import (
	"fmt"
	"log"
	"testing"
	"time"
	"unsafe"

	loader "github.com/238Studio/child-nodes-bin-loader"
)

func Test0(t *testing.T) {
	app := loader.InitBinApp()
	name, id, err := app.LoadPackage("./test")
	if err != nil {
		t.Error(err)
		return
	}
	pa := app.GetPackage(name, id)
	err = pa.Execute("Test0", nil, 0, nil)
	if err != nil {
		t.Error(err)
	}

	args := make([]uintptr, 1)
	a := "helloworld"
	args[0] = uintptr(unsafe.Pointer(&a))

	re := make([]uintptr, 1)
	var str = "外部"
	re[0] = (uintptr)(unsafe.Pointer(&str))

	var call = Call
	log.Println(unsafe.Pointer(&call))
	fmt.Printf("The address of myFunc is %p\n", call)

	testCall := (*func(string, string, uintptr) uintptr)(unsafe.Pointer(&call))
	log.Println(testCall)
	reThing := (*testCall)("package1", "function1", uintptr(unsafe.Pointer(&args)))
	log.Println(reThing)

	err = pa.Execute("Test1", args, uintptr(unsafe.Pointer(&re)), &call)
	if err != nil {
		t.Error(err)
	}

	println(pa.GetInfo("test0"))
	println(pa.GetIsPrimary())
	trigger := pa.GetTriggerCallArgs()
	println((*trigger)["package1"]["function1"][0])

	println(re)
	time.Sleep(10 * time.Second)
}

func Call(packageName, methodName string, args uintptr) (re uintptr) {
	log.Println("回调测试:", packageName, methodName)
	return args
}
