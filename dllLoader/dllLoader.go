package dllLoader

import (
	"encoding/json"
	"github.com/UniversalRobotDriveTeam/child-nodes-hdex-loader/loaderService"
	"os"
	"path"
	"strings"
	"syscall"
	"unsafe"
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

// GetFunctions 获取支持的函数列表
// 传入：无
// 传出：获得支持的函数名列表
func (dll *DllPackage) GetFunctions() []string {
	return dll.functions
}

// GetInfo 获取别的信息
// 传入：key
// 传出：value
func (dll *DllPackage) GetInfo(key string) string {
	return dll.info[key]
}

// Execute 执行函数
// 传入：方法名，参数
// 传出：返回值
// todo
func (dll *DllPackage) Execute(method string, args []uintptr) ([]uintptr, error) {
	proc, err := syscall.GetProcAddress(dll.dll, method)
	if err != nil {
		return nil, err
	}
	r, _, err := syscall.SyscallN(proc, uintptr(unsafe.Pointer(&args)))
	re := (*[]uintptr)(unsafe.Pointer(r))
	return *re, nil
}

// LoadHexPackage 根据路径加载二进制包并返回句柄
// 传入：路径
// 传出：二进制执行包
func (loader *DllLoader) LoadHexPackage(dllPath string) (*DllPackage, error) {
	// dll包对应的描述文件地址
	dllInfoPath := dllPath + ".json"
	// dll包地址
	dllPackagePath := dllPath + ".dll"
	// 获取dll包句柄
	h, err := syscall.LoadLibrary(dllPackagePath)
	if err != nil {
		return nil, err
	}
	// 加载json格式的dll信息
	content, err := os.ReadFile(dllInfoPath)
	if err != nil {
		return nil, err
	}
	var payload loaderService.HexInfo
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, err
	}
	// 初始化DllPackage类的name，dll
	dll := DllPackage{
		name:      strings.Split(path.Base(dllPackagePath), ".")[0],
		id:        0,
		functions: payload.Functions,
		dll:       h,
		info:      payload.Info,
	}
	// 根据dll计数器设置一个id
	dll.id = loader.dllCounter[dll.name]
	// 计数器自增
	loader.dllCounter[dll.name]++
	return &dll, err
}

// ReleasePackage 释放dll包
// 传入：二进制执行包
// 传出：无
func (loader *DllLoader) ReleasePackage(hexPackage *loaderService.HexPackage) error {
	_, err := (*hexPackage).Execute("release", nil)
	if err != nil {
		return err
	}
	return nil
}
