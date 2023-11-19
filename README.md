# API 文档

## `GetName() string`

### 描述
获取包名字，这个包名是全局唯一的。

### 输入
- 无

### 输出
- 类型：`string`
- 返回：名字

## `GetID() int`

### 描述
获取包ID，这个ID是分配给各个包的，同一个包可能有多个实例，它们每个都被分配了自己的ID。

### 输入
- 无

### 输出
- 类型：`int`
- 返回：ID

## `GetFunctions() []string`

### 描述
获取支持的函数列表。

### 输入
- 无

### 输出
- 类型：`[]string`
- 返回：支持的函数列表，也就是一系列的函数名。

## `GetFunctionsArgsTypes(methodName string) []string`

### 描述
获取函数传入参数类型

### 输入
- 函数名

### 输出
- 类型：`字符串数组`

## `GetFunctionReturnTypes(methodName string) []string`

### 描述
获取函数返回值类型

### 输入
- 函数名

### 输出
- 类型：`字符串数组`

## `GetInfo(key string) string`

### 描述
根据提供的键获取其他信息。

### 输入
- `key`（类型：`string`）：信息的键

### 输出
- 类型：`string`
- 返回：与提供的键对应的值

## `Execute(method string, args []uintptr) ([]uintptr, error)`
## `暂不确定`
### 描述
执行函数。

### 输入
- `method`（类型：`string`）：方法名
- `args`（类型：`[]uintptr`）：参数

### 输出
- 类型：`[]uintptr`
- 返回：执行结果

### 错误处理
- 如果执行过程中发生错误，将返回一个包含错误消息的错误对象。
