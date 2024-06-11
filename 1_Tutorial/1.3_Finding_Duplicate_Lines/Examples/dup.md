# dup

基于`Unix`系统的`uniq`命令，我们可以编写一个`Go`程序来查找重复的行。

[dup1.go](./dup1/dup1.go)程序读取标准输入，然后输出重复的行。
```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```

我们可以使用以下命令来测试`dup1.go`程序。

Windows
```shell
echo "a`nb`nc`na`nb`na" | go run dup1.go
```

linux
```shell
echo "a\nb\nc\na\nb\na" | go run dup1.go
```

## import

```go
package main
import (
    "bufio"
    "fmt"
    "os"
)
```

本例程使用了`bufio`、`fmt`、`os`包。

其中，`bufio`包提供了缓冲读取功能，`fmt`包提供了格式化输入输出功能，`os`包提供了操作系统函数。

## bufio.NewScanner

```go
input := bufio.NewScanner(os.Stdin)
```

`bufio.NewScanner`函数创建了一个`bufio.Scanner`类型的值，该值从`os.Stdin`中读取数据。

`bufio.Scanner`类型的值可以通过`Scan`方法读取数据，`Scan`方法在读取数据时会自动去掉行尾的换行符。

## map

```go
counts := make(map[string]int)
```

`counts`是一个`map`类型的变量，用于存储行的计数。

`map`是一种无序的键值对集合，`map`的键是唯一的，`map`的值可以是任意类型。

在这个例程中，`counts`的键是行的内容，`counts`的值是行的计数。

## input.Scan

```go
for input.Scan() {
    counts[input.Text()]++
}
```

`input.Scan`方法读取下一行数据，如果读取成功，则返回`true`，否则返回`false`。

在这个例程中，我们使用`for`循环读取所有行，然后将行的计数存储到`counts`中。

## input.Text()

```go
counts[input.Text()]++
```
`input.Text()`方法返回最近一次读取的行。

在这个例程中，我们使用`input.Text()`方法获取行的内容，然后将行的计数存储到`input.Text()`对应的值中

首次读到新行时，`counts[input.Text()]`的值为`0`，`counts[input.Text()]++`的值为`1`。

## for line, n := range counts

```go
for line, n := range counts {
    if n > 1 {
        fmt.Printf("%d\t%s\n", n, line)
    }
}
```

`for line, n := range counts`循环遍历`counts`中的所有键值对，并格式化输出

其中，`\t` `\n`被Go程序员称之为动词，下面的表格列出了常见的动词：

| 动词         | 说明                                           |
|------------|----------------------------------------------|
| %d         | 十进制整数                                        |
| %x, %o, %b | 十六进制，八进制，二进制整数                               |
| %f, %g, %e | 浮点数： 3.141593 3.141592653589793 3.141593e+00 |
| %t         | 布尔：true或false                                |
| %c         | 字符（rune） (Unicode码点)                         |
| %s         | 字符串                                          |
| %q         | 带双引号的字符串 "abc" 或 带单引号的字符 'c'                 |
| %v         | 变量的自然形式（natural format）                      |
| %T         | 变量的类型                                        |
| %%         | 字面上的百分号标志（无操作数）                              |

## 转义字符

在Go语言中，字符串中的转义字符有以下几种：

| 转义字符 | 含义     |
|------|--------|
| \a   | 响铃     |
| \b   | 退格     |
| \f   | 换页     |
| \n   | 换行     |
| \r   | 回车     |
| \t   | 制表符    |
| \v   | 垂直制表符  |
| \\   | 反斜杠    |
| \'   | 单引号    |
| \"   | 双引号    |
| \?   | 问号     |
| \0   | 空字符    |
| \ddd | 八进制字符  |
| \xhh | 十六进制字符 |

按照惯例，以字母`ln`结尾的格式化函数采用`fmt.Println`的格式化准则，以字母`f`结尾的格式化函数采用`fmt.Fprintf`的格式化准则。
