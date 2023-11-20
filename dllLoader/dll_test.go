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
	re := make([]string, 1)
	re[0] = ""
	err = hexPackage.Execute("Test1", args, &re)
	//	re := execute[0]
	println(re[0])
}
