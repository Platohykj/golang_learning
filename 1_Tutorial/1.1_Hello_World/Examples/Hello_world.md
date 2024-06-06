# Hello world

## 编译命令

golang的编译命令如下

```shell
go run hello.go
```
也可以先编译再运行

```shell
go build
./hello.exe
```

## 包管理

golang的代码通过包（package）组织，包类似于其它语言的库（library）或模块（module）。

包由单个目录下的一个或多个.go文件组成，***目录定义包的作用域。***

包的声明语法如下：

```go
package main
```
包的声明代表了该文件属于哪个包，包名可以与目录名不同，但是通常情况下，包名与目录名相同。

包的导入语法如下：

```go
import "fmt"
```

包的导入语句用于导入其他包，导入的包可以是标准库的包，也可以是第三方的包，也可以是自定义的包。

## 格式标准

golang中，必须先声明一个包，然后是导入的包

```go
package main
import "fmt"
```

随后是函数、变量、常量等的声明，函数的声明如下：

```go
func main() {
    fmt.Println("Hello, World!")
}
```
函数的声明中，`func`是关键字，`main`是函数名，`()`是参数列表，`{}`是函数体。 还会有返回值。

golang不需要在每行代码后面加`;`，但由于编译器会将特定符号后的换行符转换为分号，所以换行符是必须的。

golang对代码格式采取强硬态度，`gofmt`子命令把代码格式化为标准格式。

```shell
go fmt
```

该命令如果不指定包会对当前目录下所有文件应用`gofmt`命令