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

## counts

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