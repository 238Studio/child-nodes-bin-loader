package loader

// InitLoader 初始化dll二进制包
// 传入：无
// 传出：无
func InitLoader() *DllLoader {
	dllLoader := DllLoader{
		Dlls:       make(map[string]map[int]*DllPackage),
		dllCounter: make(map[string]int),
	}
	return &dllLoader
}
