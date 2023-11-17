package dllLoader

import (
	"testing"
)

func TestName(t *testing.T) {
	args := make([]uintptr, 1)
	a := "helloworld"
	args[0] = ParseStringToPtr(a)
	app := InitDllLoader()
	println("mew")
	hexPackage, err := app.LoadHexPackage("C:\\Users\\real_common_cat\\Desktop\\childNodes\\child-nodes-hex-loader\\dll")
	if err != nil {
		println(err.Error())
		return
	}
	//hexPackage.Execute("Test0", nil)
	execute, err := hexPackage.Execute("Test1", args)
	if err != nil {
		return
	}
	re := execute[0]
	println(ParsePtrToString(re))
}
