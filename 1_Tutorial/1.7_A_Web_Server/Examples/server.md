# Server

## [server1.go](./server1/server1.go)

```go
// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

```

这是一个利用go语言官方库实现的一个简单的`echo`服务

### http.HandleFunc

```go
http.HandleFunc("/", handler) // each request calls handler
```

`http.HandleFunc`函数将所有发送到`/`的请求重定向到`handler`函数.

### http.ListenAndServe

```go
log.Fatal(http.ListenAndServe("localhost:8000", nil))
```

`http.ListenAndServe`函数监听`localhost:8000`端口，接收请求并调用`handler`函数处理请求.

### handler

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
```

`handler`函数接收两个参数，一个是`http.ResponseWriter`，一个是`http.Request`，并将请求的`URL.Path`写入到`http.ResponseWriter`中.

## [server2.go](./server2/server2.go)

```go