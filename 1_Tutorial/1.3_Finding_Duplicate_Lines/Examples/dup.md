# dup

## dup1
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

### import

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

### bufio.NewScanner

```go
input := bufio.NewScanner(os.Stdin)
```

`bufio.NewScanner`函数创建了一个`bufio.Scanner`类型的值，该值从`os.Stdin`中读取数据。

`bufio.Scanner`类型的值可以通过`Scan`方法读取数据，`Scan`方法在读取数据时会自动去掉行尾的换行符。

### map

```go
counts := make(map[string]int)
```

`map` 是一个由 `make` 函数创建的数据结构的**引用**。
`map` 作为参数传递给某函数时，该函数接收这个**引用的一份拷贝**（copy，或译为副本），
被调用函数对 `map` 底层数据结构的任何修改，调用者函数都可以通过持有的 `map` 引用看到。
在我们的例子中，`countLines` 函数向 `counts` 插入的值，也会被 `main` 函数看到。

`counts`是一个`map`类型的变量，用于存储行的计数。

`map`是一种无序的键值对集合，`map`的键是唯一的，`map`的值可以是任意类型。

在这个例程中，`counts`的键是行的内容，`counts`的值是行的计数。

### input.Scan

```go
for input.Scan() {
    counts[input.Text()]++
}
```

`input.Scan`方法读取下一行数据，如果读取成功，则返回`true`，否则返回`false`。

在这个例程中，我们使用`for`循环读取所有行，然后将行的计数存储到`counts`中。

### input.Text()

```go
counts[input.Text()]++
```
`input.Text()`方法返回最近一次读取的行。

在这个例程中，我们使用`input.Text()`方法获取行的内容，然后将行的计数存储到`input.Text()`对应的值中

首次读到新行时，`counts[input.Text()]`的值为`0`，`counts[input.Text()]++`的值为`1`。

### for line, n := range counts

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

### 转义字符

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

## dup2

[dup2.go](./dup2/dup2.go)程序读取文件，然后输出重复的行。

```go
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
}
```
### os.Open

```go
f, err := os.Open(arg)
```

`os.Open`函数打开一个文件，返回两个值，第一个值是被打开的文件`*os.File`,其后被`Scanner`获取，第二个值是错误信息。

在这个例程中，我们使用`os.Open`函数打开文件，然后将文件的内容存储到`f`中。

### error

```go
if err != nil {
    fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
    continue
}
```

`os.Open`函数打开文件失败时，会返回一个错误信息,为内置的`error`类型。

如果`err`等于内置值`nil`（相当于其他语言中的`NULL`），则表示文件被成功打开。

如果`err`值不是`nil`，则表示文件打开失败，我们可以通过`fmt.Fprintf`函数将错误信息输出到标准错误输出。

```go
    fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
```

### countLines

`countLines` 函数在其声明前被调用。
函数和包级别的变量（package-level entities）可以任意顺序声明，并不影响其被调用。

## dup3

[dup1.go](./dup1/dup1.go)和[dup2.go](./dup2/dup2.go)程序都是一次性读取所有输入，然后处理输入。

[dup3.go](./dup3/dup3.go)程序则是一次性读取所有输入到内存中，一次分割为多行，然后处理输入。

```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```
`ReadFile` 函数需要文件名作为参数，因此只读指定文件，不读标准输入。

### ioutil.ReadFile

```go
data, err := ioutil.ReadFile(filename)
```

`ioutil.ReadFile`函数读取文件的所有内容，返回两个值，第一个值是文件的内容，第二个值是错误信息。

在这个例程中，我们使用`ioutil.ReadFile`函数读取文件的内容，然后将文件的内容存储到`data`中。

### strings.Split

```go
for _, line := range strings.Split(string(data), "\n") {
    counts[line]++
}
```

`strings.Split`函数将字符串分割为多个子字符串，返回一个字符串的切片。

在这个例程中，我们使用`strings.Split`函数将文件的内容分割为多行，然后将行的计数存储到`counts`中。

`Split` 的作用与前文提到的 `strings.Join` 相反。

