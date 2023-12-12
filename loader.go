package loader

// Loader loader需要实现的若干功能
// loader是一个加载本地函数的服务 它有若干个不同的实现
type Loader interface {
	// LoadBinPackage 根据路径加载二进制包并返回句柄
	// 传入：路径
	// 传出：二进制执行包
	LoadBinPackage(path string) (BinPackage, error)
	// ReleasePackage 释放dll包
	// 传入：二进制执行包
	// 传出：无
	ReleasePackage(binPackage BinPackage) error
}

// BinPackage 二进制可执行包
/*
每个包都有一个全局唯一的name
而每个函数也是全局唯一的 其通过name和function name来断定
函数调用直接传入和传出指针数组 其没有显式的类型检查
*/
type BinPackage interface {
	// GetName 获取名字
	// 传入：无
	// 传出：名字
	GetName() string

	// GetID 获取ID
	// 传入：无
	// 传出：ID

	GetID() int

	// GetFunctions 获取支持的函数列表
	// 传入：无
	// 传出：支持的函数列表
	GetFunctions() []string

	// GetFunctionsArgsTypes 获取函数传入参数类型
	// 传入：函数名
	// 传出：传入参数类型数组
	GetFunctionsArgsTypes(methodName string) ([]string, error)

	// GetFunctionReturnTypes 获得函数返回值类型列表
	// 传入：函数名
	// 传出：返回值类型列表
	GetFunctionReturnTypes(methodName string) ([]string, error)

	// GetInfo 获取别的信息
	// 传入：key
	// 传出：value
	GetInfo(key string) (string, error)

	// Execute 执行函数
	// 传入：方法名，参数，返回值指针数组的指针
	// 传出：错误
	Execute(method string, args []uintptr, re uintptr) error
}

// Parse

// BinInfo 二进制包信息
type BinInfo struct {
	// 全局唯一的名称
	Name string
	// 支持的函数 函数名
	Functions []string
	// 函数的入参类型 函数名-入参类型表
	FunctionsArgsTypes map[string][]string
	// 函数的返回值类型表 函数名-返回值类型表
	FunctionsReturnTypes map[string][]string
	// info
	Info map[string]string
}
