package loader

import "testing"

func Test0(t *testing.T) {
	app := InitBinApp()
	name, id, _ := app.LoadPackage("E:\\238Robot\\childNodes\\child-nodes-bin-loader\\test\\test")
	app.GetPackage(name, id).Execute("Test0", nil, 0)
}
