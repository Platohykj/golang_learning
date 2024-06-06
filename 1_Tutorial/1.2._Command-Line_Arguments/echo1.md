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


