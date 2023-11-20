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

## 关于hex包在机器人架构层面要实现的功能
- hex包实现的功能：hex包实现的是对现成软件包的加载，释放和调用。实际上就是对预设的函数的加载，释放和调用。
- 对于机器人而言，机器人的任务流程实际上是在"一定情况下做一定的事"，这个一定情况指的是机器人的传感器输入的状态，
自身记录的状态，或者别的判断条件，总的来说，就是机器人提供的API的返回值。
做一定的事情则是按照一定顺序来调用机器人的API，使得机器人有相应的动作，而这些是由软件包实现的。
- 在中层，通过协程来循环调用每个没有阻塞的“任务包”，这些“任务包”实际上是软件包和软件包描述这两个部分的集合。
软件包描述描述了软件包调用需要的入参类型和可以调用的方法，以及软件包的其他信息。
另外，注意到，软件包也属于机器人API的一部分，但是软件包通过组来进行了调用域上的划分。软件包可以通过中层的hex管理器来
获取同一组内其他软件包的信息，并进行调用。

