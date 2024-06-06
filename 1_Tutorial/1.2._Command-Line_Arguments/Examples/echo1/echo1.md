# echo1

本程序是Unix中`echo`命令的一个简单实现，它打印出其命令行参数。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
```
下面以该例程为蓝本，对go程序中命令行参数进行解释。

## 命令行参数

在go程序中，命令行参数是通过`os`包中的`Args`变量获取的。`os.Args`是一个字符串切片，其中第一个元素是程序的名字，随后的元素是程序的参数。

在`echo1`程序中，`os.Args`的第一个元素是程序的名字，所以我们从`os.Args[1]`开始遍历，将所有参数拼接到一个字符串中。

```go
for i := 1; i < len(os.Args); i++ {
	s += sep + os.Args[i]
	sep = " "
}
```

在这个循环中，`s`是一个字符串，`sep`是一个分隔符，`i`是一个索引。`s`用于存储所有参数的拼接结果，`sep`用于分隔参数，`i`用于遍历参数。

在每次循环中，我们将`os.Args[i]`添加到`s`中，然后将`sep`设置为空格。这样，我们就可以将所有参数拼接到一个字符串中。

最后，我们使用`fmt.Println(s)`打印出拼接结果。

## 程序结构

在这个例程中，我们使用了`package`、`import`、`func`、`var`、`for`、`fmt.Println`等关键字。

### package

`package`关键字用于声明一个包，包是go程序的基本组织单元。在这个例程中，我们声明了一个`main`包，代表了该程序是程序的入口。

### import

`import`关键字用于导入其他包，导入的包可以是标准库的包，也可以是第三方的包，也可以是自定义的包。在这个例程中，我们导入了`fmt`和`os`包。


