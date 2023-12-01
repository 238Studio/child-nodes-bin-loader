package initial

import (
	"runtime"

	dll "github.com/238Studio/child-nodes-bin-loader-dll"
	so "github.com/238Studio/child-nodes-bin-loader-so"
	"github.com/238Studio/child-nodes-bin-loader/assist"
)

// InitBinPackage 初始化二进制包
// 传入：无
// 传出：上下文，错误
func InitBinPackage() (context assist.Context, err error) {
	//初始化二进制包上下文
	context = assist.InitContext()

	//判断系统，根据系统类型使用对应的动态库类型
	sysType := runtime.GOOS
	context.SysType = sysType

	switch sysType {
	case assist.Windows:
		dllLoader := dll.InitDllLoader()
		context.DllLoader = dllLoader
	case assist.Linux:
		soLoader := so.InitSoLoader()
		context.SoLoader = soLoader
	default:
		panic("system not support")
	}

	return context, nil
}
