package api

import (
	"github.com/238Studio/child-nodes-bin-loader/assist"
)

// LoadBinPackage 加载二进制包
// 传入：上下文，路径
// 传出：包名称，包ID，错误
func LoadBinPackage(context assist.Context, path string) (name string, id int, err error) {
	systemType := context.GetSystemType()

	switch systemType {
	case assist.Windows:
		name, id, err = context.DllLoader.LoadBinPackage(path)
	case assist.Linux:
		name, id, err = context.SoLoader.LoadBinPackage(path)
	}

	return name, id, err
}

// GetFunctionArgsTypes 获取函数参数类型
// 传入：上下文，包名称，包ID，函数名称
// 传出：参数类型数组，错误
func GetFunctionArgsTypes(context assist.Context, name string, id int, methodName string) (args []string, err error) {
	systemType := context.GetSystemType()

	switch systemType {
	case assist.Windows:
		//获取包
		binPackage, err := context.DllLoader.GetBinPackage(name, id)
		if err != nil {
			return nil, err
		}
		//获取函数参数类型
		args, err = binPackage.GetFunctionsArgsTypes(methodName)

	case assist.Linux:
		//获取包
		binPackage, err := context.SoLoader.GetBinPackage(name, id)
		if err != nil {
			return nil, err
		}
		//获取函数参数类型
		args, err = binPackage.GetFunctionsArgsTypes(methodName)
	}

	return args, err
}

// GetFunctionReturnTypes 获取函数返回值类型
// 传入：上下文，包名称，包ID，函数名称
// 传出：返回值类型数组，错误
func GetFunctionReturnTypes(context assist.Context, name string, id int, methodName string) (returns []string, err error) {
	systemType := context.GetSystemType()

	switch systemType {
	case assist.Windows:
		//获取包
		binPackage, err := context.DllLoader.GetBinPackage(name, id)
		if err != nil {
			return nil, err
		}
		//获取函数返回值类型
		returns, err = binPackage.GetFunctionReturnTypes(methodName)

	case assist.Linux:
		//获取包
		binPackage, err := context.SoLoader.GetBinPackage(name, id)
		if err != nil {
			return nil, err
		}
		//获取函数返回值类型
		returns, err = binPackage.GetFunctionReturnTypes(methodName)
	}

	return returns, err
}

// GetFunctions 获取函数列表
// 传入：上下文，包名称，包ID
// 传出：函数列表，错误
func GetFunctions(context assist.Context, name string, id int) (functions []string, err error) {
	systemType := context.GetSystemType()

	switch systemType {
	case assist.Windows:
		//获取包
		binPackage, err := context.DllLoader.GetBinPackage(name, id)
		if err != nil {
			return nil, err
		}
		//获取函数列表
		functions = binPackage.GetFunctions()

	case assist.Linux:
		//获取包
		binPackage, err := context.SoLoader.GetBinPackage(name, id)
		if err != nil {
			return nil, err
		}
		//获取函数列表
		functions = binPackage.GetFunctions()
	}

	return functions, err
}

// GetInfo 获取包信息
// 传入：上下文，包名称，包ID，key
// 传出：value，错误
func GetInfo(context assist.Context, name string, id int, key string) (info string, err error) {
	systemType := context.GetSystemType()

	switch systemType {
	case assist.Windows:
		//获取包
		binPackage, err := context.DllLoader.GetBinPackage(name, id)
		if err != nil {
			return "", err
		}
		//获取info
		info, err = binPackage.GetInfo(key)

	case assist.Linux:
		//获取包
		binPackage, err := context.SoLoader.GetBinPackage(name, id)
		if err != nil {
			return "", err
		}
		//获取info
		info, err = binPackage.GetInfo(key)
	}

	return info, err
}

// Execute 执行函数
// 传入：上下文，包名称，包ID，函数名称，参数，返回值
// 传出：错误
func Execute(context assist.Context, name string, id int, methodName string, args []uintptr, re uintptr) (err error) {
	systemType := context.GetSystemType()

	switch systemType {
	case assist.Windows:
		//获取包
		binPackage, err := context.DllLoader.GetBinPackage(name, id)
		if err != nil {
			return err
		}
		//执行函数
		err = binPackage.Execute(methodName, args, re)

	case assist.Linux:
		//获取包
		binPackage, err := context.SoLoader.GetBinPackage(name, id)
		if err != nil {
			return err
		}
		//执行函数
		err = binPackage.Execute(methodName, args, re)
	}

	return err
}

// ReleasePackage 释放包
// 传入：上下文，包名称，包ID
// 传出：错误
func ReleasePackage(context assist.Context, name string, id int) (err error) {
	systemType := context.GetSystemType()

	switch systemType {
	case assist.Windows:
		err = context.DllLoader.ReleasePackage(name, id)
	case assist.Linux:
		err = context.SoLoader.ReleasePackage(name, id)
	}

	return err
}
