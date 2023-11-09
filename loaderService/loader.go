package loaderService

// Loader loader需要实现的若干功能
// loader是一个加载本地函数的服务 它有若干个不同的实现
type Loader interface {
	// LoadHexPackage 根据路径加载二进制包并返回句柄
	LoadHexPackage(path string)
}

// HexPackage 二进制可执行包
type HexPackage struct {
	// name
	name string
	// id
	id string
	// 支持的函数列表
	functions []string
}

// HexPackageExecute 二进制可执行包执行函数
type HexPackageExecute interface {
	// Execute 执行函数
	// 传入：方法名，参数
	// 传出：返回值
	Execute(method string, args []interface{}) []interface{}
}
