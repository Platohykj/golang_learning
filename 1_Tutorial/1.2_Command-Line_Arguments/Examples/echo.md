# echo

[`echo1.go`](./echo1/echo1.go)是Unix中`echo`命令的一个简单实现，它打印出其命令行参数。

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

在go程序中，命令行参数是通过`os`包中的`Args`变量获取的。
`os.Args`是一个字符串切片，其中第一个元素是程序的名字，随后的元素是程序的参数。

在[`echo1.go`](./echo1/echo1.go)程序中，`os.Args`的第一个元素是程序的名字，所以我们从`os.Args[1]`开始遍历，将所有参数拼接到一个字符串中。

```go
for i := 1; i < len(os.Args); i++ {
	s += sep + os.Args[i]
	sep = " "
}
```

在这个循环中，`s`是一个字符串，`sep`是一个分隔符，`i`是一个索引。
`s`用于存储所有参数的拼接结果，`sep`用于分隔参数，`i`用于遍历参数。

在每次循环中，我们将`os.Args[i]`添加到`s`中，然后将`sep`设置为空格。
这样，我们就可以将所有参数拼接到一个字符串中。

最后，我们使用`fmt.Println(s)`打印出拼接结果。

## 程序结构

在这个例程中，我们使用了`package`、`import`、`func`、`for`等关键字。

### package
```go
package main
```

`package`关键字用于声明一个包，包是go程序的基本组织单元。
在这个例程中，我们声明了一个`main`包，代表了该程序是程序的入口。

### import

```go
import (
	"fmt"
	"os"
)
```

`import`关键字用于导入其他包，导入的包可以是标准库的包，也可以是第三方的包，也可以是自定义的包。

在这个例程中，我们导入了`fmt`和`os`包，但并没有写成独立的`import`声明，而是把它们括起来写成列表形式。

习惯上，我们通常使用列表形式，而不是独立的`import`声明。

在列表形式中，每个包名都是一个字符串，包名之间用逗号分隔。包之间的顺序没有要求。
`gofmt`工具格式化时按照字母顺序对包名进行排序。

### 运算符

对数据类型，Go语言提供了常规的数值和逻辑运算符。
和`Python`与`JavaScript`类似，对`string`类型，Go语言提供了`+`运算符，用于拼接字符串。

### func

```go
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
```

`func`关键字用于声明一个函数，函数是go程序的基本执行单元。在这个例程中，我们声明了一个`main`函数，代表了该程序的入口。

在`main`函数中，我们声明了两个字符串变量`s`和`sep`。

对变量而言，变量会在声明时直接初始化。如果没有初始化，变量会被赋予该变量类型的零值。
比如，`int`类型的零值是`0`，`string`类型的零值是`""`。

定义变量时，变量名在前，类型在后。如果变量有初始值，可以省略类型，编译器会根据初始值推断类型。

在`main`函数中，我们使用`for`关键字声明了一个循环，用于遍历`os.Args`中的参数。
在循环中，我们将`os.Args[i]`添加到`s`中，然后将`sep`设置为空格。这样，我们就可以将所有参数拼接到一个字符串中。
最后，我们使用`fmt.Println(s)`打印出拼接结果。

循环索引变量 `i` 在 `for` 循环的第一部分中定义。
符号 `:=` 是 短变量声明（short variable declaration）的一部分，
这是定义一个或多个变量并根据它们的初始值为这些变量赋予适当类型的语句。

自增语句 `i++` 给 `i` 加 `1`；这和 `i+=1` 以及 `i=i+1` 都是等价的。
对应的还有 `i--` 给 `i` 减 `1`。它们是语句，而不像 C 系的其它语言那样是表达式。
***所以 `j=i++` 非法，而且 `++` 和 `--` 都只能放在变量名后面，因此 `--i` 也非法 (好文明！！！)***

### for

```go
for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
}
```

在Go语言中，循环语句只有`for`关键字，没有`while`和`do-while`关键字。

`for`关键字后面的第一个语句是初始化语句，第二个语句是条件语句，第三个语句是后置语句。
这三条语句都可以省略。

```go
for initialization; condition; post {
    // code
}
```

格式上，`for`循环三个部分不需要括号包围，但是循环体必须用大括号包围。
左大括号必须和`for`在同一行，不能单独成行。

`initialization`是可选的，在循环开始前执行。
`initialization`如果存在，必须是一个简单语句，即一个短变量声明、一个自增语句、一个赋值语句或函数调用。

`condition`是可选的一个布尔表达式，每次迭代开始前计算。
如果为`ture`，执行循环体；如果为`false`，退出循环。

`post`是可选的，在循环体执行完毕后执行。
之后再对`condition`求值，如果为`true`，继续执行循环体；如果为`false`，退出循环。

下面是省略了`initialization`和`post`的例子，如果省略了`initialization`和`post`，分号也必须省略。

```go
for condition {
    // ...
}
```

如果省略了`condition`，那么它默认为`true`，这样就变成了一个无限循环。

```go
for {
    // ...
}
```

可以使用`break`和`return`语句来退出循环。

## `for`循环的另一种形式： `range`关键字

见[`echo2.go`](./echo2/echo2.go)。

```go
// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
```

在本例中`for`循环在某种数据类型的区间（range）上遍历，如字符或切片。
`range`关键字返回两个值，第一个是索引，第二个是索引对应的值。
`range`关键字规定要处理元素，则必须处理索引。
那么，就有一种可能使用临时变量存储索引然后忽略它。

```go
for i, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
}
```

然而，go语言不允许使用无用的局部变量，如果使用则会导致编译错误。
在本例中。我们不需要索引，所以用`_` ***（空标识符）*** 忽略索引。

空标识符`_` 可以用来忽略值，这样可以避免声明不需要的变量。

## 短变量声明 `:=`

在Go语言中，`:=`是短变量声明（short variable declaration）的一部分。
这是定义一个或多个变量并根据它们的初始值为这些变量赋予适当类型的语句。

下面这些都等价：
```go
s := ""
var s string
var s = ""
var s string = ""
```

短变量声明只能在函数内部使用，而不能在包级别使用。
在函数内部，`:=`可以用来声明新的变量，也可以用来给已经声明的变量赋值。

## 调优

在`echo1`中，每次循环迭代字符串 `s` 的内容都会更新。
`+=` 连接原字符串、空格和下个参数，产生新字符串，并把它赋值给 `s`。
`s` 原来的内容已经不再使用，将在适当时机对它进行垃圾回收（GC）。

这种方法的问题在于，每次迭代都会产生一个新的字符串，这样会产生很多临时字符串，对性能有影响。

如果参数很多，这种方法代价很高，因为每次迭代都会产生一个新的字符串。

一种简单高效的方法是使用 `strings.Join` 函数。

[`echo3.go`](./echo3/echo3.go)：

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
```
其中，`os.Args[1:]` 表示从 `os.Args` 的第二个元素开始到最后一个元素的切片。
`" "` 是一个空格，用于分隔参数。

`strings.Join` 函数使用指定的分隔符连接字符串切片，返回一个新的字符串。

`strings.Join` 函数的第一个参数是一个字符串切片，第二个参数是一个分隔符。
