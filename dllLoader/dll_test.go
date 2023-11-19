package dllLoader

import (
	"testing"
)

func TestName(t *testing.T) {
	args := make([]uintptr, 1)
	a := "helloworld"
	args[0] = ParseStringToPtr(a)
	app := InitDllLoader()
	hexPackage, err := app.LoadHexPackage("C:\\Users\\real_common_cat\\Desktop\\childNodes\\child-nodes-hex-loader\\dll")
	if err != nil {
		println(err.Error())
		return
	}
	//hexPackage.Execute("Test0", nil)
	var re = make([]*interface{}, 1)
	println(re)
	err = hexPackage.Execute("Test1", args, &re)
	if err != nil {
		return
	}
	//	re := execute[0]
}
