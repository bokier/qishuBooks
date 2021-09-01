package main

import (
	"fmt"
	"qishuBook/engine"
	"time"
)

const (
	qiShuUrl = "https://www.kankezw.com/soft/sort010/"
)

func main() {
	t := time.Now()
	engine.RunEngineV1(qiShuUrl)
	// engine.GetBookInfosTest("https://www.kankezw.com/soft/sort011/index_2.html")
	elapsed := time.Since(t)
	fmt.Println(elapsed)
}