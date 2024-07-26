# fetchall

Go语言原生支持并发，我们可以使用`goroutine`和`channel`实现并发编程。

在这个例程中，我们使用`goroutine`和`channel`实现了并发的URL获取。

```go
// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }
    nbytes, err := io.Copy(io.Discard, resp.Body)
    resp.Body.Close() // don't leak resources
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
```

## gorutine

`goroutine`是Go语言并发编程的基本单元,相当于一种轻量化的线程。

`go`关键字会创建一个新的`goroutine`，并在这个`goroutine`中执行函数。main函数本身也运行在一个`goroutine`中。

## channel

`channel`是Go语言并发编程的通信机制，可以让`goroutine`之间进行通信。

## 通信阻塞机制

当一个 goroutine 尝试向一个 channel 发送数据时，如果没有其他 goroutine 正在从这个 channel 接收数据，那么发送操作会阻塞，直到有一个 goroutine 开始接收数据。同理，当一个 goroutine 尝试从一个 channel 接收数据时，如果没有其他 goroutine 向这个 channel 发送数据，那么接收操作会阻塞，直到有数据被发送到 channel。

通信阻塞机制较好的防止了有两个`goroutine`同时访问共享数据的情况，避免了数据竞争。