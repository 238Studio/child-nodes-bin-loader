package main_test

import (
	"testing"

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
	err = pa.Execute("Test0", nil, 0)
	if err != nil {
		t.Error(err)
	}
}
