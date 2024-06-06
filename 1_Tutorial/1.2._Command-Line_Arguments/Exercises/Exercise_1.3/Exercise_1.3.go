package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()
	fmt.Println("运行时间", end.Sub(now))

	now2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end2 := time.Now()
	fmt.Println("运行时间", end2.Sub(now2))

}
