package loader

import (
	"encoding/json"
	"errors"
	"os"
	"syscall"
	"unsafe"

	_const "github.com/238Studio/child-nodes-assist/const"
	"github.com/238Studio/child-nodes-assist/util"
)

// GetName 获取名字
// 传入：无
// 传出：包名称 这个是全局唯一的
func (dll *DllPackage) GetName() string {
	return dll.name
}

// GetID 获取ID
// 传入：无
// 传出：包ID 这个是包名称和一个局部唯一的ID组成的
func (dll *DllPackage) GetID() int {
	return dll.id
}

// GetFunctionsArgsTypes 获取函数传入参数类型
// 传入：函数名
// 传出：传入参数类型数组
func (dll *DllPackage) GetFunctionsArgsTypes(methodName string) ([]string, error) {
	functionArgs, isExist := dll.functionsArgsTypes[methodName]
	if !isExist {
		return nil, util.NewError(_const.CommonException, _const.Bin, errors.New("function not exist"))
	}
	return functionArgs, nil
}

// GetFunctionReturnTypes 获得函数返回值类型列表
// 传入：函数名
// 传出：返回值类型列表
func (dll *DllPackage) GetFunctionReturnTypes(methodName string) ([]string, error) {
	functionReturn, isExist := dll.functionsReturnTypes[methodName]
	if !isExist {
		return nil, util.NewError(_const.CommonException, _const.Bin, errors.New("function not exist"))
	}
	return functionReturn, nil
}

// GetFunctions 获取支持的函数列表
// 传入：无
// 传出：获得支持的函数列表
func (dll *DllPackage) GetFunctions() []string {
	return dll.functions
}

// GetInfo 获取别的信息
// 传入：key
// 传出：value
func (dll *DllPackage) GetInfo(key string) (string, error) {
	info, isExist := dll.info[key]
	if !isExist {
		return "", util.NewError(_const.CommonException, _const.Bin, errors.New("info not exist"))
	}
	return info, nil
}

// Execute 执行函数
// 传入：方法名，参数
// 传出：返回值（通过指针）,错误
// todo
func (dll *DllPackage) Execute(method string, args []uintptr, re uintptr) error {
	// 在dll中获得方法的句柄
	proc, err := dll.dll.FindProc(method)
	if err != nil {
		return util.NewError(_const.CommonException, _const.Bin, err)
	}

	// 如果没有参数则直接无参调用方法
	if args == nil {
		_, _, err = proc.Call()
	} else {
		// 分别传入返回值指针和变量指针
		_, _, err = proc.Call(re, uintptr(unsafe.Pointer(&args)))
	}
	if !errors.Is(err, syscall.Errno(0)) && err != nil {
		return util.NewError(_const.CommonException, _const.Bin, err)
	}
	return nil
}

// GetPackage 获取包
// 传入：name，id
// 传出：包
func (dllLoader *DllLoader) GetPackage(name string, id int) BinPackage {
	return dllLoader.Dlls[name][id]
}

// GetPackage 获取包
// 传入：name，id
// 传出：包
func (dllLoader *DllLoader) GetPackage(name string, id int) BinPackage {
	return dllLoader.Dlls[name][id]
}

// LoadPackage 根据路径加载二进制包并返回句柄
// 传入：路径
// 传出：二进制执行包
func (dllLoader *DllLoader) LoadPackage(path string) (name string, id int, err error) {
	// dll包对应的描述文件地址
	dllInfoPath := path + ".json"
	// dll包地址
	dllPackagePath := path + ".dll"
	// 获取dll包句柄
	h := syscall.MustLoadDLL(dllPackagePath)
	// 加载json格式的dll信息
	content, err := os.ReadFile(dllInfoPath)
	if err != nil {
		return "", 0, util.NewError(_const.CommonException, _const.Bin, err)
	}
	var payload BinInfo
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return "", 0, util.NewError(_const.CommonException, _const.Bin, err)
	}
	// 初始化DllPackage类的name，dll
	dll := DllPackage{
		name:                 payload.Name,
		id:                   0,
		functions:            payload.Functions,
		functionsReturnTypes: payload.FunctionsReturnTypes,
		functionsArgsTypes:   payload.FunctionsArgsTypes,
		dll:                  h,
		info:                 payload.Info,
	}
	// 是否初始化计数器
	_, ok := dllLoader.dllCounter[dll.name]
	if !ok {
		dllLoader.dllCounter[dll.name] = 0
	}
	// 根据dll计数器设置一个id
	dll.id = dllLoader.dllCounter[dll.name]
	// 计数器自增
	dllLoader.dllCounter[dll.name]++
	_, ok = dllLoader.Dlls[dll.name]
	if !ok {
		dllLoader.Dlls[dll.name] = make(map[int]*DllPackage)
	}
	dllLoader.Dlls[dll.name][dll.id] = &dll
	return dll.name, dll.id, err
}

// ReleasePackage 释放dll包
// 传入：二进制执行包
// 传出：无
func (dllLoader *DllLoader) ReleasePackage(name string, id int) (err error) {
	binPackage := dllLoader.Dlls[name][id]
	err = (binPackage).Execute("Release", nil, 0)
	//todo 常量化
	if err != nil {
		return util.NewError(_const.CommonException, _const.Bin, err)
	}
	return nil
}
