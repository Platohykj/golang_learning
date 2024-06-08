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
