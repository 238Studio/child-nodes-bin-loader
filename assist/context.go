package assist

import (
	dll "github.com/238Studio/child-nodes-bin-loader-dll"
	so "github.com/238Studio/child-nodes-bin-loader-so"
)

// Context 上下文结构
type Context struct {
	SysType   string         //系统类型
	DllLoader *dll.DllLoader //dll loader
	SoLoader  *so.SoLoader   //so loader
}

// InitContext  初始化上下文
// 传入：无
// 传出：空上下文结构体
func InitContext() Context {
	return Context{}
}

// GetSystemType 根据系统类型获取对应的system type
// 传入：无
// 传出：system type
func (context Context) GetSystemType() string {
	return context.SysType
}
