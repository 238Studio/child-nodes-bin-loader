package loader

import (
	"errors"
	"os"
	"plugin"
	"unsafe"

	_const "github.com/238Studio/child-nodes-assist/const"
	"github.com/238Studio/child-nodes-assist/util"
	jsoniter "github.com/json-iterator/go"
)

// GetName 获取name
// 传入：无
// 传出：全局唯一的包名
func (so *SoPackage) GetName() string {
	return so.name
}

// GetID 获取id
// 传入：无
// 传出：id
func (so *SoPackage) GetID() int {
	return so.id
}

// GetFunctions 获取函数列表
// 传入：无
// 传出：支持的函数列表
func (so *SoPackage) GetFunctions() []string {
	return so.functions
}

// GetFunctionsArgsTypes 获取函数入参类型
// 传入：函数名
// 传出：函数入参类型,错误
func (so *SoPackage) GetFunctionsArgsTypes(methodName string) ([]string, error) {
	functionArgs, isEXIST := so.functionsArgTypes[methodName]
	if !isEXIST {
		return nil, util.NewError(_const.CommonException, _const.Bin, errors.New("function not exist"))
	}
	return functionArgs, nil
}

// GetFunctionReturnTypes 获取函数返回值类型
// 传入：函数名
// 传出：函数返回值类型,错误
func (so *SoPackage) GetFunctionReturnTypes(methodName string) ([]string, error) {
	functionReturn, isEXIST := so.functionsReturnTypes[methodName]
	if !isEXIST {
		return nil, util.NewError(_const.CommonException, _const.Bin, errors.New("function not exist"))
	}
	return functionReturn, nil
}

// GetInfo 获取其他信息
// 传入：key
// 传出：value,错误
func (so *SoPackage) GetInfo(key string) (string, error) {
	value, isEXIST := so.info[key]
	if !isEXIST {
		return "", util.NewError(_const.CommonException, _const.Bin, errors.New("info not exist"))
	}
	return value, nil
}

// Execute 执行函数
// 传入：函数名，参数
// 传出：返回值(通过指针)，错误
func (so *SoPackage) Execute(method string, args []uintptr, re uintptr) (err error) {
	//捕获panic
	defer func() {
		if er := recover(); er != nil {
			//特例panic,级别非fatal,牵涉到cgo
			err = util.NewError(_const.CommonException, _const.Bin, errors.New(er.(string)))
		}
	}()

	//加载函数
	ptr, err := so.so.Lookup(method)
	if err != nil {
		return util.NewError(_const.CommonException, _const.Bin, err)
	}

	//如果没有参数(args==nil)，则调用无参函数
	if args == nil {
		ptr.(func())()
	} else {
		//如果有参数，则调用有参函数
		ptr.(func(uintptr, uintptr))(uintptr(unsafe.Pointer(&args)), re)
	}

	return nil
}

// GetPackage 获取包
// 传入：name，id
// 传出：包
func (soLoader *SoLoader) GetPackage(name string, id int) BinPackage {
	return soLoader.Sos[name][id]
}

// LoadBinPackage 根据路径加在二进制包
// 传入：路径
// 传出：包对象,错误
func (soLoader *SoLoader) LoadPackage(path string) (name string, id int, err error) {
	//捕获恐慌
	defer func() {
		if er := recover(); er != nil {
			//特例panic,级别非fatal,牵涉到cgo
			err = util.NewError(_const.CommonException, _const.Bin, errors.New(er.(string)))
		}
	}()

	soInfoPath := path + ".json"  //so包对应的描述文件地址
	soPackagePath := path + ".so" //so包地址

	//加载so包
	so, err := plugin.Open(soPackagePath)
	if err != nil {
		return "", 0, util.NewError(_const.CommonException, _const.Bin, err)
	}

	//加载so包对应的描述文件
	content, err := os.ReadFile(soInfoPath)
	if err != nil {
		return "", 0, util.NewError(_const.CommonException, _const.Bin, err)
	}

	//解析描述文件
	var (
		payload BinInfo
		json    = jsoniter.ConfigCompatibleWithStandardLibrary
	)
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return "", 0, util.NewError(_const.CommonException, _const.Bin, err)
	}

	//创建包对象
	soPackage := &SoPackage{
		name:                 payload.Name,
		id:                   0,
		functions:            payload.Functions,
		functionsArgTypes:    payload.FunctionsArgsTypes,
		functionsReturnTypes: payload.FunctionsReturnTypes,
		info:                 payload.Info,
		so:                   so,
	}

	//id分配
	if num, ok := soLoader.soCounter[soPackage.name]; !ok {
		soLoader.soCounter[soPackage.name] = 0 //初始化
	} else {
		soPackage.id = num
		soLoader.soCounter[soPackage.name]++
	}

	//加入集合
	if _, ok := soLoader.Sos[soPackage.name]; !ok {
		soLoader.Sos[soPackage.name] = make(map[int]*SoPackage)
	}
	soLoader.Sos[soPackage.name][soPackage.id] = soPackage

	return soPackage.name, soPackage.id, nil
}

// ReleasePackage 释放so包
// 传入：二进制执行包
// 传出：错误
func (soLoader *SoLoader) ReleasePackage(name string, id int) (err error) {
	//捕获恐慌
	defer func() {
		if er := recover(); er != nil {
			//特例panic,级别非fatal,牵涉到cgo
			err = util.NewError(_const.CommonException, _const.Bin, errors.New(er.(string)))
		}
	}()

	//通过name和id获取so package
	soPackage, isEXIST := soLoader.Sos[name][id]
	if !isEXIST {
		return util.NewError(_const.CommonException, _const.Bin, errors.New("package not exist"))
	}

	//释放so package
	err = soPackage.Execute("Release", nil, 0)
	if err != nil {
		return util.NewError(_const.CommonException, _const.Bin, err)
	}

	//从集合中移除
	delete(soLoader.Sos[name], id)

	return nil
}

// GetBinPackage 获取so包
// 传入：包名，id
// 传出：包对象,错误
func (soLoader *SoLoader) GetBinPackage(name string, id int) (soPackage *SoPackage, err error) {
	//捕获恐慌
	defer func() {
		if er := recover(); er != nil {
			//特例panic,级别非fatal,牵涉到cgo
			err = util.NewError(_const.CommonException, _const.Bin, errors.New(er.(string)))
		}
	}()

	//通过name和id获取so package
	soPackage, isEXIST := soLoader.Sos[name][id]
	if !isEXIST {
		return nil, util.NewError(_const.CommonException, _const.Bin, errors.New("package not exist"))
	}

	return soPackage, nil
}
